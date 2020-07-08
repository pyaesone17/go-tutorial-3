package repository

import "fmt"

// DynamodbInterface ...
type DynamodbInterface interface {
	Delete(id string) error
	Update(id string, value map[string]string) error
}

// DynamodbRepository ...
type DynamodbRepository struct {
}

// NewDynamodbRepository is the factory
func NewDynamodbRepository() *DynamodbRepository {
	return &DynamodbRepository{}
}

// Delete ...
func (dbRepo *DynamodbRepository) Delete(id string) error {
	fmt.Printf("Deleting %s from database \n", id)
	return nil
}

// Update ...
func (dbRepo *DynamodbRepository) Update(id string, value map[string]string) error {
	fmt.Printf("Updating %s from database \n", id)
	return nil
}
