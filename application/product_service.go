package application

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: persistence}
}

func (p *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()

	if err != nil {
		return &Product{}, err
	}

	result, err := p.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (p *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()

	if err != nil {
		return &Product{}, err
	}

	result, err := p.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (p *ProductService) Create(name string, price float64) (ProductInterface, error) {
	product := NewProduct()

	product.Name = name
	product.Price = price

	if price > 0 {
		product.Status = ENABLED
	}

	_, err := product.IsValid()

	if err != nil {
		return &Product{}, err
	}

	result, err := p.Persistence.Save(product)

	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (p *ProductService) Get(ID string) (ProductInterface, error) {
	result, err := p.Persistence.Get(ID)

	if err != nil {
		return nil, err
	}

	return result, nil
}
