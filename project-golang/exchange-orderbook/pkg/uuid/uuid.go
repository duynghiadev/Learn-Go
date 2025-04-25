package uuid

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"log"
)

type UUID string

func New() UUID {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}

	if err != nil {
		// native error handling
		panic(err)
	}
	h := sha512.New()
	h.Write(buf)
	id := hex.EncodeToString(h.Sum(nil))

	// only take 10 characters for easy readability
	return UUID(id[:10])
}
