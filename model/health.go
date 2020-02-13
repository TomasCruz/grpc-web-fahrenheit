package model

// Health checks if everything's fine
func Health() (status bool, err error) {
	return client.Health()
}
