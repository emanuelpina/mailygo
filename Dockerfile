FROM golang:alpine as build
RUN apk add --no-cache gcc musl-dev tzdata
ADD . /app
WORKDIR /app
RUN go build

FROM alpine:latest
RUN apk add --no-cache tzdata ca-certificates
COPY --from=build /app/mailygo /bin/
WORKDIR /app
EXPOSE 8080
CMD ["mailygo"]