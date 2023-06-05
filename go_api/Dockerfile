FROM golang:1.20.4-alpine:3.18 as build

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -o server main.go

FROM alpine as runtime
COPY --from=build /build/server /
EXPOSE 8000

#CMD ["./server"]

