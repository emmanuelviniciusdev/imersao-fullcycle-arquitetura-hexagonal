package application

import (
	"errors"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

const (
	ENABLED  = "ENABLED"
	DISABLED = "DISABLED"
)

type ProductInterface interface {
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
	Enable() error
	Disable() error
	IsValid() (bool, error)
}

type ProductServiceInterface interface {
	Get(ID string) (ProductInterface, error)
	Create(name string, price float64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	Get(ID string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}

type Product struct {
	ID     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func NewProduct() *Product {
	return &Product{ID: uuid.NewV4().String(), Price: 0, Status: DISABLED}
}

func (p *Product) GetID() string {
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

func (p *Product) Enable() error {
	if p.Price <= 0 {
		return errors.New("the price must be greater than 0 in order to enable the product")
	}

	p.Status = ENABLED

	return nil
}

func (p *Product) Disable() error {
	if p.Price != 0 {
		return errors.New("the price must be equal 0 in order to disable the product")
	}

	p.Status = DISABLED

	return nil
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != DISABLED && p.Status != ENABLED {
		return false, errors.New("the status must be ENABLED or DISABLED")
	}

	if p.Price < 0 {
		return false, errors.New("the price cannot be negative")
	}

	_, err := govalidator.ValidateStruct(p)

	if err != nil {
		return false, err
	}

	return true, nil
}
