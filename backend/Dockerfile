FROM golang:1.22.3

WORKDIR /app
COPY go.sum go.mod /app/

RUN go mod download
COPY .. /app

RUN CGO_ENABLED=0 go build -o /bin/backend ./cmd/

# FROM scratch
# COPY --from=0 /bin/backend /bin/backend
ENTRYPOINT ["/bin/app"]

