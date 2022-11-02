package WallpaperEngine

type User map[string]*UserConfig

type Folder struct {
	Items      map[string]bool `json:"items"`
	Subfolders []*Folder       `json:"subfolders"`
	Title      string          `json:"title"`
	Type       string          `json:"type"`
}

type Playlist struct {
	Items    []string `json:"items"`
	Name     string   `json:"name"`
	Settings struct {
		Delay         int    `json:"delay"`
		Mode          string `json:"mode"`
		Order         string `json:"order"`
		Transition    bool   `json:"transition"`
		UpdateOnPause bool   `json:"updateonpause"`
		VideoSequence bool   `json:"videosequence"`
	} `json:"settings"`
}

type Profile struct {
	Layout             int                 `json:"layout"`
	Name               string              `json:"name"`
	Profile            interface{}         `json:"profile"` // Unknown Type
	SelectedWallpapers map[string]struct { // SelectedWallpapers[display]
		File     string `json:"file"`
		Playlist struct {
			Items    []string `json:"items"`
			Name     string   `json:"name"`
			Settings struct {
				Delay         bool `json:"delay"`
				Mode          bool `json:"mode"`
				Order         bool `json:"order"`
				Transition    bool `json:"transition"`
				UpdateOnPause bool `json:"updateonpause"`
				VideoSequence bool `json:"videosequence"`
			} `json:"settings"`
		} `json:"playlist"`
	} `json:"selectedwallpapers"`
}

type PairedDevice struct {
	Alias         string `json:"alias"`
	DisplayHeight string `json:"displayHeight"`
	DisplayWidth  string `json:"displayWidth"`
	ID            string `json:"id"`
	MpkgSupport   string `json:"mpkgsupport"`
	Name          string `json:"name"`
	PIN           string `json:"pin"`
	Preset        string `json:"preset"`
	PrivateKey    string `json:"privatekey"`
}

type FilterInfo struct {
	CategoryTags struct {
		Preset    bool `json:"Preset"`
		Wallpaper bool `json:"Wallpaper"`
	} `json:"categorytags"`
	Descending bool `json:"descending"`
	RatingTags struct {
		Everyone     bool `json:"Everyone"`
		Mature       bool `json:"Mature"`
		Questionable bool `json:"Questionable"`
	} `json:"ratingtags"`
	ResolutionTags struct {
		RES_1280x720                    bool `json:"1280 x 720"`
		RES_1366x768                    bool `json:"1366 x 768"`
		RES_3840x2160                   bool `json:"3840 x 2160"`
		RES_1920x1080                   bool `json:"1920 x 1080"`
		RES_2560x1440                   bool `json:"2560 x 1440"`
		RES_Dual_3840x1080              bool `json:"Dual 3840 x 1080"`
		RES_Dual_5120x1440              bool `json:"Dual 5120 x 1440"`
		RES_Dual_7680x2160              bool `json:"Dual 7680 x 2160"`
		RES_DualStandardDefinition      bool `json:"Dual Standard Definition"`
		RES_DynamicResolution           bool `json:"Dynamic resolution"`
		RES_OtherResolution             bool `json:"Other resolution"`
		RES_Portrait1080x1920           bool `json:"Portrait 1080 x 1920"`
		RES_Portrait1440x2560           bool `json:"Portrait 1440 x 2560"`
		RES_Portrait2160x3840           bool `json:"Portrait 2160 x 3840"`
		RES_Portrait720x1280            bool `json:"Portrait 720 x 1280"`
		RES_PortraitStandardDefinition  bool `json:"Portrait Standard Definition"`
		RES_StandardDefinition          bool `json:"Standard Definition"`
		RES_Triple11520x2160            bool `json:"Triple 11520 x 2160"`
		RES_Triple4096x768              bool `json:"Triple 4096 x 768"`
		RES_Triple5760x1080             bool `json:"Triple 5760 x 1080"`
		RES_Triple7680x1440             bool `json:"Triple 7680 x 1440"`
		RES_TripleStandardDefinition    bool `json:"Triple Standard Definition"`
		RES_Ultrawide2560x1080          bool `json:"Ultrawide 2560 x 1080"`
		RES_Ultrawide3440x1440          bool `json:"Ultrawide 3440 x 1440"`
		RES_UltrawideStandardDefinition bool `json:"Ultrawide Standard Definition"`
	} `json:"resolutiontags"`
	Sort       string `json:"name"`
	SourceTags struct {
		Local    bool `json:"Local"`
		Official bool `json:"Official"`
		Workshop bool `json:"Workshop"`
	} `json:"sourcetags"`
	Tags struct {
		Abstract    bool `json:"Abstract"`
		Animal      bool `json:"Animal"`
		Anime       bool `json:"Anime"`
		CGI         bool `json:"CGI"`
		Cartoon     bool `json:"Cartoon"`
		Cyberpunk   bool `json:"Cyberpunk"`
		Fantasy     bool `json:"Fantasy"`
		Game        bool `json:"Game"`
		Girls       bool `json:"Girls"`
		Guys        bool `json:"Guys"`
		Landscape   bool `json:"Landscape"`
		MMD         bool `json:"MMD"`
		Medieval    bool `json:"Medieval"`
		Memes       bool `json:"Memes"`
		Music       bool `json:"Music"`
		Nature      bool `json:"Nature"`
		PixelArt    bool `json:"Pixel art"`
		Relaxing    bool `json:"Relaxing"`
		Retro       bool `json:"Retro"`
		SciFi       bool `json:"Sci-Fi"`
		Sports      bool `json:"Sports"`
		Technology  bool `json:"Technology"`
		Television  bool `json:"Television"`
		Unspecified bool `json:"Unspecified"`
		Vehicle     bool `json:"Vehicle"`
	} `json:"tags"`
	Type     string `json:"type"`
	TypeTags struct {
		Application bool `json:"Application"`
		Scene       bool `json:"Scene"`
		Video       bool `json:"Video"`
		Web         bool `json:"Web"`
	} `json:"typetags"`
	UtilityTags map[string]interface{} `json:"utilitytags"` // Unknown Type
}

