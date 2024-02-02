package exutils

import "os/exec"

// コマンドの存在確認
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
