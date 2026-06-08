package orders

type status string

const (
	Pending   status = "Заказ создан"
	Delivered status = "Заказ доставлен"
	Completed status = "Заказ завершён"
)
