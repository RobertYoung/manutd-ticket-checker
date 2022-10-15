package mutc

type Config struct {
	PremierLeagueOnly        bool
	MaxPrice                 int
	MinPrice                 int
	NumberOfSeats            int
	HaasUrl                  string
	HaasToken                string
	HaasNotifyDevice         string
	HaasNotificationThrottle int
}
