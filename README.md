# Go Config Parser

A simple and extensible configuration parser written in Go. This library provides an easy way to read and write configuration files with support for different types including booleans, strings, integers, and floats.

## Installation

`go get github.com/0mlml/cfgparser`

`import "github.com/0mlml/cfgparser"`

## Example Configuration File

```ini
[bool]
# Enable debugging mode
debug=false
# Enable auto-restart
auto_restart=true

[string]
# Application name
app_name=MyApp
# Database connection string
db_connection_string=host=localhost;user=user;password=pass;db=mydb

[int]
# Maximum number of retries
max_retries=5
# Port number
port=8080

[float]
# Application version
version=1.2
```

## Usage

### Parsing Configuration

```go
config := &cfgparser.Config{}
configPath := "example.cfg"
if err := config.From(configPath); err != nil {
	log.Fatalf("Error parsing config file: %v", err)
}
```

### Writing Configuration

```go
configPath := "example.cfg"
if err := config.To(configPath); err != nil {
	log.Fatalf("Error writing config file: %v", err)
}
```

## Customizing Default Configuration

You can customize the default configuration by calling the `SetDefaultConfig` function. It will be used to backfill any missing values when parsing a configuration file.

```go
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
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
