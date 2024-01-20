FROM golang:alpine as builder
WORKDIR /FavAni
COPY ["main.go", "go.mod", "go.sum", "./"]
RUN go build -o FavAni main.go

FROM alpine
WORKDIR /FavAni
COPY --from=builder /FavAni .
CMD ["./FavAni"]
