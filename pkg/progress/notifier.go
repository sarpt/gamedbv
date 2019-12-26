package progress

// Notifier is an interface that enforces implementation of progress and error handling
type Notifier interface {
	NextProgress(string)
	NextError(error)
}
