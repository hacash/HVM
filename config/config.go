package config

type VMConfig struct {
	ExternalActionKindRange [2]uint16
}

func NewVMConfig() *VMConfig {
	return &VMConfig{
		ExternalActionKindRange: [2]uint16{0, 0},
	}
}
