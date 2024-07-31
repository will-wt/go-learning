package config

// Datasource 使用 yaml/mapstructure 标签指定了 YAML 配置文件中的字段名
type Datasource struct {
	Driver            string `yaml:"driver"  mapstructure:"driver"`
	Host              string `yaml:"url"     mapstructure:"url"`
	Port              string `yaml:"port"    mapstructure:"port"`
	Params            string `yaml:"params"  mapstructure:"params"`
	Username          string `yaml:"username" mapstructure:"username"`
	Password          string `yaml:"password" mapstructure:"password"`
	Dbname            string `yaml:"dbname"   mapstructure:"dbname"`
	MaxIdle           int    `yaml:"max-idle" mapstructure:"max-idle"`
	MaxPoolSize       int    `yaml:"max-pool-size" mapstructure:"max-pool-size"`
	IdleTimeout       int    `yaml:"idle-timeout" mapstructure:"idle-timeout"`
	MaxLifetime       int    `yaml:"max-lifetime" mapstructure:"max-lifetime"`
	ConnectionTimeout int    `yaml:"connection-timeout" mapstructure:"connection-timeout"`
}

func (ds *Datasource) Dsn() string {
	return ds.Username + ":" + ds.Password + "@tcp(" + ds.Host + ":" + ds.Port + ")/" + ds.Dbname + "?" + ds.Params
}
