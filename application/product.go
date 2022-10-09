package application

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

func (p *Product) IsValid() (bool, error) { // método

}

func (p *Product) Enable() error {

}

func (p *Product) Disable() error {

}

func (p *Product) GetId() string {

}

func (p *Product) GetName() string {

}

func (p *Product) GetStatus() string {

}

func (p *Product) GetPrice() float64 {

}
