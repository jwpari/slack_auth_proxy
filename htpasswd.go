package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
)

// lookup passwords in a htpasswd file
// The entries must have been created with -s for SHA encryption
type HtEntry struct {
	Username string
	Name     string
	Hash     string
}
type HtpasswdFile struct {
	Users map[string]HtEntry
}

func NewHtpasswdFromFile(path string) (*HtpasswdFile, error) {
	log.Printf("using htpasswd file %s", path)
	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewHtpasswd(r)
}

func NewHtpasswd(file io.Reader) (*HtpasswdFile, error) {
	csv_reader := csv.NewReader(file)
	csv_reader.Comma = ':'
	csv_reader.Comment = '#'
	csv_reader.TrimLeadingSpace = true

	records, err := csv_reader.ReadAll()
	if err != nil {
		return nil, err
	}
	h := &HtpasswdFile{Users: make(map[string]HtEntry)}
	for _, record := range records {
		h.Users[record[0]] = HtEntry{Username: record[0], Hash: record[1], Name: record[2]}
	}
	return h, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *HtpasswdFile) Validate(username string, password string) bool {
	user, exists := h.Users[username]
	if !exists {
		return false
	}
	return checkPasswordHash(password, user.Hash)
}

func (h *HtpasswdFile) Fetch(username string) (*HtEntry, error) {
	user, exists := h.Users[username]
	if !exists {
		return nil, fmt.Errorf("Could not find username: %s", username)
	}
	return &user, nil
}
