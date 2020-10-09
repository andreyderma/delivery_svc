FROM golang:1.13.6 AS build-env

RUN mkdir -p /delivery-service
ENV GO111MODULE=on
WORKDIR /delivery-service
COPY . /delivery-service

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag i
RUN go test -v -count=1 ./... -coverprofile=c.out

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ./delivery

FROM alpine:3.7
WORKDIR /app
RUN apk --no-cache add tzdata su-exec ca-certificates s6 curl openssh
COPY --from=build-env /delivery-service/delivery /app/
ENTRYPOINT ["/app/delivery"]