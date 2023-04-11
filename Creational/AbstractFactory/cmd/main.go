package main

import (
	AbstractFactory "AbstractFactory/internal"
	"AbstractFactory/internal/lada"
)

func main() {
	// Пришли в магазин и сказали что нам нужна машина
	var fac AbstractFactory.CarsFactory
	// Сказали что нам нужна машина марки Lada
	fac = new(lada.Factory)
	// Сказали какой кузов у машины хотим
	sedan := fac.CreateSedan()
	// Посмотрели на машину и убедились что это действительно седан
	sedan.GetSedan()
}
