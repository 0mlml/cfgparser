package cfgparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var defaultConfig *Config

func SetDefaultConfig(config *Config) {
	defaultConfig = config
}

type Config struct {
	boolOptions   map[string]bool
	stringOptions map[string]string
	intOptions    map[string]int
	floatOptions  map[string]float64
}

func (config *Config) Literal(
	boolOptions map[string]bool,
	stringOptions map[string]string,
	intOptions map[string]int,
	floatOptions map[string]float64) {
	*config = Config{
		boolOptions:   boolOptions,
		stringOptions: stringOptions,
		intOptions:    intOptions,
		floatOptions:  floatOptions,
	}
}

func (config *Config) Bool(key string) bool {
	return config.boolOptions[key]
}

func (config *Config) SetBool(key string, value bool) {
	config.boolOptions[key] = value
}

func (config *Config) BoolKeys() []string {
	keys := make([]string, len(config.boolOptions))
	i := 0
	for key := range config.boolOptions {
		keys[i] = key
		i++
	}
	return keys
}

func (config *Config) String(key string) string {
	return config.stringOptions[key]
}

func (config *Config) SetString(key string, value string) {
	config.stringOptions[key] = value
}

func (config *Config) StringKeys() []string {
	keys := make([]string, len(config.stringOptions))
	i := 0
	for key := range config.stringOptions {
		keys[i] = key
		i++
	}
	return keys
}

func (config *Config) Int(key string) int {
	return config.intOptions[key]
}

func (config *Config) SetInt(key string, value int) {
	config.intOptions[key] = value
}

func (config *Config) IntKeys() []string {
	keys := make([]string, len(config.intOptions))
	i := 0
	for key := range config.intOptions {
		keys[i] = key
		i++
	}
	return keys
}

func (config *Config) Float(key string) float64 {
	return config.floatOptions[key]
}

func (config *Config) SetFloat(key string, value float64) {
	config.floatOptions[key] = value
}

func (config *Config) FloatKeys() []string {
	keys := make([]string, len(config.floatOptions))
	i := 0
	for key := range config.floatOptions {
		keys[i] = key
		i++
	}
	return keys
}

func (config *Config) Default() {
	*config = *defaultConfig
}

func (config *Config) From(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	*config = *defaultConfig

	scanner := bufio.NewScanner(file)
	var currentSection string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = line[1 : len(line)-1]
			continue
		}

		split := strings.SplitN(line, "=", 2)
		if len(split) != 2 {
			return fmt.Errorf("invalid line: %s", line)
		}

		key := split[0]
		value := split[1]

		switch currentSection {
		case "bool":
			if _, ok := config.boolOptions[key]; !ok {
				return fmt.Errorf("unknown bool key: %s", key)
			}
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("invalid bool value for key %s: %v", key, err)
			}
			config.boolOptions[key] = boolValue
		case "string":
			if _, ok := config.stringOptions[key]; !ok {
				return fmt.Errorf("unknown string key: %s", key)
			}
			config.stringOptions[key] = value
		case "int":
			if _, ok := config.intOptions[key]; !ok {
				return fmt.Errorf("unknown int key: %s", key)
			}
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return fmt.Errorf("invalid int value for key %s: %v", key, err)
			}
			config.intOptions[key] = intValue
		case "float":
			if _, ok := config.floatOptions[key]; !ok {
				return fmt.Errorf("unknown float key: %s", key)
			}
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Errorf("invalid float value for key %s: %v", key, err)
			}
			config.floatOptions[key] = floatValue
		default:
			return fmt.Errorf("unknown section: %s", currentSection)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (config *Config) To(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if len(config.boolOptions) != 0 {
		writer.WriteString("[bool]\n")
		for key, value := range config.boolOptions {
			line := fmt.Sprintf("%s=%v\n", key, value)
			writer.WriteString(line)
		}
	}

	if len(config.stringOptions) != 0 {
		writer.WriteString("[string]\n")
		for key, value := range config.stringOptions {
			line := fmt.Sprintf("%s=%s\n", key, value)
			writer.WriteString(line)
		}
	}

	if len(config.intOptions) != 0 {
		writer.WriteString("[int]\n")
		for key, value := range config.intOptions {
			line := fmt.Sprintf("%s=%d\n", key, value)
			writer.WriteString(line)
		}
	}

	if len(config.floatOptions) != 0 {
		writer.WriteString("[float]\n")
		for key, value := range config.floatOptions {
			line := fmt.Sprintf("%s=%f\n", key, value)
			writer.WriteString(line)
		}
	}

	return writer.Flush()
}
