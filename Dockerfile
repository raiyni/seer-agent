FROM golang:1.19.5-alpine as builder
WORKDIR /go/src/github.com/raiyni/seer-agent

COPY go.mod ./
COPY go.sum ./
COPY *.go ./

RUN go install
RUN go build -o /seer-agent


FROM alpine:3.17
COPY --from=builder /seer-agent ./

VOLUME ["/var/run/docker.sock"]
EXPOSE 3333
CMD [ "/seer-agent" ]