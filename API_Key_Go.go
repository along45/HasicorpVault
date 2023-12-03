func main() {
    apiKey := "1234567"

    // Make an API call using the hardcoded API key
    response, err := http.Get("https://api.example.com/v1/data?apiKey=" + apiKey)
    if err != nil {
        fmt.Println("Error making API call:", err)
        return
    }

    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println("Error reading API response:", err)
        return
    }

    fmt.Println("API response:", string(body))
}
