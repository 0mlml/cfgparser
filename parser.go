package cfgparser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DefaultConfig() *Config {
	return &Config{
		BoolOptions: map[string]bool{
			"debug":        false,
			"auto_restart": true,
		},
		StringOptions: map[string]string{
			"app_name":             "MyApp",
			"db_connection_string": "host=localhost;user=user;password=pass;db=mydb",
		},
		IntOptions: map[string]int{
			"max_retries": 3,
			"port":        8080,
		},
		FloatOptions: map[string]float64{
			"version": 1.0,
		},
	}
}

type Config struct {
	BoolOptions   map[string]bool
	StringOptions map[string]string
	IntOptions    map[string]int
	FloatOptions  map[string]float64
}

func ParseConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := DefaultConfig()

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
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := split[0]
		value := split[1]

		switch currentSection {
		case "bool":
			if _, ok := config.BoolOptions[key]; !ok {
				return nil, fmt.Errorf("unknown bool key: %s", key)
			}
			boolValue, err := strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf("invalid bool value for key %s: %v", key, err)
			}
			config.BoolOptions[key] = boolValue
		case "string":
			if _, ok := config.StringOptions[key]; !ok {
				return nil, fmt.Errorf("unknown string key: %s", key)
			}
			config.StringOptions[key] = value
		case "int":
			if _, ok := config.IntOptions[key]; !ok {
				return nil, fmt.Errorf("unknown int key: %s", key)
			}
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("invalid int value for key %s: %v", key, err)
			}
			config.IntOptions[key] = intValue
		case "float":
			if _, ok := config.FloatOptions[key]; !ok {
				return nil, fmt.Errorf("unknown float key: %s", key)
			}
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid float value for key %s: %v", key, err)
			}
			config.FloatOptions[key] = floatValue
		default:
			return nil, fmt.Errorf("unknown section: %s", currentSection)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func WriteConfig(filename string, config *Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	if len(config.BoolOptions) != 0 {
		writer.WriteString("[bool]\n")
		for key, value := range config.BoolOptions {
			line := fmt.Sprintf("%s=%v\n", key, value)
			writer.WriteString(line)
		}
	}

	if len(config.StringOptions) != 0 {
		writer.WriteString("[string]\n")
		for key, value := range config.StringOptions {
			line := fmt.Sprintf("%s=%s\n", key, value)
			writer.WriteString(line)
		}
	}

	if len(config.IntOptions) != 0 {
		writer.WriteString("[int]\n")
		for key, value := range config.IntOptions {
			line := fmt.Sprintf("%s=%d\n", key, value)
			writer.WriteString(line)
		}
	}

	if len(config.FloatOptions) != 0 {
		writer.WriteString("[float]\n")
		for key, value := range config.FloatOptions {
			line := fmt.Sprintf("%s=%f\n", key, value)
			writer.WriteString(line)
		}
	}

	return writer.Flush()
}
