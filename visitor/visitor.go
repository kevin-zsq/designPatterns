package visitor

import "fmt"

// 示例说明
//访问者模式允许你在结构体中添加行为， 而又不会对结构体造成实际变更。 假设你是一个代码库的维护者， 代码库中包含不同的形状结构体， 如：
//
//方形
//圆形
//三角形
//上述每个形状结构体都实现了通用形状接口。
//
//在公司员工开始使用你维护的代码库时， 你就会被各种功能请求给淹没。 让我们来看看其中比较简单的请求： 有个团队请求你在形状结构体中添加 getArea获取面积行为。
//
//解决这一问题的办法有很多。
//
//第一个选项便是将 getArea方法直接添加至形状接口， 然后在各个形状结构体中进行实现。 这似乎是比较好的解决方案， 但其代价也比较高。
//作为代码库的管理员， 相信你也不想在每次有人要求添加另外一种行为时就去冒着风险改动自己的宝贝代码。 不过， 你也一定想让其他团队的人还是用一用自己的代码库。
//
//第二个选项是请求功能的团队自行实现行为。 然而这并不总是可行， 因为行为可能会依赖于私有代码。
//第三个方法就是使用访问者模式来解决上述问题。 首先定义一个如下访问者接口：
//type visitor interface {
//	visitForSquare(square)
//	visitForCircle(circle)
//	visitForTriangle(triangle)
//}
//如果添加任何其他行为， 比如 getNumSides获取边数和 getMiddleCoordinates获取中点坐标 ， 我们将使用相同的 accept(v visitor)函数， 而无需对形状结构体进行进一步的修改。
//
//最后， 形状结构体只需要修改一次， 并且所有未来针对不同行为的请求都可以使用相同的 accept 函数来进行处理。 如果团队成员请求 getArea行为， 我们只需简单地定义访问者接口的具体实现， 并在其中编写面积的计算逻辑即可。
// 具体代码如下：

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForrectangle(t)
}

func (t *Rectangle) getType() string {
	return "rectangle"
}

// Visitor 是一个interface, 每一个具体Visitor实现了不同的算法
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForrectangle(*Rectangle)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
func (a *AreaCalculator) visitForrectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	// Calculate middle point coordinates for square.
	// Then assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) visitForrectangle(t *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
