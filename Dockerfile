FROM golang:1.14-alpine as build
RUN apk add --no-cache gcc musl-dev tzdata
ADD . /app
WORKDIR /app
RUN go test
RUN go build

FROM alpine:3.11
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /app/mailygo /bin/
WORKDIR /app
EXPOSE 8080
CMD ["mailygo"]