package main

type observer struct {
	connections map[*connection]bool

	broadcast chan interface{}

	register chan *connection

	unregister chan *connection
}

func newObserver() observer {
	return observer{
		broadcast:   make(chan interface{}),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
	}
}

func (o *observer) start() {
	for {
		select {
		case c := <-o.register:
			o.connections[c] = true
		case c := <-o.unregister:
			if _, ok := o.connections[c]; ok {
				delete(o.connections, c)
				close(c.send)
			}
		case m := <-o.broadcast:
			for c := range o.connections {
				select {
				case c.send <- m:
				default:
					delete(o.connections, c)
					close(c.send)
				}
			}
		}

	}
}
