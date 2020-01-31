package schema

type Message struct {
	Message string   `json:"message" xml:"message"`
	Path    []string `json:"path" xml:"path"`
	Code    int      `json:"code" xml:"code"`
}
