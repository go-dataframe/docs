package main

func init() { registerOutput("display", display) }

// What a notebook cell whose value is a frame shows: Display returns both
// MIME types, text/plain (String) and text/html (HTML).
func display() (string, string) {
	df := rainfall()
	d := df.Display()
	return string(d["text/plain"]), string(d["text/html"])
}
