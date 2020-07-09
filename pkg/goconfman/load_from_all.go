package goconfman

func LoadFromAll(config interface{}) {
	cDefaults, ok := config.(ConfigWithDefaults)
	if ok {
		LoadFromDefaults(cDefaults)
	}

	cAliases, ok := config.(ConfigWithAliases)
	if ok {
		LoadFromAliases(cAliases)
	}
}
