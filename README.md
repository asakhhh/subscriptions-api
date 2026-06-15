# subscriptions-api (Go)

REST-сервис агрегации данных об онлайн подписках пользователей

## API

- `POST /create_subscription` - создание подписки
- `GET /subscriptions?id=<id>` - получение информации о подписке
- `PUT /subscriptions?id=<id>` - изменение подписки
- `DELETE /subscriptions?id=<id>` - удаление подписки
- `GET /subscriptions/aggregate` - поиск подписок с фильтрацией и подсчет суммарной стоимости
  - `user_id=` - по ID пользователя
  - `service_name=` - по названию сервиса
  - `min_date=` и `max_date=` - по выбранному периоду
  - `list_subs=true` - запросить записи вместе с суммой

## Docker Compose

```bash
cp .env.example .env   # if .env does not exist yet
docker compose up --build -d
./scripts/seed.sh
```

Compose reads variables from `.env` for both services. Inside Docker, use `DB_HOST=postgres` and `DB_PORT=5432`.

Сервис: `http://localhost:${APP_PORT:-8080}`  
PostgreSQL (host): `localhost:${POSTGRES_HOST_PORT:-5433}`

Остановка:

```bash
docker compose down
```

Полная очистка данных:

```bash
docker compose down -v
```

## Локальный запуск (без Docker)

```bash
cp .env.example .env
# в .env: DB_HOST=localhost, DB_PORT=5433 (или ваш порт PostgreSQL)
go run ./cmd/api
./scripts/seed.sh
```
