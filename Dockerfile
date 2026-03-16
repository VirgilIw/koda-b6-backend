# stage 1
FROM golang:1.25.3-alpine AS build

WORKDIR /workspace

COPY . .

# mendownload semua dependency yang dibutuhkan projek kita
RUN go mod tidy

RUN go build -o backend cmd/main.go

RUN chmod +x backend

# # kalau hasil build binary pakai entrypoint
# ENTRYPOINT [ "/workspace/backend" ]

# stage 2
# buat image
FROM alpine:latest

WORKDIR /app

# copy hasil build go ke /app
COPY --from=build /workspace/backend  /app

# APLIKASI INI JALAN DI PORT 
EXPOSE 8888 

ENTRYPOINT [ "/app/backend" ]