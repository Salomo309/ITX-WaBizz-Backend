package handlers

import (
	"net/http"
	"sync"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize: 1024,
		WriteBufferSize: 1024,
	}
	connections = make(map[string]*websocket.Conn)
	connectionsLock sync.Mutex
)

func addConnection(id string, conn *websocket.Conn) {
	connectionsLock.Lock()
	defer connectionsLock.Unlock()
	connections[id] = conn
}

func removeConnection(id string) {
	connectionsLock.Lock()
	defer connectionsLock.Unlock()
	delete(connections, id)
}

func SendMessageToAll(c *gin.Context, message []byte) {
	connectionsLock.Lock()
	defer connectionsLock.Unlock()

	for id, conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			removeConnection(id)
			c.JSON(http.StatusInternalServerError, gin.H{"Message": "Failed to send chat"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Message received and saved successfully"})
}

func HandleNewWebsocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to establish Web Socket"})
		return
	}

	id := uuid.New().String()
	addConnection(id, conn)

	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}