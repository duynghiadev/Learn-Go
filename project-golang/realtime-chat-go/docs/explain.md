# Realtime Chat Go - Code Documentation

## Project Structure Overview

### 1. Main Application

The entry point of the application that:

- Initializes the hub
- Sets up HTTP routes
- Handles index page serving
- Starts the WebSocket server

Key components:

```go
func main() {
    hub := NewHub()
    go hub.run()
    http.HandleFunc("/", serveIndex)
    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(hub, w, r)
    })
}
```

### 2. Client Management (client.go)

Handles individual WebSocket connections with features:
Key Structures:

```go
type Client struct {
    id   string           // Unique client identifier
    hub  *Hub            // Reference to central hub
    conn *websocket.Conn // WebSocket connection
    send chan []byte     // Message outbound channel
}
```

Important Functions:

- `serveWs` : Establishes WebSocket connections
- `writePump` : Handles message writing
- `readPump` : Manages message reading

### 3. Message Structure (message.go)

Defines the message format:

```go
type Message struct {
    ClientID string `json:"clientID"`
    Text     string `json:"text"`
}
```

### 4. Frontend Templates Index Page ( `index.html` )

Main chat interface featuring:

- HTMX WebSocket integration
- Message input form
- Chat display area
- Tailwind CSS styling
  Key elements:

```html
<div hx-ext="ws" ws-connect="/ws">
  <div class="flex bg-gray-100 p-4">
    <ul id="chat_room" hx-swap="beforeend" hx-swap-oob="beforeend"></ul>
  </div>
  <form id="form" ws-send hx-reset-on-success>
    <!-- Message input -->
  </form>
</div>
```

Message Template ( `message.html` )
Individual message display template:

```html
<div id="chat_room" hx-swap-oob="beforeend">
  <li id="message" class="flex my-2">
    <h1 class="text-base font-bold mr-2 text-red-500">{{ .ClientID }}:</h1>
    <p class="text-base">{{ .Text }}</p>
  </li>
</div>
```

## Technical Details

### WebSocket Configuration

```go
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
```

### Connection Constants

```go
const (
    writeWait = 10 * time.Second
    pongWait = 60 * time.Second
    pingPeriod = (pongWait * 9) / 10
    maxMessageSize = 512
)
```

### Message Flow

1. Client connects via WebSocket endpoint
2. Connection upgraded to WebSocket
3. Client assigned unique UUID
4. Client registered with hub
5. Two goroutines handle communication:
   - writePump: Hub → Client communication
   - readPump: Client → Hub communication
6. Messages broadcast through hub to all clients

### Error Handling

- Connection errors logged
- Unexpected closures managed
- JSON decode errors handled
- Automatic client cleanup on disconnect

### Connection Maintenance

- Ping/pong mechanism
- Automatic timeouts
- Resource cleanup
- Message buffering

## Dependencies

- gorilla/websocket: WebSocket implementation
- google/uuid: Unique client identification
- HTMX: Frontend WebSocket integration
- Tailwind CSS: Styling

## Security Considerations

- Message size limits
- Connection timeouts
- Error logging
- Clean connection handling

This documentation provides a comprehensive overview of the codebase structure and functionality. Each component is designed to work together to create a robust real-time chat system.
