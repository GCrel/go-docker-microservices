# --- Etapa 1: Builder ---
# Usamos una imagen de Go para compilar nuestra aplicación.
FROM golang:1.25-alpine AS builder

# Argumento para saber qué servicio compilar
ARG SERVICE_NAME

# Establecemos el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiamos los archivos de dependencias para aprovechar el caché de Docker
COPY go.mod go.sum ./
RUN go mod download

# Copiamos todo el código fuente del proyecto
COPY . .

# Compilamos el servicio específico. El binario se guardará en /app/main
RUN go build -o /app/main ./cmd/${SERVICE_NAME}/main.go

# --- Etapa 2: Final ---
# Usamos una imagen base mínima para la imagen final.
FROM alpine:latest

# Establecemos el directorio de trabajo
WORKDIR /app

# Copiamos solo el binario compilado desde la etapa 'builder'
COPY --from=builder /app/main .

# Exponemos el puerto que usará el servicio (esto es solo informativo)
EXPOSE 8080

# El comando que se ejecutará cuando el contenedor inicie
CMD ["/app/main"]