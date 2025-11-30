package users

import (
	"crypto/rand"
	"encoding/hex"
)

var sessions = map[string]int{} // sessionID -> userID

func generateSessionID() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func CreateSession(userID int) (string, error) {
	id, err := generateSessionID()
	if err != nil {
		return "", err
	}
	sessions[id] = userID
	return id, nil
}

func GetUserIDBySession(sessionID string) (int, bool) {
	id, ok := sessions[sessionID]
	return id, ok
}
