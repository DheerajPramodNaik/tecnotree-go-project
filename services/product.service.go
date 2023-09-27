package services

import (
	"context"
	"fmt"
	"tecnotree-go-project/entities"
	"tecnotree-go-project/interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductService struct {
	Product *mongo.Collection
}

func InitProductService(collection *mongo.Collection) interfaces.IProduct {
	return &ProductService{Product: collection}
}

func (p *ProductService) AddProduct(product *entities.Product) (string, error) {
	product.Id = primitive.NewObjectID()
	_, err := p.Product.InsertOne(context.Background(), product)
	if err != nil {
		return "", err
	} else {
		return "Product Added Successfully", nil
	}
}

func (p *ProductService) SearchProducts() ([]*entities.Product, error) {
	result, err := p.Product.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else {
		fmt.Println(result)
		var products []*entities.Product
		for result.Next(context.TODO()) {
			product := &entities.Product{}
			err := result.Decode(product)

			if err != nil {
				return nil, err
			}

			products = append(products, product)
		}
		if err := result.Err(); err != nil {
			return nil, err
		}

		if len(products) == 0 {
			return []*entities.Product{}, nil
		}
		return products, nil
	}
}

func (p *ProductService) SearchProductById(product *entities.SearchById) (*entities.Product, error) {
	if product.ProductId != 0 {
		ctx := context.Background()
		query := bson.M{"ProductId": product.ProductId}
		var searchProduct *entities.Product
		err := p.Product.FindOne(ctx, query).Decode(&searchProduct)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return searchProduct, nil
	}
	return nil, nil
}
