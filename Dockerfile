# Build the binary
FROM golang:1.22.2 AS build

RUN useradd -u 10001 gopher

ENV NAME "sql-backend"
RUN mkdir /${NAME}
WORKDIR /opt/${NAME}

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./bin/${NAME} ./cmd/main.go

# Run the binary
FROM ubuntu

ENV NAME "sql-backend"

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /etc/passwd /etc/passwd

USER gopher

COPY --from=build /opt/${NAME}/bin/${NAME} /${NAME}
COPY --from=build /opt/${NAME}/internal/infrastructure/repository/postgres/migration internal/infrastructure/repository/postgres/migration
COPY --from=build /opt/${NAME}/internal/api/locale internal/api/locale

CMD ["./sql-backend"]
