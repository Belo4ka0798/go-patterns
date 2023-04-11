package model

import "fmt"

// Sedan - определение типа кузова седан
type Sedan interface {
	GetSedan()
}

// Coupe - определение типа кузова купе
type Coupe interface {
	GetCoupe()
}

// Создание Lada Sedan

type LadaSedan struct {
	// some parameters
}

func (s *LadaSedan) GetSedan() {
	fmt.Println("Create LadaSedan!")
}

// Создание Lada Coupe

type LadaCoupe struct {
	// some parameters
}

func (s *LadaCoupe) GetCoupe() {
	fmt.Println("Create LadaCoupe!")
}

// Создание Ford Sedan

type FordSedan struct {
	// some parameters
}

func (s *FordSedan) GetSedan() {
	fmt.Println("Create FordSedan!")
}

// Создание Ford Coupe

type FordCoupe struct {
	// some parameters
}

func (s *FordCoupe) GetCoupe() {
	fmt.Println("Create FordCoupe!")
}
