package main

// PointerPassing .
func PointerPassing(x, y *int) {

	*x, *y = *x*5, *y*4
}
