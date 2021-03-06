package models

import (
  //"fmt"
  "encoding/json"
  "publish/admin/helpers"
  "time"
  "fmt"
  //"net/http"
  "html/template"
  "strconv"
  "log"
  "database/sql"
  "strings"
  "github.com/dgrijalva/jwt-go"
)

type Content struct {
  Id int `json:"id"`
  NodeId int `json:"node_id"`
  ContentTypeNodeId int `json:"content_type_node_id,omitempty"`
  Meta map[string]interface{} `json:"meta,omitempty"`
  Node Node `json:"node,omitempty"`
  ContentType ContentType `json:"content_type,omitempty"`
  Template *Template `json:"template,omitempty"`
  Token *jwt.Token `json:"token,omitempty"`
}

func (c *Content) TemplateFunctionTest(param1 string) template.HTML {
  str := fmt.Sprintf("This is a function inside the content, accessible from the template.<br>The content id is: %d<br>and the parameter value we passed is: %s<br>It gives us a convenient way to fetch additional content, such as the information of the home node - site title, site description social options etc.", c.Id, param1);
  return template.HTML(str)
}

func (t *Content) Post(){


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
  err = tx.QueryRow(`SELECT id, path, created_by, name, node_type, created_date FROM node WHERE id=$1`, t.Node.ParentId).Scan(&id, &path, &created_by, &name, &node_type, &created_date)
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
  err = db.QueryRow(`INSERT INTO node (name, node_type, created_by, parent_id) VALUES ($1, $2, $3, $4) RETURNING id`, t.Node.Name, t.Node.NodeType, 1, t.Node.ParentId).Scan(&node_id)
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

  _, err = tx.Exec("INSERT INTO content (node_id, content_type_node_id, meta) VALUES ($1, $2, $3)", node_id, t.ContentTypeNodeId, meta)
  helpers.PanicIf(err)
  //defer r2.Close()

  if(t.Node.NodeType == 2){
    var fi FileInfo
    var fin FileNode
    if(t.ContentTypeNodeId == 16){
      fi = FileInfo{t.Node.Name, 0, 0777 , time.Now(), true}
      fin = FileNode{t.Meta["path"].(string), "", &fi, nil, "", true, ""}
      fin.Post()
    } 

    // else {
    //   fi = FileInfo{t.Node.Name, 0, 0777 , time.Time.Now(), false}
    //   fin = FileNode{t.Meta.Path, "", fi, nil, "", true, ""}
    // }
  }
  err1 := tx.Commit()
  helpers.PanicIf(err1)

  // // res, _ := json.Marshal(c)
  // // log.Println(string(res))

  // db := helpers.Db

  // meta, _ := json.Marshal(c.Meta)

  // tx, err := db.Begin()
  // helpers.PanicIf(err)
  // //defer tx.Rollback()

  // _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", c.Node.Name, c.Node.Id)
  // helpers.PanicIf(err)
  // //defer r1.Close()

  // _, err = tx.Exec(`UPDATE content 
  //   SET meta = $1 
  //   WHERE node_id = $2`, meta, c.Node.Id)
  // helpers.PanicIf(err)
  // //defer r2.Close()

  // tx.Commit()
}

type Lol struct {
  NodeId int64
  NewPath string
}

