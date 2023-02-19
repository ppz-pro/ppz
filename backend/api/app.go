package api

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type App struct{}

func (app *App) LoadConfig() _Result {
	home, err := homedir.Dir()
	if err != nil {
		return ErrResult("获取家目录失败", err)
	}

	ppzDir := filepath.Join(home, ".ppz0523")
	mkErr := os.Mkdir(ppzDir, 0600)
	if mkErr != nil {
		if os.IsExist(mkErr) {
			// 创建文件夹
		}
		// error: 创建文件失败
	}
	return NewResult("todo")
}
