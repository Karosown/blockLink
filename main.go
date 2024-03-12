package main

import (
	"blockLink/core/entity/block"
	"blockLink/core/entity/merkal"
	"log"
	"os"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	chain := block.NewBlockChain(64)
	for i := 0; i < 1024; i++ {
		f := func() func(i int) {
			return func(i int) {
				wg.Add(1)
				defer wg.Done()
				log.Println("正在处理第", i, "个协程")
				data := new(merkal.Message)
				data.Text = ("你好，世界！")
				data.Title = ("Hello,World!")
				data.CreateTime = (time.Now())
				chain.AddMsg(data)
			}
		}
		go f()(i + 1)
	}
	wg.Wait()
	a, _ := os.Create("123.json")
	a.WriteString(chain.String())
	a.Close()
}
