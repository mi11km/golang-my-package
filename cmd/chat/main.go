package main

import (
	"app/pkg/chat"
	"app/pkg/chat/middleware"
	"app/pkg/chat/templates"
	"app/pkg/chat/trace"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8080", "appのアドレス")
	flag.Parse()

	r := chat.NewRoom()
	r.Tracer = trace.New(os.Stdout)

	http.Handle("/chat", middleware.MustAuth(&templates.TemplateHandler{Filename: "chat.html"}))
	http.Handle("/login", &templates.TemplateHandler{Filename: "login.html"})
	http.HandleFunc("/auth/", middleware.LoginHandler)
	http.Handle("/room", r)

	go r.Run()

	log.Println("webサーバーを開始。ポート：", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
