package goutils

import "os"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 19:08
  @describe :
*/

var (
	_hostname = "localhost"
)

// Hostname 获取本机名
func Hostname() string {
	if _hostname != "localhost" {
		return _hostname
	}
	_hostname, _ = os.Hostname()
	return _hostname
}
