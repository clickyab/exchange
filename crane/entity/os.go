package entity

const (
	LinuxOS   = "linux"
	AndoidOS  = "android"
	WindowsOS = "windows"
)

// OS is the os
type OS struct {
	Valid  bool
	ID     int64
	Name   string
	Mobile bool
}

func IsMobileOS(name string) bool {
	switch name {
	case AndoidOS:
		return true

	default:
		return false
	}
}
