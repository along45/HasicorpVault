func main() {
    apiKey := "1234567"
    key :="JuDmGuXNc5ZghYnpMNn+oajg1xZu0ddNA57QwE2i0AaLcMhTRQISzZD3mUyPNQNe1TQPW7fM5xIFNDOU7sPHFA=="

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
