package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    // Set up the HTTP client and request.
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://<juniper-device-ip>/support/information/request", nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    req.Header.Set("User-Agent", "Mozilla/5.0")

    // Send the request and read the response body.
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Parse the RSI data.
    data := string(body)
    lines := strings.Split(data, "\n")
    for _, line := range lines {
        if strings.HasPrefix(line, "Hostname:") {
            fmt.Println(strings.TrimSpace(strings.TrimPrefix(line, "Hostname:")))
        }
        if strings.HasPrefix(line, "Model:") {
            fmt.Println(strings.TrimSpace(strings.TrimPrefix(line, "Model:")))
        }
        // Add additional parsing logic as needed for other fields.
    }
}
