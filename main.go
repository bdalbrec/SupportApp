package main

import (
	"encoding/json"
	"fmt"
	"github.com/bdalbrec/SupportApp/models"
	"github.com/bdalbrec/SupportApp/configs"
	"github.com/bdalbrec/sll"
	"html/template"
	"net/http"
)

var tpl *template.Template

var logname string

func init() {
		configs.ReadConfigs()
		logname = configs.Configs.Logname
		tpl = template.Must(template.ParseGlob("templates/*.html"))
}


func main() {

	models.InitDB("sqlserver://"+ configs.Configs.SQLUsername + ":" + configs.Configs.SQLPassword + "@" + configs.Configs.SQLServer + ":" + configs.Configs.SQLPort + "/SQLEXPRESS?database=" + configs.Configs.SQLDBName + "&connection+timeout=30")

	// all the HTTP routing
	http.HandleFunc("/", index)
	http.HandleFunc("/nav", nav)
	http.HandleFunc("/insertlink", insertLink)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))
	

	//starting the HTTP server
	sll.LogInfo("Starting the HTTP server", logname)
	err := http.ListenAndServe(":8086", nil) 
	if err != nil {
		sll.LogError("Error in serving HTTP.", logname, err)
	}

	sll.LogInfo("HTTP server listening on port 8086", logname)
}



func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		sll.LogError("Error executing index template.", logname, err)
		return
	}
	sll.LogInfo("Serving index.", logname)
}



func nav(w http.ResponseWriter, req *http.Request) {
	cats, err := models.GetCategories()
	if err != nil {
		sll.LogError("Error retrieving categories from database.", logname, err)
		return
	}
	
	if err := json.NewEncoder(w).Encode(cats); err != nil {
		sll.LogError("Could not encode navigation categories to JSON", logname, err)
	}
}


func insertLink(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		name := req.FormValue("linkName")
		address := req.FormValue("linkAddress")
		category := req.FormValue("linkCategory")
		tags := req.FormValue("linkTags")

		sll.LogInfo("Inserting " + name + address + " into" + category + " links.", logname)

		models.InsertLink(name, address, category, tags)
	}
 


	cats, err := models.GetCategories()
	if err != nil {
		sll.LogError("Error retrieving categories from database.", logname, err)
	}
	
	for _, c := range cats {
		n := c.Name
		fmt.Println(n)
	}

	err = tpl.ExecuteTemplate(w, "insertLink.html", cats)
	if err != nil {
		sll.LogError("Error executing index template.", logname, err)
		return
	}
	sll.LogInfo("Serving addLink.", logname)
}


