package chat

import (
	"sync"
	"time"

	"github.com/jellydator/ttlcache/v2"
)

// Hub maintains the set of active clients and send messages to the client
type Hub struct {
	// Registered clients.
	clients     *ttlcache.Cache
	clientsLock sync.RWMutex

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

var hub = newHub()

func newHub() *Hub {
	return &Hub{
		clients:    ttlcache.NewCache(),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() { //nolint
	for {
		select {
		case client := <-h.register:
			if err := h.Register(client.userId, client); err != nil {
				continue
			}
		case client := <-h.unregister:
			if err := h.Unregister(client.userId); err != nil {
				continue
			}
		}
	}
}

func (h *Hub) Register(key string, client *Client) error {
	return h.AddClient(key, client)
}

func (h *Hub) AddClient(key string, client *Client) error {
	h.clientsLock.Lock()
	defer h.clientsLock.Unlock()
	return h.clients.SetWithTTL(key, client, time.Hour)
}

func (h *Hub) Unregister(key string) error {
	return h.DelClient(key)
}

func (h *Hub) DelClient(key string) error {
	h.clientsLock.Lock()
	defer h.clientsLock.Unlock()
	return h.clients.Remove(key)
}
