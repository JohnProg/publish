package models

import (
    //"fmt"
  "encoding/json"
  "publish/admin/helpers"
  "time"
  "fmt"
  //"net/http"
  //"html/template"
  "strconv"
  "database/sql"
  "log"
)

type ContentType struct {
  Id int `json:"id,omitempty"`
  Node_id int `json:"node_id,omitempty"`
  Alias string `json:"alias,omitempty"`
  Description string `json:"description,omitempty"`
  Icon string `json:"icon,omitempty"`
  Thumbnail string `json:"thumbnail,omitempty"`
  ParentContentTypeNodeId int `json:"parent_content_type_node_id,omitempty"`
  Tabs []Tab `json:"tabs,omitempty"`
  Meta map[string]interface{} `json:"meta,omitempty"`
  ParentContentTypes []ContentType `json:"parent_content_types,omitempty"`
  Node *Node `json:"node,omitempty"`
}

func (t *ContentType) Post(){
  tm, err := json.Marshal(t)
  helpers.PanicIf(err)
  fmt.Println("tm:::: ")
  fmt.Println(string(tm))
  
  db := helpers.Db

  tx, err := db.Begin()
  helpers.PanicIf(err)
  //defer tx.Rollback()
  var parentNode Node
  var id, created_by, node_type int
  var path, name string
  var created_date *time.Time
  err = tx.QueryRow(`SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`, t.ParentContentTypeNodeId).Scan(&id, &path, &created_by, &name, &node_type, &created_date)
  switch {
    case err == sql.ErrNoRows:
      log.Printf("No user with that ID.")
    case err != nil:
      log.Fatal(err)
    default:
      parentNode = Node{id, path, created_by, name, node_type, created_date, 0, nil,nil, false, ""}
      //fmt.Printf("Username is %s\n", username)
  }

  // http://godoc.org/github.com/lib/pq
  // pq does not support the LastInsertId() method of the Result type in database/sql. 
  // To return the identifier of an INSERT (or UPDATE or DELETE), 
  // use the Postgres RETURNING clause with a standard Query or QueryRow call:
  
  var node_id int64
  err = db.QueryRow(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`, t.Node.Name, t.Node.NodeType, 1, t.ParentContentTypeNodeId).Scan(&node_id)
  //res, err := tx.Exec(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4)`, t.Node.Name, 3, 1, t.ParentTemplateNodeId)
  //helpers.PanicIf(err)
  //node_id, err := res.LastInsertId()
  fmt.Println(strconv.FormatInt(node_id, 10))
  if err != nil {
    //log.Println(string(res))
    log.Fatal(err.Error())
  } else {
    _, err = tx.Exec("UPDATE node SET path=$1 WHERE id=$2", parentNode.Path + "." + strconv.FormatInt(node_id, 10), node_id)
    helpers.PanicIf(err)
    //println("LastInsertId:", node_id)
  }
  //defer r1.Close()
  meta, errMeta := json.Marshal(t.Meta)
  helpers.PanicIf(errMeta)

  tabs, errTabs := json.Marshal(t.Tabs)
  helpers.PanicIf(errTabs)

  _, err = tx.Exec("INSERT INTO content_type (node_id, alias, description, icon, thumbnail, parent_content_type_node_id, meta, tabs) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", node_id, t.Alias, t.Description, t.Icon, t.Thumbnail, t.ParentContentTypeNodeId, meta, tabs)
  helpers.PanicIf(err)
  //defer r2.Close()
  err1 := tx.Commit()
  helpers.PanicIf(err1)

}

