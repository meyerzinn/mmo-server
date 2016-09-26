package network

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"errors"
)

var clients = map[string]*Client{}

type Client struct {
	id   uuid.UUID
	conn *websocket.Conn
	err  error
}

func (c *Client) Send(data []byte) {
	c.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (c *Client) Close(err error) {
	if err != nil {
		c.err = err
	} else {
		c.err = errors.New("Closed.")
	}
	delete(clients, c.id.String())
}

func (c *Client) handleConnection() {
	defer c.conn.Close()
	for {
		if c.err != nil {
			break;
		}
		mt, m, err := c.conn.ReadMessage()
		if err != nil {
			c.err = err
			break
		}
		if mt != websocket.BinaryMessage {
			c.Close(errors.New("bad protocol"))
			break
		}
	}
}

func NewClient(conn *websocket.Conn) (uuid.UUID) {
	id := uuid.NewV4()
	client := Client{
		id: id,
		conn: conn,
	}
	clients[id.String()] = &client
	go client.handleConnection();
	return id
}

func GetClient(id uuid.UUID) *Client {
	return clients[id.String()]
}
