package view
import (
"log"
"html/template"
"net/http"
"path/filepath"

)
var (
 tpIndex = parseTemplate("root.tmpl", "index.tmpl")
 tpAdminLogin = parseTemplate("root.tmpl", "admin/login.tmpl")
 tpAdminList = parseTemplate("root.tmpl", "admin/list.tmpl")
 tpAdminCreate = parseTemplate("root.tmpl", "admin/create.tmpl")
 tpAdminEdit = parseTemplate("root.tmpl", "admin/edit.tmpl")
 )
// func init() {
// 	tpIndex.Funcs(template.FuncMap{})
// 	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")
// 	 if err != nil {
// 	 	panic(err)
// 	}
// 	tpIndex = tpIndex.Lookup("root")
// }
const templateDir = "template"
func joinTemplateDir(files ...string) []string{
 r:= make([]string, len(files))
 for i, f := range files {
 	r[i] = filepath.Join(templateDir, f)
 }
 return r
}
func parseTemplate(files ...string) *template.Template{
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinTemplateDir(files...)...)
	 if err != nil {
	 	panic(err)
	}
	t = t.Lookup("root")
	return t

}
func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type","text/html; charset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}

}
