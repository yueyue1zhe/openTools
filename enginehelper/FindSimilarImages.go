package enginehelper

func FindSimilarImages(targetFile *os.File, waitPath string) (dPath []string, err error) {
	imgQuery, _, err := image.Decode(targetFile)
	if err != nil {
		return
	}
	queryHash, err := goimagehash.AverageHash(imgQuery)
	if err != nil {
		return
	}
	err = filepath.Walk(waitPath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			tmpFile, err := os.Open(path)
			if err != nil {
				return err
			}
			img, _, err := image.Decode(tmpFile)
			if err != nil {
				return err
			}
			checkHash, err := goimagehash.AverageHash(img)
			if err != nil {
				return err
			}
			distance, _ := queryHash.Distance(checkHash)
			if distance < 6 {
				dPath = append(dPath, path)
			}
		}
		return nil
	})
	return
}
