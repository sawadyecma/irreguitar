package sound

import "fmt"

func Example_absnote_Diff() {
	fmt.Println(C3.Diff(E3))
	fmt.Println(A4.Diff(A5))

	// Output:
	// 4
	// 12
}
