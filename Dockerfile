FROM golang:1.13 as builder

ENV GO111MODULE=on
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

ENV DB_URL=mongodb://localhost:27018
COPY --from=builder /app/qaa_api /app/
EXPOSE 3000
ENTRYPOINT ["/app/qaa_api"]