func GetContentTypes() (contentTypes []ContentType){
  querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
    ct.id as ct_id, ct.node_id as ct_node_id, ct.parent_content_type_node_id as ct_parent_content_type_node_id, ct.alias as ct_alias,
    ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta, res.ct_tabs as ct_tabs
    FROM node
    JOIN content_type as ct
    ON ct.node_id = node.id
    JOIN
    LATERAL
    (
      SELECT my_content_type.*,gf2.*
      FROM content_type as my_content_type, node as my_content_type_node,
      LATERAL 
      (
          SELECT okidoki.tabs as ct_tabs
          FROM (
            SELECT c.id as cid, gf.* as tabs
            FROM content_type as c, node,
          LATERAL (
              select json_agg(row1) as tabs from((
          select y.name, ss.properties
          from json_to_recordset(
          (
        select * 
        from json_to_recordset(
            (
          SELECT json_agg(ggg)
          from(
        SELECT tabs
        FROM 
        (   
            SELECT *
            FROM content_type as ct
            WHERE ct.id=c.id
        ) dsfds

          )ggg
            )
        ) as x(tabs json)
          )
          ) as y(name text, properties json),
          LATERAL (
        select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id',row.data_type, 'alias',row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
        from(
      select name, "order", data_type, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
      from json_to_recordset(properties) 
      as k(name text, "order" int, data_type int, help_text text, description text)
      JOIN data_type
      ON data_type.id = k.data_type
      )row
          ) ss
              ))row1
          ) gf
          WHERE c.id = my_content_type.id
          )okidoki
          limit 1
      ) gf2
      --
      WHERE my_content_type_node.id = my_content_type.node_id
    ) res
    ON res.node_id = ct.node_id
    WHERE node.node_type=4`

    // node
    var node_id, node_created_by, node_type int
    var node_path, node_name string
    var node_created_date time.Time

    var ct_id, ct_node_id int
    var ct_parent_content_type_node_id sql.NullString
    var ct_alias, ct_description, ct_icon, ct_thumbnail string
    var ct_tabs, ct_meta []byte

    db := helpers.Db

    rows, err := db.Query(querystr)
    helpers.PanicIf(err)
    defer rows.Close()

    for rows.Next(){
      err:= rows.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &ct_id, &ct_node_id, &ct_parent_content_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs)

      var parent_content_type_node_id int
      if ct_parent_content_type_node_id.Valid {
      // use s.String
          id, _ := strconv.Atoi(ct_parent_content_type_node_id.String)
          parent_content_type_node_id = id
      } else {
       // NULL value
      }

      ct_tabs_str := string(ct_tabs)
      //fmt.Println(":::::::::::::::::::::::::::::::::::1 ")
      //fmt.Println(ct_tabs_str)

      //fmt.Println(ct_tabs_str + " dsfjldskfj skdf")
      ct_meta_str := string(ct_meta)
      var x map[string]interface{}
      json.Unmarshal([]byte(string(ct_meta_str)), &x)
      //fmt.Println(ct_meta_str + " dsfjldskfj skdf")

      // Decode the json object

      var ctTabs []Tab
      //var tab Tab

      errlol := json.Unmarshal([]byte(ct_tabs_str), &ctTabs)
      helpers.PanicIf(errlol)

      //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
      //lol, _ := json.Marshal(ctTabs)
      //fmt.Println(string(lol))

      //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
      
      //fmt.Println("ksjdflk sdfkj: " + node_name)


      //helpers.PanicIf(err)
      switch {
          case err == sql.ErrNoRows:
                  log.Printf("No node with that ID.")
          case err != nil:
                  log.Fatal(err)
          default:
                  node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, ""}
                  contentType := ContentType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_content_type_node_id, ctTabs, x, nil, &node}
                  contentTypes = append(contentTypes,contentType)
      }
    }

    return
}


func GetContentTypeExtendedByNodeId(nodeId int) (contentType ContentType){

  querystr := `SELECT my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date,
    res.id as ct_id, res.parent_content_type_node_id as ct_parent_content_type_node_id, res.alias as ct_alias,
    res.description as ct_description, res.icon as ct_icon, res.thumbnail as ct_thumbnail, res.meta::json as ct_meta, res.ct_tabs as ct_tabs, res.parent_content_types as ct_parent_content_types
    FROM content_type 
    JOIN node as my_node 
    ON my_node.id = content_type.node_id 
    JOIN
    LATERAL
    (
      SELECT my_content_type.*,ffgd.*,gf2.*
      FROM content_type as my_content_type, node as my_content_type_node,
      LATERAL 
      (
          SELECT array_to_json(array_agg(okidoki)) as parent_content_types
          FROM (
            SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.thumbnail, c.parent_content_type_node_id, c.meta, gf.* as tabs
            FROM content_type as c, node,
          LATERAL (
              select json_agg(row1) as tabs from((
              select y.name, ss.properties
              from json_to_recordset(
            (
                select * 
                from json_to_recordset(
              (
                  SELECT json_agg(ggg)
                  from(
                SELECT tabs
                FROM 
                (   
                    SELECT *
                    FROM content_type as ct
                    WHERE ct.id=c.id
                ) dsfds

                  )ggg
              )
                ) as x(tabs json)
            )
              ) as y(name text, properties json),
              LATERAL (
            select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id',row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id',row.data_type_node_id, 'alias', row.data_type_alias,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
            from(
                select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
                from json_to_recordset(properties) 
                as k(name text, "order" int, data_type_node_id int, help_text text, description text)
                JOIN data_type
                ON data_type.node_id = k.data_type_node_id
                )row
              ) ss
              ))row1
          ) gf
            where path @> subpath(my_content_type_node.path,0,nlevel(my_content_type_node.path)-1) and c.node_id = node.id
          )okidoki
      ) ffgd,
      --
      LATERAL 
      (
          SELECT okidoki.tabs as ct_tabs
          FROM (
            SELECT c.id as cid, gf.* as tabs
            FROM content_type as c, node,
          LATERAL (
              select json_agg(row1) as tabs from((
          select y.name, ss.properties
          from json_to_recordset(
          (
        select * 
        from json_to_recordset(
            (
          SELECT json_agg(ggg)
          from(
        SELECT tabs
        FROM 
        (   
            SELECT *
            FROM content_type as ct
            WHERE ct.id=c.id
        ) dsfds

          )ggg
            )
        ) as x(tabs json)
          )
          ) as y(name text, properties json),
          LATERAL (
        select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id', row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id', row.data_type_node_id, 'alias', row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
        from(
      select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
      from json_to_recordset(properties) 
      as k(name text, "order" int, data_type_node_id int, help_text text, description text)
      JOIN data_type
      ON data_type.node_id = k.data_type_node_id
      )row
          ) ss
              ))row1
          ) gf
          WHERE c.id = my_content_type.id
          )okidoki
          limit 1
      ) gf2
      --
      WHERE my_content_type_node.id = my_content_type.node_id
    ) res
    ON res.node_id = content_type.node_id
    WHERE content_type.node_id=$1`

    // node
    var node_id, node_created_by, node_type int
    var node_path, node_name string
    var node_created_date time.Time

    var ct_id int
    var ct_parent_content_type_node_id sql.NullString

    var ct_alias, ct_description, ct_icon, ct_thumbnail string
    var ct_tabs, ct_meta []byte
    var ct_parent_content_types []byte

    db := helpers.Db

    row := db.QueryRow(querystr, nodeId)

    err:= row.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &ct_id, &ct_parent_content_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_parent_content_types)

    var parent_content_type_node_id int
    if ct_parent_content_type_node_id.Valid {
    // use s.String
        id, _ := strconv.Atoi(ct_parent_content_type_node_id.String)
        parent_content_type_node_id = id
    } else {
     // NULL value
    }

    var parent_content_types []ContentType
    var tabs []Tab
    var ct_metaMap map[string]interface{}


    json.Unmarshal(ct_parent_content_types, &parent_content_types)
    json.Unmarshal(ct_tabs, &tabs)
    json.Unmarshal(ct_meta, &ct_metaMap)

    //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
    //lol, _ := json.Marshal(ctTabs)
    //fmt.Println(string(lol))

    //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
    
    //fmt.Println("ksjdflk sdfkj: " + node_name)


    //helpers.PanicIf(err)
    switch {
        case err == sql.ErrNoRows:
                log.Printf("No node with that ID.")
        case err != nil:
                log.Fatal(err)
        default:
                node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, ""}
                
  

                contentType = ContentType{ct_id, node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_content_type_node_id, tabs, ct_metaMap, parent_content_types, &node}
                //contentType = ContentType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_content_type_node_id, ctTabs, x, nil, &node}
    }

    return
}

func GetContentTypeByNodeId(nodeId int) (contentType ContentType){
  // querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
  //   ct.id as ct_id, ct.node_id as ct_node_id, ct.parent_content_type_node_id as ct_parent_content_type_node_id, ct.alias as ct_alias,
  //   ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta, res.ct_tabs as ct_tabs
  //   FROM node
  //   JOIN content_type as ct
  //   ON ct.node_id = node.id
  //   JOIN
  //   LATERAL
  //   (
  //     SELECT my_content_type.*,gf2.*
  //     FROM content_type as my_content_type, node as my_content_type_node,
  //     LATERAL 
  //     (
  //         SELECT okidoki.tabs as ct_tabs
  //         FROM (
  //           SELECT c.id as cid, gf.* as tabs
  //           FROM content_type as c, node,
  //         LATERAL (
  //             select json_agg(row1) as tabs from((
  //         select y.name, ss.properties
  //         from json_to_recordset(
  //         (
  //       select * 
  //       from json_to_recordset(
  //           (
  //         SELECT json_agg(ggg)
  //         from(
  //       SELECT tabs
  //       FROM 
  //       (   
  //           SELECT *
  //           FROM content_type as ct
  //           WHERE ct.id=c.id
  //       ) dsfds

  //         )ggg
  //           )
  //       ) as x(tabs json)
  //         )
  //         ) as y(name text, properties json),
  //         LATERAL (
  //       select json_agg(json_build_object('name',row.name,'order',row."order",'data_type', json_build_object('id', row.data_type_id,'node_id',row.data_type, 'alias',row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
  //       from(
  //     select name, "order", data_type.id as data_type_id, data_type, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
  //     from json_to_recordset(properties) 
  //     as k(name text, "order" int, data_type int, help_text text, description text)
  //     JOIN data_type
  //     ON data_type.node_id = k.data_type
  //     )row
  //         ) ss
  //             ))row1
  //         ) gf
  //         WHERE c.id = my_content_type.id
  //         )okidoki
  //         limit 1
  //     ) gf2
  //     --
  //     WHERE my_content_type_node.id = my_content_type.node_id
  //   ) res
  //   ON res.node_id = ct.node_id
  //   WHERE node.id=$1`
  querystr := `SELECT node.id as node_id, node.path as node_path, node.created_by as node_created_by, node.name as node_name, node.node_type as node_type, node.created_date as node_created_date,
    ct.id as ct_id, ct.node_id as ct_node_id, ct.parent_content_type_node_id as ct_parent_content_type_node_id, ct.alias as ct_alias,
    ct.description as ct_description, ct.icon as ct_icon, ct.thumbnail as ct_thumbnail, ct.meta::json as ct_meta, ct.tabs as ct_tabs
    FROM node
    JOIN content_type as ct
    ON ct.node_id = node.id
    WHERE node.id=$1`

    // node
    var node_id, node_created_by, node_type int
    var node_path, node_name string
    var node_created_date time.Time

    var ct_id, ct_node_id int
    var ct_parent_content_type_node_id sql.NullString
    var ct_alias, ct_description, ct_icon, ct_thumbnail string
    var ct_tabs, ct_meta []byte

    db := helpers.Db

    row := db.QueryRow(querystr, nodeId)

    err:= row.Scan(
        &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date,
        &ct_id, &ct_node_id, &ct_parent_content_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs)

    var parent_content_type_node_id int
    if ct_parent_content_type_node_id.Valid {
    // use s.String
        id, _ := strconv.Atoi(ct_parent_content_type_node_id.String)
        parent_content_type_node_id = id
    } else {
     // NULL value
    }

    ct_tabs_str := string(ct_tabs)
    //fmt.Println(":::::::::::::::::::::::::::::::::::1 ")
    //fmt.Println(ct_tabs_str)

    //fmt.Println(ct_tabs_str + " dsfjldskfj skdf")
    ct_meta_str := string(ct_meta)
    var x map[string]interface{}
    json.Unmarshal([]byte(string(ct_meta_str)), &x)
    //fmt.Println(ct_meta_str + " dsfjldskfj skdf")

    // Decode the json object

    var ctTabs []Tab
    //var tab Tab

    errlol := json.Unmarshal([]byte(ct_tabs_str), &ctTabs)
    helpers.PanicIf(errlol)

    //fmt.Println(":::::::::::::::::::::::::::::::::::2 ")
    //lol, _ := json.Marshal(ctTabs)
    //fmt.Println(string(lol))

    //fmt.Printf("id: %d, HTML: %s, name: %s", ctTabs[0].Properties[0].DataType.Id, ctTabs[0].Properties[0].DataType.Html, ctTabs[0].Properties[0].Name)
    
    //fmt.Println("ksjdflk sdfkj: " + node_name)


    //helpers.PanicIf(err)
    switch {
        case err == sql.ErrNoRows:
                log.Printf("No node with that ID.")
        case err != nil:
                log.Fatal(err)
        default:
                node := Node{node_id,node_path,node_created_by, node_name, node_type, &node_created_date, 0, nil, nil, false, ""}
                contentType = ContentType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, parent_content_type_node_id, ctTabs, x, nil, &node}
    }

    return
}

