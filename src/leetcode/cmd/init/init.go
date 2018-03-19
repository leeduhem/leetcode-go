package init

import (
	"leetcode/util"
)

var verbose int

func SetVerbose(v bool) {
	if v {
		verbose = 1
		util.SetVerbose(v)
	} else {
		verbose = 0
		util.SetVerbose(v)
	}
}

func Do(id int) {
	println("Initiate source code directory for problem ", id)
}
