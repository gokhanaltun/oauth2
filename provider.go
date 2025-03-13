package oauth2

type Provider struct {
	AuthURL  string
	TokenURL string
}

var Providers = struct {
	Google  Provider
	Discord Provider
	Slack   Provider
}{
	Google: Provider{
		AuthURL:  "https://accounts.google.com/o/oauth2/auth",
		TokenURL: "https://oauth2.googleapis.com/token",
	},
	Discord: Provider{
		AuthURL:  "https://discord.com/api/oauth2/authorize",
		TokenURL: "https://discord.com/api/oauth2/token",
	},
	Slack: Provider{
		AuthURL:  "https://slack.com/oauth/v2/authorize",
		TokenURL: "https://slack.com/api/oauth.v2.access",
	},
}
