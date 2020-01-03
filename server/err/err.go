package err

// Error は、error を wrap した interface。
type Error interface {
	error
	GetDetail() map[string]string
}
