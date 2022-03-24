package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Article struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    Articles = []Article{
        Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Title: "larning", Desc: "larning golang", Content: "golang api"},
        Article{Title: "myblog open", Desc: "start myblog", Content: "blog write for developer"},
        Article{Title: "timecard", Desc: "start timecard project", Content: "for my company"},
        Article{Title: "beef day", Desc: "delicious", Content: "beef is"},
        Article{Title: "github error", Desc: "why", Content: "cant push"},
        Article{Title: "front is hard", Desc: "vue or react", Content: "learn basic"},
        Article{Title: "learn tailwind", Desc: "funny tailwind", Content: "front book please"},
        Article{Title: "mailsystem", Desc: "learn mailsystem,", Content: "spf DKIM"},
    }
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome browser!")
    fmt.Println("api run success!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/articles", returnAllArticles)
    log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
    handleRequests()
}