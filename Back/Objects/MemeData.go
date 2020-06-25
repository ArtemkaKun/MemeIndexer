package Objects

type MemeData struct {
	MemeFile string `json:"memeFile,omitempty"`
	MemeTags `json:",omitempty"`
}
