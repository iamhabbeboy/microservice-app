package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Models struct {
	logFileName string
	LogEntry    LogEntry
}

func NewModel() Models {
	return Models{
		logFileName: "logger.log",
		LogEntry:    LogEntry{},
	}
}

type LogEntry struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *Models) Create(entry LogEntry) error {
	f, err := os.OpenFile(m.logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	j, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	_, err = f.Write(j)
	if err != nil {
		return err
	}
	return nil
}

func (m *Models) Read() string {
	data, err := ioutil.ReadFile(m.logFileName)
	if err != nil {
		return err.Error()
	}

	return string(data)
}
