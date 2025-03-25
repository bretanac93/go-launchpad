package logger

type Config struct {
    // Output is the log output format. (allowed values = plaintext, json)
    Output string `env:"OUTPUT,default=plaintext"`
    // Level is the log level. (allowed values = DEBUG, INFO, WARN, ERROR)
    Level string `env:"LEVEL,default=DEBUG"` 
}

