FROM golang:stretch

WORKDIR /src/
COPY /src/* /src

RUN cd /src && CGO_ENABLED=0 go build -o /root/server ./main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /root/server .
COPY ./dist/ ./dist

EXPOSE ${PORT}
ENTRYPOINT ["./server"]