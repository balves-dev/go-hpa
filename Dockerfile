FROM golang:1.14-alpine as builder

WORKDIR /app
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/myapp .

FROM scratch
COPY --from=builder /app/myapp .

EXPOSE 8000

CMD [ "./myapp" ]