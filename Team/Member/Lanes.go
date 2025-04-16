package Member

var lanes = []string{
	"Top",
	"Jungle",
	"Mid",
	"Bot",
	"Support",
}

// GetLanes returns the list of lanes in League of Legends.
func GetLanes() []string {
	return lanes
}
