package dbsshmysql

import "github.com/gomodul/dbssh"

// Open ...
func Open(c dbssh.Config) (*dbssh.SSH, string, error) {
	return dbssh.Open(c, new(Driver))
}
