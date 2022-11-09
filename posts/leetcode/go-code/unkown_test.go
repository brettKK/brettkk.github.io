package main
import "testing"
import "fmt"

type element struct {
	score int
	val int
	forward []*element
}

func test_nil(t *testing.T) {
	update := make([]*element, 10)
	fmt.Printf("%+v", update[0].forward[1].forward[2])
	t.Logf("%+v", update[1])
}