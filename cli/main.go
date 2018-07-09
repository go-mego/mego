package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-mego/jennifer/jen"

	"gopkg.in/urfave/cli.v1"
)

var (
	ErrMegoExists = errors.New("mego: Mego 資訊檔案已經存在，無法再次建立新服務")
)

const (
	MegoFile = "./.mego.yml"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "以現在的目錄建立一個新的 Mego 服務",
			Action: func(c *cli.Context) error {
				if _, err := os.Stat(MegoFile); os.IsNotExist(err) {
					content := []byte(fmt.Sprintf("name: %s\nversion: 1.0.0", c.Args().Get(0)))
					return ioutil.WriteFile(MegoFile, content, 0644)
				}
				return ErrMegoExists
			},
		},
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "依照目錄中現有的 Mego 資訊建置相關的程式碼",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "產生指定程式碼並且更新 Mego 資訊檔案",
			Subcommands: []cli.Command{
				{
					Name:  "controller",
					Usage: "追加一個新的控制器",
					Action: func(c *cli.Context) error {
						f := jen.NewFile("main")
						f.Func().Id("main").Params().Block(
							jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
						)
						fmt.Printf("%#v", f)

						return nil
					},
				},
				{
					Name:  "model",
					Usage: "追加一個新的模型資料",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name:  "routes",
			Usage: "輸出此 Mego 服務所暴露的路徑與控制器方法資訊",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "transport",
			Aliases: []string{"t"},
			Usage:   "建立一個新的 HTTP 或 RPC 轉繼到內部方法",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "command",
			Aliases: []string{"c"},
			Usage:   "建立自訂的指令快捷功能",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "依照現有的程式碼更新 Mego 資訊，或者是依照變動後的 Mego 資訊更新本地程式碼",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "建置 Mego 服務為二進制檔案供正式使用",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "資料庫輔助功能",
			Subcommands: []cli.Command{
				{
					Name:  "create",
					Usage: "依照 Mego 資訊或者最後一個遷移資料來初始化資料庫結構",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "drop",
					Usage: "移除並且摧毀資料庫",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:    "migrate",
					Aliases: []string{"m"},
					Usage:   "依照目前的資料庫結構建立新的遷移紀錄資料",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "rollback",
					Usage: "回溯到上一個或是指定版本的資料庫結構",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "status",
					Usage: "顯示目前資料庫的遷移狀態",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "seed",
					Usage: "將預設的資料插入至資料庫",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "version",
					Usage: "顯示目前資料庫的版本",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "直接編譯並執行 Mego 服務",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "執行並監聽 Mego 服務的所有檔案，當檔案異動時自動重啟服務",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
