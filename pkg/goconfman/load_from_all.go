package goconfman

func LoadFromAll(config interface{}) {
	LoadFromDefaults(config)
	LoadFromAliases(config)
}
