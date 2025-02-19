package pkg

import (
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Beaver struct for handling logs
type Beaver struct {
	level    string
	filePath string
	output   string
	file     *os.File
	logger   *slog.Logger
}

// Config struct to load settings from a file
type Config struct {
	Level    string `json:"log_level" yaml:"log_level"`
	Output   string `json:"log_output" yaml:"log_output"`
	FilePath string `json:"log_file" yaml:"log_file"`
}

// NewBeaver creates a new Beaver instance
// takes in log_type, log_level, log_output
func NewBeaver(args ...string) (*Beaver, error) {
	switch args[0] {
	case "file":
		//OtherWise Open the log file for writing
		logFile, err := os.OpenFile(args[2], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}

		// Create a logger that writes to the file
		logger := slog.New(slog.NewJSONHandler(logFile, nil))

		return &Beaver{
			level:    args[1],
			filePath: args[2],
			output:   "file",
			file:     logFile,
			logger:   logger,
		}, nil
	case "remote":
		//TODO
	default: // default is console
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		return &Beaver{
			level:    args[1],
			filePath: "",
			output:   "console",
			file:     nil,
			logger:   logger,
		}, nil
	}
	return nil, errors.New("invalid input, constructor accepts, log_type, log_level, log_output_filepath")
}

// NewBeaverFromFile loads Beaver configuration from a file and initializes the Beaver
func NewBeaverFromFile(filename string) (*Beaver, error) {
	// Read and parse the configuration file
	configFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config Config

	//check if ends in .yaml or .json
	switch filename[len(filename)-4:] {

	case "yaml": // unable to open yaml file
		// if filename is .yaml, dont assume its yaml
		if err := yaml.NewDecoder(configFile).Decode(&config); err != nil {
			return nil, err
		}

	case "json":
		// if filename is .json, assume it's JSON if not yaml
		if err := json.NewDecoder(configFile).Decode(&config); err != nil {
			return nil, err
		}
	}

	// Open the log file for writing
	logFile, err := os.OpenFile(config.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	// Create a logger that writes to the file
	logger := slog.New(slog.NewJSONHandler(logFile, nil))

	return &Beaver{
		level:    config.Level,
		filePath: config.FilePath,
		output:   config.Output,
		file:     logFile,
		logger:   logger,
	}, nil
}

// TODO CONFIGURE OUTPUT based on log_output (remote services)
// Log method writes messages to the log.json file
func (l *Beaver) Log(message string) {
	switch l.output {
	case "file":
		switch l.level {
		case "info":
			l.logger.Info(message)
		case "warn":
			l.logger.Warn(message)
		case "error":
			l.logger.Error(message)
		default:
			l.logger.Info("INVALID LOG LEVEL: " + message)
		}
	case "remote":
		//TODO
	default: // default is console
		switch l.level {
		case "info":
			log.Println("INFO: " + message)
		case "warn":
			log.Println("WARN: " + message)
		case "error":
			log.Println("ERROR: " + message)
		default:
			log.Println("INVALID LOG LEVEL: " + message)
		}
	}
}

func (b *Beaver) Info(message string) {
	b.logger.Info(message)
}

func (b *Beaver) Error(message string) {
	b.logger.Error(message)
}

func (b *Beaver) Warn(message string) {
	b.logger.Warn(message)
}

func (b *Beaver) Close() {
	b.file.Close()
}

func (b *Beaver) GetLevel() string {
	return b.level
}

func (b *Beaver) GetFilePath() string {
	return b.filePath
}

func LoggingMiddleware(beaver *Beaver, next http.Handler) http.Handler {
	defer beaver.Close()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logMessage := "Method: " + r.Method + " | Path: " + r.URL.Path + " | Duration: " + time.Since(start).String()
		beaver.Log(logMessage)
	})
}
