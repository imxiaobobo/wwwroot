package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	f, err := os.Create("./day03/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	_, err = f.WriteString("我是你大爷")
	if err != nil {
		fmt.Println(err)
		return
	}
}

//func main() {
//	var smap sync.Map
//	smap.Store("name", "wangbo")
//	smap.Store("age", 33)
//	smap.Store("phone", "17551053834")
//
//	smap.Range(func(key, value interface{}) bool {
//		fmt.Println(key, value)
//		return true
//	})
//}

func Foo() {
	fmt.Println("hello world")
}

func mapTest() {
	var once sync.Once
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(Foo)
		}()
	}
	time.Sleep(10 * time.Second)
}

func MutexTest() {
	/**
	  sync.Mutex:互斥锁
	    -- 互斥锁保证同一时刻,只有一个协程可以访问某对象
	    -- mutex的初始值为解锁状态,通常作为其他结构体的匿名字段来使用,使其结构体具有上锁和解锁的功能
	    -- 线程安全,但是对没有加锁的变量进行解锁操作,会引发panic
	*/
	var wg sync.WaitGroup //等待组
	var m sync.Mutex      //互斥锁

	wg.Add(1000) //定义等待组能有1000次的加操作
	var num = 0
	for i := 0; i < 1000; i++ {
		go func() {
			m.Lock()         //数据需要修改前加锁
			defer m.Unlock() //数据修改完解锁,即可保证数据完整
			num += 2
			wg.Done() //每一个协程执行完毕,就会去减1
		}()
	}
	wg.Wait() //循环外就可以判断是否wg为0,阻塞着main函数
	fmt.Println(num)
}

func PoolTest() {
	/**
	  sync.pool:临时对象池
		-- 用于存储那些被分配了但是还没有使用,而未来可能会使用的值,以减少垃圾回收的压力
	    -- 协程安全的,应该用于管理协程共享的变量,不推荐用于非协程之间的对象管理
	    -- 调用new函数,将使用函数创建一个新对象返回
	    -- 从pool中取出对象时,如果pool中没有对象,则执行new,如果没有对象对new进行赋值,则返回nil
		-- 和栈类似,先进后出
	*/
	//代码示例
	bufPool := sync.Pool{New: func() interface{} {
		return new(bytes.Buffer)
	}}
	buf := bufPool.Get().(*bytes.Buffer) //使用的时候获取
	_, err := buf.WriteString("我是你爸爸")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(buf.String())
	bufPool.Put(buf) //使用之后存一个
}
