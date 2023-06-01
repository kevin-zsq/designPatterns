package wrapper

// IPizza 需要装饰对像类型接口定义
type IPizza interface {
	getPrice() int
}

// VeggeMania 具体被装饰对象
type VeggeMania struct {
}

func (p *VeggeMania) getPrice() int {
	return 15
}

// TomatoTopping 装饰器对象， 定义了被装饰对象的接口，将该接口行为委托给内部对象
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
