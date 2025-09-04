package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	root_cmd = &cobra.Command{
		Use:   "game-xjson",
		Short: "game-xjson是一款由浊水楼台团队开源的，帮助游戏开发者快速将xlsx数据转为json数据的工具",
		Long: `Ciallo～(∠・ω< )⌒☆

game-xjson是一款由浊水楼台团队开源的，帮助游戏开发者快速将xlsx数据转为json数据的工具，目前仅支持xlsx文件转json文件

我是开发者陈汐胤，很高兴能和大家在这里见面，这是本项目的开源地址（https://github.com/ZSLTChenXiYin/game-xjson）

关注浊水楼台喵，关注浊水楼台谢谢喵`,
	}
)

func init() {
	Version = fmt.Sprintf("%s %s/%s", Version, runtime.GOOS, runtime.GOARCH)

	root_cmd.Version = Version
}
