FROM golang:1.18-alpine

WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go install .


CMD [ "errorcodes" ]
