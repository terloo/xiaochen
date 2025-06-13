package db

type Config struct {
	Path            string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

var C Config

func init() {
	C = Config{
		Path:            "./xiaochen.db",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 3600,
	}
}
