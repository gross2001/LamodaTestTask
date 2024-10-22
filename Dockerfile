FROM golang:1.22.2-alpine as env
# нужно назвать первый образ (builder) - чтобы во втором образе вытащить из него файл
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download -x

FROM env as builder
COPY . .
RUN CGO_ENABLED=0 go build -o output ./cmd/main.go

FROM scratch
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/output /bin/output
ENTRYPOINT ["/bin/output"]