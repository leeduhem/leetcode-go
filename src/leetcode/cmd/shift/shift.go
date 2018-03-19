package shift

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"sort"
	"strconv"
	"strings"

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
	pDir := fmt.Sprintf("p%03d", id)
	srcDir, err := util.SrcDirOf(pDir)
	if err != nil {
		log.Fatal(err)
	}
	if verbose > 0 {
		log.Printf("Source directory for problem #%d: %s\n", id, srcDir)
	}

	existedSrcDirs := make([]string, 0)
	dirs, err := ioutil.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		if strings.HasPrefix(dir.Name(), pDir) {
			existedSrcDirs = append(existedSrcDirs, dir.Name())
		}
	}

	solutionNumber := 1

	if len(existedSrcDirs) > 1 {
		sort.Strings(existedSrcDirs)
		last := existedSrcDirs[len(existedSrcDirs)-1]
		idx := strings.LastIndex(last, "_")
		if idx == -1 {
			log.Fatalf("Invalid source directory name %s\n", last)
		}
		n, err := strconv.Atoi(last[idx+1:])
		if err != nil {
			log.Printf("Invalid source directory name %s\n", last)
			n = 0
		}
		solutionNumber = n + 1
	}

	copyPdir(srcDir, pDir, solutionNumber)

	// TODO: Update copied source directory
}

func copyPdir(srcDir, pDir string, num int) {
	newPdir := fmt.Sprintf("%s_%d", pDir, num)
	srcPath := path.Join(srcDir, pDir)
	dstPath := path.Join(srcDir, newPdir)

	files, err := ioutil.ReadDir(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()
		newFileName := fileName
		if strings.HasPrefix(fileName, pDir) {
			newFileName = fmt.Sprintf("%s_%d%s", pDir, num, fileName[len(pDir):])
		}

		util.Copy(path.Join(dstPath, newFileName),
			path.Join(srcPath, fileName))
	}
}
