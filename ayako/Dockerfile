FROM golang:1.13-alpine3.10 AS builder

ENV GO111MODULE="on"

RUN apk add --no-cache bash make git

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

FROM scratch
WORKDIR /root/
COPY --from=builder /ayako ./server
COPY ./ayako/config.yaml .

EXPOSE 2400
CMD ["/root/server"]
