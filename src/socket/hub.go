package socket

type Hub struct {
	conns map[*Conn]bool
}

func NewHub() *Hub {
	return &Hub{}
}

func (h *Hub) Register(conn *Conn) error {
	_, exists := h.conns[conn]
	if !exists {
		h.conns[conn] = true
	}
	return nil
}

func (h Hub) Broadcast(conn *Conn, msg []byte) error {
	for c, _ := range h.conns {
		if c != conn {
			err := c.Write(msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (h Hub) BroadcastString(conn *Conn, msg string) error {
	return h.Broadcast(conn, []byte(msg))
}