type UserConfig struct {
	General struct {
		Browser *struct {
			AdvertiseExplore       bool          `json:"advertiseexplore"`
			AdvertiseSendToMobile  bool          `json:"advertisesendtomobile"`
			AdvertiseWorkshop      bool          `json:"advertiseworkshop"`
			AdvertiseWorkshopPopup bool          `json:"advertiseworkshoppopup"`
			AgeGateConfirmed       bool          `json:"agegateconfirmed"`
			AuthorBlockListNames   []interface{} `json:"authorblocklistnames"` // Unknown Type
			DefaultFilterConfig    struct {
				Anime                 bool `json:"Anime"`
				RecommendedResolution bool `json:"recommendedresolution"`
				ShowOtherResolution   bool `json:"showotherresolution"`
			} `json:"defaultfilterconfig"`
			Explore struct {
				Custom    []interface{} `json:"custom"`
				Disliked  []interface{} `json:"disliked"`
				Favorites []interface{} `json:"favorites"`
				Liked     []interface{} `json:"liked"`
			} `json:"explore"`
			FilterInfo struct {
				Installed *FilterInfo `json:"installed"`
				Workshop  *FilterInfo `json:"workshop"`
			} `json:"filterinfo"`
			Folders                     []*Folder       `json:"folders"`
			LastSelectedMonitor         string          `json:"lastselectedmonitor"`
			PairedDevices               []*PairedDevice `json:"paireddevices"`
			PairingGUID                 string          `json:"pairingguid"`
			ResultsPerPage              int             `json:"resultsperpage"`
			SeasonalDialogDisabled      interface{}     `json:"seasonaldialogdisabled"` // Unknow Type
			SeasonalDialogTimestamp     int             `json:"seasonaldialogtimestamp"`
			ShowMonitorSelectionOnStart bool            `json:"showmonitorselectiononstart"`
			ViewIconSize                string          `json:"viewiconsize"`
		} `json:"browser"`
		Editor *struct {
			AutoSavingEnabled       bool                   `json:"autosavingenabled"`
			HasShownWorkshopRules   bool                   `json:"hasshownworkshoprules"`
			PreviewRatio            string                 `json:"previewratio"`
			RecentFiles             []string               `json:"recentfiles"`
			SavedLayouts            map[string]interface{} `json:"savedlayouts"` // Unstructured
			ScriptConsoleAutoScroll bool                   `json:"scriptconsoleautoscroll"`
			ShowGrid                bool                   `json:"showgrid"`
			ShowOnStartup           bool                   `json:"showonstartup"`
			ShowSelectionBox        bool                   `json:"showselectionbox"`
			ShowStats               bool                   `json:"showstats"`
			ShowStatsMode           string                 `json:"showstatsmode"`
			SuggestImportResolution bool                   `json:"suggestimportresolution"`
		} `json:"editor"`
		Playlists []*Playlist `json:"playlists"`
		Profiles  []*Profile  `json:"profiles"`
		Shared    struct {
			Hotkeys []*struct {
				Action string `json:"action"`
				Keys   []int  `json:"keys"`
				Value  string `json:"value"`
			} `json:"hotkeys"`
		} `json:"shared"`
		User *struct {
			AdjustDMWColor              bool                   `json:"adjustdwmcolor"`
			AdjustDMWColormode          string                 `json:"adjustdwmcolormode"`
			AntiCrash                   interface{}            `json:"anticrash"` // Unknown Type
			AppRules                    interface{}            `json:"apprules"`  // Unknown Type
			AudioInputDevice            string                 `json:"audioinputdevice"`
			AudioInputThreshold         int                    `json:"audioinputthreshold"`
			AudioinputVolume            int                    `json:"audioinputvolume"`
			AutoStart                   bool                   `json:"autostart"`
			AutoStartScheduler          bool                   `json:"autostartscheduler"`
			AutoStartx64                bool                   `json:"autostartx64"`
			CefCommandLine              string                 `json:"cefcommandline"`
			ExtremeSleepHack            bool                   `json:"extremesleephack"`
			FPS                         int                    `json:"fps"`
			HasShownHighPriorityWarning bool                   `json:"hasshownhighprioritywarning"`
			HasShownWelcomeDialog       bool                   `json:"hasshownwelcomedialog"`
			HighPrecisionTimer          bool                   `json:"highprecisiontimer"`
			IconOpacity                 int                    `json:"iconopacity"`
			Language                    string                 `json:"language"`
			LogLevel                    string                 `json:"loglevel"`
			MediaBlockList              interface{}            `json:"mediablocklist"` // Unknown Type
			MediaIntegration            bool                   `json:"mediaintegration"`
			MonitorDetection            string                 `json:"monitordetection"`
			MSAA                        string                 `json:"msaa"`
			OverrideLockscreen          bool                   `json:"overridelockscreen"`
			OverrideWallpaper           bool                   `json:"overridewallpaper"`
			PauseVRAM                   bool                   `json:"pausevram"`
			PlaybackAudio               string                 `json:"playbackaudio"`
			PlaybackFocus               string                 `json:"playbackfocus"`
			PlaybackFullscreen          string                 `json:"playbackfullscreen"`
			PlaybackMaximized           string                 `json:"playbackmaximized"`
			PlaybackOnBattery           string                 `json:"playbackonbattery"`
			PlaybackSleep               string                 `json:"playbacksleep"`
			PluginDelay                 int                    `json:"plugindelay"`
			Plugins                     map[string]interface{} `json:"plugins"` // Not enough information to structure
			PostProcessing              string                 `json:"postprocessing"`
			Preset                      string                 `json:"preset"`
			ProcessPriority             string                 `json:"processpriority"`
			Reflection                  bool                   `json:"reflection"`
			ReloadAudio                 bool                   `json:"reloadaudio"`
			Resolution                  string                 `json:"resolution"`
			SafeMode                    bool                   `json:"safemode"`
			SlideshowKiller             bool                   `json:"slideshowkiller"`
			SteamLanguage               string                 `json:"steamlanguage"`
			UiCMD                       string                 `json:"uicmd"`
			UiEffects                   bool                   `json:"uieffects"`
			UiHardwareAcceleration      bool                   `json:"uihardwareacceleration"`
			UiQuality                   string                 `json:"uiquality"`
			UiSkin                      string                 `json:"uiskin"`
			UiSkinSeasonal              bool                   `json:"uiskinseasonal"`
			UnPauseAero                 bool                   `json:"unpauseaero"`
			UserMonitors                map[string]struct {
				VideoAudioOutput bool `json:"videoaudiooutput"`
			} `json:"usermonitors"`
			VDesktopEnabled           bool   `json:"vdesktopenabled"`
			VideoAudioOutput          bool   `json:"videoaudiooutput"`
			VideoHardwareAcceleration bool   `json:"videohardwareacceleration"`
			VideoLoopMode             string `json:"videoloopmode"`
			VideoMfStutterHack        bool   `json:"videomfstutterhack"`
			VideoReadMode             string `json:"videoreadmode"`
			WebmFramework             bool   `json:"webmframework"`
			Windows7Mode              string `json:"windows7mode"`
			WindowUpdateRate          int    `json:"windowupdaterate"`
		} `json:"user"`
	} `json:"general"`
	Version int8 `json:"version"`
}
