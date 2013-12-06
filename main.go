package main

import (
	"io"
	"net/http"
	"os"
    "time"
    "log"
    "flag"
    "fmt"
)

func watcher(host, user, password string){
    url := "http://" + host + "/snapshot.cgi"
    log.Printf("watching(\"%s\", \"%s\", \"%s\")\n", url, user, password)
    for i:= 0 ; ; i++ {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            panic(err)
        }
        req.SetBasicAuth(user, password)

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            panic(err)
        }
        if resp.StatusCode != 200{
            val, exists := resp.Header["Content-Type"]
            if exists && val[0] == "text/html"{
                // call log.Fatal with contents of resp.Body
            }
            log.Println("not a 200 response")
        }

        f, err := os.Create(fmt.Sprintf("out_%05d.jpg", i))
        if err != nil {
            panic(err)
        }

        written, err := io.Copy(f, resp.Body)
        if err != nil {
            panic(err)
        }
        log.Println("Wrote", written, "bytes")
        time.Sleep(3 * time.Second)
    }
}

func main() {
    host     := flag.String("h", "ipcam_000000000000_0", "host"     )
    user     := flag.String("u", "admin"               , "user name")
    password := flag.String("p", "admin"               , "password" )
    flag.Parse()
    watcher(*host, *user, *password)
}
