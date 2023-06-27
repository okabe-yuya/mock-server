package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "os"
)

const SETTING_JSON_PATH = "./server_setting.json"

type ServerSetting struct {
    Port int `json:"port"`
    Endpoints[]EndPoint `json:"endpoints"`
}

type EndPoint struct {
    Path string `json:"path"`
    Status int `json:"status"`
}

func main() {
    var s ServerSetting
    txt, err := ioutil.ReadFile(SETTING_JSON_PATH)
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }
    json.Unmarshal([]byte(txt), &s)

    for _, ep := range s.Endpoints {
        http.HandleFunc(ep.Path, func(w http.ResponseWriter, r *http.Request) {
            if ep.Status >= 100 && ep.Status < 600 {
                w.WriteHeader(ep.Status)
            } else {
                log.Fatalf("The specified %d is not a valid status.", ep.Status)
                os.Exit(1)
            }
        })
    }

    log.Printf("listening server on port %d", s.Port)
    port := fmt.Sprintf(":%d", s.Port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}
