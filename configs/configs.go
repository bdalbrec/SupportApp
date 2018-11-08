package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/bdalbrec/sll"
)


type Config struct {
	SQLServer   string `json:"server"`
	SQLPort     string `json:"port"`
	SQLUsername     string `json:"username"`
	SQLPassword string `json:"password"`
	SQLDBName string `json:"database"`
	Logname string `json:"logname"`
}

// initialize a variable to unmarshall into
var Configs Config


func ReadConfigs() {
	logfile := "logs/applicationStartupLog.log"

	sll.LogInfo("Application startup: Reading configs.", logfile)

	// open JSON file
	jf, err := os.Open("Config.json")
	if err != nil {
		fmt.Printf("Failed to open config file. %v/n", err)
		sll.LogError("Error in ReadConfigs. Failed to open config file.", logfile, err)
	}

	// defer closing of JSON file so that we can read it
	defer jf.Close()

	// initialize a variable to unmarshall into
	//var Configs Config

	// read the JSON file into a slice of bytes
	bs, err := ioutil.ReadAll(jf)
	if err != nil {
		fmt.Printf("Could not read config file. %v/n", err)
		sll.LogError("Error in ReadConfigs. Unable to read config file.", logfile, err)
	}

	//unmarshal the byte slice into the con variable of type Conn
	err = json.Unmarshal(bs, &Configs)
	if err != nil {
		fmt.Printf("Could not unmarshal JSON. %v", err)
		sll.LogError("Error in ReadConfigs. Unable unmarshal JSON.", logfile, err)
	}

	// Boom! We've unmarshalled the JSON into our program.
	sll.LogInfo("Successfully read configs from file.", logfile)

}