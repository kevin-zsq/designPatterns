代理模式： proxy, 是一种结构型设计模式， 让你能够提供对象的替代品或其占位符。 代理控制着对于原对象的访问， 并允许在将请求提交给对象前后进行一些处理。


代理模式适合场景：
（1）延迟初始化 （虚拟代理）。 如果你有一个偶尔使用的重量级服务对象， 一直保持该对象运行会消耗系统资源时， 可使用代理模式。python中叫做lazySetting
（2）访问控制 （保护代理）。 如果你只希望特定客户端使用服务对象， 这里的对象可以是操作系统中非常重要的部分， 而客户端则是各种已启动的程序 （包括恶意程序）， 此时可使用代理模式
（3）本地执行远程服务 （远程代理）。 适用于服务对象位于远程服务器上的情形。 在这种情形中， 代理通过网络传递客户端请求， 负责处理所有与网络相关的复杂细节
（4）记录日志请求 （日志记录代理）。 适用于当你需要保存对于服务对象的请求历史记录时
（5）缓存请求结果 （缓存代理）。 适用于需要缓存客户请求结果并对缓存生命周期进行管理时， 特别是当返回结果的体积非常大时
（6）智能引用。 可在没有客户端使用某个重量级对象时立即销毁该对象

实现方式：
1、 创建一个代理类，代理类继承服务类或者实现服务类的所有接口
2、 代理类必须包含一个存储指向服务的引用的成员变量。通常情况下，代理负责创建服务并维护服务对象的整个生命周期管理
3、 根据需求实现代理方法。 在大部分情况下， 代理在完成一些任务后应将工作委派给服务对象
4、 可以考虑新建一个构建方法来判断客户端可获取的是代理还是实际服务。 你可以在代理类中创建一个简单的静态方法， 也可以创建一个完整的工厂方法
5、 可以考虑为服务对象实现延迟初始化

优点：
（1）在客户端毫不知情的情况下控制服务对象
（2）如果客户端对服务对象的生命周期没有特殊要求，可以由代理控制管理
（3）开闭原则：可以不对服务和客户端做出修改的情况下创建新代理