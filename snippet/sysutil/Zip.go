package fileutil

import (
	"archive/zip"
	"bytes"
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ZipPack(outPath string, inputPaths ...string) error {
	if err := os.MkdirAll(path.Dir(outPath), os.ModePerm); err != nil {
		return err
	}
	archive, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = archive.Close()
	}()
	zipWriter := zip.NewWriter(archive)
	defer func() {
		_ = zipWriter.Close()
	}()

	for _, inputPath := range inputPaths {
		inputPath = strings.TrimSuffix(inputPath, string(os.PathSeparator))
		err = filepath.Walk(inputPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}
			header.Method = zip.Deflate

			header.Name, err = filepath.Rel(filepath.Dir(inputPath), path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				header.Name += string(os.PathSeparator)
			}

			headerWriter, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer func() {
				_ = f.Close()
			}()
			_, err = io.Copy(headerWriter, f)
			return err
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func ZipUnpack(zipPath, outPath string) error {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer func() {
		_ = reader.Close()
	}()
	for _, file := range reader.File {
		if err := unzipFile(file, outPath); err != nil {
			return err
		}
	}
	return nil
}

func unzipFile(file *zip.File, outPath string) error {
	filePath := path.Join(outPath, file.Name)
	if file.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer func() {
		_ = rc.Close()
	}()

	w, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = w.Close()
	}()
	_, err = io.Copy(w, rc)
	return err
}

type ZipWriteFile struct {
	Name, Body string
}

func ZipWrite(writer io.Writer, files []ZipWriteFile) error {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			return fmt.Errorf("创建文件【%v】失败：%v", file.Name, err.Error())
		}
		if _, err = f.Write([]byte(file.Body)); err != nil {
			return fmt.Errorf("写入文件【%v】失败：%v", file.Name, err.Error())
		}
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("压缩包操作异常：%v", err.Error())
	}
	if _, err := buf.WriteTo(writer); err != nil {
		return fmt.Errorf("压缩包输出异常：%v", err.Error())
	}
	return nil
}

func ZipGetWriteFilesByEmbed(raw embed.FS, basePath string) ([]ZipWriteFile, error) {
	dir, err := raw.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("源文件目录[%v]读取异常：%v", basePath, err.Error())
	}
	return embedReadZipWriteFilesEach(raw, basePath, dir, []ZipWriteFile{})
}
func embedReadZipWriteFilesEach(raw embed.FS, useBasePath string, dir []fs.DirEntry, useList []ZipWriteFile) ([]ZipWriteFile, error) {
	for _, entry := range dir {
		useName := useBasePath + "/" + entry.Name()
		readDir := func() error {
			tmpDir, err := raw.ReadDir(useName)
			if err != nil {
				return fmt.Errorf("源文件子目录[%v]获取异常：%v", useName, err.Error())
			}
			tmpUseList, err := embedReadZipWriteFilesEach(raw, useName, tmpDir, []ZipWriteFile{})
			if err != nil {
				return err
			}
			useList = append(useList, tmpUseList...)
			return nil
		}
		readFile := func() error {
			tmpData, err := raw.ReadFile(useName)
			if err != nil {
				return fmt.Errorf("源文件[%v]获取异常:%v", useName, err.Error())
			}
			useList = append(useList, ZipWriteFile{
				Name: useName,
				Body: string(tmpData),
			})
			return err
		}
		if entry.IsDir() {
			if err := readDir(); err != nil {
				return nil, err
			}
		} else {
			if err := readFile(); err != nil {
				return nil, err
			}
		}
	}
	return useList, nil
}
