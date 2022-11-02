package WallpaperEngine

import (
	"encoding/json"
	"io/ioutil"
)

func Read(path_to_config string, variable interface{}) error {
	dat, err := ioutil.ReadFile(path_to_config)
	if err != nil {
		return err
	}

	return json.Unmarshal(dat, variable)
}

func Save(path_to_config string, data interface{}) error {
	dat, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path_to_config, dat, 0777)
}
