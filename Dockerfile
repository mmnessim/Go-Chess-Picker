FROM golang:latest
WORKDIR /build
COPY ./go.mod go.mod
RUN go mod tidy
COPY  . .
RUN go build -o ./app

FROM debian:bookworm-slim
WORKDIR /chess
COPY --from=0 /build .
ENTRYPOINT ["./app"]
