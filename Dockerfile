# This Dockerfile builds an image from the statically linked
# application, assuming that it exists in the path.
# The image is based on "scratch" for size reasons and should be
# approximately 5 MB light.

FROM scratch
COPY go-webget /
ENTRYPOINT ["/go-webget"]
