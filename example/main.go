package example

import (
	"fmt"
	"log"

	"github.com/0mlml/cfgparser"
)

func main() {
	defaultConfig := &cfgparser.Config{}
	defaultConfig.Literal(
		map[string]bool{
			"debug":        false,
			"auto_restart": true,
		},
		map[string]string{
			"app_name":             "MyApp",
			"db_connection_string": "host=localhost;user=user;password=pass;db=mydb",
		},
		map[string]int{
			"max_retries": 3,
			"port":        8080,
		},
		map[string]float64{
			"version": 1.0,
		},
	)

	// Set the default configuration
	cfgparser.SetDefaultConfig(defaultConfig)

	// Create a new Config object and parse the configuration file
	config := &cfgparser.Config{}
	configPath := "example.cfg"
	if err := config.From(configPath); err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Print the current configuration
	fmt.Println("Current Configuration:")
	printConfig(config)

	// Modify a value
	config.SetBool("debug", true)
	config.SetString("app_name", "UpdatedApp")

	// Write the modified configuration to a new file
	newConfigPath := "updated_example.cfg"
	if err := config.To(newConfigPath); err != nil {
		log.Fatalf("Error writing config file: %v", err)
	}

	fmt.Printf("Updated configuration written to %s\n", newConfigPath)
}

func printConfig(config *cfgparser.Config) {
	fmt.Println("[bool]")
	for _, key := range config.BoolKeys() {
		fmt.Printf("%s=%v\n", key, config.Bool(key))
	}

	fmt.Println("[string]")
	for _, key := range config.StringKeys() {
		fmt.Printf("%s=%s\n", key, config.String(key))
	}

	fmt.Println("[int]")
	for _, key := range config.IntKeys() {
		fmt.Printf("%s=%d\n", key, config.Int(key))
	}

	fmt.Println("[float]")
	for _, key := range config.FloatKeys() {
		fmt.Printf("%s=%f\n", key, config.Float(key))
	}
}
