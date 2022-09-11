package temp

import (
	"fmt"
	"os"
)

func RemoveFile(file *os.File) {
	fileName := file.Name()
	err := file.Close()
	if err != nil {
		panic(fmt.Sprintf("Can't close %s file", fileName))
	}
	err = os.Remove(fileName)
	if err != nil {
		panic(fmt.Sprintf("Can't remove %s file", fileName))
	}
}
