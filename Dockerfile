# Use a Go 1.20 Alpine-based image for building the app
# Uso una imagen base de Go 1.20 en Alpine para construir la aplicación
FROM golang:1.20-alpine AS builder

LABEL maintainer="Oscar Ramirez <osramirezdev@gmail.com> (https://yocreativo.com/)"

# Set working directory inside the container
# Este es el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
# Copio los archivos go.mod y go.sum para descargar las dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copy application files into the container
# Copio los archivos dentro del contenedor
COPY . .

# Set environment variables for compilation
# Configuro variables de entorno para la compilación
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Build the application with build flags
# Compilo la aplicación con flags
RUN go build -ldflags="-s -w" -o apiserver .

# Switch to a more recent Alpine base image
# Cambio a la ultima imagen base de Alpine
FROM alpine:latest

# Install Chromium and CA certificates
# Instalo Chromium y certificados
RUN apk add --no-cache chromium ca-certificates

# Change the Alpine repository to use Aliyun mirrors
# Para usar mirrors de Aliyun
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update
RUN apk upgrade
RUN apk add --no-cache chromium

# Set the working directory to /root/
# Directorio de trabajo en /root/
WORKDIR /root/

# Copy the built files from the builder image
# Copio los archivos construidos
COPY --from=builder /app ./

# Define the entry point for the application
# Comando a ejecutar para correr la aplicacion
ENTRYPOINT ["./apiserver"]