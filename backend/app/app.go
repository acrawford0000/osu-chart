package app

import (
	"context"
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
func (a *App) SetOsuAuthCredentials(id string, secret string) {

	os.Setenv("OSU_CLIENT_ID", id)
	os.Setenv("OSU_CLIENT_SECRET", secret)
}

// Display whether or not client ID and Secret exist
func (a *App) AreOsuAuthCredentialsSet() bool {
	return os.Getenv("OSU_CLIENT_ID") != "" && os.Getenv("OSU_CLIENT_SECRET") != ""
}

// Check for credentials on startup and if they exist, create a new client
func NewClientOnStartup() {

	// Check env vars
	id := os.Getenv("OSU_CLIENT_ID")
	secret := os.Getenv("OSU_CLIENT_SECRET")

	// Create client if credentials set
	if id != "" && secret != "" {
		CreateNewClient()
	}

}

// Initialize and create a new client
// Get auth credentials for the client
var Client *client.OsuClient

func CreateNewClient() {
	Client, _ = client.NewClient(getOsuAuthCredentials())
}

// getOsuAuthCredentials returns client ID and client secret
func getOsuAuthCredentials() (int, string) {
	clientSecret := os.Getenv("OSU_CLIENT_SECRET")
	clientID, err := strconv.Atoi(os.Getenv("OSU_CLIENT_ID"))
	if err != nil {
		log.Fatalln(err)
	}

	if clientID == 0 || clientSecret == "" {
		log.Fatalln("client credentials not set in environment")
	}

	return clientID, clientSecret
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

	key := enum.UserKeyUsername
	opts := opts.GetUserOpts{
		Username: username,
		Mode:     &gameMode,
		Key:      &key,
	}

	user, err = Client.GetUser(opts)

	return user, err
}
