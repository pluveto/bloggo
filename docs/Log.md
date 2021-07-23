2021/07/23 Friday

打算这样分层：

+ API: 负责直接处理用户的请求
+ Service：领域业务，负责和数据仓库交互
+ Repo：数据仓库，负责操作数据库、缓存等

今天在考虑怎么把数据库连接传递给 repo 层。一种方法是使用全局变量，类似单例模式。使用 `repo.dbConn` 或者 `repo.GetConn()` 获取连接。

但是缺点也是明显的，无法隔离数据库，任何地方都可以访问它。而且过于分散。还有就是如果同时连接不同种数据库，

还有个方法就是用个结构体把连接包起来，然后用这个结构体作为接收器。缺点就是耦合太重了。

还有个方法是把连接一级一级地传下去。但是这样用不到它的层也要传递它。

还有一种做法是用一个容器将各种连接、服务注入进去，然后依赖它的类将这些实例取出。这和全局对象没有本质的差别。

还有利用 gin 的 context 的 K/V 特性做容器的，但是这样会让底层依赖 gin，我觉得没必要。

还是先梳理一下关系：


```
            gin
-------------|------------------
            API             
-------------|------------------
            Service             
-------------|------------------
            Repo             
             |
            / \
       MySQL   Redis             
--------------------------------

依赖关系：

Repo -> MySQL + Redis
Service -> Repo
API -> Service, gin
```

可以看到，真正操作数据库连接的是 Repo 层。应该尽量把数据库的一切都放到 Repo 层。

而 main 函数负责启动各个层。至于启动的细节，应该下放到各层内部实现。

Repo -> database/sql + IRepo
Service -> IRepo
API -> IService, gin

数据库字段长度问题：

+ 名称字段：varchar(200)
+ 较长的名称字段/简介字段：varchar(500)
+ 特别长的描述字段： varchar(2000)
+ 超过2000中文字的字段：text

https://galaxyyao.github.io/2019/07/30/MySQL-%E6%B2%A1%E6%9C%89%E5%BF%85%E8%A6%81%E7%9A%84varchar-255-%E9%95%BF%E5%BA%A6%E5%8F%8A%E5%AD%98%E5%82%A8%E6%B1%89%E5%AD%97%E9%97%AE%E9%A2%98%E6%B1%87%E6%80%BB/

在 SQLite 中，则不必：

> SQLite does not enforce the length of a VARCHAR. You can declare a VARCHAR(10) and SQLite will be happy to store a 500-million character string there. And it will keep all 500-million characters intact. Your content is never truncated. SQLite understands the column type of "VARCHAR(N)" to be the same as "TEXT", regardless of the value of N.
> https://www.sqlite.org/faq.html#q9

现在又遇到的问题是控制器调用服务。

自然可以在控制器里面创建一个服务，但是这样每个请求都要产生一堆内存分配。

也可以把服务层写成静态的。但是这样不利于扩展。

也可以搞依赖注入，但是耦合就呵呵了。

还可以用单例模式，但是单例模式不便于带有运行时信息地创建，也难以适应更加复杂的构造。

最后还是决定仿照 Repo 层的写法。用 SetupService 一次性在内存创建所有服务。

刚刚偷瞄了一下屑站的源码，发现他们竟然用全局变量，醉了醉了：

https://github.com/1315402725/go-common/blob/master/app/interface/main/web/http/http.go

果然再精妙的设计都不如方便实在？看着日益复杂的代码结构，我悟了。

最终决定采用这种架构：

```go

type Service struct {
	repo *repository.Repository
}

func New(c *bloggo.Config) *Service {
	var srv = &Service{}
	srv.repo = repository.New(c)
	return srv
}

func (s *Service) GetSetting(key string) interface{} {
	return s.repo.GetSetting(key)
}
```

也即全部用 struct 构成各层实例和层实例的私有、公有属性，如此可以保证只有 Serivce 层能访问到 `repo`，同时避免全局变量。