func (c *Content) Update(){

  // res, _ := json.Marshal(c)
  // log.Println(string(res))

  db := helpers.Db

  meta, _ := json.Marshal(c.Meta)

  tx, err := db.Begin()
  helpers.PanicIf(err)
  //defer tx.Rollback()

  _, err = tx.Exec("UPDATE node SET name = $1 WHERE id = $2", c.Node.Name, c.Node.Id)
  helpers.PanicIf(err)
  //defer r1.Close()

  _, err = tx.Exec(`UPDATE content 
    SET meta = $1 
    WHERE node_id = $2`, meta, c.Node.Id)
  helpers.PanicIf(err)
  //defer r2.Close()
  if(c.Node.NodeType == 2){
    //originalPath := "media\\Another Image Folder"
    //originalNodeName := "Another Image Folder"
    originalNodeName := c.Node.OldName
    fmt.Println("Original Node Name: " + originalNodeName);

    // rename filesystem folder that has this original url (btw make a hidden input field holding the old url) with c.Node.Name
    folderNode := GetFilesystemNodeById("media", originalNodeName)
    folderNode.FullPath = c.Meta["path"].(string)
    //folderNode.OldPath = originalPath
    //folderNode.FullPath = "media\\Another Image Folder1"
    folderNode.Update()
    fmt.Println("TEST ::: TEST ::: ERR (node_id: ")
    fmt.Println(c.Node.Id)

    // if content is of media type: folder
    if(c.ContentTypeNodeId == 16){

      // check if node has children (SQL SELECT QUERY USING LTREE PATH)
      rows, err101 := tx.Query(`SELECT content.node_id as node_id, meta as content_meta 
        FROM content 
        JOIN node 
        ON node.id = content.node_id 
        WHERE node.path <@ '1.` + strconv.Itoa(c.Node.Id) + `' AND node.path != '1.` + strconv.Itoa(c.Node.Id) + `'`)
      //, strconv.Itoa(c.Node.Id), strconv.Itoa(c.Node.Id)
      // if has children, iterate them
      if err101 != nil {
        log.Fatal(err101)
      }
      defer rows.Close()
      var res []Lol
      // foreach child node
      fmt.Println("TEST ::: TEST ::: ERR1")
      for rows.Next() {
        fmt.Println("TEST ::: TEST ::: ERR2")
        var node_id int64
        var content_meta_byte_arr []byte

        if err := rows.Scan(&node_id, &content_meta_byte_arr); err != nil {
          log.Fatal(err)
        }

        var content_meta map[string]interface{}
        json.Unmarshal([]byte(string(content_meta_byte_arr)), &content_meta)

        var path string = content_meta["path"].(string)
        var newPath string = strings.Replace(path, folderNode.OldPath, folderNode.FullPath, -1)
        // update node's content.meta.url part where substing equals oldurl - with c.Meta.url
        fmt.Println("TEST ::: TEST ::: ERR3")

        res = append(res,Lol{node_id, newPath})
        // _, err102 := tx.Exec(`UPDATE content 
        //   SET meta = json_object_update_key(meta::json, 'url', '$1'::text)::jsonb 
        //   WHERE node_id=$2`, newUrl, node_id)
        // helpers.PanicIf(err102)
      }
      if err101 := rows.Err(); err101 != nil {
        log.Fatal(err101)
      }
      fmt.Println("TEST ::: TEST ::: ERR4")
      for i := 0; i < len(res); i++ {
        fmt.Println(fmt.Sprintf("newpath: %s, node id: %v", res[i].NewPath, res[i].NodeId))
        _, err102 := tx.Exec(`UPDATE content 
          SET meta = json_object_update_key(meta::json, 'path', $1::text)::jsonb 
          WHERE node_id=$2`, string(res[i].NewPath), res[i].NodeId)
        helpers.PanicIf(err102)
      }
      

      
      
      
      
      
    }
  }

  tx.Commit()
}

