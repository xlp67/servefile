FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init servefile
RUN go mod tidy
RUN go build -o servefile

FROM golang:alpine
WORKDIR /app
COPY --from=builder /app/servefile .
COPY --from=builder /app/config.json .
COPY --from=builder /app/.env .
CMD [ "./servefile" ]