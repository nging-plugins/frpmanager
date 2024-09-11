package cmder

import (
	"strings"

	"github.com/coscms/webcore/initialize/backend"
	"github.com/coscms/webcore/library/common"
	"github.com/coscms/webcore/library/config"
	"github.com/webx-top/echo"
)

func onServerConfigChange(file string) error {
	id := config.FromCLI().GenerateIDFromConfigFileName(file, true)
	if len(id) == 0 {
		return common.ErrIgnoreConfigChange
	}
	if !config.FromCLI().IsRunning(`frpserver.` + id) {
		return common.ErrIgnoreConfigChange
	}
	if cm, err := GetServer(); err == nil {
		return cm.RestartBy(id)
	}
	return nil
}

func onClientConfigChange(file string) error {
	id := config.FromCLI().GenerateIDFromConfigFileName(file, true)
	if len(id) == 0 {
		return common.ErrIgnoreConfigChange
	}
	if !config.FromCLI().IsRunning(`frpclient.` + id) {
		return common.ErrIgnoreConfigChange
	}
	if cm, err := GetClient(); err == nil {
		return cm.RestartBy(id)
	}
	return nil
}

func init() {
	backend.OnConfigChange(func(filePath string) (err error) {
		if strings.Contains(filePath, echo.FilePathSeparator+`frp`+echo.FilePathSeparator+`server`+echo.FilePathSeparator) {
			err = onServerConfigChange(filePath)
		}
		return
	})
	backend.OnConfigChange(func(filePath string) (err error) {
		if strings.Contains(filePath, echo.FilePathSeparator+`frp`+echo.FilePathSeparator+`client`+echo.FilePathSeparator) {
			err = onClientConfigChange(filePath)
		}
		return
	})
}
