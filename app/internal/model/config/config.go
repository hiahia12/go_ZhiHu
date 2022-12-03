package config

type Config struct {
	Logger     *Logger     `mapstructure:"logger"yaml:"logger"`
	Database   *Database   `mapstructure:"database" yaml:"database"`
	Server     *Server     `mapstructure:"database" yaml:"database"`
	App        *App        `mapstructure:"app" yaml:"app"`
	Middleware *Middleware `mapstructure:"middleware" yaml:"middleware"`
}
