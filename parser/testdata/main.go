package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	// prototext.Marshal をインポート

	"github.com/bufbuild/protocompile/parser"
	"github.com/bufbuild/protocompile/reporter"
	"google.golang.org/protobuf/proto"
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
	result, err := parser.ResultFromAST(fileNode, true, handler)
	if err != nil {
		fmt.Println("Error converting AST to result:", err)
		return
	}

	// ダンプ結果をテキスト形式に変換
	resultProto := result.FileDescriptorProto()         // 修正: ポインタを取得しない
	resultProtoBytes, err := proto.Marshal(resultProto) // 修正: 変数名を変更
	if err != nil {
		fmt.Println("Error marshaling result to text:", err)
		return
	}

	// ダンプ結果をファイルに書き込む
	err = ioutil.WriteFile("hoge.proto", resultProtoBytes, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("AST has been dumped to hoge.proto")
}
