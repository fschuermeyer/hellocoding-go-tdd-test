package gotddtest

import (
	"fmt"

	"github.com/fschuermeyer/hellocoding-go-tdd-test/internal/example"
)

func Main() {
	a := 1
	b := 2

	response := example.Sum(a, b)

	fmt.Println(response)
}
