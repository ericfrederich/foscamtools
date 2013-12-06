package main

import (
    "flag"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "time"
)

func watcher(host, user, password string, interval time.Duration) {
    url := "http://" + host + "/snapshot.cgi"
    log.Printf("watching(\"%s\", \"%s\", \"%s\")\n", url, user, password)
    i := 0
    for _ = range time.Tick(interval) {
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            panic(err)
        }
        req.SetBasicAuth(user, password)

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            panic(err)
        }
        if resp.StatusCode != 200 {
            val, exists := resp.Header["Content-Type"]
            if exists && val[0] == "text/html" {
                b, err := ioutil.ReadAll(resp.Body)
                if err != nil {
                    panic(err)
                }
                log.Fatal(string(b))
            }
            log.Println("not a 200 response")
        }

        fname := fmt.Sprintf("out_%05d.jpg", i)
        f, err := os.Create(fname)
        if err != nil {
            panic(err)
        }

        written, err := io.Copy(f, resp.Body)
        if err != nil {
            panic(err)
        }
        log.Printf("Wrote %s (%d bytes)\n", fname, written)
        i += 1
    }
}

func main() {
    host := flag.String("h", "ipcam_000000000000_0", "host")
    user := flag.String("u", "admin", "user name")
    password := flag.String("p", "admin", "password")
    interval := flag.Duration("i", 24*time.Second, "interval")
    flag.Parse()
    watcher(*host, *user, *password, *interval)
}
