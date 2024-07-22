package gojpcal

type Config struct {
	ConnpassGroups []string
}

func LoadConfig() *Config {
	return &Config{
		ConnpassGroups: []string{
			"asakusago",
			"ehimego",
			"gdgchugoku",
			"go-online",
			"gocon",
			"golangtokyo",
			"gotalk",
			"kamakurago",
			"kanazawago",
			"kobego",
			"kyotogo",
			"nobishii-go",
			"sendaigo",
			"tenntenn",
			"tinygo-keeb",
			"umedago",
			"womenwhogo-tokyo",
			"yokohama-go-reading",
		},
	}
}
