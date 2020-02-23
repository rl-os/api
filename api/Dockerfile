FROM golang:1.13-alpine3.10 AS builder

ENV GO111MODULE="on"

RUN apk add --no-cache bash make gcc g++ git

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/osu-lazer/api
COPY . ./
RUN make install
RUN make build-prod && mv ./bin/* /

FROM scratch
COPY --from=builder /www ./www
COPY config.yaml config.yaml

EXPOSE 2100
ENTRYPOINT ["./www"]
