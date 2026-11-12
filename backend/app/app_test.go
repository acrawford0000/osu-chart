package app

import (
	"errors"
	"os"
	"path/filepath"
	"project/backend/api/client/opts"
	"project/backend/api/model"
	"strings"
	"sync"
	"testing"
)

type fakeClient struct{}

func (*fakeClient) GetUser(o opts.GetUserOpts) (*model.User, error) {
	return &model.User{UserCompact: model.UserCompact{ID: 1, Username: o.Username}}, nil
}
func testApp(t *testing.T) *App {
	t.Helper()
	a := NewApp()
	a.credentialsPath = filepath.Join(t.TempDir(), "config", "credentials.json")
	a.legacyPath = ""
	a.newClient = func(int, string) (userClient, error) { return &fakeClient{}, nil }
	return a
}
func TestCredentialSaveKeepsWorkingFileAndClientOnFailure(t *testing.T) {
	a := testApp(t)
	if err := a.SaveOsuAuthCredentials("1", "old-test-secret"); err != nil {
		t.Fatal(err)
	}
	original, _ := os.ReadFile(a.credentialsPath)
	originalClient := a.client
	a.newClient = func(int, string) (userClient, error) {
		return nil, errors.New("AUTH_INVALID: Rejected test credentials")
	}
	if err := a.SaveOsuAuthCredentials("2", "invalid-test-secret"); err == nil {
		t.Fatal("invalid credentials were accepted")
	}
	after, _ := os.ReadFile(a.credentialsPath)
	if string(after) != string(original) || a.client != originalClient {
		t.Fatal("failed save changed working credentials")
	}
	a.newClient = func(int, string) (userClient, error) { return &fakeClient{}, nil }
	if err := a.SaveOsuAuthCredentials("3", "new-test-secret"); err != nil {
		t.Fatal(err)
	}
	saved, err := a.GetCredentials()
	if err != nil || saved.ClientID != "3" {
		t.Fatal("valid replacement did not persist")
	}
}
func TestCredentialWriteFailureIsRecoverable(t *testing.T) {
	a := testApp(t)
	a.client = &fakeClient{}
	blocked := filepath.Join(t.TempDir(), "blocked")
	if err := os.WriteFile(blocked, []byte("file"), 0600); err != nil {
		t.Fatal(err)
	}
	a.credentialsPath = filepath.Join(blocked, "credentials.json")
	if err := a.SaveOsuAuthCredentials("1", "test-secret"); err == nil || !strings.Contains(err.Error(), "CONFIG_ERROR") {
		t.Fatalf("expected recoverable config error, got %v", err)
	}
	if !a.IsClientValid() {
		t.Fatal("failed write removed working client")
	}
}
func TestLegacyCredentialsMigrateOutsideWorkingDirectory(t *testing.T) {
	a := testApp(t)
	a.legacyPath = filepath.Join(t.TempDir(), "legacy.json")
	if err := os.WriteFile(a.legacyPath, []byte(`{"client_id":"42","client_secret":"test-only"}`), 0600); err != nil {
		t.Fatal(err)
	}
	data, err := a.GetCredentials()
	if err != nil || data.ClientID != "42" {
		t.Fatalf("migration failed: %v", err)
	}
	if _, err := os.Stat(a.credentialsPath); err != nil {
		t.Fatal("new config file missing")
	}
	if _, err := os.Stat(a.legacyPath); err != nil {
		t.Fatal("original copy was deleted")
	}
	// Once migrated, a changed legacy file cannot overwrite the saved configuration.
	os.WriteFile(a.legacyPath, []byte(`{"client_id":"99","client_secret":"other"}`), 0600)
	data, err = a.GetCredentials()
	if err != nil || data.ClientID != "42" {
		t.Fatal("legacy data overwrote migrated configuration")
	}
}
func TestMissingAndMalformedCredentials(t *testing.T) {
	a := testApp(t)
	data, err := a.GetCredentials()
	if err != nil || data.ClientID != "" || a.AreOsuAuthCredentialsSet() {
		t.Fatal("missing file should be a normal unconfigured state")
	}
	os.MkdirAll(filepath.Dir(a.credentialsPath), 0700)
	os.WriteFile(a.credentialsPath, []byte(`not json`), 0600)
	if _, err := a.GetCredentials(); err == nil || !strings.Contains(err.Error(), "CONFIG_ERROR") {
		t.Fatal("malformed file was silently ignored")
	}
}
func TestConcurrentRequestsKeepTheirOwnMode(t *testing.T) {
	a := testApp(t)
	var wg sync.WaitGroup
	var mu sync.Mutex
	observed := map[int]int{}
	a.loadHistory = func(id, mode int) ([]*model.UserStats, error) {
		mu.Lock()
		observed[id] = mode
		mu.Unlock()
		return []*model.UserStats{}, nil
	}
	for id, mode := range map[int]string{1: "standard", 2: "mania", 3: "taiko", 4: "catch"} {
		wg.Add(1)
		go func(id int, mode string) {
			defer wg.Done()
			if _, err := a.GetStatsUpdates(id, mode); err != nil {
				t.Error(err)
			}
		}(id, mode)
	}
	wg.Wait()
	for id, want := range map[int]int{1: 0, 2: 3, 3: 1, 4: 2} {
		if observed[id] != want {
			t.Fatalf("player %d got wrong mode", id)
		}
	}
	if _, err := a.GetUser("demo", "standard"); err == nil || !strings.Contains(err.Error(), "AUTH_REQUIRED") {
		t.Fatal("missing client was not handled")
	}
	if _, err := a.GetStatsUpdates(1, "invalid"); err == nil {
		t.Fatal("invalid mode accepted")
	}
}

func TestConcurrentProfilesKeepTheirOwnMode(t *testing.T) {
	a := testApp(t)
	ready := make(chan struct{}, 4)
	release := make(chan struct{})
	expected := map[string]string{"one": "osu", "two": "mania", "three": "taiko", "four": "fruits"}
	a.client = userClientFunc(func(o opts.GetUserOpts) (*model.User, error) {
		ready <- struct{}{}
		<-release
		if string(*o.Mode) != expected[o.Username] {
			t.Errorf("wrong mode for %s", o.Username)
		}
		return &model.User{UserCompact: model.UserCompact{ID: 1, Username: o.Username}}, nil
	})
	var wg sync.WaitGroup
	for name, mode := range map[string]string{"one": "standard", "two": "mania", "three": "taiko", "four": "catch"} {
		wg.Add(1)
		go func(name, mode string) {
			defer wg.Done()
			if _, err := a.GetUser(name, mode); err != nil {
				t.Error(err)
			}
		}(name, mode)
	}
	for range 4 {
		<-ready
	}
	close(release)
	wg.Wait()
}

type userClientFunc func(opts.GetUserOpts) (*model.User, error)

func (f userClientFunc) GetUser(o opts.GetUserOpts) (*model.User, error) { return f(o) }
