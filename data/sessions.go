package data

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/anthonydroberts/gologger/state"
)

func CreateSession(sessionName string) {
	sName := strings.ToLower(sessionName) + "/"

	errMkSession := os.Mkdir(filepath.Join(state.Glog.GologgerHome, "sessions/", sName), 0755)
	errMkData := os.Mkdir(filepath.Join(state.Glog.GologgerHome, "sessions/", sName, "data/"), 0755)
	if errMkSession != nil || errMkData != nil {
		log.Fatal("Failed to create a Gologger session")
	}
}

func SessionExists(sessionName string) bool {
	if _, err := os.Stat(filepath.Join(state.Glog.GologgerHome, "sessions/", sessionName)); os.IsNotExist(err) {
		return false
	}

	return true
}

func GetSessions() []string {
	items, err := os.ReadDir(filepath.Join(state.Glog.GologgerHome, "sessions/"))
	if err != nil {
		log.Fatal("Failed to read sessions directory")
	}

	var sessions []string
	for _, i := range items {
		if i.IsDir() {
			sessions = append(sessions, i.Name())
		}
	}

	return sessions
}

func GetSessionLastModifiedTime(sessionName string) time.Time {
	sessionDir, err := os.Open(filepath.Join(state.Glog.GologgerHome, "sessions/", sessionName))
	if err != nil {
		log.Fatal(err)
	}

	stat, err := sessionDir.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return stat.ModTime()
}

func GetSessionEntryPaths(sessionName string) []string {
	items, err := os.ReadDir(filepath.Join(state.Glog.GologgerHome, "sessions/", sessionName))
	if err != nil {
		log.Fatal("Failed to read sessions directory")
	}

	var entries []string
	for _, i := range items {
		if !i.IsDir() {
			entries = append(entries, filepath.Join(filepath.Join(state.Glog.GologgerHome, "sessions/", sessionName), i.Name()))
		}
	}

	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i] > entries[j]
	})

	return entries
}

func UpdateActiveSession(sessionName string) {
	state.UpdateConfigActiveSession(sessionName)
}

func DeleteSession(sessionName string) {
	err := os.RemoveAll(filepath.Join(state.Glog.GologgerHome, "sessions/", sessionName))
	if err != nil {
		log.Fatalf("Failed to remove session %s: %s\n", sessionName, err)
	}
}

func DeleteEntries(sessionName string) {
	DeleteSession(sessionName)
	CreateSession(sessionName)
}
