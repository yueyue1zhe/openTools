package file

import (
	"fmt"
	"openTools/y/compute"
	"openTools/y/global"
	"time"
)

type makePath struct {
	AttachmentPath func(string2 string) string
}

func (y *File) MakePath() *makePath {
	return &makePath{
		AttachmentPath: y.AttachmentPath,
	}
}
func (y *makePath) Video(ext string) (path string, err error) {
	return y.makeAttachPath("videos", ext, "")
}
func (y *makePath) Audio(ext string) (path string, err error) {
	return y.makeAttachPath("audios", ext, "")
}
func (y *makePath) Image(ext string) (path string, err error) {
	return y.makeAttachPath("images", ext, "")
}
func (y *makePath) Excel(name string) (path string, err error) {
	return y.makeAttachPath("excel", ".xlsx", name)
}

func (y *makePath) MakePemPath(name string) (path string) {
	path = fmt.Sprintf("%v/cert/", global.NewGlobal().IaRoot)
	if err := pathExists(path); err != nil {
		return ""
	}
	return path + name
}

func (y *makePath) MakeVersionPath(version string) (path string) {
	path = fmt.Sprintf("%v/patch/%v/", global.NewGlobal().IaRoot, version)
	if err := pathExists(path); err != nil {
		return ""
	}
	return path
}

func (y *makePath) makeAttachPath(pre, ext, name string) (path string, err error) {
	path = pre + time.Now().Format("/2006/01/")
	if err = pathExists(y.AttachmentPath(path)); err != nil {
		return "", err
	}
	yRandom := compute.NewCompute().NewRandom()
	if name != "" {
		path += name + time.Now().Format("20060102150405") + yRandom.String(2)
	} else {
		path += yRandom.String(30)
	}
	return path + ext, nil
}
func (y *makePath) MakeTmpPath(ext string) (path string, err error) {
	return y.makeAttachPath("tmp", ext, "")
}
