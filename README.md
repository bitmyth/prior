# prior


This package provide a easy and efficient framework for arranging all initialization functions in a golang project, like init() function, but we should prefer explicitly calling initialization function which makes initialization process more clear and easy to understand instead of spread into init() functions everywhere.

There are several advantages to explictly state initialization call and gather all initialization function calls in one file.

First,this will produce easy and understandable code, it clearly express the dependancy relations of them.

Second, this will make writing test more easy and flexible. if initialization call are placed into init() functions everywhere in the codebase, we lost control to disabling some of these initialization calls.
If we want to disable some of them,we need to temporarily comment out those calling lines which is tidious and ugly.

All initialization functions forms a dependency tree. 
Ideally, We want parent node start running before its children node.Besides we want the nodes at the same level to start concurrently to make overall time consumed for initialization more short.


这个库用于解决 Go 项目初始化操作的启动问题。提供了一个简单且高效的组织初始化操作的框架

我们应该更喜欢显式调用初始化函数，这使得初始化过程更加清晰易懂，而不是分散到各处的 init() 函数中。

常见的初始化操作比如读配置文件，建立MySQL数据库连接等。


所有初始化操作是树形结构，父节点需要先运行，子节点后运行，
同级节点可以并发运行。
合理的依赖项启动模式应该允许同级初始化操作同时运行，这样可以减少启动总时间。

显式声明初始化调用并将所有初始化函数调用收集在一个文件中有几个优点。

首先，这将产生简单易懂的代码，它清楚地表达了它们的依赖关系。

其次，这将使编写测试变得更加容易和灵活。 如果初始化调用被放入代码库中各处的 init() 函数中，我们就会失去对禁用其中一些初始化调用的控制。 如果我们想禁用其中一些，我们需要暂时注释掉那些整洁且丑陋的调用行。

所有初始化函数形成依赖树。 理想情况下，我们希望父节点在其子节点之前开始运行。此外，我们希望同一级别的节点同时启动，以使初始化所消耗的总时间更短。


## demo
```go
package main

import (
	"github.com/bitmyth/prior/pkg/prior"
	"log"
	"time"
)

type BootConfig struct{}

func (b BootConfig) Initialize() error {
	time.Sleep(time.Second)
	log.Println("Config Initialized")
	return nil
}

type BootMySQL struct{}

func (b BootMySQL) Initialize() error {
	time.Sleep(time.Second)
	log.Println("MySQL Initialized")
	return nil
}

type BootRedis struct{}

func (b BootRedis) Initialize() error {
	time.Sleep(time.Second)
	log.Println("Redis Initialized")
	return nil
}

func main() {
	var bootConfig BootConfig
	var bootMysql BootMySQL
	var bootRedis BootRedis

	prior.Register(bootConfig).
		Register(bootMysql, bootRedis)

	prior.Boot()

	println("BOOTED !")
}
```
