package config

import "os"

func StaticUrl() string {
	if IsProductionEnv() {
		return "https://asset.lekcije.com/static"
	} else if IsDevelopmentEnv() {
		return "http://asset.local.lekcije.com/static"
	} else {
		return "/static"
	}
}

func EnvString() string {
	return os.Getenv("NODE_ENV")
}

func IsProductionEnv() bool {
	return EnvString() == "production"
}

func IsDevelopmentEnv() bool {
	return EnvString() == "development"
}

func IsLocalEnv() bool {
	return EnvString() == "local"
}
