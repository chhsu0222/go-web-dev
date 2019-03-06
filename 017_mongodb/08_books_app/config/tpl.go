package config

import "html/template"

// TPL handles all the template related functionalities.
var TPL *template.Template

func init() {
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}
