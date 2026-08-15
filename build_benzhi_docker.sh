#!/bin/bash
# Ensure the script exits immediately if any command fails.
set -e

# Use the first argument as the image name, defaulting to "my-project".
IMAGE_NAME=${1:-my-project}
# Use the second argument as the target platform, defaulting to "linux/amd64".
DOCKER_PLATFORM=${2:-linux/amd64}

# Build with the evaluation Dockerfile rather than any project Dockerfile.
docker build --platform $DOCKER_PLATFORM -f benzhi.Dockerfile -t $IMAGE_NAME .

echo ""
echo "Docker image '$IMAGE_NAME' built successfully!"
echo ""
echo "Next steps (for testing):"
echo "  * Interactive shell: docker run -it $IMAGE_NAME:latest"
echo ""
