package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)


var URL string = ""
var Login string = ""
var Password string = ""
var Objects1C []Object1C
var Port string

type OData struct {
	OdataMetadata string `json:"odata.metadata"`
	//Value         []struct {
	//	Name string `json:"name"`
	//	URL  string `json:"url"`
	//} `json:"value"`
	Value []Object1C `json:"value"`
}

type Object1C struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type specificHandler struct {
	Objects1C []Object1C
}


func main() {
	LoadINI()


	DownloadOData()

	//URL1 := `/`
	router := mux.NewRouter()
	router.HandleFunc("/{url1c}", OpenURL)
	router.HandleFunc(`/`, ServeHTTP)
	http.Handle(`/`,router)

	http.ListenAndServe(":" + Port, nil)

}

func OpenURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	URL1C := vars["url1c"]
	//response := fmt.Sprintf("Product category=%s id=%s", cat, id)
	//fmt.Fprint(w, URL1C)
	//print(URL1C)

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL + `/odata/standard.odata/` + URL1C + `?$format=json`, nil)
	req.SetBasicAuth(Login, Password)
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)

	html := `<a href="http://localhost:8080">Назад</a><BR><BR>`
	//PrettyS := string(pretty.Pretty([]byte(s)))
	PrettyS := strings.ReplaceAll(s, "\n", "<BR>")
	html = html + PrettyS

	//pos1 := strings.Index(s, "[")
	//if pos1>0 {
	//	s = s[pos1+1:]
	//}
	//
	//pos2 := strings.LastIndex(s, "]")
	//if pos2>0 {
	//	s = s[:pos2]
	//}

	//fmt.Fprint(w, URL1C)
	//fmt.Fprint(w, s)

	//var MapAll map[string]interface{}
	//err = json.Unmarshal([]byte(s), &MapAll)
	//if err != nil {
	//	fmt.Fprint(w, "error convert to map")
	//	return
	//}

	//fmt.Fprint(w, MapAll["value"])

	//html := "<table><TR><TD>Имя</TD><TD>Значение</TD></TR>"
	//for key, value := range MapAll {
	//	html = html + "<TR>"
	//	html = html + "<TD>" + key + "</TD>"
	//	html = html + "<TD>" + fmt.Sprintf("%v", value) + "</TD>"
	//	html = html + "</TR>"
	//}
	//html = html + "</table>"
	fmt.Fprint(w, html)

}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "Объекты 1С:")

	HTML := "<table>"
	for _, Object1 := range Objects1C {
		HTML = HTML +  "<TR>"

		//HTML = HTML +  "<TD>"
		//HTML = HTML + Object1.Name
		//HTML = HTML +  "</TD>"

		HTML = HTML +  `<TD><a href="http://localhost:8080/`
		HTML = HTML + Object1.URL
		HTML = HTML +  `">` + Object1.Name + `</a></TD>`

		HTML = HTML +  "</TR>"
	}

	HTML = HTML +  "</table>"

	fmt.Fprintf(w, HTML)

}

func WebRoot(w http.ResponseWriter, r *http.Request) {

}

func DownloadOData() {

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL + "/odata/standard.odata?$format=json", nil)
	req.SetBasicAuth(Login, Password)
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s := string(bodyText)
	//print(s)

	var spisok OData
	err = json.Unmarshal([]byte(s), &spisok)

	//var Objects1C []Object1С
	Objects1C = make([]Object1C, len(spisok.Value))

	//var Object1 = Object1C{"11", "22"}
	//Objects1C = append(Objects1C, Object1)

	copy(Objects1C, spisok.Value)

	for i := range Objects1C {
		Name := Objects1C[i].Name
		Name = strings.ReplaceAll(Name, "Catalog_", "Справочник ")
		Name = strings.ReplaceAll(Name, "Document_", "Документ ")
		Name = strings.ReplaceAll(Name, "_", ".")
		Objects1C[i].Name = Name


	}

	//fmt.Print(Objects1C)
}

func LoadINI() {
	ProgramDir := ProgramDir()
	cfg, err := ini.Load(ProgramDir + "settings.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	SectionMain := cfg.Section("Main")

	URL = SectionMain.Key("URL").String()
	Login = SectionMain.Key("Login").String()
	Password = SectionMain.Key("Password").String()
	Port = SectionMain.Key("Port").String()
	//print(Login)
}

func ProgramDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		dir = ""
	}
	dir = dir + "\\"
	return dir
}

// GetStringInBetween Returns empty string if no start string found
func GetStringBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	if e == -1 {
		return
	}
	return str[s:e]
}