package binance

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

func StreamBinancePrices(clientConn *websocket.Conn) {
    log.Println("Attempting to connect to Binance...")
    
    // Add required headers
    headers := http.Header{}
    headers.Add("User-Agent", "Mozilla/5.0")
    headers.Add("Origin", "https://stream.binance.com")

    // Try the public testnet endpoint first
    binanceConn, resp, err := websocket.DefaultDialer.Dial(
        "wss://testnet.binance.vision/ws/btcusdt@trade",
        headers,
    )
    if err != nil {
        log.Printf("Failed to connect to Binance: %v", err)
        if resp != nil {
            log.Printf("Response status: %s", resp.Status)
            log.Printf("Response headers: %v", resp.Header)
        }
        return
    }
    defer binanceConn.Close()
    
    log.Println("Connected to Binance successfully!")

    for {
        _, message, err := binanceConn.ReadMessage()
        if err != nil {
            log.Printf("Error reading from Binance: %v", err)
            break
        }

        log.Printf("Received message from Binance: %s", string(message))

        // Forward the message to the client
        if err := clientConn.WriteMessage(websocket.TextMessage, message); err != nil {
            log.Printf("Error writing to client: %v", err)
            break
        }
    }
}