package controllers

import (
    "fmt"
    "net/http"
    "publish/admin/helpers"
    "publish/admin/models"
    "strconv"
    "encoding/json"
)

type DataTypeController struct {
    Controller
}

func (this *DataTypeController) Post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    dataType := models.DataType{}

    err := json.NewDecoder(r.Body).Decode(&dataType)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    dataType.Post()

}

func (this *DataTypeController) GetDataTypes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    dataTypes := models.GetDataTypes()

    res, err := json.Marshal(dataTypes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *DataTypeController) GetDataTypeByNodeId(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    nodeIdStr := r.URL.Query().Get(":nodeId")
    nodeId, _ := strconv.Atoi(nodeIdStr)

    dataType := models.GetDataTypeByNodeId(nodeId)

    res, err := json.Marshal(dataType)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)

}

func (this *DataTypeController) PutDataType(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // nodeIdStr := r.URL.Query().Get(":nodeId")
    // nodeId, _ := strconv.Atoi(nodeIdStr)

    dataType := models.DataType{}

    err := json.NewDecoder(r.Body).Decode(&dataType)

    if err != nil {
        http.Error(w, "Bad Request", 400)
    }

    //fmt.Println(dataType)
    dataType.Update()

    // fmt.Println("Id:")
    // fmt.Println(dataType.Id)
    // fmt.Println("\n")
    // fmt.Println("Node Id:")
    // fmt.Println(dataType.NodeId)
    // fmt.Println("\n")
    // fmt.Println("Alias: ")
    // fmt.Println(dataType.Alias)
}