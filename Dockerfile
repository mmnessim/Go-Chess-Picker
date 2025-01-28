FROM golang:1.23
WORKDIR /src
COPY . .
RUN go build

FROM scratch
COPY --from=0 ./go-chess /bin
CMD ["/bin/go-chess"]
