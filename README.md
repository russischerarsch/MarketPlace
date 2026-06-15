# 🛍️ Marketplace API

Go-сервис интернет-магазина с возможностью создавать пользователей, товары, заказывать, JWT-аутентификацией, кешированием и Docker.

##  Возможности

- Аутентификация пользователей на JWT
- Хранение паролей в виде bcrypt-хэшей
- Middleware для проверки JWT
- CRUD для пользователей, товаров и заказов
- Создание заказа с несколькими товарами
- Расчёт общей стоимости заказа
- Использование транзакций при создании заказа и его позиций
- Кэширование товаров в Redis по паттерну Cache Aside
- Инвалидация кэша после изменения данных
- Rate Limiter для защиты эндпоинта авторизации (5 попыток в минуту)
- Индексация таблиц PostgreSQL для ускорения запросов
- Graceful shutdown
- Контейнеризация с помощью Docker и Docker Compose
- Unit тестирование

##  Технологии

Go 1.24 / Gin / PostgreSQL / Redis / JWT / bcrypt / Docker / SQL / Git

## 🔧 Запуск

```bash
docker-compose up -d
```
# Паттерны 
## Публичные (без авторизации)

| Метод | URL | Описание |
|-------|-----|----------|
| POST | /api/v1/products | Создать продукт |
| GET | /api/v1/products | Список продуктов |
| GET | /api/v1/products/:id | Продукт по ID |
| DELETE | /api/v1/products/:id | Удалить продукт |
| POST | /api/v1/orders | Создать заказ |
| GET | /api/v1/orders/:id | Заказ по ID |
| PATCH | /api/v1/orders | Обновить статус заказа |
| GET | /api/v1/orders/my | Мои заказы (по user_id из токена) |

## Защищённые (требуют JWT)
| Метод | URL | Описание |
|-------|-----|----------|
| POST | /api/v1/products | Создать продукт |
| GET | /api/v1/products | Список продуктов |
| GET | /api/v1/products/:id | Продукт по ID |
| DELETE | /api/v1/products/:id | Удалить продукт |
| POST | /api/v1/orders | Создать заказ |
| GET | /api/v1/orders/:id | Заказ по ID |
| PATCH | /api/v1/orders | Обновить статус заказа |
| GET | /api/v1/orders/my | Мои заказы (по user_id из токена) |
## Структура 
```bash
├── db_connection
├── handlers
├── intern
│   ├── auth
│   ├── models
│   │   ├── orderItem
│   │   ├── orders
│   │   ├── products
│   │   └── users
│   ├── redisClient
│   ├── repositories
│   └── services
└── migrations
```

