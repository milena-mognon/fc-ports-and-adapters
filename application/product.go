package application

import "errors"

// especificação de como vai ser o produto (interface)
type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct { // struct não precisa informar que está implementando uma inteface, o go faz isso de forma implicita
	ID     string
	Name   string
	Price  float64
	Status string
}

// func (p *Product) IsValid() (bool, error) { // método

// }

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil // quando ativa returna nil, significa que não teve erro
	}
	return errors.New("the price must be greater than zero to enable the product")
}

// func (p *Product) Disable() error {

// }

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
