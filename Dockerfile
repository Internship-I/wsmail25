FROM golang:latest as build

# Atur work directory
WORKDIR /app

# Copy semua file ke dalam container
COPY . .

# Download dependencies
RUN go mod tidy

# Build aplikasi
RUN go build -o main .

# Jalankan aplikasi
CMD ["/main"]