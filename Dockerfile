FROM golang:1.15-alpine AS builder

ENV GO111MODULE="on"

RUN apk add --no-cache bash make git curl gcc g++
RUN curl -fsSL -o /dbmate https://github.com/amacneil/dbmate/releases/download/v1.8.0/dbmate-linux-amd64
RUN chmod +x /dbmate

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/rl-os/api
COPY go.mod go.sum Makefile ./
RUN make install

WORKDIR $GOPATH/src/github.com/rl-os/api
COPY . .
RUN make build-prod && mv bin/* /


FROM alpine:3.8
RUN apk add --no-cache bash curl

RUN curl -fsSL -o /bin/dbmate https://github.com/amacneil/dbmate/releases/download/v1.8.0/dbmate-linux-musl-amd64
RUN chmod +x /bin/dbmate

COPY --from=builder /server server
COPY ./config.example.yaml config.yaml
COPY ./migrations migrations
COPY ./docker-init.sh docker-init.sh

EXPOSE 2400
CMD ["./docker-init.sh", "&&", "./server"]
