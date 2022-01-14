package data

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/anthonydroberts/gologger/state"
	"github.com/anthonydroberts/gologger/terminal"
	"gopkg.in/yaml.v2"
)

type Entry struct {
	Command      string `yaml: "command"`
	TimeStamp    string `yaml: "timestamp"`
	ExitCode     int    `yaml: "exitcode"`
	DataFileName string `yaml: "datafilename"`
}

func CreateEntry(commandRan string, data []byte, exitCode int, startTime time.Time) {
	// Remove illegal filename characters
	re := regexp.MustCompile(`[\.\ \\/:"*?<>|]+`)
	commandRanFiltered := re.ReplaceAllString(commandRan, "_")

	timeStampFN := startTime.Format("2006-01-02_15-04-05.000000")
	entryName := timeStampFN + "-" + commandRanFiltered

	dataFileName := entryName + ".log"

	entry := &Entry{Command: commandRan, TimeStamp: startTime.Format("2006-01-02 15:04:05"), ExitCode: exitCode, DataFileName: dataFileName}
	entryFilePath := filepath.Join(state.Glog.GologgerHome, "sessions/", state.Glog.ActiveSession, entryName+".yml")
	entryData, yamlMarshalErr := yaml.Marshal(entry)
	if yamlMarshalErr != nil {
		log.Fatal(yamlMarshalErr)
	}

	writeFile(entryData, entryFilePath)
	writeFile(data, filepath.Join(state.Glog.GologgerHome, "sessions/", state.Glog.ActiveSession, "data/", entryName+".log"))

	terminal.Msg("print", fmt.Sprintf("Created entry '%s' in session '%s'", entryName, state.Glog.ActiveSession))
}

func GetEntries() []*Entry {
	entryYamlPaths := GetSessionEntryPaths(state.Glog.ActiveSession)

	var entryList []*Entry
	for _, entryYamlFile := range entryYamlPaths {
		entryList = append(entryList, GetEntry(entryYamlFile))
	}

	return entryList
}

func GetEntry(entryPath string) *Entry {
	entryFileData, readConfErr := os.ReadFile(entryPath)
	if readConfErr != nil {
		fmt.Printf("Failed to read entry YAML file at: %s\n", entryPath)
		log.Fatal(readConfErr)
	}

	entryContent := &Entry{}
	unmarshalErr := yaml.Unmarshal(entryFileData, entryContent)
	if unmarshalErr != nil {
		fmt.Printf("Possibly malformed YAML in file at: %s\n", entryPath)
		log.Fatal(readConfErr)
	}

	return entryContent
}

func OpenEntryData(entry *Entry, editor string) {
	if editor == "terminal-output" {
		data, err := os.ReadFile(filepath.Join(state.Glog.GologgerHome, "sessions/", state.Glog.ActiveSession, "data/", entry.DataFileName))
		if err != nil {
			log.Fatalf("Error reading file %s\n", err)
		}

		fmt.Printf("\n%s\n", data)
	} else {
		cdErr := os.Chdir(filepath.Join(state.Glog.GologgerHome, "sessions/", state.Glog.ActiveSession, "data/"))
		if cdErr != nil {
			log.Fatalf("Failed to change working directory: %s\n", cdErr)
		}

		cmd := exec.Command(editor, filepath.Base(entry.DataFileName))
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		runErr := cmd.Run()
		if runErr != nil {
			terminal.Msg("fail", fmt.Sprintf("%s returned error: %s", editor, runErr))
		}
	}
}

func DeleteEntry(entry *Entry) {
	entryName := strings.TrimSuffix(filepath.Base(entry.DataFileName), filepath.Ext(entry.DataFileName)) + ".yml"

	err1 := os.Remove(filepath.Join(state.Glog.GologgerHome, "sessions/", state.Glog.ActiveSession, "data/", entry.DataFileName))
	if err1 != nil {
		log.Fatalf("Failed to remove file: %s\n", err1)
	}
	err2 := os.Remove(filepath.Join(state.Glog.GologgerHome, "sessions/", state.Glog.ActiveSession, entryName))
	if err2 != nil {
		log.Fatalf("Failed to remove file: %s\n", err2)
	}
}

func writeFile(data []byte, filePath string) {
	f, fileCreateErr := os.Create(filePath)
	if fileCreateErr != nil {
		log.Fatal(fileCreateErr)
	}

	defer f.Close()

	_, fileWriteErr := f.Write(data)
	if fileWriteErr != nil {
		log.Fatal(fileWriteErr)
	}
}
