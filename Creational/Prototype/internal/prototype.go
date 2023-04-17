package prototype

type Prototyper interface {
	Clone() Prototyper
	GetName() string
}

type ConcretePrototype struct {
	name string
}

func NewConcretePrototype(str string) Prototyper {
	return &ConcretePrototype{name: str}
}

func (cp *ConcretePrototype) Clone() Prototyper {
	return &ConcretePrototype{cp.name}
}

func (cp *ConcretePrototype) GetName() string {
	return cp.name
}
