package sqlite

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type DatabaseExecutor interface {
	PrepareContext(ctx gin.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx gin.Context, query string, args ...interface{}) (*sql.Rows, error)
}
