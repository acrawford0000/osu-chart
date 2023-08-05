package auth

/* import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Config struct {
    ClientID     string
    ClientSecret string
    AccessToken  string
}

type AccessToken struct {
    Token     string
    ExpiresAt time.Time
}

var accessToken AccessToken

func GetAccessToken(clientID string, clientSecret string) error {
    // Set up the request body
    requestBody := url.Values{}
    requestBody.Set("client_id", clientID)
    requestBody.Set("client_secret", clientSecret)
    requestBody.Set("grant_type", "client_credentials")
    requestBody.Set("scope", "public")

    // Make the POST request
    response, err := http.PostForm("https://osu.ppy.sh/oauth/token", requestBody)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    // Parse the response
    var responseBody map[string]interface{}
    err = json.NewDecoder(response.Body).Decode(&responseBody)
    if err != nil {
        return err
    }

    // Extract the access token and its expiration time
    accessToken.Token = responseBody["access_token"].(string)
    expiresIn := responseBody["expires_in"].(float64)
    accessToken.ExpiresAt = time.Now().Add(time.Duration(expiresIn) * time.Second)

    return nil
}

func CheckAccessToken() error {
    // Check if the access token is expired
    if time.Now().After(accessToken.ExpiresAt) {
        // Get a new access token
        err := GetAccessToken(Config.ClientID, Config.ClientSecret)
        if err != nil {
            return err
        }
    }
    return nil
}

func LoadConfig() (Config, error) {
    // Set up an empty Config struct
    var config Config

    // Open the config file
    configFile, err := os.Open("config.json")
    if err != nil {
        return config, err
    }
    defer configFile.Close()

    // Parse the config file
    err = json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        return config, err
    }

    return config, nil
}

func SaveConfig(config Config) error {
    // Create the config file
    configFile, err := os.Create("config.json")
    if err != nil {
        return err
    }
    defer configFile.Close()

    // Write the config data to the file
    err = json.NewEncoder(configFile).Encode(config)
    if err != nil {
        return err
    }

    return nil
}
*/