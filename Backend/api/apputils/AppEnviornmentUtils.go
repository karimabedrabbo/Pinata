package apputils

import (
	k "github.com/karimabedrabbo/eyo/api/constants"
	"log"
	"os"
)

func GetAppEnvIsProduction() bool {
	appenv := GetAppEnv()
	if appenv == k.AppEnvProduction {
		return true
	} else if appenv == k.AppEnvDevelopment {
		return false
	}
	log.Fatalf("unknown app enviornment: %s", appenv)
	return false
}

func GetAppContextIsLocal() bool {
	appcontext := GetAppContext()
	if appcontext == k.AppContextRemote {
		return false
	} else if appcontext == k.AppContextLocal {
		return true
	} else if appcontext == k.AppContextLocalContainer {
		return false
	}
	log.Fatalf("unknown app context: %s", appcontext)
	return false
}

func getAppContextKeyPrefix() string {
	appcontext := GetAppContext()
	if appcontext == k.AppContextRemote {
		return k.AppContextKeyPrefixRemote
	} else if appcontext == k.AppContextLocalContainer {
		return k.AppContextKeyPrefixLocalContainer
	} else if appcontext == k.AppContextLocal {
		return k.AppContextKeyPrefixLocal
	}
	log.Fatalf("unknown app context: %s", appcontext)
	return ""
}

func getAppEnvKeyPrefix() string {
	if GetAppEnvIsProduction() {
		return k.AppEnvKeyPrefixProduction
	} else {
		return k.AppEnvKeyPrefixDevelopment
	}
}

func GetAppUrl() string {
	if GetAppEnvIsProduction() {
		return GetProductionUrl()
	} else {
		return GetDevelopmentUrl()
	}
}

func GetProductionUrl() string {
	return k.AppProductionUrl
}

func GetDevelopmentUrl() string {
	return k.AppDevelopmentUrl
}

func GetAppEnv() string {
	return os.Getenv("APP_ENV")
}

func GetAppContext() string {
	return os.Getenv("APP_CONTEXT")
}

func GetAppName() string {
	return k.AppName
}

func GetAppApiSecret() []byte {
	return []byte(os.Getenv("API_SECRET"))
}

func GetApiPort() string {
	return os.Getenv("API_PORT")
}

func GetGoogleCloudServiceEmail() string {
	return os.Getenv("GCLOUD_SERVICE_EMAIL")
}

func GetGoogleCloudServicePrivateKey() []byte {
	return []byte(os.Getenv("GCLOUD_SERVICE_PRIVATE_KEY"))
}

func GetSendgridApiKey() string {
	return os.Getenv("SENDGRID_API_KEY")
}

func getRedisKeyPrefix() string {
	return k.RedisKeyInitialPrefix + getAppContextKeyPrefix()
}

func GetRedisHost() string {
	return os.Getenv(getRedisKeyPrefix() + k.RedisKeySuffixHost)
}

func GetRedisPort() string {
	return os.Getenv(getRedisKeyPrefix() + k.RedisKeySuffixPort)
}

func getDatabaseKeyPrefix() string {
	return k.DatabaseKeyInitialPrefix + getAppContextKeyPrefix() + getAppEnvKeyPrefix()
}

func GetDatabaseHost() string {
	return os.Getenv(getDatabaseKeyPrefix() + k.DatabaseKeySuffixHost)
}

func GetDatabasePort() string {
	return os.Getenv(getDatabaseKeyPrefix() + k.DatabaseKeySuffixPort)
}

func GetDatabaseName() string {
	return os.Getenv(getDatabaseKeyPrefix() + k.DatabaseKeySuffixName)
}

func GetDatabaseUser() string {
	return os.Getenv(getDatabaseKeyPrefix() + k.DatabaseKeySuffixUser)
}

func GetDatabasePassword() string {
	return os.Getenv(getDatabaseKeyPrefix() + k.DatabaseKeySuffixPassword)
}
