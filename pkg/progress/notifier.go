package progress

// Notifier is an interface that enforces implementation of progress and error handling.
type Notifier interface {
	NextStatus(Status)
	NextError(error)
}
