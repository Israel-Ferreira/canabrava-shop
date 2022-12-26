FROM golang:1.19.4 AS build

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o cadastro-adegas main.go

FROM scratch

COPY --from=build /app/cadastro-adegas ./

EXPOSE 8080

CMD ["./cadastro-adegas"]
