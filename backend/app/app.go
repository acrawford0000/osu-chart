package app

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"project/backend/api/apierror"
	"project/backend/api/client"
	"project/backend/api/client/opts"
	"project/backend/api/enum"
	"project/backend/api/model"
	"project/backend/api/osutrack"
	"strconv"
	"strings"
	"sync"
)

type userClient interface {
	GetUser(opts.GetUserOpts) (*model.User, error)
}

type App struct {
	ctx           context.Context
	mu            sync.RWMutex
	credentialsMu sync.Mutex
	client        userClient
	// Explicit paths and factories allow isolated tests without real keys or APIs.
	credentialsPath string
	legacyPath      string
	newClient       func(int, string) (userClient, error)
	loadHistory     func(int, int) ([]*model.UserStats, error)
}

func NewApp() *App {
	return &App{
		legacyPath:  "osu_auth_credentials.json",
		newClient:   func(id int, secret string) (userClient, error) { return client.NewClient(id, secret) },
		loadHistory: osutrack.GetStatsUpdates,
	}
}
func (a *App) Startup(ctx context.Context) { a.ctx = ctx }

type osuAuthCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (a *App) configPath() (string, error) {
	if a.credentialsPath != "" {
		return a.credentialsPath, nil
	}
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", apierror.New("CONFIG_ERROR", "Could not locate the application settings folder.")
	}
	return filepath.Join(configDir, "osu-chart", "osu_auth_credentials.json"), nil
}

func validateCredentials(data osuAuthCredentials) (int, error) {
	id, err := strconv.Atoi(strings.TrimSpace(data.ClientID))
	if err != nil || id <= 0 || strings.TrimSpace(data.ClientSecret) == "" {
		return 0, apierror.New("AUTH_INVALID", "Enter a positive numeric Client ID and a non-empty Client Secret.")
	}
	return id, nil
}

// Caller holds credentialsMu. Migrate the existing file without deleting the user's copy.
func (a *App) readCredentials() (osuAuthCredentials, error) {
	path, err := a.configPath()
	if err != nil {
		return osuAuthCredentials{}, err
	}
	raw, err := os.ReadFile(path)
	migrate := false
	if errors.Is(err, os.ErrNotExist) && a.legacyPath != "" {
		raw, err = os.ReadFile(a.legacyPath)
		migrate = err == nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return osuAuthCredentials{}, nil
	}
	if err != nil {
		return osuAuthCredentials{}, apierror.New("CONFIG_ERROR", "Could not read your saved credentials. Check access to the application settings folder.")
	}
	var data osuAuthCredentials
	if json.Unmarshal(raw, &data) != nil {
		return data, apierror.New("CONFIG_ERROR", "The saved credentials file is not valid JSON. Save your credentials again in Settings.")
	}
	if _, err := validateCredentials(data); err != nil {
		return data, err
	}
	if migrate {
		if err := writeCredentials(path, data); err != nil {
			return data, err
		}
	}
	return data, nil
}

// Write a temporary sibling then rename it; a failed validation/write leaves the old file intact.
func writeCredentials(path string, data osuAuthCredentials) error {
	fail := func() error {
		return apierror.New("CONFIG_ERROR", "Could not save your credentials. Your previous settings have been kept.")
	}
	if os.MkdirAll(filepath.Dir(path), 0700) != nil {
		return fail()
	}
	raw, err := json.Marshal(data)
	if err != nil {
		return fail()
	}
	file, err := os.CreateTemp(filepath.Dir(path), ".credentials-*")
	if err != nil {
		return fail()
	}
	defer os.Remove(file.Name())
	if err = file.Chmod(0600); err == nil {
		_, err = file.Write(raw)
	}
	if err == nil {
		err = file.Sync()
	}
	closeErr := file.Close()
	if err != nil || closeErr != nil {
		return fail()
	}
	if os.Rename(file.Name(), path) != nil {
		return fail()
	}
	return nil
}

func (a *App) GetCredentials() (osuAuthCredentials, error) {
	a.credentialsMu.Lock()
	defer a.credentialsMu.Unlock()
	return a.readCredentials()
}
func (a *App) AreOsuAuthCredentialsSet() bool {
	data, err := a.GetCredentials()
	if err != nil {
		return false
	}
	_, err = validateCredentials(data)
	return err == nil
}
func (a *App) IsClientValid() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.client != nil
}
func (a *App) CreateNewClient() error {
	a.credentialsMu.Lock()
	defer a.credentialsMu.Unlock()
	data, err := a.readCredentials()
	if err != nil {
		return err
	}
	id, err := validateCredentials(data)
	if err != nil {
		return err
	}
	candidate, err := a.newClient(id, data.ClientSecret)
	if err != nil {
		return err
	}
	a.mu.Lock()
	a.client = candidate
	a.mu.Unlock()
	return nil
}
func (a *App) SaveOsuAuthCredentials(id, secret string) error {
	a.credentialsMu.Lock()
	defer a.credentialsMu.Unlock()
	data := osuAuthCredentials{ClientID: strings.TrimSpace(id), ClientSecret: strings.TrimSpace(secret)}
	numericID, err := validateCredentials(data)
	if err != nil {
		return err
	}
	candidate, err := a.newClient(numericID, data.ClientSecret)
	if err != nil {
		return err
	}
	path, err := a.configPath()
	if err != nil {
		return err
	}
	if err = writeCredentials(path, data); err != nil {
		return err
	}
	a.mu.Lock()
	a.client = candidate
	a.mu.Unlock()
	return nil
}

func parseMode(mode string) (enum.GameMode, int, error) {
	switch mode {
	case "standard":
		return enum.GameModeStandard, 0, nil
	case "taiko":
		return enum.GameModeTaiko, 1, nil
	case "catch":
		return enum.GameModeCtb, 2, nil
	case "mania":
		return enum.GameModeMania, 3, nil
	default:
		return "", 0, apierror.New("INVALID_MODE", "Choose a supported game mode.")
	}
}

func (a *App) GetUser(username, mode string) (*model.MyUserStruct, error) {
	gameMode, _, err := parseMode(mode)
	if err != nil {
		return nil, err
	}
	username = strings.TrimSpace(username)
	if username == "" {
		return nil, apierror.New("INVALID_USERNAME", "Enter an osu! username first.")
	}
	a.mu.RLock()
	currentClient := a.client
	a.mu.RUnlock()
	if currentClient == nil {
		return nil, apierror.New("AUTH_REQUIRED", "Set and verify your API credentials in Settings before adding players.")
	}
	key := enum.UserKeyUsername
	user, err := currentClient.GetUser(opts.GetUserOpts{Username: username, Mode: &gameMode, Key: &key})
	if err != nil {
		return nil, err
	}
	if user == nil || user.ID == 0 {
		return nil, apierror.New("NOT_FOUND", "No matching osu! player was found.")
	}
	return client.GetMyUserData(user), nil
}
func (a *App) GetStatsUpdates(id int, mode string) ([]*model.UserStats, error) {
	_, trackMode, err := parseMode(mode)
	if err != nil {
		return nil, err
	}
	if id <= 0 {
		return nil, apierror.New("INVALID_PLAYER", "This player does not have a valid osu! ID.")
	}
	return a.loadHistory(id, trackMode)
}
