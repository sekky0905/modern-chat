package errs

import "fmt"

// Error は、error を wrap した interface。
type Error interface {
	error
	GetDetail() map[string]string
}

// NoSuchDataError は、指定したデータが存在しないことを表すエラー。
type NoSuchDataError struct {
	Detail  map[string]string
	BaseErr error
}

// NewNoSuchDataError は、NoSuchDataError を生成して、返す。
func NewNoSuchDataError(detail map[string]string, baseErr error) *NoSuchDataError {
	if baseErr != nil {
		return &NoSuchDataError{
			Detail:  detail,
			BaseErr: baseErr,
		}
	}

	return &NoSuchDataError{
		Detail: detail,
	}
}

// GetDetail は、詳細を取得する。
func (e *NoSuchDataError) GetDetail() map[string]string {
	if e == nil {
		return nil
	}
	return e.Detail
}

// Error は、エラーを返す。
func (e *NoSuchDataError) Error() string {
	if e.GetDetail() == nil {
		return "no such data error nil"
	}

	return fmt.Sprintf("no nosuch data error: %+v", e.GetDetail())
}
