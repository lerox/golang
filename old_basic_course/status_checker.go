package old_basic_course

import (
    "time"
    "net/http"
    "fmt"
)

func statusChecker() {
    links := []string {
        "http://google.com",
        "http://facebook.com",
        "http://stackoverflow.com",
        "http://golang.org",
        "http://amazon.com",
    }

    c := make(chan string)

    for _, link := range links {
        go get(link, c)
    }

    // for {
    //     go get(<-c, c)
    // }

    for l := range c {
        go func (link string) { // literal func
            time.Sleep(time.Second * 5)
            get(link, c)
        }(l)
    }
}

func get(link string, c chan string) {
    fmt.Println(link)

    _, err := http.Get(link)

    if err != nil {
        fmt.Println(link, "might be down!")
        c <- link
        return
    }

    fmt.Println(link, "is up!")

    c <- link
}
