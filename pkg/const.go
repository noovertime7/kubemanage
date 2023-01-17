package pkg

import "strconv"

const (
	ValidatorKey  = "ValidatorKey"
	TranslatorKey = "TranslatorKey"
)

const (
	LoginURL     = "/api/user/login"
	LogoutURL    = "/api/user/logout"
	WebShellURL  = "/api/k8s/pod/webshell"
	HostWebShell = "/api/cmdb/webshell"
)

var (
	AdminDefaultAuth      uint = 111
	AdminDefaultAuthStr        = strconv.Itoa(int(AdminDefaultAuth))
	UserDefaultAuth       uint = 222
	UserDefaultAuthStr         = strconv.Itoa(int(UserDefaultAuth))
	UserSubDefaultAuth    uint = 2221
	UserSubDefaultAuthStr      = strconv.Itoa(int(UserSubDefaultAuth))
)
