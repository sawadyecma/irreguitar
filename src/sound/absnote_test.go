package sound

import "fmt"

func Example_absnote_Diff() {
	fmt.Println(C3.Diff(E3))
	fmt.Println(A4.Diff(A5))

	// Output:
	// 4
	// 12
}

func Example_absnote_Name() {
	fmt.Println(C3.Name())
	fmt.Println(C4.Name())
	fmt.Println(E5.Name())

	// Output:
	// C
	// C
	// E
}
