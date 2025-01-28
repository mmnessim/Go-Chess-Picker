FROM golang:1.23
WORKDIR /src
COPY . .
RUN go build -o /bin/go-chess ./main.go

FROM alpine:latest
COPY --from=0 /bin/go-chess /app/go-chess
CMD ["/app/go-chess"]
