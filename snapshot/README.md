备忘录模式： 快照、Snapshot、Memento，是一种行为设计模式， 允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。


适用场景：
（1）当你需要创建对象状态快照来恢复其之前的状态时， 可以使用备忘录模式
备忘录模式允许你复制对象中的全部状态 （包括私有成员变量）， 并将其独立于对象进行保存。 
尽管大部分人因为 “撤销” 这个用例才记得该模式， 但其实它在处理事务 （比如需要在出现错误时回滚一个操作） 的过程中也必不可少
（2）当直接访问对象的成员变量、 获取器或设置器将导致封装被突破时， 可以使用该模式
备忘录让对象自行负责创建其状态的快照。 任何其他对象都不能读取快照， 这有效地保障了数据的安全性
 
实现方式：
1、确定担任原发器角色的类（originator,可以产生备忘录（memento））。 重要的是明确程序使用的一个原发器中心对象， 还是多个较小的对象
2、创建备忘录类，逐一声明对应每个原发器成员变量的备忘录成员变量
3、将备忘录类设为不可变。 备忘录只能通过构造函数一次性接收数据。 该类中不能包含设置器
4、如果你所使用的编程语言支持嵌套类， 则可将备忘录嵌套在原发器中； 如果不支持， 那么你可从备忘录类中抽取一个空接口， 
然后让其他所有对象通过接口来引用备忘录。 你可在该接口中添加一些元数据操作， 但不能暴露原发器的状态
5、在原发器中添加一个创建备忘录的方法。 原发器必须通过备忘录构造函数的一个或多个实际参数来将自身状态传递给备忘录
6、在原发器类中添加一个用于恢复自身状态的方法。 该方法接受备忘录对象作为参数。 如果你在之前的步骤中抽取了接口， 
那么可将接口作为参数的类型。 在这种情况下， 你需要将输入对象强制转换为备忘录， 因为原发器需要拥有对该对象的完全访问权限
7、无论负责人是命令对象、 历史记录或其他完全不同的东西， 它都必须要知道何时向原发器请求新的备忘录、 如何存储备忘录以及何时使用特定备忘录来对原发器进行恢复
8、负责人与原发器之间的连接可以移动到备忘录类中。 在本例中， 每个备忘录都必须与创建自己的原发器相连接。 恢复方法也可以移动到备忘录类中， 
但只有当备忘录类嵌套在原发器中， 或者原发器类提供了足够多的设置器并可对其状态进行重写时， 这种方式才能实现


优点：
你可以在不破坏对象封装情况的前提下创建对象状态快照
你可以通过让负责人维护原发器状态历史记录来简化原发器代码

缺点：
如果客户端过于频繁地创建备忘录， 程序将消耗大量内存
负责人必须完整跟踪原发器的生命周期， 这样才能销毁弃用的备忘录
绝大部分动态编程语言 （例如 PHP、 Python 和 JavaScript） 不能确保备忘录中的状态不被修改
