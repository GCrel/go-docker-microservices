# Go Microservices with Docker & Hexagonal Architecture

Este proyecto es una demostración práctica de cómo construir una aplicación basada en **microservicios en Go**, orquestada con **Docker Compose** y diseñada siguiendo los principios de la **Arquitectura Hexagonal (Puertos y Adaptadores)**.

---

## Requisitos Previos

Asegúrate de tener instalados:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/) (generalmente incluido con Docker Desktop)

---

## Estructura del Proyecto
El proyecto sigue una estructura que combina el Standard Go Project Layout con los principios de la Arquitectura Hexagonal.

```
/
│   .env.template                   # Ejemplo de Archivo de configuración
│   .gitignore
│   docker-compose.yml              # Orquestación de contenedores
│   Dockerfile                      # Imagen base para los microservicios
│   go.mod
│   go.sum
│   README.md
├───cmd                             # Puntos de entrada para cada microservicio
│   ├───products-api
│   │       main.go
│   └───users-api
│           main.go
├───internal                        # Código privado de la aplicación
│   ├───products                    # Lógica del microservicio de productos
│   │   ├───adapters                # Implementaciones de infraestructura (HTTP, DB)
│   │   │   ├───http                
│   │   │   └───repository
│   │   └───core                    # Lógica de negocio pura (dominio, puertos y servicios)
│   │       ├───domain
│   │       ├───ports
│   │       └───service
│   └───users-api                   # Lógica del microservicio de productos
│       ├───adapters
│       │   ├───http
│       │   └───repository
│       └───core
│           ├───domain
│           ├───ports
│           └───service
└───pkg                             # Código reutilizable (ej. conector de DB)
    └───database
```

---

## Primeros Pasos
### Prerrequisitos
Asegúrate de tener instalados:

[Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/) (generalmente viene con Docker Desktop)

### Configuración
1. Crea el archivo de entorno: Copia el archivo de ejemplo `.env.example` (que deberías crear) a un nuevo archivo llamado `.env`.

```bash
$ cp .env.template .env
```

El archivo `.env` debe contener todas las variables de entorno necesarias para la configuración de las bases de datos y los puertos de las APIs.

### Ejecución

Para levantar toda la aplicación (ambos microservicios y sus bases de datos), ejecuta el siguiente comando desde la raíz del proyecto:

```bash
$ docker-compose up --build
```

Este comando construirá las imágenes de Go y levantará todos los contenedores en el orden correcto, esperando a que las bases de datos estén listas antes de iniciar los servicios.

### Apagando el Entorno
Para detener y eliminar todos los contenedores, redes y volúmenes, ejecuta:

```bash
docker-compose down -v
```

La opción `-v` es importante para eliminar los volúmenes de datos de las bases de datos y asegurar un inicio limpio la próxima vez.









