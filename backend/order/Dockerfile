FROM golang:1.23 AS build-stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /service ./cmd/server/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /service /service

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/service"]