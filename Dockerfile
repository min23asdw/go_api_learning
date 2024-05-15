FROM golang:1.22.2 AS build-stage
  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  # Copy all Go source files, including those in subdirectories
  COPY **/*.go ./
  
  # RUN CGO_ENABLED=0 GOOS=linux go build -o /api
  # RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
  RUN CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  go build -o /api

  # Run the tests in the container
  RUN chmod -x /api
  RUN chmod +x /api
FROM build-stage AS run-test-stage
  RUN go test -v ./...

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage
  WORKDIR /
  


  COPY --from=build-stage /api /api

  EXPOSE 8080
  
  ENTRYPOINT ["/api"]

# syntax=docker/dockerfile:1

# FROM golang:1.22.2
# RUN mkdir -p /app
# # Set destination for COPY
# WORKDIR /app

# # Download Go modules
# COPY go.mod go.sum ./
# RUN go mod download

# # Copy the source code. Note the slash at the end, as explained in
# # https://docs.docker.com/reference/dockerfile/#copy
# # COPY **/*.go ./
# COPY . .
# # Build

# # RUN CGO_ENABLED=0 GOOS=linux go build  -o /go_api_learning
# # RUN chmod -R 777 /go_api_learning
# ENV GOPATH /app

# RUN go build
# # Optional:
# # To bind to a TCP port, runtime parameters must be supplied to the docker command.
# # But we can document in the Dockerfile what ports
# # the application is going to listen on by default.
# # https://docs.docker.com/reference/dockerfile/#expose
# EXPOSE 8080

# # Run
# ENTRYPOINT /app/bin/main
# # CMD ["/go_api_learning"]