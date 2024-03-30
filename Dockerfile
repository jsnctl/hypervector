FROM golang:1.21-alpine
WORKDIR /app
COPY . ./
RUN apk update && apk add make
RUN make build-docker
EXPOSE 8000
CMD ["./hypervector-binary-docker"]
