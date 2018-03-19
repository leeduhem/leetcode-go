package util

var verbose int

func SetVerbose(v bool) {
	if v {
		verbose = 1
	} else {
		verbose = 0
	}
}
