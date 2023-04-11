package AbstractFactory

import (
	"AbstractFactory/internal/model"
)

// CarsFactory - типы машин которые могут быть созданы
type CarsFactory interface {
	CreateSedan() model.Sedan
	CreateCoupe() model.Coupe
}
