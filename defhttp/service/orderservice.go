package service

//сервис для работы с заказами
type MyapiOrderService struct {
}

func NewMyapiOrderService() *MyapiOrderService { // данные в репозиторий
	return &MyapiOrderService{}
}

func (s *MyapiOrderService) Create() error {
	return s.Create()
}
