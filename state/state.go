package state

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type GoLoggerState struct {
	GologgerHome       string
	GologgerConfigFile string
	ActiveSession      string
}

type Conf struct {
	ActiveSession string `yaml: "activesession"`
}

var gologgerHomeName = ".gologger"
var gologgerConfigFileName = "gologger-config.yml"
var gologgerHomeDir = filepath.Join(getHomeDir(), gologgerHomeName)
var configFilePath = filepath.Join(gologgerHomeDir, gologgerConfigFileName)

var default_config_value_active_session = "default-session"

var Glog = newGoLoggerState()

func UpdateConfigActiveSession(sessionName string) {
	configContent := getConfigContent(gologgerHomeDir)

	configContent.ActiveSession = strings.ToLower(sessionName)

	writeConfig(configContent)
}

func newGoLoggerState() *GoLoggerState {
	configContent := getConfigContent(gologgerHomeDir)

	writeConfig(configContent)

	return &GoLoggerState{GologgerHome: gologgerHomeDir, GologgerConfigFile: configFilePath, ActiveSession: configContent.ActiveSession}
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return homeDir
}

func getConfigContent(gologgerHome string) *Conf {
	if _, err := os.Stat(gologgerHome); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(gologgerHome, 0755)
		if mkdirErr != nil {
			fmt.Println("Could not create gologger home directory!")
			log.Fatal(mkdirErr)
		}
	}

	if _, err := os.Stat(filepath.Join(gologgerHome, "sessions/")); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(filepath.Join(gologgerHome, "sessions/"), 0755)
		if mkdirErr != nil {
			fmt.Println("Could not create gologger sessions directory!")
			log.Fatal(mkdirErr)
		}
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		createNewConfigFile()
	}

	configFileData, readConfErr := os.ReadFile(configFilePath)
	if readConfErr != nil {
		fmt.Printf("Failed to read Gologger config file at: %s\n", configFilePath)
		log.Fatal(readConfErr)
	}

	configContent := &Conf{}
	unmarshalErr := yaml.Unmarshal(configFileData, configContent)
	if unmarshalErr != nil {
		fmt.Printf("Possibly malformed YAML in config file at: %s\n", configFilePath)
		log.Fatal(readConfErr)
	}

	// Ensure active session exists
	if _, err := os.Stat(filepath.Join(gologgerHome, "sessions/", configContent.ActiveSession)); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(filepath.Join(gologgerHome, "sessions/", configContent.ActiveSession), 0755)
		if mkdirErr != nil {
			fmt.Println("Could not create gologger sessions active session directory!")
			log.Fatal(mkdirErr)
		}
	}
	if _, err := os.Stat(filepath.Join(gologgerHome, "sessions/", configContent.ActiveSession, "data/")); os.IsNotExist(err) {
		mkdirErr := os.Mkdir(filepath.Join(gologgerHome, "sessions/", configContent.ActiveSession, "data/"), 0755)
		if mkdirErr != nil {
			fmt.Println("Could not create gologger sessions data directory!")
			log.Fatal(mkdirErr)
		}
	}

	return configContent
}

func createNewConfigFile() {
	fmt.Println("Gologger config file not found, doing first time setup...")

	config := &Conf{ActiveSession: default_config_value_active_session}
	writeConfig(config)

	fmt.Printf("Created Gologger config at: %s\n", configFilePath)
}

func writeConfig(config *Conf) {
	data, yamlMarshalErr := yaml.Marshal(config)
	if yamlMarshalErr != nil {
		log.Fatal(yamlMarshalErr)
	}

	f, fileCreateErr := os.Create(configFilePath)
	if fileCreateErr != nil {
		log.Fatal(fileCreateErr)
	}

	defer f.Close()

	_, fileWriteErr := f.Write(data)
	if fileWriteErr != nil {
		log.Fatal(fileWriteErr)
	}
}
