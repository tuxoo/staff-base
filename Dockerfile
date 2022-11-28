FROM golang:1.19.3-alpine

RUN go version
ENV GOPATH=/
ENV APP_PATH=/home/src

EXPOSE ${HTTP_PORT}
VOLUME $APP_PATH/logs
WORKDIR $APP_PATH

COPY ./ $APP_PATH

RUN go mod download
RUN go build -o staff-base $APP_PATH/cmd/main.go

ENTRYPOINT ["./staff-base"]