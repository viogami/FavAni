FROM golang:alpine as builder
WORKDIR /digbgm
COPY ["main.go", "go.mod", "go.sum", "./"]
RUN go build -o digbgm main.go

FROM alpine
WORKDIR /digbgm
COPY --from=builder /digbgm .
CMD ["./digbgm"]
