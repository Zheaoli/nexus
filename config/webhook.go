package config

type WebhookConfig struct {
	URL          string              `toml:"url"`
	Source       string              `toml:"source"`
	Target       string              `toml:"target"`
	Wechat       *WechatWebhookConfig `toml:"wechat_config"`
	TemplatePath string              `toml:"template_path"`
}

type WechatWebhookConfig struct {
	GroupName    string `toml:"group_name"`
	GroupID      string `toml:"group_id"`
	GroupOwner   string `toml:"group_owner"`
	GroupMembers []string `toml:"group_members"`
}
