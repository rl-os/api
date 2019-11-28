FROM golang:1.13-alpine3.10 AS builder

ENV GO111MODULE="on"

RUN apk add --no-cache bash

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/deissh/osu-api-server
COPY . ./
RUN ./scripts/install.sh
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /server .

FROM scratch
COPY --from=builder /server ./server
COPY config.yaml config.yaml

EXPOSE 2100
ENTRYPOINT ["./server"]
