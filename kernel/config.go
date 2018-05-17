package kernel

var (
	// Qwer probe
	Qwer = "qwerty"

	// Conf srv
	Conf Config
)

// Config for app
type Config struct {
	Addr string
	Port int
	Ssl  bool
}

// Load config
func Load(addr string, port int, ssl bool) {
	Conf = Config{
		Addr: addr,
		Port: port,
		Ssl:  ssl,
	}
}
