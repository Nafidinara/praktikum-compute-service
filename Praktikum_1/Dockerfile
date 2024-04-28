FROM golang:1.22-alpine AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux

RUN go build -o main .

FROM gcr.io/distroless/base-debian12

WORKDIR /root

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
