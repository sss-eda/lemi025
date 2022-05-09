package lemi025

// SystemStatus TODO
type SystemStatus bool

// String TODO
func (systemStatus SystemStatus) String() string {
	if systemStatus {
		return "Running"
	}

	return "Stopped"
}
