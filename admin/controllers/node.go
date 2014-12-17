package controllers

import (
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    "publish/admin/helpers"
    "publish/admin/models"
    "strconv"
    "github.com/gorilla/schema"
    "encoding/json"
    //"strings"
)

type NodeController struct {
    Controller
}

func (this *NodeController) GetNodes(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    queryStrParams := r.URL.Query()

    nodes := models.GetNodes(queryStrParams)

    res, err := json.Marshal(nodes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *NodeController) GetNodeById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    idStr := r.URL.Query().Get(":id")
    id, _ := strconv.Atoi(idStr)

    node := models.GetNodeById(id)

    res, err := json.Marshal(node)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *NodeController) GetNodeByIdChildren(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    idStr := r.URL.Query().Get(":id")
    id, _ := strconv.Atoi(idStr)

    nodes := models.GetNodeByIdChildren(id)

    res, err := json.Marshal(nodes)
    helpers.PanicIf(err)

    fmt.Fprintf(w,"%s",res)
}

func (this *NodeController) Post(w http.ResponseWriter, r *http.Request) {
    path := r.FormValue("path")
    created_by,err := strconv.Atoi(r.FormValue("created_by"))
    name := r.FormValue("name")
    node_type,err := strconv.Atoi(r.FormValue("node_type"))



    

    // body, err := ioutil.ReadAll(r.Body)
    // helpers.PanicIf(err)
    // log.Println(string(body))
    // var t test_struct
    // err = json.Unmarshal(body, &t)
    // helpers.PanicIf(err)
    // log.Println(t.path)

    // decoder := json.NewDecoder(r.Body)
    // var t test_struct   
    // err := decoder.Decode(&t)
    // helpers.PanicIf(err)
    // log.Println(t)

    //body, err := ioutil.ReadAll(r.Body)
    //helpers.PanicIf(err)
    //log.Println(r.Body)
    //id := r.URL.Query().Get(":id")
    db := helpers.Db
    //var created_date time.Time = time.Now()

    //t := test_struct{path, created_by, name,node_type,created_date}
    //log.Println(t)
    // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
    querystr := fmt.Sprintf("INSERT INTO node (path, created_by, name, node_type) VALUES ('%s', %d, '%s', %d)", path, created_by, name, node_type)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)
    // JSON(w, r.Body)
}

func (this *NodeController) Put(w http.ResponseWriter, r *http.Request) {
    // path := r.FormValue("path")
    // created_by,err := strconv.Atoi(r.FormValue("created_by"))
    // name := r.FormValue("name")
    // node_type,err := strconv.Atoi(r.FormValue("node_type"))


    templol := r.URL.Query().Get(":id")
    rofl,err1 := strconv.Atoi(templol)
    helpers.PanicIf(err1)

    parm_id := rofl

    t := new(models.Node)

    err := r.ParseForm()

    helpers.PanicIf(err)

    decoder := schema.NewDecoder()
    // r.PostForm is a map of our POST form values
    decoder.Decode(t, r.PostForm)


fmt.Println(r.PostForm)
fmt.Println(t.Path)


    // Do something with person.Name or person.Phone




    // Do something with person.Name or person.Phone

    

    // body, err := ioutil.ReadAll(r.Body)
    // helpers.PanicIf(err)
    // log.Println(string(body))
    // var t test_struct
    // err = json.Unmarshal(body, &t)
    // helpers.PanicIf(err)
    // log.Println(t.path)

    // decoder := json.NewDecoder(r.Body)
    // var t test_struct   
    // err := decoder.Decode(&t)
    // helpers.PanicIf(err)
    // log.Println(t)

    //body, err := ioutil.ReadAll(r.Body)
    //helpers.PanicIf(err)
    //log.Println(r.Body)
    //id := r.URL.Query().Get(":id")
    db := helpers.Db
    //var created_date time.Time = time.Now()

    //t := test_struct{path, created_by, name,node_type,created_date}
    //log.Println(t)
    // http://stackoverflow.com/questions/244243/how-to-reset-postgres-primary-key-sequence-when-it-falls-out-of-sync
    fmt.Println(fmt.Sprintf("path: %s, created_by: %d, name: %s, node type: %d", t.Path, t.CreatedBy, t.Name, t.NodeType))


    querystr := fmt.Sprintf("UPDATE node SET (path, created_by, name, node_type) = ('%s', %d, '%s', %d) WHERE id=%d", t.Path, t.CreatedBy, t.Name, t.NodeType, parm_id)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)
    
    // JSON(w, r.Body)
}

func (this *NodeController) Delete(w http.ResponseWriter, r *http.Request) {
    // path := r.FormValue("path")
    // created_by,err := strconv.Atoi(r.FormValue("created_by"))
    // name := r.FormValue("name")
    // node_type,err := strconv.Atoi(r.FormValue("node_type"))


    templol := r.URL.Query().Get(":id")
    rofl,err1 := strconv.Atoi(templol)
    helpers.PanicIf(err1)

    parm_id := rofl

   
    db := helpers.Db
    
    querystr := fmt.Sprintf("DELETE FROM node WHERE id=%d", parm_id)
    res, err := db.Exec(querystr)
    helpers.PanicIf(err)
    fmt.Println(res)
    
    // JSON(w, r.Body)
}