package services

type servive struct {
}

type IService interface {
}

func New() IService {
	return &servive{}
}
