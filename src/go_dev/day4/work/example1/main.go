package main

import (
	"fmt"
)

func multi() {
	for i := 0; i < 9; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", (i + 1), j+1, (i+1)*(j+1))
		}
		fmt.Println()
	}
}

func main() {
	multi()
}

//把终端恢复默认值
// {
// 	"workbench.colorTheme":"Monokai",
// 	"edit.fontSize":16,
// 	"workbench.iconTheme":"vscode-icons",
// 	"window.zoomLevel":0.8,
// 	"files.autoSave":"afterDelay",
// 	"editor.wordWrap":"on"
//  }
