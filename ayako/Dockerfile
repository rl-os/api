FROM golang:1.13-alpine3.10 AS builder

ENV GO111MODULE="on"

RUN apk add --no-cache bash make git curl
RUN curl -fsSL -o /dbmate https://github.com/amacneil/dbmate/releases/download/v1.8.0/dbmate-linux-amd64
RUN chmod +x /dbmate

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/osu-lazer/ayako
COPY ./ayako/go.mod ./ayako/go.sum ./ayako/Makefile ./
RUN make install

# Copy .git folder (using to get versions and last commit id)
WORKDIR $GOPATH/src/github.com/deissh/osu-lazer
COPY .git .git

WORKDIR $GOPATH/src/github.com/deissh/osu-lazer/ayako
COPY ./ayako .
RUN make build-prod && mv bin/* /


FROM alpine:3.8
RUN apk add --no-cache bash curl

RUN curl -fsSL -o /bin/dbmate https://github.com/amacneil/dbmate/releases/download/v1.8.0/dbmate-linux-musl-amd64
RUN chmod +x /bin/dbmate

COPY --from=builder /ayako server
COPY ./ayako/config.yaml .

COPY ./ayako/docker-init.sh docker-init.sh

EXPOSE 2400
CMD ["./docker-init.sh", "&&", "./server"]
