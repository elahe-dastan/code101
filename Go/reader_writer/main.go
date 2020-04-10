package main

import (
	"fmt"
	"time"
)

func main() {
	sharedMemory := newShared()

	for i := 0; i < 100; i++ {
		for j :=  0; j < 10; j++ {
			go sharedMemory.write()
		}
		go sharedMemory.read()
	}

	time.Sleep(2 * time.Second)
	fmt.Println(sharedMemory.sharedMemory)

}

type shared struct {
	sharedMemory int
	readLock chan int
	writeLock chan int
	readerNumber int
}

func newShared() shared {
	return shared {
		sharedMemory: 0,
		readLock:     make(chan int, 1),
		writeLock:    make(chan int, 1),
		readerNumber: 0,
	}
}

func (s *shared) write()  {
	s.writeLock<- 0
	s.sharedMemory += 1
	<-s.writeLock
}

func (s *shared) read()  {
	s.readLock<- 0
	s.readerNumber++

	if s.readerNumber == 1 {
		s.writeLock<- 0
	}
	<-s.readLock

	fmt.Println(s.sharedMemory)

	s.readLock<- 0
	s.readerNumber--

	if s.readerNumber == 0{
		<-s.writeLock
	}

	<-s.readLock
}

