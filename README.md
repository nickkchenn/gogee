# gogee
webframework in go for practice

## DAY1
day1 实现了一个基本的struct Engine,这个struct 通过实现net/http中的handler接口，将所有地址的路由工作承接了起来

当然，只实现了最基本的静态地址的路由，其实就是将特定地址和对应handler进行绑定

API 是以post/get的形式设计的，后面可以看一下gin的API 设计，这种设计应该是比较易用

## DAY2

DAY1 版的实现需要用户去处理request 和 response 的所有内容，意味着需要处理请求header里类似状态码，content-type之类的信息，这类信息反复处理是冗余的，所以我们要通过新的结构体，把这一类反复使用的信息的编解码封装起来，提供出接口，简化使用。

这里设计了context结构体进行这种封装

    针对使用场景，封装*http.Request和http.ResponseWriter的方法，简化相关接口的调用，只是设计 Context 的原因之一。对于框架来说，还需要支撑额外的功能。例如，将来解析动态路由/hello/:name，参数:name的值放在哪呢？再比如，框架需要支持中间件，那中间件产生的信息放在哪呢？Context 随着每一个请求的出现而产生，请求的结束而销毁，和当前请求强相关的信息都应由 Context 承载。因此，设计 Context 结构，扩展性和复杂性留在了内部，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。



