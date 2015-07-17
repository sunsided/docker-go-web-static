# go-webget

A simple Go application that issues an HTTPS GET request to [https://www.google.com](https://www.google.com) and prints the response. The application is built and statically linked using the [centurylink/golang-builder](https://github.com/CenturyLinkLabs/golang-builder) docker image which optionally also creates a slim docker image `FROM scratch` containing the application.

## Build and statically link the application

Issue the following command to build the application and place it in the working directory.

```bash
./build
```

## Build a docker image containing the statically linked application

Issue the following command to build the application and additionally create a new docker image based on the [Dockerfile](./Dockerfile) provided in the working directory.

```bash
./build-image
```

### Run the docker container

In order to run a container off the built image, the following command can be issued:

```bash
docker run --rm go-webget
```
