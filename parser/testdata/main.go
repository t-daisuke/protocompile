package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
)

func main() {
	// ファイルを読み込む
	data, err := ioutil.ReadFile("./largeproto.proto")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// パーサーハンドラを作成
	handler := reporter.NewHandler(nil)

	// ASTに変換
	fileNode, err := parser.Parse("largeproto.proto", bytes.NewReader(data), handler)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// ASTをダンプ
	_, err = parser.ResultFromAST(fileNode, true, handler) //
	if err != nil {
		fmt.Println("Error converting AST to result:", err)
		return
	}

	fmt.Println("AST has been dumped to hoge.proto")
}
