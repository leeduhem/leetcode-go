package init

var verbose int

func SetVerbose(v bool) {
	if v {
		verbose = 1
	} else {
		verbose = 0
	}
}

func Do(id int) {
	println("Initiate source code directory for problem ", id)
}
