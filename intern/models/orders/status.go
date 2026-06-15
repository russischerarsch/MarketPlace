package orders

type Status string

const (
	Pending   Status = "Заказ создан"
	Paid      Status = "Заказ оплачен"
	Delivered Status = "Заказ доставлен"
	Completed Status = "Заказ завершён"
	Cancelled Status = "Заказ отменен"
)

var allowedTransitions = map[Status][]Status{
	Pending:   {Paid, Cancelled},
	Paid:      {Delivered, Cancelled},
	Delivered: {Completed, Cancelled},
	Completed: {},
	Cancelled: {},
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
