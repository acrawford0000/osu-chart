package app

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"project/backend/go-osu/client"
	"project/backend/go-osu/client/opts"
	"project/backend/go-osu/enum"
	"project/backend/go-osu/model"
	"strconv"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	NewClientOnStartup()
}

// Setting up an api client ID and Secret from the front end
func (a *App) SetOsuAuthCredentials(id string, secret string) error {
	err := saveOsuAuthCredentials(id, secret)
	if err != nil {
		return err
	}
	err = CreateNewClient()
	if err != nil {
		return err
	}
	return nil
}

// Display whether or not client ID and Secret exist
func (a *App) AreOsuAuthCredentialsSet() bool {
	id, secret := getOsuAuthCredentials()
	return id != 0 && secret != ""
}

// Check for credentials on startup and if they exist, create a new client
func NewClientOnStartup() {

	// Check credentials from JSON file
	id, secret := getOsuAuthCredentials()

	// Create client if credentials set
	if id != 0 && secret != "" {
		CreateNewClient()
	}

}

// Initialize and create a new client
// Get auth credentials for the client
var Client *client.OsuClient

func CreateNewClient() error {
	client, err := client.NewClient(getOsuAuthCredentials())
	if err != nil {
		return errors.New("failed to create a new client. Please verify your credentials are correct")
	}
	Client = client
	log.Printf("Client set: %v", Client)
	return nil
}

type osuAuthCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

const osuAuthCredentialsFile = "osu_auth_credentials.json"

func saveOsuAuthCredentials(clientID string, clientSecret string) error {
	data := osuAuthCredentials{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(osuAuthCredentialsFile, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// getOsuAuthCredentials returns client ID and client secret
func getOsuAuthCredentials() (int, string) {
	jsonData, err := os.ReadFile(osuAuthCredentialsFile)
	if err != nil {
		log.Printf("Failed to read osu auth credentials file: %v", err)
		return 0, ""
	}
	var data osuAuthCredentials
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("Failed to unmarshal osu auth credentials: %v", err)
		return 0, ""
	}
	clientID, err := strconv.Atoi(data.ClientID)
	if err != nil {
		log.Printf("Failed to convert client ID to int: %v", err)
		return 0, ""
	}
	return clientID, data.ClientSecret
}

// Set default Game Mode and then function to set it
var gameMode enum.GameMode = enum.GameModeStandard

func (a *App) SetGameMode(mode string) {

	switch mode {
	case "standard":
		gameMode = enum.GameModeStandard
	case "mania":
		gameMode = enum.GameModeMania
	case "taiko":
		gameMode = enum.GameModeTaiko
	case "catch":
		gameMode = enum.GameModeCtb

	}

}

// Main function to get user profile details and return them to the frontend
func (a *App) GetUser(username string) (user *model.User, err error) {
	log.Printf("GetUser: Client = %v", Client)

	key := enum.UserKeyUsername
	opts := opts.GetUserOpts{
		Username: username,
		Mode:     &gameMode,
		Key:      &key,
	}

	user, err = Client.GetUser(opts)

	if err != nil {
		return nil, err
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("user.json", userJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return user, err
}
