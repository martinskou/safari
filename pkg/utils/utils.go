package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/* Logs function execution time, use defer */
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

/* Pretty-print anything which can be marshalled */
func Pprint(v interface{}) {
	s, e := json.MarshalIndent(v, "", "  ")
	if e != nil {
		fmt.Printf("Pprint error : %s\n", e)
	} else {
		fmt.Printf("%s\n", s)
	}
}

/* Return true if path exists */
func ExistsPath(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

/* Generate a simple password */
func SimplePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghjkmnpqrstuvwxyz23456789ABCDEFGHJKMNPQRSTUVWXYZ")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
