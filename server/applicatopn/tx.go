package applicatopn

import "github.com/sekky0905/modern-chat/server/domain/repository"

// CloseTransaction は、トランザクションの後処理。
type CloseTransaction func(tx repository.DB, err error) error
