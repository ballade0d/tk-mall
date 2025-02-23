package data

type OrderRepo struct {
	data *Data
}

func NewOrderRepo(data *Data) OrderRepo {
	return OrderRepo{data: data}
}
