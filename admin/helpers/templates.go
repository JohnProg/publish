package helpers

import(
	//"net/http"
    //"database/sql"
    //_ "github.com/lib/pq"
	"html/template"
	// "net/http"
	// "fmt"
)


var Templates map[string]*template.Template

// renderTemplate is a wrapper around template.ExecuteTemplate.
// func RenderTemplate(w http.ResponseWriter, name string, data map[string]interface{}) error {
//     // Ensure the template exists in the map.
//     tmpl, ok := Templates[name]
//     if !ok {
//         return fmt.Errorf("The template %s does not exist.", name)
//     }
//     fmt.Print(Templates)
//     w.Header().Set("Content-Type", "text/html; charset=utf-8")
//     tmpl.ExecuteTemplate(w, "base", data)

//     return nil
// }
