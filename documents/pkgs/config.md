# 配置文件工具的使用

[TOC]

## 技术选型

主要依赖的库为`yaml`库，故所有配置文件均采用该格式进行编写

## 使用规范

1. 在第一级的`config`目录下，创建yaml格式的配置文件
2. 在`config/configStruct`目录下，创建与该配置文件同名的go文件，其内容为读取该yaml文件的结构体定义
3. 调用`configurator.InitConfig`函数，参数分别为一个空的结构体和配置文件名，返回值即配置信息的结构体。

```go
var config configStruct.RedisConfig
configurator.InitConfig(
    &config, "redis.yaml",
)

redisCaches, _ := InitRedis(
    config.Host+":"+config.Port, config.Password, config.Databases,
)

fmt.Println(redisCaches)
```

## 其他

### 为什么没有使用`viper`库

`viper`相对功能更强大，但是多出来的功能可能不是特别能用到，现有方案是一个非常简单的封装，如果确实遇到对更强大的功能的需求，再考虑使用`viper`。

当配置文件中全部为确定内容时，`viper`库不依赖于某个确定的结构体，但是在读取配置时，需要用key-value的形式从返回的配置对象中去读取值，科学的做法是将这些key定义成常量，相比于定义某个确定的结构体来说，并不会轻松很多。

当配置文件中存在不确定项时，`viper`库同样需要定义一个结构体，与现有情况相同。

综上，暂时先完成了这个比较简单的配置文件库。
