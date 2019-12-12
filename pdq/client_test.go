package pdq

import (
	"log"
	"os"
	"testing"
)

var c *SimpleClient

func TestMain(m *testing.M) {
	mondayAPIToken, ok := os.LookupEnv("MONDAY_API_TOKEN")
	if !ok {
		log.Println("could not get monday api token from env")
		return
	}
	c = NewSimpleClient(mondayAPIToken)
	m.Run()
}
