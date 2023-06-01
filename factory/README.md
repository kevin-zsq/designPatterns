工厂方法模式是一种**创建型**设计模式， 给使用者提供一个创建对象的方法， 允许使用者决定实例化对象的类型（编写接口逻辑）。
思想在go语言中很常见，虽然go中没有继承的概念，但是interface接口可以限制类型的方法。在实现接口的时候其实就是重写子类的方法，达到自定义子类的目的。
实现接口的类型也是工厂模式生产的子类。


适用场景1：
后续代码可能需要兼容更多的场景和类别，出于对扩展性考量：
（1）在实现工厂模式时，创建产品的代码和实际产品的代码进行分离。如果需要添加新产品，只需要开发新的类别，新类别满足interface接口即可
（2）对于基础框架代码或者扩展软件库，如果希望给使用者提供重写的方式，在C++或者python里面，常采用继承的方式来实现。对于go语言，可以定义interface,
给用户提供重写接口的方式达到目的
（3）连接池：对于处理大型资源密集型对象，比如数据库连接、文件系统和网络资源时，我们常常需要复用现有对象。我们可以提供一个函数既能创建新对象，又能重用
现有方法。这个函数就是工厂函数的构造方法


实现方式：
1、所有产品都遵循同一接口，该接口声明对所有产品都有意义的方法，接口定义，抽象类的抽象方法或者go语言中的interface
2、在创建类中添加一个空的工厂方法。该方法返回类型遵循通用的产品接口，生产函数。你可能需要新增一些额外的参数来区分具体需要返回哪些产品类型
3、定义每一种产品的创建者子类。然后将对应创建方法移到工厂方法中
4、对于子类过多的情况，可以分类创建基类

优点：
1、通过工厂方法解耦创建者和具体类型
2、单一职责原则(SRP：如果一个类有多于一个的动机被改变，那么这个类就具有多于一个的职责)：将产品创建代码放在程序的单一位置，入口确定，代码易维护
3、开闭原则（OCP: 对于扩展是开放的，但是对于修改是封闭的）：无需修改客户端代码，就可以在程序中引入新的产品类型

由于 Go 中缺少类和继承等 OOP 特性， 所以无法使用 Go 来实现经典的工厂方法模式。 不过， 我们仍然能实现模式的基础版本， 即简单工厂。
