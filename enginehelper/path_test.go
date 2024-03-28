package tests

import (
	"e.coding.net/zhechat/magic/taihao/core"
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestOpenFile(t *testing.T) {
	path := filepath.Join(core.IaRoot(), "/diy-avatar/2024032802/381-123-张.jpg")
	fmt.Println(path)
	exec.Command(`cmd`, `/c`, `explorer`, path).Start()
}
