# WallpaperEngineConfigStruct
Structure of Wallpaper Engine's config file.

# Installation
```
go get github.com/yxl-prz/WallpaperEngineConfigStruct
```
# Possible Usage Cases
## Manual
```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"github.com/yxl-prz/WallpaperEngineConfigStruct"
)

type Config struct {
	InstallDirectory string                      `json:"?installdirectory"`
	MyComputerUser   *WallpaperEngine.UserConfig `json:"MyComputerUser"`
}

func main() {
	var config_path string = "/path/to/config.json"
	dat, err := ioutil.ReadFile(config_path)
	if err != nil {
		// Handle Error
	}

	var WallpaperEngineConfig *Config

	err = json.Unmarshal(dat, &WallpaperEngineConfig)
	if err != nil {
		// Handle Error
	}

	WallpaperEngineConfig.MyComputerUser.General.Browser.ViewIconSize = WallpaperEngine.ViewIconSize_Medium

	dat, err = json.MarshalIndent(WallpaperEngineConfig, "", "\t")
	if err != nil {
		// Handle Error
	}

	err = ioutil.WriteFile(config_path, dat, 0777)
	if err != nil {
		// Handle Error
	}
}

```
## Integrated
```go
import "github.com/yxl-prz/WallpaperEngineConfigStruct"

type Config struct {
	InstallDirectory string                      `json:"?installdirectory"`
	MyComputerUser   *WallpaperEngine.UserConfig `json:"MyComputerUser"`
}

func main() {
	var config_path string = "/path/to/config.json"

	var WallpaperEngineConfig *Config

	err := WallpaperEngine.Read(config_path, &WallpaperEngineConfig)
	if err != nil {
		// Handle Error
	}
    
	WallpaperEngineConfig.MyComputerUser.General.Browser.ViewIconSize = WallpaperEngine.ViewIconSize_Medium

	err = WallpaperEngine.Save(path_to_config, WallpaperEngineConfig)

	if err != nil {
		// Handle Error
	}
}
```