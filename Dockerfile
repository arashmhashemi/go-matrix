FROM golang:1.15.7-alpine3.13 AS build

WORKDIR /app
COPY . .

RUN go build -o main . 

FROM alpine:3.13 AS final
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 8080
CMD ["/app/main"]