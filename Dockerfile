FROM golang:latest as build
WORKDIR /app
COPY ./main.go ./main_test.go ./go.mod /app/
RUN go build -o goapp

FROM alpine
WORKDIR /binApp

COPY --from=build /app/goapp ./
