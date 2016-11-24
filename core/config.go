package core

type IConfig interface {
	GetSlack() Slack
}

type Slack struct {
	Key string
}

type config struct {
	Slack *Slack
}

func NewConfig() (IConfig, error) {
	c := &config{}
	if err := c.parse(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *config) parse() error {
	return nil
}

func (c *config) GetSlack() Slack {
	return *c.Slack
}
