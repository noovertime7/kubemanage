package runtime

import "gorm.io/gorm"

type Pager interface {
	GetPage() int
	GetPageSize() int
	Fitter
}

type Fitter interface {
	IsFitter() bool
	Do(tx *gorm.DB)
}
