FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY ./ ./
RUN go build -o api

FROM alpine
COPY --from=builder /app/api .
CMD [ "./api" ]