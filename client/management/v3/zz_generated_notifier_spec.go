package client

const (
	NotifierSpecType                             = "notifierSpec"
	NotifierSpecFieldClusterID                   = "clusterId"
	NotifierSpecFieldDescription                 = "description"
	NotifierSpecFieldDisplayName                 = "displayName"
	NotifierSpecFieldMicrosoftTeamsWebhookConfig = "microsoftTeamsWebhookConfig"
	NotifierSpecFieldPagerdutyConfig             = "pagerdutyConfig"
	NotifierSpecFieldSMTPConfig                  = "smtpConfig"
	NotifierSpecFieldSlackConfig                 = "slackConfig"
	NotifierSpecFieldWebhookConfig               = "webhookConfig"
	NotifierSpecFieldWechatConfig                = "wechatConfig"
)

type NotifierSpec struct {
	ClusterID                   string                       `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Description                 string                       `json:"description,omitempty" yaml:"description,omitempty"`
	DisplayName                 string                       `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	MicrosoftTeamsWebhookConfig *MicrosoftTeamsWebhookConfig `json:"microsoftTeamsWebhookConfig,omitempty" yaml:"microsoftTeamsWebhookConfig,omitempty"`
	PagerdutyConfig             *PagerdutyConfig             `json:"pagerdutyConfig,omitempty" yaml:"pagerdutyConfig,omitempty"`
	SMTPConfig                  *SMTPConfig                  `json:"smtpConfig,omitempty" yaml:"smtpConfig,omitempty"`
	SlackConfig                 *SlackConfig                 `json:"slackConfig,omitempty" yaml:"slackConfig,omitempty"`
	WebhookConfig               *WebhookConfig               `json:"webhookConfig,omitempty" yaml:"webhookConfig,omitempty"`
	WechatConfig                *WechatConfig                `json:"wechatConfig,omitempty" yaml:"wechatConfig,omitempty"`
}
