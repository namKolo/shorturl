package storage

// Storage is inteface
type Storage interface {
	Save(string) (string, error)
	Load(string) (string, error)
}
