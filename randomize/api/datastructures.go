package api

type Configuration struct {
	MinListLength   int
	MaxListLength   int
	MinMapLength    int
	MaxMapLength    int
	MinStringLength int
	MaxStringLength int
}

func DefaultConfiguration() Configuration {
	return Configuration{
		MinListLength:   1,
		MaxListLength:   10,
		MinMapLength:    1,
		MaxMapLength:    5,
		MinStringLength: 5,
		MaxStringLength: 15,
	}
}
