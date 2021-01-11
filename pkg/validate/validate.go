package validate

import (
	"net"
	"regexp"
)

const (
	RegExpLogin  = `^[a-zA-Z0-9]+$`
	RegExpIPv4   = `^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`
	RegExpDomain = `^([a-z0-9]+(-[a-z0-9]+)*\.)+[a-z]{2,}$`
	RegExpPort   = `^([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])$`
)

var (
	domain = regexp.MustCompile(RegExpDomain)
	login  = regexp.MustCompile(RegExpLogin)
	ipv4   = regexp.MustCompile(RegExpIPv4)
	port   = regexp.MustCompile(RegExpPort)
)

func IsDomain(s string) bool {
	return domain.MatchString(s)
}

func IsLogin(s string) bool {
	return login.MatchString(s)
}

func IsIPs(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil
}
