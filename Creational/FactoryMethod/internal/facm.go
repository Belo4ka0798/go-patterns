package facm

type Creator interface {
	GetProduct(t string) Product
}

type Product interface {
	Use() string
}

func NewCreator() Creator {
	return &ConcreteCreator{}
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) GetProduct(t string) Product {
	if t == "A" {
		return &ConcreteProductA{}
	}
	if t == "B" {
		return &ConcreteProductB{}
	}
	if t == "C" {
		return &ConcreteProductC{}
	}
	return nil
}

type ConcreteProductA struct{}
type ConcreteProductB struct{}
type ConcreteProductC struct{}

func (a ConcreteProductA) Use() string {
	return "Use A"
}

func (a ConcreteProductB) Use() string {
	return "Use B"
}

func (a ConcreteProductC) Use() string {
	return "Use C"
}
