package main

// nolint[gocyclo]

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	docker "github.com/fsouza/go-dockerclient"
)

type options struct {
	dockerHost                    string
	sweepInterval                 int
	removeImages                  bool
	removeHealthyContainersExited bool
}

var opts options

func init() {
	flag.StringVar(&opts.dockerHost, "docker-host", "unix:///var/run/docker.sock", "-docker-host=unix:///var/run/docker.sock")
	flag.IntVar(&opts.sweepInterval, "sweep-interval", 60, "-sweep-interval=60")
	flag.BoolVar(&opts.removeImages, "remove-images", true, "-remove-images=true")
	flag.BoolVar(&opts.removeHealthyContainersExited, "remove-healthy-containers-exited", true, "-remove-healthy-containers-exited=true")
}

func main() {
	flag.Parse()

	dc, err := docker.NewClient(opts.dockerHost)
	if err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGTERM, syscall.SIGINT)

	fmt.Fprintf(os.Stdout, "gcd: [info]: (Time: %v)\n", time.Now().String())
	fmt.Fprintf(os.Stdout, "gcd: [info]: (Docker Host: %v)\n", opts.dockerHost)
	fmt.Fprintf(os.Stdout, "gcd: [info]: (Sweep Interval: %vs)\n", opts.sweepInterval)
	fmt.Fprintf(os.Stdout, "gcd: [info]: (Remove Images: %v)\n", opts.removeImages)
	fmt.Fprintf(os.Stdout, "gcd: [info]: (Remove Healthy Containers Exited: %v)\n", opts.removeHealthyContainersExited)

	for {
		select {
		case <-s:
			os.Exit(0)
		case <-time.Tick(time.Duration(opts.sweepInterval) * time.Second):
			run(dc, opts)
		}
	}
}

func run(dc *docker.Client, opts options) {
	fmt.Fprintf(os.Stdout, "\ngcd: [info]: (Time: %v)\n", time.Now().String())
	containers, err := dc.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	var wgContainers sync.WaitGroup
	for _, container := range containers {
		wgContainers.Add(1)
		go func(container docker.APIContainers) {
			defer wgContainers.Done()
			exitCodeFromContainer := "(-)"
			if splitedStatus := strings.Split(container.Status, " "); len(splitedStatus) > 1 {
				exitCodeFromContainer = splitedStatus[1]
			}
			if container.State != "running" {
				if (opts.removeHealthyContainersExited && exitCodeFromContainer == "(0)") || exitCodeFromContainer != "(0)" {
					fmt.Fprintf(os.Stdout, "gcd: [trying remove container]: (Id: %v, Labels: %v)\n", container.ID, container.Labels)
					if err := dc.RemoveContainer(docker.RemoveContainerOptions{
						ID:            container.ID,
						RemoveVolumes: true,
						Force:         true,
					}); err != nil {
						fmt.Fprintf(os.Stderr, "gcd: [error]: (when try remove container, reason: %v)\n", err.Error())
					} else {
						fmt.Fprintf(os.Stdout, "gcd: [removed container]: (Id: %v, Labels: %v)\n", container.ID, container.Labels)
					}
				}
			}
		}(container)
	}
	wgContainers.Wait()
	if opts.removeImages {
		var wgImages sync.WaitGroup
		images, err := dc.ListImages(docker.ListImagesOptions{})
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		for _, image := range images {
			wgImages.Add(1)
			go func(image docker.APIImages) {
				defer wgImages.Done()
				fmt.Fprintf(os.Stdout, "gcd: [trying remove image]: (Id: %v, Tags: %v)\n", image.ID, image.RepoTags)
				if err := dc.RemoveImage(image.ID); err != nil {
					fmt.Fprintf(os.Stderr, "gcd: [error]: (when try remove image, reason: %v)\n", err.Error())
				} else {
					fmt.Fprintf(os.Stdout, "gcd: [removed image]: (Id: %v, Tags: %v)\n", image.ID, image.RepoTags)
				}
			}(image)
		}
		wgImages.Wait()
	}
}
