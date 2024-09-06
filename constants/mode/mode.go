package mode

// Mode shows whether the application is in development mode or production mode.
type Mode uint

const (
	_ Mode = iota
	Development
	Production
)

func (m Mode) String() string {
	return []string{"", "Development", "Production"}[m]
}
