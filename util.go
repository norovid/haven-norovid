package main

import (
	"fmt"
	"io"
	"encoding/json"
	"os"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func nPubToPubkey(nPub string) string {
	_, v, err := nip19.Decode(nPub)
	if err != nil {
		panic(err)
	}
	return v.(string)
}

type WriteWhitelist struct {
	Pubkeys []string `json:"pubkeys"`
}

func loadWriteWhitelist() (*WriteWhitelist, error) {
	// Try opening "whitelist.json" first
	file, err := os.Open("whitelist.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Fallback to "write_whitelist.json" if "whitelist.json" does not exist
			file, err = os.Open("write_whitelist.json")
			if err != nil {
				return nil, fmt.Errorf("could not open file: %w", err)
			}
		} else {
			return nil, fmt.Errorf("could not open file: %w", err)
		}
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	var writeWhitelist WriteWhitelist
	if err := json.Unmarshal(bytes, &writeWhitelist); err != nil {
		return nil, fmt.Errorf("could not parse JSON: %w", err)
	}

	return &writeWhitelist, nil
}

type ReadWhitelist struct {
	Pubkeys []string `json:"pubkeys"`
}

func loadReadWhitelist(filename string) (*ReadWhitelist, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	var readWhitelist ReadWhitelist
	if err := json.Unmarshal(bytes, &readWhitelist); err != nil {
		return nil, fmt.Errorf("could not parse JSON: %w", err)
	}

	return &readWhitelist, nil
}