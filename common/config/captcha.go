package config

type Captcha struct {
	IsEnable  bool 	`mapstructure:"is-enable" json:"isEnable" yaml:"is-enable"`
	KeyLong   int 	`mapstructure:"key-long" json:"keyLong" yaml:"key-long"`
	ImgWidth  int 	`mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`
	ImgHeight int 	`mapstructure:"img-height" json:"imgHeight" yaml:"img-height"`
}
