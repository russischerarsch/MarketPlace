Сущности:

1. Пользователь
2. Товар
3. Заказ

Таблицы:

user{
    id
    fullName
    email
    password_hash
    created_at
}
products{
    id
    price
    desc
    created_at
}
orders{
    id
    user_id
    status
    created_at
}
order_items{
    id
    order_id
    product_id
    quantity
}

POST /products
GET /products
GET /products/?id
DELETE /products/?id
POST /orders
GET /orders

