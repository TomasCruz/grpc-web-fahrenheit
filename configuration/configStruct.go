package configuration

// Routes contains routes of this service
type Routes struct {
	HealthRoute string
	C2FRoute    string
	F2CRoute    string
}

// Config holds environment variable values, it's populated on startup
type Config struct {
	Port     string
	GrpcPort string
	GrpcHost string
}
