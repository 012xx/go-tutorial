package main

import (
	"fmt"
	"log"
	"net/http"
)

// フォームのリクエストを処理するハンドラ関数
func formHandler(w http.ResponseWriter, r *http.Request) {
	// フォームデータの解析
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "エラー: %v", err)
		return
	}
	fmt.Fprintf(w, "POSTリクエストが成功しました\n")
	// フォームの値を取得
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "名前 = %s\n", name)
	fmt.Fprintf(w, "住所 = %s\n", address)
}

// "/hello" パスに対する GET リクエストを処理するハンドラ関数
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// URLパスが "/hello" でない場合は 404 エラーを返す
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	// メソッドが GET でない場合はエラーを返す
	if r.Method != "GET" {
		http.Error(w, "メソッドがサポートされていません", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "こんにちは！")
}

func main() {
	// "static" ディレクトリ内のファイルを提供するためのハンドラを作成
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	// "/form" パスに対するリクエストを formHandler ハンドラで処理
	http.HandleFunc("/form", formHandler)

	// "/hello" パスに対するリクエストを helloHandler ハンドラで処理
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("サーバーをポート 8080 で起動中\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
