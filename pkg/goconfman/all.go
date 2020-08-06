package goconfman

func LoadFromAll(config interface{}, configMap map[string]interface{}, configFilePath string, envVarPrefix string) {
	LoadFromDefaults(config)
	if configMap != nil {
		err := LoadFromMap(config, configMap)
		if err != nil {
			panic(err)
		}
	}
	if configFilePath != "" {
		err := LoadFromFile(config, configFilePath)
		if err != nil {
			panic(err)
		}
	}
	LoadFromEnvVars(config, envVarPrefix)
	LoadFromAliases(config)
}
