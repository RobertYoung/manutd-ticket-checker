package mutc

type Config struct {
	PremierLeagueOnly        bool
	MaxPrice                 int
	HaasUrl                  string
	HaasToken                string
	HaasNotifyDevice         string
	HaasNotificationThrottle int
}
