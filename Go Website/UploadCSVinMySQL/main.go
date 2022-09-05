package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"uploadcsv/mysqlfun"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
	if r.Method == http.MethodPost {
		file, handle, err := r.FormFile("filename")
		if err != nil {
			fmt.Fprintf(w, ""+err.Error())
			return
		}
		defer file.Close()
		fp := csv.NewReader(file)
		data, _ := fp.ReadAll()
		check := mysqlfun.InsertInMySQL(data)
		fileupload, err := os.OpenFile("assets/"+handle.Filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			fmt.Fprintf(w, ""+err.Error())
			return
		}
		defer fileupload.Close()
		fpupload := csv.NewWriter(fileupload)
		f, _ := fileupload.Stat()
		if f.Size() == 0 {
			fpupload.Write([]string{"sid", "name", "course"})
		}
		value := "<table>"
		for _, line := range data {
			err := fpupload.Write([]string{line[0], line[1], line[2]})
			if err != nil {
				panic(err)
			}
			value = value + "<tr>"
			value = value + "<td>" + line[0] + "</td>"
			value = value + "<td>" + line[1] + "</td>"
			value = value + "<td>" + line[2] + "</td>"
			value = value + "</tr>"
		}
		value = value + "</table>"
		fpupload.Flush()

		fmt.Fprintf(w, value)
		if check {
			fmt.Fprintf(w, "Record inserted in MySQL")
		} else {
			fmt.Fprintf(w, "Record not inserted in MySQL! Check connection")
		}
	}

}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":9999", nil)
}
