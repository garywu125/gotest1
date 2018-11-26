package main

import (
	"flag"
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// log.SetLevel(log.WarnLevel)
}

func main() {
	name := flag.String("name", "", "name")
	flag.Parse()

	if err := greet(*name); err != nil {
		// fmt.Printf("Failed to greet you: %v", err)
		log.Warn("Failed to greet you:  ", err)

	}

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Warn("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

func greet(name string) error {
	if name == "" {
		return errors.New("no name provided")
	}
	log.Printf("Hello %s! I'm not in the $GOPATH!\n", name)
	return nil
}
