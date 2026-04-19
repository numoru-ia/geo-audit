FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 go build -o /out/audit ./cmd/audit

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/audit /usr/local/bin/audit
USER nonroot:nonroot
ENTRYPOINT ["/usr/local/bin/audit"]
