package fileutil

import (
	"os"
)

func Path2DirAllFileList(path string) (list []string) {
	return path2dirAllFileListDose(path, []string{})
}

func path2dirAllFileListDose(path string, input []string) []string {
	dir, err := os.ReadDir(path)
	if err != nil {
		return input
	}
	for _, entry := range dir {
		if entry.IsDir() {
			children := path2dirAllFileListDose(path+"/"+entry.Name(), []string{})
			for _, child := range children {
				input = append(input, child)
			}
		} else {
			input = append(input, path+"/"+entry.Name())
		}
	}
	return input
}
