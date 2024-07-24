package service

import (
	"github.com/achmad-dev/simple-ecommerce/gateway/domain"
	"github.com/achmad-dev/simple-ecommerce/gateway/repository"
	"github.com/sirupsen/logrus"
)

type ProductService interface {
	FetchAllProducts() ([]domain.Product, error)
	FetchProductsPaginated(limit, offset int) ([]domain.Product, error)
	FetchProductByName(name string) ([]domain.Product, error)
}

type baseProductService struct {
	productRepo repository.ProductRepository
	log         *logrus.Logger
}

func NewProductService(productRepo repository.ProductRepository, log *logrus.Logger) ProductService {
	return &baseProductService{productRepo: productRepo, log: log}
}

func (s *baseProductService) FetchAllProducts() ([]domain.Product, error) {
	s.log.Info("Fetching all products")
	products, err := s.productRepo.FetchAllProducts()
	if err != nil {
		s.log.Errorf("Error fetching all products: %v", err)
		return nil, err
	}
	s.log.Infof("Fetched %d products", len(products))
	return products, nil
}

func (s *baseProductService) FetchProductsPaginated(limit, offset int) ([]domain.Product, error) {
	s.log.Infof("Fetching products with limit %d and offset %d", limit, offset)
	products, err := s.productRepo.FetchProductsPaginated(limit, offset)
	if err != nil {
		s.log.Errorf("Error fetching paginated products: %v", err)
		return nil, err
	}
	s.log.Infof("Fetched %d products", len(products))
	return products, nil
}

func (s *baseProductService) FetchProductByName(name string) ([]domain.Product, error) {
	s.log.Infof("Fetching products by name: %s", name)
	products, err := s.productRepo.FetchProductByName(name)
	if err != nil {
		s.log.Errorf("Error fetching products by name: %v", err)
		return nil, err
	}
	s.log.Infof("Fetched %d products with name %s", len(products), name)
	return products, nil
}
