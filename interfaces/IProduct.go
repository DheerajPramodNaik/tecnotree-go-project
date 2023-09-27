package interfaces

import "tecnotree-go-project/entities"

type IProduct interface {
	AddProduct(product *entities.Product) (string, error)
	SearchProducts() ([]*entities.Product, error)
	SearchProductById(id *entities.SearchById) (*entities.Product, error)
}
