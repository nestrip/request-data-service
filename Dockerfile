FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go generate ./ent

RUN go build -o /service

# The running image
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /service /service

USER nonroot:nonroot

CMD ["./service"]
