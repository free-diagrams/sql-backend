package httpresp

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
