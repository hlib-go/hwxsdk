package hwxsdk

type Config struct {
	ServiceUrl  string
	AppId       string
	AppSecret   string
	AccessToken string
}

func (c *Config) GetServiceUrl() string {
	if c.ServiceUrl == "" {
		c.ServiceUrl = "https://api.weixin.qq.com"
	}
	return c.ServiceUrl
}
