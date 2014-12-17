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

type ContentTypeController struct {
    Controller
}

func (this *ContentTypeController) Post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    contentType := models.ContentType{}

    err := json.NewDecoder(r.Body).Decode(&contentType)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    contentType.Post()

}

// func (this *ContentTypeController) GetContentTypes(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     contentTypes := models.GetContentTypes()

//     res, err := json.Marshal(contentTypes)
//     helpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)

// }

// func (this *ContentTypeController) GetContentTypeExtendedByNodeId(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     var nodeIdStr string = ""
//     nodeIdStr = r.URL.Query().Get(":nodeId")

//     nodeId, _ := strconv.Atoi(nodeIdStr)
//     content := models.GetContentTypeExtendedByNodeId(nodeId)
//     res, err := json.Marshal(content)
//     helpers.PanicIf(err)

//     fmt.Fprintf(w,"%s",res)
// }

func (this *ContentTypeController) GetContentTypeByNodeId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var nodeIdStr string = ""
    nodeIdStr = r.URL.Query().Get(":nodeId")

    nodeId, _ := strconv.Atoi(nodeIdStr)

    var extended bool = false
    extended, _ = strconv.ParseBool(r.URL.Query().Get("extended"))

    //extended, _ := strconv.Atoi(extendedStr)

    if(!extended){
        content := models.GetContentTypeByNodeId(nodeId)
        res, err := json.Marshal(content)
        helpers.PanicIf(err)

        fmt.Fprintf(w,"%s",res)
    } else {
        content := models.GetContentTypeExtendedByNodeId(nodeId)
        res, err := json.Marshal(content)
        helpers.PanicIf(err)

        fmt.Fprintf(w,"%s",res)
    }

    
}

// func (this *ContentTypeController) GetContentTypeByNodeId(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     var nodeIdStr string = ""
//     nodeIdStr = r.URL.Query().Get(":nodeId")
    

//     if(len(nodeIdStr)>0){
//         fmt.Println("lol1")
//         nodeId, _ := strconv.Atoi(nodeIdStr)
//         content := models.GetContentTypeByNodeId(nodeId)
//         res, err := json.Marshal(content)
//         helpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }else{
//         fmt.Println("lol2")
//         contentTypes := models.GetContentTypes()
//         res, err := json.Marshal(contentTypes)
//         helpers.PanicIf(err)

//         fmt.Fprintf(w,"%s",res)
//     }
// }

func (this *ContentTypeController) PutContentType(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    contentType := models.ContentType{}

    err := json.NewDecoder(r.Body).Decode(&contentType)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    // b, err := json.Marshal(contentType)
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    

    contentType.Update()
}