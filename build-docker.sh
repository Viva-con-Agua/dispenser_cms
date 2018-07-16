set -ex

# SET THE FOLLOWING VARIABLES
# docker hub username
USERNAME=vivaconagua
# image name
IMAGE=ripple-docker

docker build -t $USERNAME/$IMAGE:latest .
