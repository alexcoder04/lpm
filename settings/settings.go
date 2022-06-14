package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/alexcoder04/arrowprint"
)

type Config struct {
	FirstRun  bool `json:"firstRun"`
	DebugMode bool `json:"debugMode"`
}

const DEFAULT_REPO = "https://raw.githubusercontent.com/alexcoder04/LeoConsole-repo-main/main/index.json"

var Folders map[string]string = make(map[string]string)
var CurrentConfig Config = Config{}

func UpdateConfig() error {
	arrowprint.Suc0("loading config")
	configFile := path.Join(Folders["config"], "config.json")
	cont, err := ioutil.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(configFile)
			if err != nil {
				return err
			}
			defer f.Close()
			_, err = f.Write([]byte(`{"firstRun":false}`))
			if err != nil {
				return err
			}
		}
	}
	err = json.Unmarshal(cont, &CurrentConfig)
	if err != nil {
		return err
	}
	return nil
}
