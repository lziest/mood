// package mood hsows a example of a exported method whose behavior
// is controlled by package exported variable.
package mood

import (
	"flag"
	"fmt"
)

// Status is a simple string
var Status = "happy"

// Show shows mood status
func Show() {
	fmt.Printf("mood.Status ptr = %p\n", &Status)
	fmt.Println("I'm", Status)
}
