# syntax=docker/dockerfile:1

FROM golang:1.18.4-alpine

WORKDIR /opt/udm_pro_exporter

COPY LICENSE ./
COPY go.mod ./
COPY go.sum ./
COPY unifi_api ./unifi_api
COPY main.go ./

RUN go mod download
RUN go build

EXPOSE 9182

CMD [ "./udm_pro_exporter" ]