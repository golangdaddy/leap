FROM golang:1.21 as base

WORKDIR /app/

COPY ./go.mod /app/
COPY ./go.sum /app/
COPY ./build/* /app/

RUN GONOPROXY=github.com/golangdaddy go get -d

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .

CMD ["./main"]