FROM golang:latest as build

WORKDIR /build
COPY . . 
RUN mkdir -p ./bin && go build -o ./bin/tibia-api-exporter && chmod +x ./bin/tibia-api-exporter

FROM debian:latest

COPY --from=build /build/bin/tibia-api-exporter /usr/bin/tibia-api-exporter
WORKDIR /app/

CMD ["tibia-api-exporter"]
