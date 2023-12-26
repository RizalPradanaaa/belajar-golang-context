package belajargolangcontext

import (
	"context"
	"fmt"
	"testing"
)

// Membuat Context
func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}


// Context With Value
func TestContextWithValue(t *testing.T) {

	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "f", "F")

	contextF := context.WithValue(contextC, "f", "F")


	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)



	// Context Get Value
	fmt.Println(contextF.Value("f"))	// Dapat
	fmt.Println(contextF.Value("c"))	// Dapat Milik Parent
	fmt.Println(contextF.Value("b"))	// Tidak dapat, beda parent
	fmt.Println(contextA.Value("b"))	// Tidak dapat mengambil data child
}
