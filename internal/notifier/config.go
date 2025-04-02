package notifier

type Channel string

const (
	Signal Channel = "signal"
	Email  Channel = "email"
)

type Config struct {
	Channel Channel `yaml:"channel"`

	// Signal
	SignalFrom string `yaml:"signal_from"`
	SignalTo   string `yaml:"signal_to"`

	// Email
	SMTPServer string `yaml:"smtp_server"`
	SMTPPort   int    `yaml:"smtp_port"`
	SMTPUser   string `yaml:"smtp_user"`
	SMTPPass   string `yaml:"smtp_pass"`
	EmailFrom  string `yaml:"email_from"`
	EmailTo    string `yaml:"email_to"`
}