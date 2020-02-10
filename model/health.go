package model

// Health checks if everything's fine
func Health() (status bool, err error) {
	status, err = client.Health()
	if err != nil {
		err = ErrClient
	}

	return
}
