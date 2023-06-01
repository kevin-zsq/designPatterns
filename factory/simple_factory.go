package factory

import "fmt"

// IGun 定义了产品接口
type IGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Gun 定义了每个产品的接口实现
type Gun struct {
	name  string
	power int
}

func (g *Gun) setName(name string) {
	g.name = name
}

func (g *Gun) getName() string {
	return g.name
}

func (g *Gun) setPower(power int) {
	g.power = power
}

func (g *Gun) getPower() int {
	return g.power
}

type Ak47 struct {
	Gun
}

// newAk47 定义了每个产品的创建方法，将该创建方法聚合到工厂方法中
func newAk47() IGun {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// getGun 工厂方法，暴露给调用者，添加额外参数gunType来获取具体的产品
func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
