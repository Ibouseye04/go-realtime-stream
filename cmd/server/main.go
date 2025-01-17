package main

import (
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Videos []Video
}

type Video struct {
    ID          string
    Title       string
    Description string
}

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Main page handler
    http.HandleFunc("/", handleHome)

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Videos: []Video{
            {
                ID:          "EzGPmg4fFL8",
                Title:       "Epic Snowboarding",
                Description: "GoPro: Best of 2023",
            },
            {
                ID:          "SCuY6osbOTs",
                Title:       "Mountain Biking",
                Description: "Trail Riding Highlights",
            },
        },
    }

    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}