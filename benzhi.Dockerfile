# Complete Go toolchain remains available for in-container development and tests.
FROM golang:1.22

WORKDIR /app

# This project uses only the Go standard library, so it has no go.sum or
# external modules to download.
COPY . .

# Prime the compiler cache and verify the initial source builds.
RUN go build ./...

CMD ["bash"]
