// The uuid package generates a random formatted uuid string.
// mostly copied from http://www.ashishbanerjee.com/home/go/go-generate-uuid
package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// Generates a UUID string like that:
// B2E2D4C4-B24C-4BC3-8966-D28A22BEDA48
func GenUUID() string {
	uuid := make([]byte, 24)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		log.Fatal("can't generate uuid")
		return ""
	}
	// TODO: verify the two lines implement RFC 4122 correctly
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7

	str := hex.EncodeToString(uuid)
	return fmt.Sprintf("%s-%s-%s-%s-%s", str[:8], str[9:13], str[14:18], str[19:23], str[29:])
}
