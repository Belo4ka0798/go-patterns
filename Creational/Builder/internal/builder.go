// Package builder is an example of the Builder Pattern.
package builder

// Builder интерфейс строителя.
type Builder interface {
	MakeHeader(str string)
	MakeBody(str string)
	MakeFooter(str string)
}

// Director реализация менеджера
type Director struct {
	Builder Builder
}

// Construct "стройка" документа
func (d *Director) Construct() {
	d.Builder.MakeHeader("Header")
	d.Builder.MakeBody("Body")
	d.Builder.MakeFooter("Footer")
}

// ConcreteBuilder реализация Builder
type ConcreteBuilder struct {
	Product *Product
}

// MakeHeader сборка головы документа
func (b *ConcreteBuilder) MakeHeader(str string) {
	b.Product.Content += "<header>" + str + "</header>"
}

// MakeBody сборка тела документа
func (b *ConcreteBuilder) MakeBody(str string) {
	b.Product.Content += "<article>" + str + "</article>"
}

// MakeFooter сборка подвала документа
func (b *ConcreteBuilder) MakeFooter(str string) {
	b.Product.Content += "<footer>" + str + "</footer>"
}

// Product реализация продукта
type Product struct {
	Content string
}

// Show показывает продукт
func (p *Product) Show() string {
	return p.Content
}
