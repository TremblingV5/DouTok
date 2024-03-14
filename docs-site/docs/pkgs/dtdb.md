# dtdb

> `dtdb` 意为 `DouTok Database`, 对数据库相关功能的二次封装。

## 统一事务开启方式

### 事务相关结构体

#### txKey

用于标识事务对象在 `context.Context` 中的键名。

#### TX

标识一个事务对象对应的接口，目前包括`Commit`和`Rollback`两个方法。

#### TXHandle

TX控制器，主要用于确定创建事务的具体方法。由于DouTok使用微服务的架构，不同服务的数据库事务对象不一致，所以使用`TXHandle`来确定具体的事务对象。

### 事务相关方法

#### NewTXHandle

用于创建一个新的`TXHandle`对象。需要传入具体的事务对象创建方法。例如：

```golang
txHandle := NewTXHandle(func() TX) {
    return query.Q.Begin()
})
```

#### WithTx

向`context.Context`中添加一个事务对象。例如：

```golang
ctx = WithTx(ctx, tx)
```

#### Tx

从`context.Context`中获取事务对象。例如：

```golang
tx := Tx(ctx)
```

#### persist

具体的持久化方法，不向外暴露。主要流程为从传入的`context.Context`中获取事务对象，然后调用`Commit`方法，如果失败则调用`Rollback`方法。

#### WithTXPersist

在服务的service目录中使用，开启事务。例如：

```golang
func doSomething(ctx context.Context) (err error) {
	ctx, persist := txHandle.WithTXPersist(ctx)
	defer func() {
		persist(err)
    }()
}
```

### 事务使用示例

1. 在`service`目录中使用`WithTXPersist`开启事务。

```golang
func doSomething(ctx context.Context) (err error) {
	ctx, persist := txHandle.WithTXPersist(ctx)
	defer func() {
		persist(err)
    }()
}
```

2. 在`repository`目录中使用`Tx`获取事务对象。
3. 对事务对象进行操作。

```golang
func (p *PersistRepository) LoadUserListByIds(ctx context.Context, ids ...uint64) ([]*model.User, error) {
	tx := dtdb.Tx(ctx).(*query.QueryTx)
	return tx.User.Where(query.User.ID.In(ids...)).Find()
}
```
