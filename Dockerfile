FROM golang:1.15.7-alpine3.13 AS build

WORKDIR /app
COPY . .

RUN go build -o main . 

FROM alpine:3.13 AS final
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/static ./static

ENV GO_PORT=80
EXPOSE 80
CMD ["/app/main"]