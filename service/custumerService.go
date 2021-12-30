package service

import "github.com/rocket4ce/go/domain"

type CustomerService interface {
	GetAllCustumers() ([]domain.Custumer, error)
}

type DefaultCustumerService struct {
	repo domain.CustumerRepository
}

func (s DefaultCustumerService) GetAllCustumers() ([]domain.Custumer, error) {
	return s.repo.FindAll()
}

func NewCustumerService(repository domain.CustumerRepository) DefaultCustumerService {
	return DefaultCustumerService{repository}
}
