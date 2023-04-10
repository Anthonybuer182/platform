package mysqld

import (
	"platform/pkg"
)

type DBEngine interface {
	pkg.DB
	Configure(...Options) DBEngine
}
