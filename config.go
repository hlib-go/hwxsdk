package hwxsdk

// 微信公众号配置信息
type Config struct {
	ServiceUrl string `json:"serviceUrl" bson:"serviceUrl"`
	Appid      string `json:"appid" bson:"appid"`
	Secret     string `json:"secret" bson:"secret"`
}

func (c *Config) GetServiceUrl() string {
	if c.ServiceUrl == "" {
		c.ServiceUrl = "https://api.weixin.qq.com"
	}
	return c.ServiceUrl
}
