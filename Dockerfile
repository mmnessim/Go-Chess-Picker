FROM golang:1.23
WORKDIR /app
COPY . .
RUN go build -o bin

FROM scratch
COPY --from=0 /app/bin /
ENTRYPOINT ["/bin"]