// STILL NEEDS SOME WORK REGARDING TABS vs THE DATA TYPE ID/WHOLE OBJECT PROBLEM

func (ct *ContentType) Update(){
  db := helpers.Db

  meta, _ := json.Marshal(ct.Meta)

  tabs, _ := json.Marshal(ct.Tabs)

  var testMapSlice []map[string]interface{}
  err1 := json.Unmarshal(tabs,&testMapSlice)
  helpers.PanicIf(err1)
  
  // //tabs, _ := json.Marshal(ct.Tabs)
  // for i := 0; i < len(testMapSlice); i++ {
  //   var properties []interface{} = testMapSlice[i]["properties"].([]interface{})
  //   for j := 0; j < len(properties); j++ {
  //     //fmt.Println(properties[j])
  //     var property map[string]interface{} = properties[j].(map[string]interface{})
  //     //var dt interface{} = property.data_type
  //     fmt.Println("property begin ---")
  //     fmt.Println(property)
  //     fmt.Println("property end ---\n")
  //     var dt map[string]interface{} = property["data_type"].(map[string]interface{})
  //     fmt.Println(dt)
  //     //property["data_type"] = dt["id"]
  //   }
    
  // }

  res, _ := json.Marshal(testMapSlice)
  log.Println(string(res))

  // //b, _ := json.Marshal(testMap)
  // fmt.Println(testMapSlice)
  // fmt.Println(testMapSlice[0]["name"])
  // fmt.Println(testMapSlice[0]["properties"])
  // //fmt.Println(testMapSlice[name])
  // //fmt.Println(testMapSlice['name'])
  // //fmt.Println(testMapSlice[[`name`])

  tx, err := db.Begin()
  helpers.PanicIf(err)
  //defer tx.Rollback()

  _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", ct.Node.Name, ct.Node.Id)
  helpers.PanicIf(err)
  //defer r1.Close()

  _, err = tx.Exec(`UPDATE content_type 
    SET alias = $1, description = $2, icon = $3, thumbnail = $4, meta = $5, tabs = $6 
    WHERE node_id = $7`, ct.Alias, ct.Description, ct.Icon, ct.Thumbnail, meta, tabs, ct.Node.Id)
  helpers.PanicIf(err)
  //defer r2.Close()

  tx.Commit()
}