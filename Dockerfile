FROM golang:1.23
WORKDIR /src
COPY . .
RUN go build -o /bin/chess

FROM scratch
COPY --from=0 /bin/chess /bin/chess
CMD ["/bin/chess"]


