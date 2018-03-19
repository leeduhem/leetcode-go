package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func Copy(dstPath, srcPath string) {
	srcFile, err := os.Stat(srcPath)
	if err != nil {
		log.Fatal(err)
	}

	if srcFile.IsDir() {
		copyDir(dstPath, srcPath)
	} else {
		copyFile(dstPath, srcPath)
	}
}

func copyFile(dstPath, srcPath string) {
	copy := func(dstPath, srcPath string) {
		src, err := os.Open(srcPath)
		if err != nil {
			log.Fatal(err)
		}
		defer src.Close()

		dst, err := os.Create(dstPath)
		if err != nil {
			log.Fatal(err)
		}
		defer dst.Close()

		size, err := io.Copy(dst, src)
		if verbose > 0 {
			log.Printf("Copy %d bytes to %s from %s\n", size, dstPath, srcPath)
		}
	}

	dstFile, err := os.Stat(dstPath)
	if os.IsNotExist(err) {
		dir, _ := path.Split(dstPath)
		if len(dir) > 0 {
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
		copy(dstPath, srcPath)
		return
	} else {
		log.Fatal(err)
	}

	if dstFile.IsDir() {
		_, srcFile := path.Split(srcPath)
		dstPath = path.Join(dstPath, srcFile)
	} else if !dstFile.Mode().IsRegular() {
		log.Fatalf("%s is neither a directory nor a regular file", dstPath)
	}

	copy(dstPath, srcPath)
}

func copyDir(dstPath, srcPath string) {
	copy := func(dstPath, srcPath string) {
		srcs, err := ioutil.ReadDir(srcPath)
		if err != nil {
			log.Fatal(err)
		}
		for _, src := range srcs {
			Copy(path.Join(dstPath, src.Name()),
				path.Join(srcPath, src.Name()))
		}
	}

	dstFile, err := os.Stat(dstPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(dstPath, 0777)
		if err != nil {
			log.Fatal(err)
		}

		copy(dstPath, srcPath)
		return
	} else {
		log.Fatal(err)
	}

	if !dstFile.IsDir() {
		log.Fatalf("%s is not a directory", dstPath)
	}

	copy(dstPath, srcPath)
}
