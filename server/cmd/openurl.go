package cmd

// 打开系统默认浏览器

import (
	"os/exec"
)

func Open(url string) error {
	cmd := exec.Command("cmd", "/c", "start "+url)
	return cmd.Start()
}
