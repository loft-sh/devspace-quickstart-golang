FROM golang:1.12

# Create go project directory
RUN mkdir -p /go/src/app && ln -s /go/src/app /app && cd /go/src/app

# Add source
ADD . /go/src/app

# Build application
RUN go get ./... && go build -o app . && chmod +x app && cd /app 

WORKDIR /app

EXPOSE 8080

# Run application
CMD ["/go/src/app/app"]
