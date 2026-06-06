Сущности:

1. Пользователь
2. Товар
3. Заказ

Таблицы:

user{
    id
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