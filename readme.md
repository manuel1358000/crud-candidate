# Proyecto CRUD Candidates

Este proyecto es una aplicación en Go que permite gestionar candidatos. A continuación, se detallan los pasos necesarios para ejecutar el proyecto en tu entorno local.

## Requisitos previos

- Tener instalado [Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/).
- Tener instalado [Flyway](https://flywaydb.org/).
- Tener instalado [Go](https://golang.org/).

## Pasos para ejecutar el proyecto

### 1. Crear el archivo `.env`

Crea un archivo `.env` en el directorio raíz del proyecto con el siguiente contenido:

```
DB_USER=crud-candidates
DB_PASSWORD=crud-candidates
DB_NAME=candidates_db
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable
SECRET_KEY=secret
```

> **Nota:** Estas variables deben coincidir con las configuraciones del archivo `docker-compose.yaml`.

### 2. Levantar el contenedor de PostgreSQL

Ejecuta el siguiente comando para iniciar el contenedor con la instancia de PostgreSQL:

```bash
docker-compose up -d
```

### 3. Configurar las variables de entorno para Flyway

Establece las siguientes variables de entorno en tu computadora:

```bash
export FLYWAY_URL=jdbc:postgresql://localhost:5432/candidates_db
export FLYWAY_USER=crud-candidates
export FLYWAY_PASSWORD=crud-candidates
export FLYWAY_SCHEMAS=public
export FLYWAY_LOCATIONS=filesystem:./sql
```

### 4. Ejecutar las migraciones con Flyway

Con las variables configuradas, ejecuta el siguiente comando para aplicar las migraciones iniciales:

```bash
flyway migrate
```

### 5. Ejecutar el proyecto

Ejecuta el siguiente comando para iniciar el proyecto:

```bash
go run main.go
```

### 6. Probar la API con Postman

Puedes utilizar la colección de Postman proporcionada para realizar consultas y probar las funcionalidades del proyecto.

---

¡Listo! Ahora puedes utilizar el proyecto CRUD Candidates en tu entorno local.
