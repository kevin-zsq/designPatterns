生成器模式也称建造者模式，Builder。属于创建型设计模式，使你能够分步骤创建复杂对象。允许你使用相同的代码创建不同的类型的对象

适用场景：
产品构造参数复杂，绝大多数情况下参数没有使用。对象构造方式繁琐。在go语言中，strings.Builder 就是生成器模式。
使用生成器模式可以避免重复构造函数的出现，因为生成器允许你分步骤生成对象

实现方式：
（1）生成器Builder接口声明，接口声明所有类型生成器的产品构造方式，比如针对不同产品，需要声明所有可能需要构造对象的方法
（2）具体生成器：提供构造过程的不同实现
（3）主管类（Director）可以认为掌管具体产品需要执行哪些步骤的角色。需要和Builder生成器关联。因为需要使用Builder执行构造的每一个子步骤。
另外，Director定制了不同产品使用Builder的不同步骤
（4）客户端（Client）负责将Director和Builder关联起来

优点:
（1）生成器关注如何分步生成复杂对象。抽象工厂专门用于生产一系列相关对象。抽象工厂会马上返回产品，生成器允许你在获取产品前执行一些额外的构造步骤。
举个例子：gorm库中的Statement对象提供了WriteString/WriteByte/WriteQuoted等方法用于拼装最后的sql,Statement对象就像是我们的生成器






