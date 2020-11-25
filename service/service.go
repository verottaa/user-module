package service

import (
	"github.com/verottaa/user-module/common"
)

type Service interface {
	common.Destroyable
	common.Reader
	common.Writer
}