func GetBackendContentByNodeId(nodeId int) (content Content){
  db := helpers.Db
  queryStr := `SELECT content_node.id AS node_id, content_node.path AS node_path, content_node.created_by AS node_created_by, content_node.name AS node_name, content_node.node_type AS node_type, content_node.created_date AS node_created_date, content_node.parent_id AS node_parent_id,
  content.id AS content_id, content.node_id AS content_node_id, content.content_type_node_id AS content_content_type_node_id, content.meta AS content_meta,
  modified_content_type.id AS ct_id, modified_content_type.node_id AS ct_node_id, modified_content_type.parent_content_type_node_id AS ct_parent_content_type_node_id, modified_content_type.alias AS ct_alias,
  modified_content_type.description AS ct_description, modified_content_type.icon AS ct_icon, modified_content_type.thumbnail AS ct_thumbnail, modified_content_type.meta::json AS ct_meta, modified_content_type.ct_tabs AS ct_tabs, modified_content_type.parent_content_types AS ct_parent_content_types
FROM content
JOIN node AS content_node 
ON content_node.id = content.node_id
JOIN
LATERAL
(
  SELECT ct.*,pct.*,ct_tabs_with_dt.*
  FROM content_type AS ct, node AS ct_node,
  -- Parent content types
  LATERAL 
  (
    SELECT array_to_json(array_agg(res1)) AS parent_content_types
    FROM 
    (
      SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.thumbnail, c.parent_content_type_node_id, c.meta, gf.* AS tabs
      FROM content_type AS c, node,
      LATERAL 
      (
        SELECT json_agg(row1) AS tabs 
        FROM 
        (
          SELECT y.name, ss.properties
          FROM json_to_recordset (
            (
              SELECT * 
              FROM json_to_recordset(
                (
                  SELECT json_agg(ggg)
                  FROM 
                  (
                    SELECT ct.tabs
                    FROM content_type AS ct
                    WHERE ct.id=c.id
                  )ggg
                )
              ) AS x(tabs json)
            )
          ) AS y(name text, properties json),
          LATERAL 
          (
            SELECT json_agg
            (
              json_build_object
              (
                'name',row.name,
                'order',row."order",
                'data_type_node_id',row.data_type_node_id,
                'data_type', json_build_object
                  (
                    'id',row.data_type_id, 
                    'node_id',row.data_type_node_id, 
                    'alias', row.data_type_alias,'html', 
                    row.data_type_html
                  ), 
                'help_text', row.help_text, 
                'description', row.description
              )
            ) AS properties
            FROM 
            (
              SELECT name, "order", data_type.id AS data_type_id, data_type_node_id, data_type.alias AS data_type_alias, data_type.html AS data_type_html, help_text, description
              FROM json_to_recordset(properties) 
              AS k(name text, "order" int, data_type_node_id int, help_text text, description text)
              JOIN data_type
              ON data_type.node_id = k.data_type_node_id
            )row
          ) ss
        )row1
      ) gf
      where path @> subpath(ct_node.path,0,nlevel(ct_node.path)-1) and c.node_id = node.id
    )res1
  ) pct,
  -- Tabs
  LATERAL 
  (
    SELECT res2.tabs AS ct_tabs
    FROM 
    (
      SELECT c.id AS cid, gf.* AS tabs
      FROM content_type AS c, node,
      LATERAL 
      (
        SELECT json_agg(row1) AS tabs 
        FROM
        (
          SELECT y.name, ss.properties
          FROM json_to_recordset
          (
            (
              SELECT * 
              FROM json_to_recordset(
                (
                  SELECT json_agg(ggg)
                  FROM
                  (
                    SELECT ct.tabs
                    FROM content_type AS ct
                    WHERE ct.id=c.id
                  )ggg
                )
              ) AS x(tabs json)
            )
          ) AS y(name text, properties json),
          LATERAL 
          (
            SELECT json_agg
            (
              json_build_object
              (
                'name',row.name,
                'order',row."order",
                'data_type_node_id', row.data_type_node_id,
                'data_type', json_build_object
                  (
                    'id',row.data_type_id, 
                    'node_id', row.data_type_node_id, 
                    'alias', row.data_type_alias, 
                    'html', row.data_type_html
                  ), 
                'help_text', row.help_text, 
                'description', row.description
              )
            ) AS properties
            FROM
            (
              SELECT name, "order", data_type.id AS data_type_id, data_type_node_id, data_type.alias AS data_type_alias, data_type.html AS data_type_html, help_text, description
              FROM json_to_recordset(properties) 
              AS k(name text, "order" int, data_type_node_id int, help_text text, description text)
              JOIN data_type
              ON data_type.node_id = k.data_type_node_id
            )row
          ) ss
        )row1
      ) gf
      WHERE c.id = ct.id
    )res2
    limit 1
  ) ct_tabs_with_dt
  --
  WHERE ct_node.id = ct.node_id
) modified_content_type
ON modified_content_type.node_id = content.content_type_node_id
WHERE content_node.id=$1`
  // queryStr :=
  // `SELECT my_node.id as node_id, my_node.path as node_path, my_node.created_by as node_created_by, my_node.name as node_name, my_node.node_type as node_type, my_node.created_date as node_created_date, my_node.parent_id as node_parent_id,
  //   content.id as content_id, content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
  //   res.id as ct_id, res.node_id as ct_node_id, res.parent_content_type_node_id as ct_parent_content_type_node_id, res.alias as ct_alias,
  //   res.description as ct_description, res.icon as ct_icon, res.thumbnail as ct_thumbnail, res.meta::json as ct_meta, res.ct_tabs as ct_tabs, res.parent_content_types as ct_parent_content_types
  //   FROM content
  //   JOIN node as my_node 
  //   ON my_node.id = content.node_id
  //   JOIN
  //   LATERAL
  //   (
  //     SELECT my_content_type.*,ffgd.*,gf2.*
  //     FROM content_type as my_content_type, node as my_content_type_node,
  //     LATERAL 
  //     (
  //         SELECT array_to_json(array_agg(okidoki)) as parent_content_types
  //         FROM (
  //           SELECT c.id, c.node_id, c.alias, c.description, c.icon, c.thumbnail, c.parent_content_type_node_id, c.meta, gf.* as tabs
  //           FROM content_type as c, node,
  //         LATERAL (
  //             select json_agg(row1) as tabs from((
  //             select y.name, ss.properties
  //             from json_to_recordset(
  //           (
  //               select * 
  //               from json_to_recordset(
  //             (
  //                 SELECT json_agg(ggg)
  //                 from(
  //               SELECT tabs
  //               FROM 
  //               (   
  //                   SELECT *
  //                   FROM content_type as ct
  //                   WHERE ct.id=c.id
  //               ) dsfds

  //                 )ggg
  //             )
  //               ) as x(tabs json)
  //           )
  //             ) as y(name text, properties json),
  //             LATERAL (
  //           select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id',row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id',row.data_type_node_id, 'alias', row.data_type_alias,'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
  //           from(
  //               select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
  //               from json_to_recordset(properties) 
  //               as k(name text, "order" int, data_type_node_id int, help_text text, description text)
  //               JOIN data_type
  //               ON data_type.node_id = k.data_type_node_id
  //               )row
  //             ) ss
  //             ))row1
  //         ) gf
  //           where path @> subpath(my_content_type_node.path,0,nlevel(my_content_type_node.path)-1) and c.node_id = node.id
  //         )okidoki
  //     ) ffgd,
  //     --
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
  //       select json_agg(json_build_object('name',row.name,'order',row."order",'data_type_node_id', row.data_type_node_id,'data_type', json_build_object('id',row.data_type_id, 'node_id', row.data_type_node_id, 'alias', row.data_type_alias, 'html', row.data_type_html), 'help_text', row.help_text, 'description', row.description)) as properties
  //       from(
  //     select name, "order", data_type.id as data_type_id, data_type_node_id, data_type.alias as data_type_alias, data_type.html as data_type_html, help_text, description
  //     from json_to_recordset(properties) 
  //     as k(name text, "order" int, data_type_node_id int, help_text text, description text)
  //     JOIN data_type
  //     ON data_type.node_id = k.data_type_node_id
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
  //   ON res.node_id = content.content_type_node_id
  //   WHERE my_node.id=$1`
  

  var node_id, node_created_by, node_type int
  var node_path, node_name string
  var node_created_date time.Time
  var node_parent_id sql.NullString

  var content_id, content_node_id, content_content_type_node_id int
  var content_meta []byte
  

  var ct_id, ct_node_id int
  var ct_parent_content_type_node_id sql.NullString

  var ct_alias, ct_description, ct_icon, ct_thumbnail string
  var ct_tabs, ct_meta []byte
  var ct_parent_content_types []byte

  row := db.QueryRow(queryStr, nodeId)

  err:= row.Scan(
      &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date, &node_parent_id,
      &content_id, &content_node_id, &content_content_type_node_id, &content_meta,
      &ct_id, &ct_node_id, &ct_parent_content_type_node_id, &ct_alias, &ct_description, &ct_icon, &ct_thumbnail, &ct_meta, &ct_tabs, &ct_parent_content_types)

  helpers.PanicIf(err)

  var content_type_parent_node_id int
  if ct_parent_content_type_node_id.Valid {
    // use s.String
    id, _ := strconv.Atoi(ct_parent_content_type_node_id.String)
    content_type_parent_node_id = id
  } else {
     // NULL value
  }

  var content_parent_node_id int
  if node_parent_id.Valid {
    // use s.String
    id, _ := strconv.Atoi(node_parent_id.String)
    content_parent_node_id = id
  } else {
     // NULL value
  }

  node := Node{node_id, node_path, node_created_by, node_name, node_type, &node_created_date, content_parent_node_id, nil, nil, false, ""}

  var parent_content_types []ContentType
  var tabs []Tab
  var ct_metaMap map[string]interface{}
  var content_metaMap map[string]interface{}

  json.Unmarshal(ct_parent_content_types, &parent_content_types)
  json.Unmarshal(ct_tabs, &tabs)
  json.Unmarshal(ct_meta, &ct_metaMap)
  json.Unmarshal(content_meta, &content_metaMap)

  content_type := ContentType{ct_id, ct_node_id, ct_alias, ct_description, ct_icon, ct_thumbnail, content_type_parent_node_id, tabs, ct_metaMap, parent_content_types, nil}

  content = Content{content_id, content_node_id, content_content_type_node_id, content_metaMap, node, content_type, nil, nil}

  return
}

