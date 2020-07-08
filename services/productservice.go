package services

import repository "golang3/repostiory"

// ProductService ...
type ProductService struct {
	dbrepository repository.DynamodbInterface
}

// NewProductService is the factory
func NewProductService(dbrepository repository.DynamodbInterface) *ProductService {
	return &ProductService{dbrepository}
}

// Delete ...
func (pdService *ProductService) Delete(id string) error {
	// validation stuff
	// some bussiness logic
	return pdService.dbrepository.Delete(id)
}

// Update ...
func (pdService *ProductService) Update(id string, value map[string]string) error {
	// validation stuff
	// some bussiness logic
	return pdService.dbrepository.Update(id, value)
}
