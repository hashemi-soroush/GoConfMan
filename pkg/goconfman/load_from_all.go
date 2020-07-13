package goconfman

func LoadFromAll(config interface{}, envVarPrefix string) {
	LoadFromDefaults(config)
	LoadFromAliases(config)
	LoadFromEnvVars(config, envVarPrefix)
}
