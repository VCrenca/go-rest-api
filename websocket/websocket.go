package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Hub -
type Hub struct {
	clients   map[*websocket.Conn]bool
	broadcast chan *socketMsg
}

type socketMsg struct {
	content string
	msgType int
}

// NewHub -
func NewHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan (*socketMsg)),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Run -
func (h *Hub) Run() {
	for {
		select {
		case newMsg := <-h.broadcast:
			for conn := range h.clients {
				conn.WriteMessage(newMsg.msgType, []byte(newMsg.content))
			}
		}
	}
}

// Handler -
func (h *Hub) Handler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("upgrade : %s", err.Error())
		return
	}

	// Add the new connection to the hub
	h.clients[ws] = true

	log.Println(h.clients)

	for {
		t, msg, err := ws.ReadMessage()
		if err != nil {
			delete(h.clients, ws)
			log.Printf("reading message : %s", err.Error())
			break
		}

		h.broadcast <- &socketMsg{
			content: string(msg),
			msgType: t,
		}
	}
}
