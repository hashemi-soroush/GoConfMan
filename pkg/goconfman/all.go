package goconfman

func LoadFromAll(config interface{}, configFilePath string, envVarPrefix string) {
	LoadFromDefaults(config)
	if configFilePath != "" {
		err := LoadFromFile(config, configFilePath)
		if err != nil {
			panic(err)
		}
	}
	LoadFromEnvVars(config, envVarPrefix)
	LoadFromAliases(config)
}
