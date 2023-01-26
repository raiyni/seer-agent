FROM golang:1.19.5-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /seer-agent

EXPOSE 3333
CMD [ "/seer-agent" ]