package config

type LoggingConfig struct {
	logPath string `toml:"log_path"`
}

func (lc *LoggingConfig) GetLogPath() string {
	return lc.logPath
}

