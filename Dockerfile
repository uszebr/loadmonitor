# Stage 1 - Build
FROM golang:1.23.2-alpine AS build

WORKDIR /app

# Copy all files to the container
COPY . .

# Download and verify Go dependencies
RUN go mod download && go mod verify

# Build the application binary
RUN go build -o /app/loadmonitor ./cmd/loadmonitor.go

# Stage 2 - Distroless
FROM gcr.io/distroless/base-debian11

# Copy the binary from the builder stage to the distroless image
COPY --from=build /app/loadmonitor /loadmonitor
COPY --from=build /app/assets /assets

# Set the entrypoint
ENTRYPOINT ["/loadmonitor"]