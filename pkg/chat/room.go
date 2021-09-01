package chat

import (
	"app/pkg/chat/middleware"
	"app/pkg/chat/trace"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan *message
	join    chan *client
	leave   chan *client
	clients map[*client]bool
	Tracer  trace.Tracer // 操作ログ
}

func NewRoom() *room {
	return &room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		Tracer:  trace.Off(),
	}
}

func (r *room) Run() {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			r.Tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.Tracer.Trace("クライアントが退出しました")
		case msg := <-r.forward:
			r.Tracer.Trace("メッセージを受信しました：", msg.Message)
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージ送信
					r.Tracer.Trace("-- クライアントに送信されました")
				default:
					// 送信に失敗
					delete(r.clients, client)
					close(client.send)
					r.Tracer.Trace("-- 送信に失敗しました。クライアントをクリーンアップします")
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan *message, messageBufferSize),
		room:   r,
	}
	userData := middleware.DecodeUserInfo(req)
	if userData != nil {
		client.userData = userData
	}

	r.join <- client
	defer func() { r.leave <- client }()
	go client.write()
	client.read()
}
