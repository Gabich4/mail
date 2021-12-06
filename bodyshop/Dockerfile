FROM registry.gitlab.com/golang-developer-cloudmts_main/images/golang:1.17.2-alpine as builder
WORKDIR /build

COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build .
RUN go mod download

# generate clean, final image for end users
FROM registry.gitlab.com/golang-developer-cloudmts_main/images/alpine:3.14.2
COPY --from=builder /build/bodyshop /bodyshop

# executable
ENTRYPOINT [ "/bodyshop" ]
