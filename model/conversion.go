package model

// C2F converts degrees celsius to degrees fahrenheit
func C2F(c float64) (f float64, err error) {
	return client.C2F(c)
}

// F2C converts degrees fahrenheit to degrees celsius
func F2C(f float64) (c float64, err error) {
	return client.F2C(f)
}
