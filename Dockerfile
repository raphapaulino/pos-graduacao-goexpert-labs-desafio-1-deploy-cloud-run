FROM golang:1.22.2 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/cloudrun .
EXPOSE 8080
ENTRYPOINT ["./cloudrun"]