package tools

import (
	"log"
	"os"
    "io"
)

type File struct {
}

func (f *File) Copy(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		log.Println(err.Error())
		return
	}

	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)

}
