package models

import (
  _ "fmt"
  "net/url"
  //"encoding/json"
  "time"
  "publish/admin/helpers"
)

type Node struct {
  Id int `json:"id,omitempty"`
  Path string `json:"path,omitempty"`
  CreatedBy int `json:"created_by,omitempty"`
  Name string `json:"name,omitempty"`
  NodeType int `json:"node_type,omitempty"`
  CreatedDate *time.Time `json:"created_date,omitempty"`
  ParentId int `json:"parent_id,omitempty"`
  ParentNodes []Node `json:"parent_nodes,omitempty"`
  ChildNodes []Node `json:"child_nodes,omitempty"`
  Show bool `json:"show,omitempty"`
  OldName string `json:"old_name,omitempty"`
}

func GetNodes(queryStringParams url.Values) (nodes []Node){

	db := helpers.Db
  sql := `SELECT id, path, created_by, name, node_type, created_date FROM node`

  // if ?node-type=x&levels=x(,x..)
  // else if ?node-type=x
  // else if ?levels=x(,x..)
  if(queryStringParams.Get("node-type") != "" && queryStringParams.Get("levels") != ""){
      sql = sql + ` WHERE node_type=` + queryStringParams.Get("node-type") + ` and node.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
  } else if(queryStringParams.Get("node-type") != "" && queryStringParams.Get("levels")==""){
      sql = sql + ` WHERE node_type=` + queryStringParams.Get("node-type")
  } else if(queryStringParams.Get("node-type") == "" && queryStringParams.Get("levels") != ""){
      sql = sql + ` WHERE node.path ~ '1.*{`+queryStringParams.Get("levels") +`}'`
  }
  
  //fmt.Println(sql)
  rows, err := db.Query(sql)
  helpers.PanicIf(err)
  defer rows.Close()

  var id, created_by, node_type int
  var path, name string
  var created_date time.Time

  for rows.Next(){
      err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date)
      helpers.PanicIf(err)
      node := Node{id, path, created_by, name, node_type, &created_date, 0, nil,nil,false, ""}
      nodes = append(nodes,node)
  }
  return
}

func GetNodeById(id int) (node Node){
	db := helpers.Db
  querystr := `SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`

  row := db.QueryRow(querystr,id)

  var created_by, node_type int
  var path, name string
  var created_date time.Time

  err := row.Scan(&id, &path, &created_by, &name, &node_type, &created_date)
  helpers.PanicIf(err)
  node = Node{id, path, created_by, name, node_type, &created_date, 0, nil, nil, false, ""}
    
    return
}

func GetNodeByIdChildren(id int) (nodes []Node){
  db := helpers.Db

  querystr := "SELECT id, path, created_by, name, node_type, created_date FROM node WHERE parent_id=$1"

  rows, err := db.Query(querystr, id)
  helpers.PanicIf(err)
  defer rows.Close()

  var created_by, node_type int
  var path, name string
  var created_date time.Time

  //nodes = []models.Node{}

  for rows.Next(){
      err := rows.Scan(&id, &path, &created_by, &name, &node_type, &created_date)
      helpers.PanicIf(err)
      node := Node{id,path,created_by, name, node_type, &created_date, 0, nil, nil, false, ""}

      nodes = append(nodes, node)
      // fmt.Fprintf(w, "Id: %d, Path: %s, Created by: %d, Name: %s, Node type: %d, Created date: %s\n", id, path, created_by, name, node_type, created_date)
  }
  // slice := []int{1,2,3}
  return
}