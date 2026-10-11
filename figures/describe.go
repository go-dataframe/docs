package main

func init() { registerOutput("describe", describe) }

// df.describe(): count, mean, std, min, quartiles and max of each numeric column.
func describe() (string, string) {
	d := rainfall().Describe()
	return d.String(), d.HTML()
}
