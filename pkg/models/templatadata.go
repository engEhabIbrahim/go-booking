package models

// TemplateData sets the data to be sent from handler to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string // Success Message
	Warning   string
	Error     string
}
