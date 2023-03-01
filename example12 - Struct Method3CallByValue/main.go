package main

import (
	"fmt"
	"time"
)

type email struct {
	to   string
	from string
}

func (e email) From(s string) email {
	e.from = s
	return e
}

func (e email) To(s string) email {
	e.to = s
	return e
}

func (e email) Send() {
	fmt.Printf("From %s, To:%s\n", e.from, e.to)
}

func main() {
	e := &email{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			e.From(fmt.Sprintf("example%02d@example.com", i)).
				To(fmt.Sprintf("example%02d@example.com", i+1)).
				Send()
		}(i)
	}
	time.Sleep(1 * time.Second)

}
