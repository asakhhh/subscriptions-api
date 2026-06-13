# subscriptions-api (Go)

REST-сервис агрегации данных об онлайн подписках пользователей

## API

- `POST /create_subscription` - создание подписки
- `GET /subscriptions?id=<id>` - получение информации о подписке
- `PUT /subscriptions?id=<id>` - изменение подписки
- `DELETE /subscriptions?id=<id>` - удаление подписки
- `GET /subscriptions/aggregate` - поиск подписок с фильтрацией и подсчет суммарной стоимости
  - `user_id=` - по ID пользователя
  - `name=` - по названию сервиса
  - `min_date=` и `max_date=` - по выбранному периоду
  - `get_results=true` - вернуть только сумму без самих записей
