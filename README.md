# GoDockerHttpServer

## Description
This project simplifies the process of building and running a Docker image for a Go HTTP server using a Makefile.

## Targets
The following targets are available:

- `build`: Build a Docker image for the Go HTTP server.
- `run`: Run a Docker container from the built image.
- `clean`: Remove the Docker image.
- `help`: Show the available targets and their descriptions.

## Environment Variables
The project uses a `.env` file to store environment variables. These variables are used to configure the Docker image and container.

## How to Use
1. Create a `.env` file in the project root with the required environment variables.
2. Run `make build` to build the Docker image.
3. Run `make run` to start a Docker container from the built image.
4. Run `make clean` to remove the Docker image.

## Note
This project assumes that you have Docker installed on your system. If you don't have Docker installed, you can download it from the official Docker website.
