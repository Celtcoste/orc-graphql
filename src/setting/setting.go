package setting

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type App struct {
	PrefixUrl string

	RuntimeRootPath string

	LogSavePath        string
	LogSaveName        string
	LogFileExt         string
	TimeFormat         string
	RunMode 		   string
}

var AppSetting = &App{}

// Setup initialize the configuration instance
func Setup() {
	if GetenvStr("APP_ENV") == "TEST" {
		err:= godotenv.Overload()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	AppSetting.RuntimeRootPath = GetenvStr("RUNTIME_ROOT_PATH")
	AppSetting.LogSavePath = GetenvStr("LOG_SAVE_PATH")
	AppSetting.LogSaveName = GetenvStr("LOG_SAVE_NAME")
	AppSetting.LogFileExt = GetenvStr("LOG_FILE_EXT")
	AppSetting.TimeFormat = GetenvStr("TIME_FORMAT")
	AppSetting.RunMode = GetenvStr("RUN_MODE")
}

func GetenvStr(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatal("Environment variable %s doesn't exist", key)
	}
	return v
}

func getenvInt(key string) int {
	s := GetenvStr(key)
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func getenvBool(key string) bool {
	s := GetenvStr(key)
	v, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
