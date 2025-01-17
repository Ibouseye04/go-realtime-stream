package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
    "your-project/internal/binance"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins for development
    },
}

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    // WebSocket endpoint
    http.HandleFunc("/ws/prices", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Printf("Failed to upgrade connection: %v", err)
            return
        }
        defer conn.Close()

        binance.StreamBinancePrices(conn)
    })

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}