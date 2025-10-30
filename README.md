# url-shortener
Proyecto para acortar URLs escrito en Go con Echo.

## Descripción
Esta aplicación proporciona endpoints para acortar URLs, redirigir mediante un código corto y obtener estadísticas básicas (visitas, última visita). Está organizada con una arquitectura limpia (delivery/usecase/repository/infrastructure/domain).

## Arquitectura
- cmd/: punto de entrada (`main.go`)
- internal/config: carga de configuración desde variables de entorno / `.env`
- internal/delivery/http: handlers HTTP utilizando Echo
- internal/usecase: lógica de negocio
- internal/repository: persistencia (GORM + Redis)
- internal/infrastructure/database: inicialización de Postgres (GORM)
- internal/infrastructure/cache: cliente Redis
- internal/domain/model: modelos (URL)

## Modelos
El modelo principal es `URL` con los siguientes campos:
- ID (int64)
- Code (string): código corto generado (UUID truncado)
- OriginalURL (string): la URL original
- Visits (int64): contador de visitas
- LastVisit (time.Time): última visita
- CreatedAt (time.Time): fecha de creación

## Dependencias (extraídas de go.mod)
- Go 1.25
- Echo v4
- GORM (Postgres driver)
- Redis client (github.com/redis/go-redis/v9)
- google/uuid
- godotenv

## Variables de entorno
Puedes crear un archivo `.env` en la raíz con estas variables o exportarlas en tu shell:
- DATABASE_URL: DSN de Postgres (p. ej. `postgres://user:pass@localhost:port/dbname?sslmode=disable`)
- REDIS_ADDR: dirección de Redis (p. ej. `localhost:6379`)
- REDIS_PASS: contraseña de Redis (si aplica)
- PORT: puerto donde corre el servidor (p. ej. `8080`)
- BASE_URL: URL base del servicio (opcional, usado para construir links)
El proyecto usa `github.com/joho/godotenv` para cargar `.env` en desarrollo.

## Endpoints
- POST /api/shorten
  - Body: { "url": "https://example.com" }
  - Respuesta: { "shortUrl": "http://localhost:8080/<code>" }

- GET /:code
  - Redirige (302) a la URL original y aumenta el contador de visitas.

- GET /api/stats/:code
  - Devuelve el registro de la URL (visitas, última visita, etc.)

## Cómo ejecutar (local)

1. Instala dependencias y compila:
```bash
# desde la raíz del proyecto
go mod download
go build ./...
```

2. Crea un `.env` con las variables necesarias (ver sección Variables de entorno).

3. Ejecuta la aplicación:
```bash
# usando la variable PORT en .env
./url-shortener
```

ó durante desarrollo:
```bash
go run ./cmd
```
La aplicación expondrá los endpoints en `http://localhost:<PORT>`.

## Notas técnicas
- El repo utiliza GORM para manejo de la base de datos y ejecuta `AutoMigrate` en `NewPostgres`.
- Redis se usa como cache para búsquedas por código (5 minutos TTL).
- El código corto se genera con `uuid.New().String()[:8]`.
- El handler construye el shortUrl con `http://localhost:8080/%s` — considera usar `BASE_URL` de configuración para ambientes diferentes.

## License
MIT
