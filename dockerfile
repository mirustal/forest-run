FROM golang:latest as builder

WORKDIR /app

COPY ./go.work .
COPY ./go.work.sum .


COPY ./main-server/ ./main-server/
COPY ./common/ ./common/
RUN go work use
WORKDIR /app/main-server
RUN go work use
RUN go mod tidy

WORKDIR /app/common
RUN go work use
RUN go mod tidy

WORKDIR /app/main-server


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -p $(nproc) -a -installsuffix cgo -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main-server/main .
COPY ./main-server/conf.yaml .
COPY ./main-server/docs/swagger.json /docs/
COPY ./data/ /data/

EXPOSE 8080

CMD ["./main"]
