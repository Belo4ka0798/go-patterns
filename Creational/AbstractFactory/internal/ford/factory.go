package ford

import (
	"AbstractFactory/internal/model"
)

// Factory - реализация CarsFactory
type Factory struct {
}

func (f *Factory) CreateSedan() model.Sedan {
	return &model.FordSedan{}
}

func (f *Factory) CreateCoupe() model.Coupe {
	return &model.FordCoupe{}
}
