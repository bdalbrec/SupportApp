package main

import (
	"github.com/bdalbrec/SupportApp/models"
	"github.com/bdalbrec/SupportApp/configs"
	"github.com/bdalbrec/sll"
	"html/template"
	"net/http"
	"strconv"
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
	http.HandleFunc("/insertlink", insertLink)
	http.HandleFunc("/insertPhone", insertPhone)
	http.HandleFunc("/insertCategory", insertCategory)
	http.HandleFunc("/CiscoGenesys", ciscoGenesys)
	http.HandleFunc("/FrontDoorPhoneMenu", frontDoorPhoneMenu)
	http.HandleFunc("/YAS_Klarity", yasKlarity)
	http.HandleFunc("/WindowsShortcuts", windowsShortcuts)
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

	type Link struct {
		Name string
		Address string
	}
	
	type Section struct {
		Name string
		Links []Link
	}

	
	type Phone struct {
		Name string
		Number string
	}

	type O struct {
		Sections []Section
		Phones []Phone
	}

	var Output O

	cats, err := models.GetCategories()
	if err != nil {
		sll.LogError("Error retrieving categories from database.", logname, err)
		return
	}

	for _, c := range cats {
		var s Section
		s.Name = c.Name

		// grab all the links under the category of s
		links, err := models.GetLinks(s.Name)
		var lnk Link
		if err != nil {
			sll.LogError("Error retrieving links from datbase.", logname, err)
			return
		}

		// iterate over the links and add them to the Section.Links array
		for _, l := range links {
			lnk.Name = l.Name
			lnk.Address = l.Address
			s.Links = append(s.Links, lnk)
		}

		// append the newly created section and its links to the Sections array of Output variable O
		 Output.Sections = append(Output.Sections, s) 
	}


	// implement getting the phone numbers from the database. Will need to change the output to a struct to get it to the template
	ps, err := models.GetPhones()
	if err != nil {
		sll.LogError("Error retrieving phones from database.", logname, err)
		return
	}

	for _, p := range ps {
		var phone Phone
		phone.Name = p.Name
		phone.Number = p.Number

		Output.Phones = append(Output.Phones, phone)
	}

	// execute the template and send it to the client
	err = tpl.ExecuteTemplate(w, "index.html", Output)
	if err != nil {
		sll.LogError("Error executing index template.", logname, err)
		return
	}
	sll.LogInfo("Serving index.", logname)
}


func insertCategory(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		name := req.FormValue("categoryName")
		number := req.FormValue("categoryNumber")
		sll.LogInfo("Inserting " + name + " into" + " category.", logname)

		n, err := strconv.Atoi(number)
		if err != nil {
			sll.LogError("Error converting category number string to int", logname, err)
			return
		}

		models.InsertCategory(name, n)
	}

	err := tpl.ExecuteTemplate(w, "insertCategory.html", nil)
	if err != nil {
		sll.LogError("Error executing insertCategory template.", logname, err)
		return
	}
	sll.LogInfo("Serving insertCategory.", logname)
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
	
	err = tpl.ExecuteTemplate(w, "insertLink.html", cats)
	if err != nil {
		sll.LogError("Error executing index template.", logname, err)
		return
	}
	sll.LogInfo("Serving insertLink.", logname)
}

func insertPhone(w http.ResponseWriter, req *http.Request) {
	
	if req.Method == http.MethodPost {
		name := req.FormValue("phoneName")
		number := req.FormValue("phoneNumber")

		sll.LogInfo("Inserting " + name + number + " into" + " phone.", logname)

		models.InsertPhone(name, number)
	}

	err := tpl.ExecuteTemplate(w, "insertPhone.html", nil)
	if err != nil {
		sll.LogError("Error executing insertPhone template.", logname, err)
		return
	}
	sll.LogInfo("Serving insertPhone.", logname)

}


// handlers for static pages

func ciscoGenesys(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "CiscoGenesys.html", nil)
	if err != nil {
		sll.LogError("Error executing CiscoGenesys template.", logname, err)
		return
	}
	sll.LogInfo("Serving CiscoGenesys.", logname)
}

func frontDoorPhoneMenu(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "FrontDoorPhoneMenu.html", nil)
	if err != nil {
		sll.LogError("Error executing FrontDoorPhoneMenu template.", logname, err)
		return
	}
	sll.LogInfo("Serving FrontDoorPhoneMenu.", logname)
}

func yasKlarity(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "YAS_Klarity.html", nil)
	if err != nil {
		sll.LogError("Error executing YAS_Klarity template.", logname, err)
		return
	}
	sll.LogInfo("Serving YAS_Klarity.", logname)
}

func windowsShortcuts(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "WindowsShortcuts.html", nil)
	if err != nil {
		sll.LogError("Error executing windows_shortcuts template.", logname, err)
		return
	}
	sll.LogInfo("Serving windows_shortcuts.", logname)
}