package templates

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTmpl() {

	tmpl, tempErr := template.ParseGlob("./templates/*.html")

	if tempErr != nil {
		fmt.Printf("erreur avec le temp : %v", tempErr.Error())
		os.Exit(02)
	}

	Temp = tmpl

}
