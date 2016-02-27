# Dockerfile extending the generic Go image with application files for a
# single application.
FROM gcr.io/google_appengine/base

RUN apt-get update \
    && apt-get install -y \
       curl gcc git libc6-dev make \
       --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

ENV GO_VERSION 1.5
ENV GO_WRAPPER_COMMIT 6ea1f29b1fe7e6b0b8eb89493ed5e06bac454654

RUN curl -sSL https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz \
    | tar -v -C /usr/local -xz

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go:/go/src/app/_gopath
ENV GO15VENDOREXPERIMENT 1

RUN mkdir -p /go/src/app /go/bin && chmod -R 777 /go

RUN curl https://raw.githubusercontent.com/docker-library/golang/${GO_WRAPPER_COMMIT}/1.5/go-wrapper \
    -o /usr/local/bin/go-wrapper \
    && chmod 755 /usr/local/bin/go-wrapper

RUN ln -s /go/src/app /app
WORKDIR /go/src/app

COPY . /go/src/app

# Hacky since I had problem with GOPATH :(
RUN mv /go/src/app/src/webp_demo /go/src/app/_gopath/src/webp_demo

RUN go-wrapper install -tags appenginevm

# ENTRYPOINT go-wrapper run main.go
CMD ["go-wrapper", "run"]

# Open up the port where the app is running.
EXPOSE 8080
