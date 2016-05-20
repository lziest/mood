package mood

var Status = "sad"

func init() {
	// Only define flag once.
	if flag.Lookup("mood") == nil {
		flag.StringVar(&Status, "mood", "sad", "mood is a string")
	}
	fmt.Println("mood.Status ptr = %p\n", &Status)
}
