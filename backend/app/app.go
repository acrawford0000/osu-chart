package app

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"project/backend/api/client"
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
	"project/backend/api/osutrack"
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
	// NewClientOnStartup()
}

// osu!api v2 setup and functions
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

func (a *App) CreateNewClient() error {
	client, err := client.NewClient(getOsuAuthCredentials())
	if err != nil {
		return errors.New("failed to create a new client. Please verify your credentials are correct")

	}
	Client = client
	log.Printf("Client set: %v", Client)
	return nil
}

// IsClientValid checks if the osu! API client was successfully created
func (a *App) IsClientValid() bool {
	return Client != nil
}

type osuAuthCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

const osuAuthCredentialsFile = "osu_auth_credentials.json"

// Display whether or not client ID and Secret exist
func (a *App) AreOsuAuthCredentialsSet() bool {
	id, secret := getOsuAuthCredentials()
	if id != 0 && secret != "" {
		return true
	}
	return false
}

func (a *App) GetCredentials() (osuAuthCredentials, error) {
	jsonData, err := os.ReadFile(osuAuthCredentialsFile)
	if err != nil {
		log.Printf("Failed to read osu auth credentials file: %v", err)
		return osuAuthCredentials{}, err
	}
	var data osuAuthCredentials
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("Failed to unmarshal osu auth credentials: %v", err)
		return osuAuthCredentials{}, err
	}
	return data, nil
}

func (a *App) SaveOsuAuthCredentials(id string, secret string) error {
	data := osuAuthCredentials{
		ClientID:     id,
		ClientSecret: secret,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(osuAuthCredentialsFile, jsonData, 0644)
	if err != nil {
		return err
	}

	// Create a new client with the saved credentials
	err = CreateNewClient()
	if err != nil {
		// Return an error if the client could not be created
		return errors.New("failed to create a new client with the saved credentials: %w")
	}

	return nil
}

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

var userData *model.User

// Main function to get user profile details and return them to the frontend
func (a *App) GetUser(username string) (user *model.MyUserStruct, err error) {

	key := enum.UserKeyUsername
	opts := opts.GetUserOpts{
		Username: username,
		Mode:     &gameMode,
		Key:      &key,
	}

	userData, err = Client.GetUser(opts)
	if err != nil {
		return nil, err
	}

	user = client.GetMyUserData(userData)

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

// Set default Game Mode and then function to set it
var gameMode enum.GameMode = enum.GameModeStandard

func (a *App) SetGameMode(mode string) {

	switch mode {
	case "standard":
		gameMode = enum.GameModeStandard
		trackmode = 0
	case "mania":
		gameMode = enum.GameModeMania
		trackmode = 3
	case "taiko":
		gameMode = enum.GameModeTaiko
		trackmode = 1
	case "catch":
		gameMode = enum.GameModeCtb
		trackmode = 2

	}
}

// osutrack setup and functions
var trackmode int = 0

func (a *App) GetStatsUpdates(id int) (stats []*model.UserStats, err error) {
	statsUpdates, err := osutrack.GetStatsUpdates(id, trackmode)
	if err != nil {
		log.Print("Failed to get user stats.", err)
		return nil, err
	}

	userstatsJSON, err := json.Marshal(statsUpdates)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("userstats.json", userstatsJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return statsUpdates, err

}

func (a *App) SavePlayerData(playerData string) error {
	// Write player data to a file
	err := os.WriteFile("players.json", []byte(playerData), 0644)
	if err != nil {
		return err
	}
	return nil
}
