go test -v -cover $(go list ./... | grep -v /vendor/)