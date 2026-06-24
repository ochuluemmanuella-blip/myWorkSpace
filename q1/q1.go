package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Increment() {
	c.count++
}
func (c Counter) Value() int {
	return c.count

}
func main() {
	counter := &Counter{count: 0}
	counter.Increment()
	counter.Increment()
	counter.Increment()

	fmt.Println(counter.Value())
}
