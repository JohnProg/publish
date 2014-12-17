package controllers

import (
    "fmt"
    "net/http"
    //"time"
    //"database/sql"
    _ "github.com/lib/pq"
    "publish/admin/helpers"
    "publish/admin/models"
    "strconv"
    //"log"
    //"github.com/gorilla/schema"
    "encoding/json"
    //"log"
    //"io/ioutil"
    //"path/filepath"
    //"strings"
    //"html/template"
)

type TemplateController struct {
    Controller
}

func (this *TemplateController) PostTemplate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    template := models.Template{}

    err := json.NewDecoder(r.Body).Decode(&template)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    template.Post()

}


func (this *TemplateController) PutTemplate(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    template := models.Template{}

    err := json.NewDecoder(r.Body).Decode(&template)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    template.Update()

}

func (this *TemplateController) GetTemplates(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    templates := models.GetTemplates()

    res, err := json.Marshal(templates)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *TemplateController) GetTemplateByNodeId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    nodeIdStr := r.URL.Query().Get(":nodeId")
    nodeId, _ := strconv.Atoi(nodeIdStr)

    template := models.GetTemplateByNodeId(nodeId)

    res, err := json.Marshal(template)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}