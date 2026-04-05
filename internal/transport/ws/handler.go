package ws

import(
	"net/http"
	"github.com/gorilla/websocket"
)
// Settings Door
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	//////////////////////////////////////////////////////////
	// ATTENTION!!! Need permissions react-app for connection 
	//////////////////////////////////////////////////////////
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Hub) WSHandler(w http.ResponseWriter, r *http.Request, userID uint){
	// Upgrade connection from http to WebSocket
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		return
	}
	h.Register(userID,conn)
}