package systeminfo

type SystemInfo interface {
	HomeDir() (string, error)
	EOL() (string, error)
	CPUs() (string, error)
	Username() (string, error)
	Architecture() (string, error)
}
