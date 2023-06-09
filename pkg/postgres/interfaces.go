package postgres

import (
	"platform/pkg"
)

type DBEngine interface {
	pkg.DB
	Configure(...Option) DBEngine
}
