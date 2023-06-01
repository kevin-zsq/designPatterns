策略模式strategy是一种行为设计模式， 它能让你定义一系列算法， 并将每种算法分别放入独立的类中， 以使算法的对象能够相互替换。

使用场景：
（1） 当你想使用对象中各种不同的算法变体， 并希望能在运行时切换算法时， 可使用策略模式
（2） 当你有许多仅在执行某些行为时略有不同的相似类时， 可使用策略模式：
策略模式让你能将不同行为抽取到一个独立类层次结构中， 并将原始类组合成同一个， 从而减少重复代码
（3） 如果算法在上下文的逻辑中不是特别重要， 使用该模式能将类的业务逻辑与其算法实现细节隔离开来
（4） 当类中使用了复杂条件运算符以在同一算法的不同变体中切换时， 可使用该模式：
    策略模式将所有继承自同样接口的算法抽取到独立类中， 因此不再需要条件语句。 原始对象并不实现所有算法的变体， 而是将执行工作委派给其中的一个独立算法对象

实现方式：
1、从上下文类中找出修改频率较高的算法 （也可能是用于在运行时选择某个算法变体的复杂条件运算符）
2、声明该算法所有变体的通用策略接口
3、将算法抽取到各自类中，每个算法类必须实现策略接口
4、在上下文类中添加一个成员变量用于保存对于策略对象的引用。 然后提供设置器以修改该成员变量。 上下文仅可通过策略接口同策略对象进行交互， 如有需要还可定义一个接口来让策略访问其数据
5、客户端必须将上下文类与相应策略进行关联， 使上下文可以预期的方式完成其主要工作


优缺点：
1、开闭原则。 你无需对上下文进行修改就能够引入新的策略
2、在运行时切换对象内的算法