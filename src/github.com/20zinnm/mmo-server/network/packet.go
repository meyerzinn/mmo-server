package network

import "errors"

var BadProtocolError = errors.New("bad protocol")

func handle(id uint8, data *[]byte, client *Client) {
	switch (id) {
	case 0:
		// handle respawning
		break
	default:
		client.Close(BadProtocolError)
	}
}