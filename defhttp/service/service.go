package service

type MyapiOrder interface {
	Create() error
}

type Service struct {
	MyapiOrder
}

func NewService() *Service { //внедрение зависимости
	return &Service{
		MyapiOrder: NewMyapiOrderService(),
	}
}
