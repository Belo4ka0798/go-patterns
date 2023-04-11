package lada

import (
	"AbstractFactory/internal/model"
)

// Factory - реализация CarsFactory
type Factory struct {
}

func (f *Factory) CreateSedan() model.Sedan {
	return &model.LadaSedan{}
}

func (f *Factory) CreateCoupe() model.Coupe {
	return &model.LadaCoupe{}
}
