# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.14.1

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/distelli/client_service

#build the application
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN cd /go/src/github.com/distelli/client_service &&  go build

# Run the command by default when the container starts.
ENTRYPOINT /go/src/github.com/distelli/client_service/client_service

# Document that the service listens on port 8080
EXPOSE 8080
