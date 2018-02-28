package ethrpc

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	err := godotenv.Load()
	if err != nil && os.Getenv("GO_ENV") != "test" {
		log.Fatal("Error loading .env file")
	}
}

func TestRPCClient(t *testing.T) {
	Convey("Success", t, func() {
		c := NewRPCClient(os.Getenv("GETH_ENDPOINT"))
		_, err := c.GetBlockNumber()
		So(err, ShouldBeNil)
	})
}

func TestGetBlockNumber(t *testing.T) {
	Convey("Success", t, func() {
		c := NewRPCClient(os.Getenv("GETH_ENDPOINT"))
		_, err := c.GetBlockNumber()
		So(err, ShouldBeNil)
	})
}
