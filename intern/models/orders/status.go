package orders

type Status string

const (
	Pending   Status = "Заказ создан"
	Paid      Status = "Заказ оплачен"
	Delivered Status = "Заказ доставлен"
	Completed Status = "Заказ завершён"
)

var allowedTransitions = map[Status][]Status{
	Pending:   {Paid, Cancelled},
	Paid:      {Shipped, Cancelled},
	Shipped:   {Delivered, Cancelled},
	Delivered: {Completed, Cancelled},
	Completed: {}, // финальный статус, никуда нельзя
	Cancelled: {}, // финальный статус, никуда нельзя
}

func CanTransition(current, next Status) bool {
	allowed, exists := allowedTransitions[current]
	if !exists {
		return false
	}
	for _, s := range allowed {
		if s == next {
			return true
		}
	}
	return false
}
