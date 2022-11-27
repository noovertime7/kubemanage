package pkg

import "strconv"

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
)

var (
	AdminDefaultAuth      uint = 111
	AdminDefaultAuthStr        = strconv.Itoa(int(AdminDefaultAuth))
	UserDefaultAuth       uint = 222
	UserDefaultAuthStr         = strconv.Itoa(int(UserDefaultAuth))
	UserSubDefaultAuth    uint = 2221
	UserSubDefaultAuthStr      = strconv.Itoa(int(UserSubDefaultAuth))
)
