package p093

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var restore func (s string, p int, ip string, ips *[]string);
	restore = func (s string, p int, ip string, ips *[]string) {
		if p >= 4 && s == "" {
			*ips = append(*ips, ip)
			return
		}

		if p >= 4 || s == "" {
			return
		}

		if ip != "" {
			ip += "."
		}

		restore(s[1:], p+1, ip + s[0:1], ips)
		if len(s) >= 2 && s[0] != '0' {
			restore(s[2:], p+1, ip + s[0:2], ips)
			if len(s) >= 3 {
				n, ok := strconv.Atoi(s[0:3])
				if ok == nil && n <= 255 {
					restore(s[3:], p+1, ip + s[0:3], ips)
				}
			}
		}
	}

	ips := []string {}
	restore(s, 0, "", &ips)

	return ips
}
