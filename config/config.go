package config

type VMConfig struct {
	ExternalActionKindRange [2]uint8
}

func NewVMConfig() *VMConfig {
	return &VMConfig{
		ExternalActionKindRange: [2]uint8{0, 0},
	}
}
