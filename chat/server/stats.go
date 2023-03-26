package main

type varUpdate struct {
	// Name of the variable to update
	varname string
	// Value to publish (int, float, etc.)
	value interface{}
	// Treat the count as an increment as opposite to the final value.
	inc bool
}

func statsInc(name string, val int) {
	if 
}
