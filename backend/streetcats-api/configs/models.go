package configs

type SWAGGER struct {
	Enabled bool `json:"enabled" yaml:"enabled"`
	Auth    struct {
		Enabled bool   `json:"enabled" yaml:"enabled"`
		User    string `json:"user" yaml:"user"`
		Passwd  string `json:"passwd" yaml:"passwd"`
	} `json:"auth" yaml:"auth"`
}

type API struct {
	Host        string  `json:"host" yaml:"host"`
	Port        uint16  `json:"port" yaml:"port"`
	User        string  `json:"user" yaml:"user"`
	Name        string  `json:"name" yaml:"name"`
	Version     string  `json:"version" yaml:"version"`
	Description string  `json:"description" yaml:"description"`
	Password    string  `json:"password" yaml:"password"`
	Swagger     SWAGGER `json:"swagger" yaml:"swagger"`
}

type NATS struct {
	Enabled  bool   `json:"enabled" yaml:"enabled"`
	Host     string `json:"host" yaml:"host"`
	Port     uint16 `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"passwd" yaml:"passwd"`
}

type POSTGIS struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Host    string `json:"host" yaml:"host"`
	DbName  string `json:"dbname" yaml:"dbname"`
	User    string `json:"user" yaml:"user"`
	Port    uint16 `json:"port" yaml:"port"`
	Passwd  string `json:"passwd" yaml:"passwd"`
	Logger  struct {
		Enable bool   `json:"enable" yaml:"enable"`
		Level  string `json:"level" yaml:"level"`
	} `json:"logger" yaml:"logger"`
}

type REDIS struct {
	Enabled bool   `json:"enabled" yaml:"enabled"`
	Host    string `json:"host" yaml:"host"`
	Port    uint16 `json:"port" yaml:"port"`
	Passwd  string `json:"passwd" yaml:"passwd"`
}

type KEYCLOAK struct {
	Enabled      bool   `json:"enabled" yaml:"enabled"`
	Host         string `json:"host" yaml:"host"`
	Port         uint16 `json:"port" yaml:"port"`
	ClientId     string `json:"client_id" yaml:"client_id"`
	ClientSecret string `json:"client_secret" yaml:"client_secret"`
	Realm        string `json:"realm" yaml:"realm"`
	BasicAuth    *struct {
		User   string `json:"user" yaml:"user"`
		Passwd string `json:"passwd" yaml:"passwd"`
	} `json:"basic_auth" yaml:"basic_auth"`
}

type LOG struct {
	Level             string `json:"level" yaml:"level"`
	TimeFormat        string `json:"time_format" yaml:"time_format"`
	DisableStacktrace bool   `json:"disable_stacktrace" yaml:"disable_stacktrace"`
}

type ConfigModel struct {
	Api      API      `json:"api" yaml:"api"`
	Postgis  POSTGIS  `json:"core_db" yaml:"core_db"`
	Redis    REDIS    `json:"redis_db" yaml:"redis_db"`
	Log      LOG      `json:"log" yaml:"log"`
	Nats     NATS     `json:"nats" yaml:"nats"`
	Keycloak KEYCLOAK `json:"keycloak" yaml:"keycloak"`
}
