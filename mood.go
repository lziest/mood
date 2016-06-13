// package mood hsows a example of a exported method whose behavior
// is controlled by package exported variable.
package mood

import (
	"fmt"
)

// Status is a simple string
var Status = "happy"

func init() {
	fmt.Printf("init(): mood.Status ptr = %p\n", &Status)
}

// Show shows mood status
func Show() {
	fmt.Println("I'm", Status)
}
