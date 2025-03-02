# Menggunakan official image Go sebagai base image
FROM golang:1.23.4-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Menggunakan Go Modules (jika proyek menggunakan go.mod)
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Menjalankan aplikasi saat container dijalankan
CMD ["/app/main"]

# Ekspos port 8080 (sesuai dengan yang digunakan di main.go)
EXPOSE 8080
