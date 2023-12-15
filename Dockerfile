# syntax=docker/dockerfile:1

####################################
## Build the application from source
####################################
FROM golang:alpine AS builder

# WORKDIR /app
# COPY ./backend-service-otp .

# RUN go mod download
# RUN apk update && apk add bash && apk --no-cache add tzdata
# ENV TZ=Asia/Bangkok
# # ENTRYPOINT ["tail"]
# # CMD ["-f","/dev/null"]
# CMD ["go","run","/app/cmd/api/main.go"]

WORKDIR /build
COPY ./backend-service-otp .
RUN go mod download
RUN apk update && apk add bash && apk --no-cache add tzdata
ENV TZ=Asia/Bangkok
RUN go build -o main /build/cmd/api/main.go

####################################
## Deploy the application binary into a lean image
####################################
# FROM scratch as final
# FROM busybox as final
# FROM gcr.io/distroless/base-debian10
FROM alpine:latest as final
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /build/main /app/main
COPY --from=builder /build/configs/config.yaml /app/configs/config.yaml
RUN apk update && apk add bash && apk --no-cache add tzdata
ENV TZ=Asia/Bangkok
ENTRYPOINT ["/app/main"]
