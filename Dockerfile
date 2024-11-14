FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build .

FROM debian:bookworm-slim AS run

COPY --from=build /app/uptime-server /bin/uptime-server
EXPOSE 8080
CMD [ "uptime-server" ]
