package syncx

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"
)

var rwMutex *sync.RWMutex
var wg *sync.WaitGroup

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func TestMutex(t *testing.T) {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
	//wg.Add(2)
	//
	////多个同时读取
	//go readData(1)
	//go readData(2)
	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)
	wg.Wait()
	fmt.Println("main..over...")
}
func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始写：write start。。")
	rwMutex.Lock() //写操作上锁
	fmt.Println(i, "正在写：writing。。。。")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写结束：write over。。")
}

func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始读：read start。。")
	rwMutex.RLock() //读操作上锁
	fmt.Println(i, "正在读取数据：reading。。。")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读结束：read over。。。")
}
