package model

// Client is client interface
type Client interface {
	Health() (status bool, err error)
	Shutdown() (err error)
	C2F(c float64) (f float64, err error)
	F2C(f float64) (c float64, err error)
}
