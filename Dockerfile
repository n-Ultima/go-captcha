FROM golang:alpine
COPY . /app
WORKDIR /app
RUN go build -o app cmd/server/main.go
CMD [ "./app" ]