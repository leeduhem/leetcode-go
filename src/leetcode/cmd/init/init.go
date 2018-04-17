package init

import (
	"leetcode/util"
	"log"
	"os"
	"path"
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
	pn := util.Pid2Pkg(id)
	pp := path.Join(util.SrcDir(), pn)
	if err := os.MkdirAll(pp, 0777); err != nil {
		log.Fatal(err)
	}

	// Default code
	code, err := util.DefaultCode(id, "golang")
	if err != nil {
		log.Fatal(err)
	}

	pkg, err := os.Create(path.Join(pp, pn) + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer pkg.Close()

	_, err = pkg.WriteString(code)
	if err != nil {
		log.Fatal(err)
	}

	// Corresponding test code
	code_test, err := util.DefaultTestCode(id, code, "golang")
	if err != nil {
		log.Fatal(err)
	}

	pkg_test, err := os.Create(path.Join(pp, pn) + "_test.go")
	if err != nil {
		log.Fatal(err)
	}
	defer pkg_test.Close()

	_, err = pkg.WriteString(code_test)
	if err != nil {
		log.Fatal(err)
	}
}
