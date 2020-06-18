package main

import "fmt"

func main() {
	sum := 0

	m := New()
	d := make(chan string)

	for i := 0; i < 1000; i++ {
		go func(mutex Mutex, done chan string) {
			mutex.Lock()
			sum++
			if sum == 1000 {
				d<- "done"
			}
			mutex.Unlock()
		}(m, d)
	}

	<-d

	fmt.Println(sum)
}

type Mutex struct {
	l chan string
}

func New() Mutex {
	return Mutex{l:make(chan string, 1)}
}
func (m Mutex) Lock()  {
	m.l<- "lock"
}

func (m Mutex) Unlock()  {
	<-m.l
}