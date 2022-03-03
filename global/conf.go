package global

// 全局配置
var Conf config

type config struct {
	App      appConf        `mapstructure:"app"`
	Log      logConf        `mapstructure:"log"`
	Redis    []redisConf    `mapstructure:"redis"`
	Postgres []postgresConf `mapstructure:"postgres"`
	Tdengine []TdengineConf `mapstructure:"tdengine"`
	Mqtt     []MqttConf     `mapstructure:"mqtt"`
}

// app配置
type appConf struct {
	AppId      string `mapstructure:"app_id"`
	AppSecret  string `mapstructure:"app_secret"`
	Env        string `mapstructure:"env"`
	ServerAddr string `mapstructure:"server_addr"`
	RootPath   string
}

// 日志配置
type logConf struct {
	MainPath string `mapstructure:"main_path"`
}

// 日志配置
type redisConf struct {
	InsName      string `mapstructure:"ins_name"`
	Addr         string `mapstructure:"addr"`
	Auth         string `mapstructure:"auth"`
	Db           int    `mapstructure:"db"`
	ConnTimeout  int    `mapstructure:"conn_timeout"`  //客户端连接超时时间 单位：毫秒
	ReadTimeout  int    `mapstructure:"read_timeout"`  //客户端读超时时间 单位：毫秒
	WriteTimeout int    `mapstructure:"write_timeout"` //客户端写超时时间 单位：毫秒
	MaxIdle      int    `mapstructure:"max_idle"`      //客户端最大空闲连接
	MaxActive    int    `mapstructure:"max_active"`    //客户端最大活跃连接
	IdleTimeout  int    `mapstructure:"idle_timeout"`  //空闲连接超时时间
	MaxConnAge   int    `mapstructure:"max_conn_age"`  //客户端连接做大超时时间
}

// postgres配置
type postgresConf struct {
	InsName      string `mapstructure:"ins_name"`
	Addr         string `mapstructure:"addr"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max_idle_time"`
	MaxLifetime  int    `mapstructure:"max_life_time"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

// tdengine配置
type TdengineConf struct {
	InsName      string `mapstructure:"ins_name"`
	Driver       string `mapstructure:"driver"`
	Network      string `mapstructure:"network"`
	Fqdn         string `mapstructure:"fqdn"`
	Port         int    `mapstructure:"port"`
	RestfulPort  string `mapstructure:"restful_port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"db_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxIdleTime  int    `mapstructure:"max_idle_time"`
	MaxLifeTime  int    `mapstructure:"max_life_time"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

//mqtt配置
type MqttConf struct {
	InsName  string `mapstructure:"ins_name"`
	Addr     string `mapstructure:"addr"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	ClientId string `mapstructure:"client_id"`
}
