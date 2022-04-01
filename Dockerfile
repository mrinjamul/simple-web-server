# Start from the latest golang base image
FROM golang:1.18-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /src
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Add the source from the current directory to the Working Directory inside the container
ADD . .
# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest


# Add Maintainer Info
LABEL maintainer="Injamul Mohammad Mollah <mrinjamul@gmail.com>"
# Set the Current Working Directory inside the container
WORKDIR /home/app
# Add CA Certificates
RUN apk --no-cache add ca-certificates
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /src/main /usr/local/bin/sws
# Copy config file
COPY --from=builder /src/config.json.example /etc/sws/config.json
# Command to run the executable
CMD ["sws"]