func GetFrontendContentByNodeId(nodeParamId int) (content Content) {
  db := helpers.Db

  queryStr := `SELECT cn.id AS node_id, cn.path AS node_path, cn.created_by AS node_created_by, cn.name AS node_name, cn.node_type AS node_type, 
  cn.created_date AS node_created_date, cn.parent_id AS node_parent_id,
  content.id AS content_id, content.node_id AS content_node_id, content.content_type_node_id AS content_content_type_node_id, content.meta AS content_meta,
  tpl.parent_template_node_id AS parent_template_node_id, tpl.alias AS template_alias, tpl.partial_template_nodes,
  tn.id AS template_node_id, tn.parent_template_nodes AS parent_template_nodes, tn.name AS template_node_name
FROM content
JOIN node AS cn
ON content.node_id = cn.id
JOIN 
(
  SELECT my_node.*, res1.*
  FROM node AS my_node,
  LATERAL 
  (
    -- SELECT array_to_json(array_agg(node)) AS parent_template_nodes
    SELECT json_agg((SELECT x FROM (SELECT node.id, node.path, node.name, node.node_type, node.created_by, node.parent_id) x)) AS parent_template_nodes
    FROM node
    WHERE path @> subpath(my_node.path,0,nlevel(my_node.path)-1) AND node_type=3 
    ORDER BY my_node.path ASC
  ) res1
  WHERE my_node.node_type = 3
) AS tn
ON (content.meta->>'template_node_id')::int = tn.id
JOIN 
(
  SELECT template.*, res2.* 
  FROM template,
  LATERAL
  (
    SELECT json_agg((SELECT x FROM (SELECT node.id, node.path, node.name, node.node_type, node.created_by, node.parent_id) x)) AS partial_template_nodes
    FROM node
    WHERE node.id = ANY(template.partial_template_node_ids)
    --WHERE node.id IN (SELECT unnest(template.partial_template_node_ids))
    ORDER BY template.node_id ASC
  ) res2 
) AS tpl
ON tpl.node_id = tn.id
WHERE content.node_id=$1`

//   queryStr := `SELECT cn.id as node_id, cn.path as node_path, cn.created_by as node_created_by, cn.name as node_name, cn.node_type as node_type, cn.created_date as node_created_date, cn.parent_id as node_parent_id,
//   content.id as content_id, content.node_id as content_node_id, content.content_type_node_id as content_content_type_node_id, content.meta as content_meta,
//   bummelum.parent_template_node_id as parent_template_node_id, bummelum.alias as template_alias, bummelum.partial_template_nodes,
//   tn.id as template_node_id, tn.parent_template_nodes as parent_template_nodes, tn.name as template_node_name
//   FROM content
//   JOIN node as cn
//   ON content.node_id = cn.id
//   JOIN 
//   (
//   SELECT my_node.*, ffgd.*
//   from node as my_node,
//   LATERAL 
//   (
//       --SELECT array_to_json(array_agg(node)) as parent_template_nodes
//       SELECT json_agg((SELECT x FROM (SELECT node.id, node.path, node.name, node.node_type, node.created_by, node.parent_id) x)) as parent_template_nodes
//       from node
//       where path @> subpath(my_node.path,0,nlevel(my_node.path)-1) and node_type=3 
//       order by my_node.path asc
//   ) ffgd
//   where my_node.node_type = 3
//   )as tn
//   ON (content.meta->>'template_node_id')::int = tn.id
//   JOIN 
//   LATERAL 
//   (SELECT template.*, rofl.* 
//   FROM template,
//   LATERAL
//     (
//         SELECT json_agg((SELECT x FROM (SELECT node.id, node.path, node.name, node.node_type, node.created_by, node.parent_id) x)) as partial_template_nodes
//         from node
//         where node.id = ANY(template.partial_template_node_ids)
//         order by template.node_id asc
//     ) rofl 
//     --where template.node_id = tn.id
//   )bummelum
//     ON bummelum.node_id = tn.id
//   -- template
// --   ON tn.id = template.node_id
//   WHERE content.node_id=$1`


  // node
  var node_id, node_created_by, node_type int
  var node_path, node_name string
  var node_created_date time.Time
  var node_parent_id sql.NullString

  // content
  var content_id, content_node_id, content_content_type_node_id int
  var content_meta []byte

  // template node
  var template_id, template_node_id int
  var parent_template_nodes []byte
  var template_node_name string
  var template_is_partial bool

  // template
  var parent_template_node_id int
  var template_alias string
  var partial_template_nodes []byte

  // master template
  //var master_template_name string


  row := db.QueryRow(queryStr, nodeParamId)

  row.Scan(
      &node_id, &node_path, &node_created_by, &node_name, &node_type, &node_created_date, &node_parent_id,
      &content_id, &content_node_id, &content_content_type_node_id, &content_meta,
      &parent_template_node_id, &template_alias, &partial_template_nodes,
      &template_node_id, &parent_template_nodes, &template_node_name)

  /* THIS IS IMPORTANT TO ACTIVATE AGAIN AT SOME POINT AND HANDLE ALL NULLS PROPERLY!!! */
  //helpers.PanicIf(err)

  var content_parent_node_id int
  if node_parent_id.Valid {
    // use s.String
    id, _ := strconv.Atoi(node_parent_id.String)
    content_parent_node_id = id
  } else {
     // NULL value
  }

  var parent_template_nodes_final []Node
  var partial_template_nodes_slice []Node
  var meta map[string]interface{}

  json.Unmarshal(parent_template_nodes, &parent_template_nodes_final)
  json.Unmarshal(content_meta, &meta)
  json.Unmarshal(partial_template_nodes, &partial_template_nodes_slice)
  //helpers.PanicIf(myerr)

  //fmt.Println("TEST::: BEGIN ::: ")
  fmt.Println(string(partial_template_nodes))
  //fmt.Println("THIS IS::: WEIRD!!!! ::: ")
  fmt.Println(partial_template_nodes_slice)
  //fmt.Println("TEST::: END :::")

  contentNode := Node{node_id, node_path, node_created_by, node_name, node_type, &node_created_date, content_parent_node_id, nil, nil, false, ""}
  templateNode := Node{template_node_id," ",0, template_node_name,0,&time.Time{}, 0, parent_template_nodes_final, nil, false, ""}
  template := Template{template_id, template_node_id, template_alias, parent_template_node_id, "", nil, partial_template_nodes_slice, nil, template_is_partial, &templateNode}
  //templateNode := Node{template_node_id," ",0, template_node_name,0,time.Time{},parent_template_nodes_final, nil, false}
  //template := &Template{}
  content = Content{content_id, content_node_id, content_content_type_node_id, meta, contentNode, ContentType{}, &template, nil}

  return

  // var jsonString string = ""
  // if(template_partial_templates == nil && parent_template_nodes == nil){
  //     jsonString = fmt.Sprintf(`{"node_id":%d, "node_path": "%s", "node_created_by":%d, "node_name": "%s", "node_type":%d, "node_created_date": "%s", "content_id":%d, "content_node_id":%d, "content_meta":%v, "template_node_id":%d, "parent_template_node_id":%d, "template_alias": "%s"}`, node_id, node_path, node_created_by, node_name, node_type, node_created_date, content_id, content_node_id, string(content_meta), template_node_id, parent_template_node_id, template_alias)
  // }else if(template_partial_templates == nil){
  //     jsonString = fmt.Sprintf(`{"node_id":%d, "node_path": "%s", "node_created_by":%d, "node_name": "%s", "node_type":%d, "node_created_date": "%s", "content_id":%d, "content_node_id":%d, "content_meta":%v, "template_node_id":%d, "parent_template_node_id":%d, "template_alias": "%s", "parent_template_nodes": %v}`, node_id, node_path, node_created_by, node_name, node_type, node_created_date, content_id, content_node_id, string(content_meta), template_node_id, parent_template_node_id, template_alias, string(parent_template_nodes))
  // }else if(parent_template_nodes == nil){
  //     jsonString = fmt.Sprintf(`{"node_id":%d, "node_path": "%s", "node_created_by":%d, "node_name": "%s", "node_type":%d, "node_created_date": "%s", "content_id":%d, "content_node_id":%d, "content_meta":%v, "template_node_id":%d, "parent_template_node_id":%d, "template_alias": "%s", "partial_templates": %v}`, node_id, node_path, node_created_by, node_name, node_type, node_created_date, content_id, content_node_id, string(content_meta), template_node_id, parent_template_node_id, template_alias, string(template_partial_templates))
  // } else{
  //     jsonString = fmt.Sprintf(`{"node_id":%d, "node_path": "%s", "node_created_by":%d, "node_name": "%s", "node_type":%d, "node_created_date": "%s", "content_id":%d, "content_node_id":%d, "content_meta":%v, "template_node_id":%d, "parent_template_node_id":%d, "template_alias": "%s", "partial_templates": %v, "parent_template_nodes": %v}`, node_id, node_path, node_created_by, node_name, node_type, node_created_date, content_id, content_node_id, string(content_meta), template_node_id, parent_template_node_id, template_alias, string(template_partial_templates), string(parent_template_nodes))
  // }

  
  // byt := []byte(jsonString)
  // var data map[string]interface{}
  //helpers.PanicIf(err)
  // switch {
  //     case err == queryStr.ErrNoRows:
  //             log.Printf("No node with that ID.")
  //     case err != nil:
  //             log.Fatal(err)
  //     default:
  //             // fmt.Println("byt tostring: ")
  //             // fmt.Println(string(byt))

  //             // if err := json.Unmarshal(byt, &data); err != nil {
  //             //     fmt.Println("unmarshal error")
  //             //     panic(err)
  //             // }

  //             // fmt.Println("data unmarshal: ")
  //             // fmt.Println(data)
  // }
  //fmt.Println("just before return! ")
  //return data
}


