package ws

import(
	"encoding/json"
	"sync"
 	"github.com/gorilla/websocket"
)

// Message from frontend
type WSMessage struct {
	Type     string `json:"type"`      // "chat", "call_offer", "call_answer"
	TargetID uint   `json:"target_id"` // Кому отправить
	Content  string `json:"content"`   // Зашифрованная "абракадабра" от фронта
	SenderID uint   `json:"sender_id"` // Кто отправил (проставим на бэке)
}

type Hub struct{
	Clients map[uint]*websocket.Conn
	mu sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[uint]*websocket.Conn),
	}
}

func (h *Hub) Register(userID uint, conn *websocket.Conn){
	h.mu.Lock()
	h.Clients[userID] = conn
	h.mu.Unlock()

	// Start listener user
	go h.listen(userID, conn)
}

func (h *Hub) Unregister(userID uint){
	h.mu.Lock()
	if conn, ok := h.Clients[userID]; ok {
		conn.Close()
		delete(h.Clients, userID)
	}
	h.mu.Lock()
}

// Reader for everyone user
func(h *Hub) listen(userID uint, conn *websocket.Conn){
	defer h.Unregister(userID)

	for{
		// Read incoming JSON from Frontend
		_, payload, err := conn.ReadMessage()
		if err != nil {
			break // If user is exit
		}
		// Unpacked JSON
		var msg WSMessage
		if err := json.Unmarshal(payload, &msg); err != nil {
			continue // If garbage - ignore
		}
		msg.SenderID = userID // 

		h.routeMessage(msg)
	}
}

func (h *Hub) routeMessage(msg WSMessage){
	h.mu.RLock()
	targetConn, ok := h.Clients[msg.TargetID]
	h.mu.Unlock()
	data, _ := json.Marshal(msg)
	if ok {
		targetConn.WriteMessage(websocket.TextMessage, data)
	} else{
		// TO DO Offline user - add save in BD for Repository
	}
}

