package repository

import (
	"github.com/verottaa/user-module/common"
)

type Repository interface {
	common.Destroyable
	common.Reader
	common.Writer
}
