FROM golang:1.13.1 AS build-env
WORKDIR /RoomStatus
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY go.mod /RoomStatus/go.mod
COPY go.sum /RoomStatus/go.sum
RUN go mod download
COPY . /RoomStatus
RUN CGO_ENABLED=0 GOOS=linux go build -o build/RoomStatus ./RoomStatus


FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /RoomStatus/build/RoomStatus /
ENTRYPOINT ["/RoomStatus"]
CMD ["up", "--grpc-port=80"]
