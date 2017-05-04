FROM golang:1.8 AS build-env
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ADD . /go/src/github.com/cpuguy83/dockerhub-webhook-listener
WORKDIR /go/src/github.com/cpuguy83/dockerhub-webhook-listener/hub-listener
RUN go get && go build

FROM scratch 
ADD ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/src/github.com/cpuguy83/dockerhub-webhook-listener/hub-listener/hub-listener /hub-listener
ENTRYPOINT ["/hub-listener"]
CMD ["-listen", "0.0.0.0:80"]

