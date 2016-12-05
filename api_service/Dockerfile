# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.7

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/distelli/api_service

#build the application
RUN go get github.com/tools/godep
RUN cd /go/src/github.com/distelli/api_service && godep go build

# Run the command by default when the container starts.
ENTRYPOINT /go/src/github.com/distelli/api_service/api_service

# Document that the service listens on port 8001
EXPOSE 8001
