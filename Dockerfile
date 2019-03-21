FROM golang:1.12

# Create go project directory
RUN mkdir -p /go/src/app && ln -s /go/src/app /app

# Add source
ADD . /go/src/app

# Build application
RUN cd /go/src/app && go get ./... && go build -o app && chmod +x app

WORKDIR /app

EXPOSE 8080

# Run application
CMD ["/go/src/app/app"]
