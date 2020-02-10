package client

func (g grpcClient) Shutdown() (err error) {
	return g.conn.Close()
}
