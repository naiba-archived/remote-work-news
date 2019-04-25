FROM golang:alpine AS binarybuilder
# Install build deps
RUN apk --no-cache --no-progress add --virtual build-deps build-base git linux-pam-dev
WORKDIR /rwn/
COPY . .
RUN go build -o rwn -ldflags="-s -w -X github.com/naiba/remote-work-news.BuildVersion=`git rev-parse HEAD`" app/main.go

FROM alpine:latest
RUN echo http://dl-2.alpinelinux.org/alpine/edge/community/ >>/etc/apk/repositories && apk --no-cache --no-progress add \
  tzdata \
  libstdc++ \
  ca-certificates
# Copy binary to container
WORKDIR /rwn
COPY resource ./resource
COPY --from=binarybuilder /rwn/rwn .
# Configure Docker Container
VOLUME ["/rwn/data"]
EXPOSE 8080
CMD ["/rwn/rwn"]