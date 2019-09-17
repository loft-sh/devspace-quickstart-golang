FROM golang:1.12

# Create go project directory
RUN mkdir -p /go/src/app && ln -s /go/src/app /app

# Add source
ADD . /go/src/app

# Build application
RUN cd /go/src/app && go get ./... && go build -o /tmp/my-app && chmod +x /tmp/my-app

# Install CompileDaemon for hot reloading
RUN go get github.com/githubnemo/CompileDaemon

WORKDIR /go/src/app

EXPOSE 8080

# Start CompileDaemon for hot reloading (will watch for file changes and then rebuild & restart the application)
ENTRYPOINT CompileDaemon -log-prefix=false -color=true -build="go build -o /tmp/my-app" -command="/tmp/my-app"