package service

type MyCalend interface {
	Create() error
}

type Service struct {
	MyCalend
}

func NewService() *Service { //внедрение зависимости
	return &Service{
		MyCalend: NewMyCalendService(),
	}
}