// type content Content

// func (c *Content) UnmarshalJSON(b []byte) (err error) {
// 	j, i, ni := content{}, 0, 0
// 	var m map[string]interface{}

// 	if err = json.Unmarshal(b, &j); err == nil {
// 		*c = Content(j)
// 		return
// 	}
//   if err = json.Unmarshal(b, &i); err == nil {
//     c.Id = i
//     return
//   }
//   if err = json.Unmarshal(b, &ni); err == nil {
//     c.NodeId = ni
//     return
//   }
// 	// if err = json.Unmarshal(b, &n); err == nil {
// 	// 	d.NodeId = n
// 	// 	return
// 	// }
// 	if err = json.Unmarshal(b, &m); err == nil {
// 		c.Meta = m
// 		return
// 	}
// 	return
// }

// func (c *Content) MarshalJSON() ([]byte, error) {
//     if c.Id != 0 && c.NodeId != 0 {
//         return json.Marshal(map[string]interface{}{
//             "id": c.Id,
//             //"node_id": d.NodeId,
//             "node_id": c.NodeId,
//             "content_type_node_id": c.ContentTypeNodeId,
//             "meta": c.Meta,
//         })
//     }
//     if c.Id != 0 {
//         return json.Marshal(c.Id)
//     }
//     if c.NodeId != 0 {
//         return json.Marshal(c.NodeId)
//     }
//     if c.ContentTypeNodeId != 0 {
//         return json.Marshal(c.ContentTypeNodeId)
//     }
//     return json.Marshal(nil)
// }