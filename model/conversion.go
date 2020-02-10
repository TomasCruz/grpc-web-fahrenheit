package model

// C2F converts degrees celsius to degrees fahrenheit
func C2F(c float64) (f float64, err error) {
	f, err = client.C2F(c)
	if err != nil {
		err = ErrClient
	}

	return
}

// F2C converts degrees fahrenheit to degrees celsius
func F2C(f float64) (c float64, err error) {
	c, err = client.F2C(f)
	if err != nil {
		err = ErrClient
	}

	return
}
