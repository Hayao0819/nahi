package nconf

type Nconf[T any] struct {
	dirs      []string
	filenames []string
	loaded    bool
	data      T
}
