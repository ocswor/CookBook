The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled,
 all Contexts derived from it are also canceled.
 
 在userip中通过 withValue 派生出的新的context. 这些值构成一棵树.
 当 一个context 被取消的时候,所有的派生的context都会被取消.
 
 Background is the root of any Context tree; it is never canceled:
 
 background返回的context 是这棵树的根节点,永远不会被取消.
  
  
 WithCancel and WithTimeout return derived Context values that can be canceled sooner than the parent Context. The Context associated with an incoming request is typically canceled when the request handler returns.
 WithCancel is also useful for canceling redundant requests when using multiple replicas. 
 WithTimeout is useful for setting a deadline on requests to backend servers:
 
 
 WithValue provides a way to associate request-scoped values with a Context:
 
[官方博客](https://blog.golang.org/context)

结合 handlers中 handleSearch 的方法,可以了解,WithCancel,WithTimeout,WithValue,这三个派生出新的context的具体作用.

WithCancel 可以手动触发,取消函数
WihtTimeout 用于管理请求超时,请求截止的上下文
WithValue 可以在传递上下文的时候,传递请求作用域值,在多个函数之前传递用户的标示,做出对应的处理.


 
 