FROM golang:1.22.2 AS build-stage
  WORKDIR /app

  COPY go.mod go.sum ./
  RUN go mod download

  # Copy all Go source files, including those in subdirectories
  COPY **/*.go ./
  
  # RUN CGO_ENABLED=0 GOOS=linux go build -o /api
  RUN CGO_ENABLED=0 GOOS=windows go build -o /api .

  # Run the tests in the container
  RUN chmod +x /api
FROM build-stage AS run-test-stage
  RUN go test -v ./...

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage
  WORKDIR /


  COPY --from=build-stage /api /api

  EXPOSE 8080
  
  ENTRYPOINT ["/api"]