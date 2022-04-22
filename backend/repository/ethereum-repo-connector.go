package repository

type EthereumRepoConnector struct {
	dupa string
}

// GetConnector implements RepoConnector
func (c *EthereumRepoConnector) GetConnector() *EthereumRepoConnector {
	return c
}

// InitConnection implements RepoConnector
func (c *EthereumRepoConnector) InitConnection(name, connectionString string) error {
	return nil
}
