# BENZHI Evaluation Instructions

## Project

`altitude-converter` is a Go command-line utility for converting altitude values between meters (`m`), feet (`ft`), kilometers (`km`), and nautical miles (`nm`). It accepts a single terminal value, text-file batches, and CSV-file batches.

## Dependencies

- Go 1.22 or later
- No third-party Go modules; the project uses only the Go standard library
- Docker Desktop for container-based evaluation

## Standard commands

Run these from the repository root:

```sh
go build ./...
go test ./...
go run ./cmd/altitude-converter -value 8848.86 -from m -to ft
```

## Docker build and run

The image intentionally keeps the complete Go toolchain available for editing, building, and testing inside the container.

```sh
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh altitude-converter linux/amd64
docker run -it altitude-converter:latest
```

For Apple Silicon validation, also build the arm64 image:

```sh
./build_benzhi_docker.sh altitude-converter linux/arm64
```

Inside the container, verify the project with:

```sh
go version
go build ./...
go test ./...
```

Neither the Dockerfile nor the project contains architecture-specific binaries or downloads, so the official Go base image supports both `linux/amd64` and `linux/arm64`.
