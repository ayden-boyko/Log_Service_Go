package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ayden-boyko/Log_Service_Go/pkg"
)

const test_file = "./test.json"

type testLog struct {
	Time  string `json:"time"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

type testLogs []testLog

var testBeaver *pkg.Beaver

func init() {
	testBeaver, _ = pkg.NewBeaver("file", "info", test_file)
	// clear test.json file
	os.Truncate(test_file, 0)
}

// test beaver creation
func TestBeaverCreated(t *testing.T) {
	if testBeaver == nil {
		t.Errorf("Beaver not created")
	}
}

func TestBeaverConfigured(t *testing.T) {
	if testBeaver.GetLevel() != "info" {
		t.Errorf("Beaver not configured, level")
	}
	if testBeaver.GetFilePath() != test_file {
		t.Errorf("Beaver not configured, filepath")
	}
}

func TestBeaverCreatedWithYAML(t *testing.T) {
	testYAMLPath := "test_yaml.yaml"
	tbeaver, err := pkg.NewBeaverFromFile(testYAMLPath)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
}

func TestBeaverCreatedWithJSON(t *testing.T) {
	testJSONPath := "test_json.json"
	tbeaver, err := pkg.NewBeaverFromFile(testJSONPath)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
}

// test info log
func TestBeaverInfo(t *testing.T) {
	tbeaver, err := pkg.NewBeaver("file", "info", test_file)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
	tbeaver.Info("TestBeaverInfo")

	file, err := os.OpenFile(test_file, os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	var logs testLogs

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var log testLog
		if err := decoder.Decode(&log); err != nil {
			t.Errorf("Error decoding log: %v", err)
		}
		logs = append(logs, log)
	}

	if len(logs) != 1 {
		t.Errorf("Incorrect number of logs: %v", len(logs))
	}
	if logs[0].Level != "INFO" {
		t.Errorf("Incorrect log level: %v", logs[0].Level)
	}
	if logs[0].Msg != "TestBeaverInfo" {
		t.Errorf("Incorrect log message: %v", logs[0].Msg)
	}

}

// test warn log
func TestBeaverWarn(t *testing.T) {
	tbeaver, err := pkg.NewBeaver("file", "warn", test_file)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
	tbeaver.Warn("TestBeaverWarn")

	file, err := os.OpenFile(test_file, os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	var logs testLogs

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var log testLog
		if err := decoder.Decode(&log); err != nil {
			t.Errorf("Error decoding log: %v", err)
		}
		logs = append(logs, log)
	}

	if len(logs) != 2 {
		t.Errorf("Incorrect number of logs: %v", len(logs))
	}
	if logs[1].Level != "WARN" {
		t.Errorf("Incorrect log level: %v", logs[0].Level)
	}
	if logs[1].Msg != "TestBeaverWarn" {
		t.Errorf("Incorrect log message: %v", logs[0].Msg)
	}
}

// test error log
func TestBeaverError(t *testing.T) {
	tbeaver, err := pkg.NewBeaver("file", "error", test_file)
	if err != nil {
		t.Errorf("Error creating beaver: %v", err)
	}
	if tbeaver == nil {
		t.Errorf("Beaver not created")
	}
	tbeaver.Error("TestBeaverError")

	file, err := os.OpenFile(test_file, os.O_RDONLY, 0644)
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	var logs testLogs

	decoder := json.NewDecoder(file)
	for decoder.More() {
		var log testLog
		if err := decoder.Decode(&log); err != nil {
			t.Errorf("Error decoding log: %v", err)
		}
		logs = append(logs, log)
	}

	if len(logs) != 3 {
		t.Errorf("Incorrect number of logs: %v", len(logs))
	}
	if logs[2].Level != "ERROR" {
		t.Errorf("Incorrect log level: %v", logs[0].Level)
	}
	if logs[2].Msg != "TestBeaverError" {
		t.Errorf("Incorrect log message: %v", logs[0].Msg)
	}
}
