package runner

type Message struct {
	Results []Result `json:"results"`
	Fails   int      `json:"fails"`
}
