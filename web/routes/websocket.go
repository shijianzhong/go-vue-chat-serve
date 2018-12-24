package routes

import (
	"fmt"
	"github.com/kataras/iris/websocket"
)
func  HandleConnection(c websocket.Connection)  {
		c.On("chat", func(msg string) {
			// Print the message to the console, c.Context() is the iris's http context.
			fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
			// Write message back to the client message owner with:
			// c.Emit("chat", msg)
			// Write message to all except this client with:
			c.To(websocket.Broadcast).Emit("chat", msg)
		})
}
