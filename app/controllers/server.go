package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/taki-21/go-todo-app/config"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)

	// 第二引数をnilにした場合のデフォルトのマルチプレクサは登録されていないURLにアクセスしたら、デフォルトで 404 Page Not Found を返す
	// 第二引数がnilだから DefaultServeMux が使用される
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
