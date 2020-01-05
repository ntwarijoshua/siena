FROM golang:latest

LABEL maintainer="Ntwari Joshua <ntwarijoshua@gmail.com>"
WORKDIR /app
COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ .
RUN go build -o cmd/app cmd/main.go

RUN rm .env && mv .env.container .env

EXPOSE 8090

CMD [ "./cmd/app" ]