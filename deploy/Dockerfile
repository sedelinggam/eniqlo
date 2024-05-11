# build
FROM golang:1.22.0-alpine as build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main_pinginciuman cmd/main.go

# deploy
FROM alpine:latest

WORKDIR /app

COPY .env .
COPY --chown=65534:65534 --from=build usr/src/app/main_pinginciuman .
USER 65534

EXPOSE 8080

ENTRYPOINT [ "./main_pinginciuman" ]