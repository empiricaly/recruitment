package storage

// Store is an abstract KV store interface (Modeled after Badger...)
type Store interface {
	Txn(txn func(Transaction) error, writes bool) error
	Close() error
	DropAll() error
}

// Transaction is an abstract KV store transaction.
type Transaction interface {
	Delete([]byte) error
	Get([]byte) (Item, error)
	Set(key, val []byte) error
	Commit() error
}

// Item is a single value returned by a KV store.
type Item interface {
	Key() []byte
	String() string
	ValueCopy(dst []byte) ([]byte, error)
}

// Conn represents a datastore connection.
type Conn struct {
	Store
}

// Connect creates a connection to a messaging service with the given config.
func Connect(config *Config) (c *Conn, err error) {
	if config.Badger.Enabled {
		c, err = ConnectBadger(config)
	}

	// if config.Redis.Enabled {
	// 	c, err = ConnectRedis(config)
	// }

	return c, err
}
