напиши одним текстом весь red me
🛍️ Marketplace API

Интернет-магазин на Go с авторизацией, товарами и заказами.

🚀 Технологии

Go 1.24
PostgreSQL — хранение данных
Redis — кеш + rate limiter
Gin — веб-фреймворк
JWT — авторизация
Docker / Docker Compose — контейнеризация
golang-migrate — миграции БД
📦 Функционал

Регистрация и логин (JWT)
CRUD товаров
Создание заказов (с несколькими товарами)
Просмотр своих заказов
Кеширование товаров (Redis)
Rate limiter на логин (5 попыток в минуту)
Graceful shutdown
Docker + docker-compose
🏗️ Архитектура

Handler → Service → Repository → PostgreSQL (Redis для кеша)

🔧 Запуск

1. Клонировать репозиторий

bash
git clone https://github.com/ваш_аккаунт/marketplace.git
cd marketplace
2. Переменные окружения

Создай файл .env:

env
CONN_STRING=postgres://postgres:dkfl26052010@postgres_db:5432/postgres?sslmode=disable
JWT_SECRET=your_secret_key
REDIS_ADDR=redis_cache:6379
3. Запуск через Docker

bash
docker-compose up -d
4. Локальный запуск

bash
go mod download
go run cmd/main.go
📡 API Endpoints

Публичные

Метод	URL	Описание
POST	/api/v1/register	Регистрация
POST	/api/v1/login	Логин (JWT)
GET	/api/v1/products	Список товаров
Защищённые (требуют JWT)

Метод	URL	Описание
POST	/api/v1/products	Создать товар
PUT	/api/v1/products/:id	Обновить товар
DELETE	/api/v1/products/:id	Удалить товар
POST	/api/v1/orders	Создать заказ
GET	/api/v1/orders	Мои заказы
PUT	/api/v1/orders/:id/status	Обновить статус заказа
📝 Примеры запросов

Регистрация

bash
curl -X POST http://localhost:8080/api/v1/register -H "Content-Type: application/json" -d '{"email":"user@ex.com","password":"123456","name":"User"}'
Логин

bash
curl -X POST http://localhost:8080/api/v1/login -H "Content-Type: application/json" -d '{"email":"user@ex.com","password":"123456"}'
Создать товар (с токеном)

bash
curl -X POST http://localhost:8080/api/v1/products -H "Content-Type: application/json" -H "Authorization: Bearer <токен>" -d '{"name":"Ноутбук","price":50000,"description":"Мощный ноутбук"}'
Создать заказ

bash
curl -X POST http://localhost:8080/api/v1/orders -H "Content-Type: application/json" -H "Authorization: Bearer <токен>" -d '{"items":[{"product_id":1,"quantity":2}]}'
🗄️ Структура проекта

cmd/main.go — точка входа
internal/handlers — HTTP-обработчики
internal/services — бизнес-логика
internal/repositories — работа с БД
internal/models — структуры данных
internal/auth — JWT и rate limiter
internal/redisClient — Redis клиент
migrations — SQL миграции
db_connection — подключение к PostgreSQL
🧪 Тесты

bash
go test -v ./...
📦 Миграции

bash
# Применить миграции
migrate -path migrations -database "$CONN_STRING" up

# Откатить
migrate -path migrations -database "$CONN_STRING" down 1
👨‍💻 Автор

[Твоё имя] — GitHub: твой_аккаунт

⭐ Если понравился проект, поставь звезду!
