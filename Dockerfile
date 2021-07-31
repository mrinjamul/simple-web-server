# Start from the latest golang base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /src
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
# Add the source from the current directory to the Working Directory inside the container
ADD . .
# Build the Go app
RUN go build -o main .


######## Start a new stage from scratch #######
FROM scratch

# SET environments
ENV DIR=static
# Add Maintainer Info
LABEL maintainer="Injamul Mohammad Mollah <mrinjamul@gmail.com>"
# Set working directory
WORKDIR /root/
# Copy the Pre-built binary file from the previous stage
COPY --from=builder /src/main .
# Command to run the executable
CMD ["./main"]
