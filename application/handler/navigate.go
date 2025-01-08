package handler

import (
	"github.com/coscms/webcore/library/navigate"
	"github.com/webx-top/echo"
)

var LeftNavigate = &navigate.Item{
	Display: true,
	Name:    echo.T(`内网穿透`),
	Action:  `frp`,
	Icon:    `road`,
	Children: &navigate.List{
		{
			Display: true,
			Name:    echo.T(`服务端配置`),
			Action:  `server_index`,
		},
		{
			Display: false,
			Name:    echo.T(`查看服务端日志`),
			Action:  `server_log`,
		},
		{
			Display: false,
			Name:    echo.T(`添加服务端配置`),
			Action:  `server_add`,
		},
		{
			Display: false,
			Name:    echo.T(`修改服务端配置`),
			Action:  `server_edit`,
		},
		{
			Display: false,
			Name:    echo.T(`删除服务端配置`),
			Action:  `server_delete`,
		},

		{
			Display: true,
			Name:    echo.T(`账号管理`),
			Action:  `account`,
		},
		{
			Display: false,
			Name:    echo.T(`添加FRP账号`),
			Action:  `account_add`,
		},
		{
			Display: false,
			Name:    echo.T(`修改FRP账号`),
			Action:  `account_edit`,
		},
		{
			Display: false,
			Name:    echo.T(`删除FRP账号`),
			Action:  `account_delete`,
		},

		{
			Display: true,
			Name:    echo.T(`客户端配置`),
			Action:  `client_index`,
		},
		{
			Display: false,
			Name:    echo.T(`查看客户端日志`),
			Action:  `client_log`,
		},
		{
			Display: false,
			Name:    echo.T(`添加客户端配置`),
			Action:  `client_add`,
		},
		{
			Display: false,
			Name:    echo.T(`修改客户端配置`),
			Action:  `client_edit`,
		},
		{
			Display: false,
			Name:    echo.T(`删除客户端配置`),
			Action:  `client_delete`,
		},
		{
			Display: true,
			Name:    echo.T(`分组管理`),
			Action:  `group_index`,
		},
		{
			Display: false,
			Name:    echo.T(`添加分组`),
			Action:  `group_add`,
		},
		{
			Display: false,
			Name:    echo.T(`修改分组`),
			Action:  `group_edit`,
		},
		{
			Display: false,
			Name:    echo.T(`删除分组`),
			Action:  `group_delete`,
		},
		{
			Display: false,
			Name:    echo.T(`统计图表`),
			Action:  `dashboard`,
			Target:  `_blank`,
		},
		{
			Display: false,
			Name:    echo.T(`重启服务端`),
			Action:  `server_restart`,
		},
		{
			Display: false,
			Name:    echo.T(`关闭服务端`),
			Action:  `server_stop`,
		},
		{
			Display: false,
			Name:    echo.T(`重启客户端`),
			Action:  `client_restart`,
		},
		{
			Display: false,
			Name:    echo.T(`关闭客户端`),
			Action:  `client_stop`,
		},
		{
			Display: false,
			Name:    echo.T(`配置表单`),
			Action:  `addon_form`,
		},
	},
}
