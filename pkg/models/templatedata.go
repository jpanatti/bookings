package models

type TemplateData struct {
	StringMap 	map[string]string
	IntMap 		map[string]int
	FloatMap 	map[string]float32
	Data 		map[string]interface{}
	CSRFToken string //para Forms
	Flash string  //Messages al usuario
	Warning string
	Error string
}

