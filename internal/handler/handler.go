package handler

import (
	"html/template"
	"net/http"
	"strconv"

	gen "github.com/diofanto33/password-gen/internal/generator"
)

func GeneratorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		/* obtain the values from the form and convert them as needed */
		length, _ := strconv.Atoi(r.FormValue("length"))
		useUpper := r.FormValue("useUpper") == "on"
		useLower := r.FormValue("useLower") == "on"
		useNumbers := r.FormValue("useNumbers") == "on"
		useSpecial := r.FormValue("useSpecial") == "on"
		filter := r.FormValue("filter")

		/* create the structure with the user input */
		request := PasswordRequest{
			Length:     length,
			UseUpper:   useUpper,
			UseLower:   useLower,
			UseNumbers: useNumbers,
			UseSpecial: useSpecial,
			Filter:     filter,
		}

		/* call the function that generates the password and get the generated password */
		password := gen.GeneratePassword(request.Length, request.UseUpper, request.UseLower, request.UseNumbers, request.UseSpecial, request.Filter)

		/* create the structure with the generated password */
		pageVariables := PageVariables{
			GeneratedPassword: password,
		}

		/* show the template with the generated data */
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, pageVariables)
		return
	}

	/* otherwise, show the page */
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}
