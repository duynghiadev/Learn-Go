# WebSocket Client Implementation

## Overview

This file implements the WebSocket client functionality for the real-time chat application. It handles individual client connections, message processing, and connection maintenance.

## Structure

### Client Type

```go
type Client struct {
    id   string           // Unique identifier for each client
    hub  *Hub            // Reference to the central hub
    conn *websocket.Conn // WebSocket connection
    send chan []byte     // Buffered channel for outbound messages
}
```

### Constants

- writeWait : 10 seconds - Maximum time allowed to write a message
- pongWait : 60 seconds - Maximum time to wait for pong response
- pingPeriod : 54 seconds - Interval between ping messages
- maxMessageSize : 512 bytes - Maximum allowed message size

## Main Functions

### serveWs

- Handles new WebSocket connection requests
- Creates new client instances
- Initiates read and write pumps
- Registers client with the hub

### writePump

- Manages writing messages to the WebSocket connection
- Handles:
  - Message broadcasting
  - Ping/pong for connection maintenance
  - Connection closure
  - Message queuing

### readPump

- Manages reading messages from the WebSocket connection
- Features:
  - Message size limiting
  - Connection timeout handling
  - JSON message decoding
  - Error handling
  - Client unregistration on disconnect

## Message Flow

1. Client connects via WebSocket
2. Unique UUID assigned to client
3. Client registered with hub
4. Two goroutines started:
   - writePump: Hub → Client
   - readPump: Client → Hub
5. Messages processed and broadcast through hub

## Error Handling

- Connection errors logged
- Unexpected closure handling
- JSON decode error management
- Automatic client cleanup on disconnect

## Connection Maintenance

- Ping/pong mechanism for connection health
- Automatic timeout handling
- Resource cleanup on connection close

## Usage Example

```go
hub := NewHub()
http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    serveWs(hub, w, r)
})
```
