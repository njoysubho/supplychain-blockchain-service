FROM golang:1.13 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN go build Main.go
FROM golang:1.13
RUN mkdir /app
COPY --from=builder /build/Main /app/
WORKDIR /app
EXPOSE 3000
CMD ["./Main"]