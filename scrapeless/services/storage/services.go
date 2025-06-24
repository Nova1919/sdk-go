package storage

type Storage struct {
	*Dataset
	*KV
	*Object
	*Queue
}

var (
	defaultStorage Storage
)

func init() {
	// todo Judge whether it is online environment IS_ONLINE according to environment variables
	defaultStorage = Storage{
		Dataset: &Dataset{},
		KV:      &KV{},
		Object:  &Object{},
		Queue:   &Queue{},
	}
}

func NewStorage(env ...string) Storage {
	return defaultStorage
}

func (s *Storage) Close() error {
	return nil
}
