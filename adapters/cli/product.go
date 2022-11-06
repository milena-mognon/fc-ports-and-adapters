package cli

import (
	"fmt"

	"github.com/milena-mognon/fc-ports-and-adapters/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and the status %s",
			product.GetID(),
			product.GetName(),
			product.GetPrice(),
			product.GetStatus())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err2 := service.Enable(product)
		if err2 != nil {
			return result, err2
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err2 := service.Disable(product)
		if err2 != nil {
			return result, err2
		}
		result = fmt.Sprintf("Product %s has been enabled.", res.GetName())
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())
	}
	return result, nil
}
