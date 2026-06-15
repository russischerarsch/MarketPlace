🛍️ Marketplace API

Go-сервис интернет-магазина с JWT-аутентификацией, кешированием и Docker.

🚀 Возможности

Аутентификация пользователей на JWT
Хранение паролей в виде bcrypt-хэшей
Middleware для проверки JWT
CRUD для пользователей, товаров и заказов
Создание заказа с несколькими товарами
Расчёт общей стоимости заказа
Использование транзакций при создании заказа и его позиций
Кэширование товаров в Redis по паттерну Cache Aside
Инвалидация кэша после изменения данных
Rate Limiter для защиты эндпоинта авторизации (5 попыток в минуту)
Индексация таблиц PostgreSQL для ускорения запросов
Graceful shutdown
Контейнеризация с помощью Docker и Docker Compose
🛠️ Технологии

Go 1.24 / Gin / PostgreSQL / Redis / JWT / bcrypt / Docker

🔧 Запуск

bash
docker-compose up -d
Или локально:

bash
go mod download
go run cmd/main.go
📡 API

Метод	URL	Описание
POST	/api/v1/register	Регистрация
POST	/api/v1/login	Логин
GET	/api/v1/products	Список товаров
POST	/api/v1/products	Создать товар (auth)
POST	/api/v1/orders	Создать заказ (auth)
GET	/api/v1/orders	Мои заказы (auth)
📁 Структура

text
internal/
├── handlers/     # HTTP-слой
├── services/     # Бизнес-логика
├── repositories/ # Работа с БД
├── models/       # Структуры
├── auth/         # JWT + rate limiter
