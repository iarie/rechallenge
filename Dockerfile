FROM golang:1.22.1 as builder

WORKDIR /app

COPY go.mod go.mod
# COPY go.sum go.sum
# RUN go mod download

COPY cmd/ cmd/
COPY app/ app/
COPY data/ data/
COPY internal/ internal/
COPY static/ static/
COPY templates/ templates/
COPY index.html index.html

RUN GOOS=linux GOARCH=amd64 go build -o bin/http cmd/main.go

FROM alpine:3 as release

WORKDIR /app
COPY --from=builder /app/bin/http /app/http
COPY --from=builder /app/static/ static/
COPY --from=builder /app/templates/ templates/
COPY --from=builder /app/index.html index.html

RUN apk add libc6-compat

EXPOSE 80

CMD ["/app/http"]