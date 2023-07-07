package controllers

import (
	"net/http"

	"github.com/taki-21/go-todo-app/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)

	// 第二引数をnilにした場合のデフォルトのマルチプレクサは登録されていないURLにアクセスしたら、デフォルトで 404 Page Not Found を返す
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
