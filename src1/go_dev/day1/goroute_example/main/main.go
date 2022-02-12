package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

func main() {
	var pipe chan int
	pipe = make(chan int, 1)
	go goroute.Add(100, 300, pipe)
	sum := <-pipe
	fmt.Println("sum=", sum)
}

/*
https://golang.org/可在在线编辑go代码
https://tour.go-zh.org/flowcontrol/1可在在线编辑go代码
https://go-zh.org/pkg/->Packages->ctrl+f搜索time模块->Index->func Now() Time->func (t Time) Unix() int64
shift+Alt+↓复制当前行到下一行
G:\project>go  build  -o  bin/goroute_test.exe  go_dev/day1/goroute_example/main

G:\project>dir  bin

G:\project\bin 的目录
2020/06/24  22:20    <DIR>          .
2020/06/24  22:20    <DIR>          ..
2020/06/23  23:35         2,077,696 calc
2020/06/23  23:39         2,077,696 calC_test.exe
2020/06/24  22:20         2,078,720 goroute_test.exe
               3 个文件      6,234,112 字节
			   2 个目录 19,996,143,616 可用字节

F1->Preferences:Open Settings(JSON)
settings.json
{
  "explorer.confirmDelete": false,
  "workbench.tree.renderIndentGuides": "none",
  "explorer.compactFolders": false,
  "search.followSymlinks": false,
  "git.autorefresh": false,
  "files.autoSave": "off",
  "git.confirmSync": false,
  "explorer.confirmDragAndDrop": false,
  "editor.formatOnType": false,
  "editor.formatOnSave": true,
  "go.useCodeSnippetsOnFunctionSuggest": true,
  "[go]": {
    "editor.insertSpaces": false,
    "editor.formatOnSave": false,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }
}
*/
