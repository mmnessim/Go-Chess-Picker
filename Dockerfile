FROM golang:1.23
WORKDIR /src
COPY . .
RUN go build -o /bin/go-chess ./main.go

FROM scratch
COPY --from=0 /bin/go-chess /bin/go-chess
CMD ["/bin/go-chess"]
