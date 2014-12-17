package main

import (
	//"fmt"
	"flag" // https://gobyexample.com/command-line-flags
	"net"
	"log"
	"html/template"
	"io/ioutil"
	"net/http"
    _ "net/url"
    _ "database/sql"
    _ "github.com/lib/pq"
    _ "time"
    "publish/admin/controllers"
    "publish/admin/helpers"
    "publish/admin/models"
    //"encoding/json"
    "github.com/bmizerany/pat"
    //"os"
    //"path/filepath"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
    cc := controllers.ContentController{}
    content := models.Content{}
    cc.RenderTemplate(w, "admin.tmpl", &content)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    cc := controllers.ContentController{}
    cc.RenderTemplate(w, "index.tmpl", nil)
}

// func init is run before main. All packages can have an init function
// func init(){}


func main() {
	// After all flags are defined, call flag.parse() to parse the command line into the defined flags. 
	flag.Parse()

    // program will exit when the main() function finishes. This is likely to happen before your goroutine has time to run and print its output
    // fmt.Println("testing123")
    
    if helpers.Templates == nil {
        helpers.Templates = make(map[string]*template.Template)
    }
    helpers.Templates["admin.tmpl"] = template.Must(template.ParseFiles("admin/views/includes/admin.tmpl", "admin/views/layouts/base.tmpl"))
    // helpers.Templates["Home.tmpl"] = template.Must(template.ParseFiles("views/Home.tmpl", "views/Layout.tmpl"))
    // templates["Page.tmpl"] = template.Must(template.ParseFiles("views/Page.tmpl", "views/base.tmpl"))

    m := pat.New()
    
    // Controllers
    nodeController := controllers.NodeController{}
    contentController := controllers.ContentController{}
    contentTypeController := controllers.ContentTypeController{}
    dataTypeController := controllers.DataTypeController{}
    templateController := controllers.TemplateController{}
    directoryController := controllers.DirectoryController{}
    

    // Entity routes

    // m.Get("/api/entity/:nodeId") ?node-type=2&section=myplugin ???????????? l8r
    m.Post("/api/content/:nodeId", http.HandlerFunc(contentController.Post))
    m.Get("/api/content/:nodeId", http.HandlerFunc(contentController.GetBackendContentByNodeId))
    m.Put("/api/content/:nodeId", http.HandlerFunc(contentController.PutContent))
    m.Get("/api/media/:nodeId", http.HandlerFunc(contentController.GetBackendContentByNodeId))

    // m.Get("/api/content-type/extended/:nodeId", http.HandlerFunc(contentTypeController.GetContentTypeExtendedByNodeId))
    m.Get("/api/content-type/:nodeId", http.HandlerFunc(contentTypeController.GetContentTypeByNodeId))
    //m.Get("/api/content-type/", http.HandlerFunc(contentTypeController.GetContentTypeByNodeId)) // not sure about this
    //m.Get("/api/content-type/", http.HandlerFunc(contentTypeController.GetContentTypes)) // not sure about this
    m.Put("/api/content-type/:nodeId", http.HandlerFunc(contentTypeController.PutContentType))
    m.Post("/api/content-type/:nodeId", http.HandlerFunc(contentTypeController.Post))
    // m.Put("/api/content-type/:nodeId", http.HandlerFunc(contentTypeController.PutContentType))
    // m.Post("/api/content-type", http.HandlerFunc(contentTypeController.PostContentType))
    // m.Delete("/api/content-type/nodeId", http.HandlerFunc(contentTypeController.DeleteContentType))

    m.Get("/api/data-type/:nodeId", http.HandlerFunc(dataTypeController.GetDataTypeByNodeId))
    m.Get("/api/data-type/", http.HandlerFunc(dataTypeController.GetDataTypes)) // not sure about this
    m.Put("/api/data-type/:nodeId", http.HandlerFunc(dataTypeController.PutDataType))
    m.Post("/api/data-type/:nodeId", http.HandlerFunc(dataTypeController.Post))
    // m.Post("/api/data-type", http.HandlerFunc(dataTypeController.PostDataType))
    // m.Delete("/api/data-type/nodeId", http.HandlerFunc(dataTypeController.DeleteDataType))

    m.Get("/api/template/:nodeId", http.HandlerFunc(templateController.GetTemplateByNodeId))
    m.Get("/api/template/", http.HandlerFunc(templateController.GetTemplates)) // not sure about this
    m.Put("/api/template/:nodeId", http.HandlerFunc(templateController.PutTemplate))
    m.Post("/api/template/:nodeId", http.HandlerFunc(templateController.PostTemplate))
    // m.Put("/api/template/:nodeId", http.HandlerFunc(templateController.PutTemplate))
    // m.Post("/api/template", http.HandlerFunc(templateController.PostTemplate))
    // m.Delete("/api/template/nodeId", http.HandlerFunc(templateController.DeleteTemplate))

    // TODO:::

    // script
    // stylesheet

    // Node routes

    // m.Get("/api/node/node-type/:nodeTypeId/content-type/:contentTypeId", http.HandlerFunc(nodeController.GetContentNodesOfContentType)) // not important now
    

    // MAYBE NOT SO GOOD AN IDEA TO GET TOO SPECIFIC HERE... should use query parameters for filtering instead!!!
    //######### m.Get("/api/node/node-type/:nodeTypeId/path-level/:levels", http.HandlerFunc(nodeController.GetNodesByPathLevels)) // WITH THIS
    // m.Get("/api/node/node-type/:nodeTypeId", http.HandlerFunc(nodeController.GetNodesOfType)) // this is used for all other than content nodes

    // REPLACED ROUTES
    // /////m.Get("/api/node/content/type/:id", http.HandlerFunc(nodeController.GetContentNodesOfType)) // GETS REPLACED (by /api/node/node-type/1/content-type/:contentTypeId)
    ////////m.Get("/api/node/media", http.HandlerFunc(nodeController.GetMediaNodes)) // GETS REPLACED (by /api/node/node-type/2/content-type/:contentTypeId)
    
    


    //######### m.Get("/api/node/content-type", http.HandlerFunc(nodeController.GetContentTypeNodes)) // GETS REPLACED (with /api/node/node-type/4)
    

    // Media type nodes
    //######### m.Get("/api/node/media-type", http.HandlerFunc(nodeController.GetMediaTypeNodes)) // GETS REPLACED (with /api/node/node-type/7)

    //######### m.Get("/api/node/data-type", http.HandlerFunc(nodeController.GetDataTypeNodes)) // GETS REPLACED (with /api/node/node-type/11)
    
    //######### m.Get("/api/node/template", http.HandlerFunc(nodeController.GetTemplateNodes)) // GETS REPLACED (with /api/node/node-type/3)
    //m.Get("/api/node/template/all", http.HandlerFunc(nodeController.GetAllTemplateNodes))
    

    // Generic node routes

    m.Get("/api/node", http.HandlerFunc(nodeController.GetNodes))
    m.Get("/api/node/:id", http.HandlerFunc(nodeController.GetNodeById))
    m.Get("/api/node/:id/children", http.HandlerFunc(nodeController.GetNodeByIdChildren))


    m.Post("/api/directory/upload-file-test", http.HandlerFunc(directoryController.UploadFileTest))
    m.Post("/api/directory/:rootdir/:name", http.HandlerFunc(directoryController.Post))
    m.Put("/api/directory/:rootdir/:name", http.HandlerFunc(directoryController.Put))

    m.Get("/api/directory/:rootdir", http.HandlerFunc(directoryController.Get))
    m.Get("/api/directory/:rootdir/:name", http.HandlerFunc(directoryController.GetById))

    m.Post("/api/node", http.HandlerFunc(nodeController.Post))
    m.Put("/api/node/:id", http.HandlerFunc(nodeController.Put))
    m.Del("/api/node/:id", http.HandlerFunc(nodeController.Delete))

    



    //user
    userController := controllers.UserController{}
    m.Post("/api/user/login", http.HandlerFunc(userController.Login))
    m.Get("/api/user", http.HandlerFunc(userController.Get))
    m.Post("/api/user", http.HandlerFunc(userController.Post))


    // m.Get("/section/:id", http.HandlerFunc(HelloServer))

    // Frontend rendering routes

    
    
	http.HandleFunc("/admin/", adminHandler)
    
    http.HandleFunc("/index/", indexHandler)

    m.Get("/:nodeId", http.HandlerFunc(contentController.RenderContent))

    //http://localhost:18019/umbraco/backoffice/UmbracoTrees/DataTypeTree/GetNodes?id=-1&application=developer&tree=&isDialog=false
    //http://localhost:18019/umbraco/backoffice/UmbracoApi/DataType/GetById?id=-43
    //http://localhost:18019/umbraco/backoffice/UmbracoTrees/LegacyTree/GetNodes?rnd=b2e67db4ea0d447b82ab5eeab24aaf97&id=-1&treeType=users&contextMenu=true&isDialog=false

    // https://gobyexample.com/url-parsing
    // get sections
    
    // s := controllers.SectionController{}
    // http.HandleFunc("/api/section", s.Get)
    // [
    //     {"name":"Content","cssclass":"traycontent","alias":"content"},
    //     {"name":"Media","cssclass":"traymedia","alias":"media"},
    //     {"name":"Settings","cssclass":"traysettings","alias":"settings"},
    //     {"name":"Developer","cssclass":"traydeveloper","alias":"developer"},
    //     {"name":"Users","cssclass":"trayuser","alias":"users"},
    //     {"name":"Members","cssclass":"traymember","alias":"member"},
    //     {"name":"Tea Commerce","cssclass":"icon-shopping-basket-alt-2","alias":"teacommerce"}
    // ]

    // ?section=content&id=xxxx (er den virkelig nødvendig/smart eller overflødig?)
    // http,HandleFunc("/api/tree/getnodes")

    // ?section=content
    // ?section=media
    // ?section=settings/developer iscontainer=true
    // ?section=users iscontainer=true
    

    //m.PathPrefix("/public").Handler(http.FileServer(http.Dir("./public/")))
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./admin/public"))))
    http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir("./media"))))
    log.Println("Registered a handler for static files.")
    http.Handle("/", m)


    // http.HandleFunc("/api/Media/GetChildren?id=-1&pageNumber=0&pageSize=0&orderBy=SortOrder&orderDirection=Ascending&filter=")
    // http.handleFunc("/api/Media/GetById?id=1088")

    // ?id=xx&type=xx
    // http.HandleFunc("/api/Entity/GetAncestors?id=1078&type=document", treeHandler)


    if *addr {
        l, err := net.Listen("tcp", "127.0.0.1:0")
        if err != nil {
            log.Fatal(err)
        }
        err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
        if err != nil {
            log.Fatal(err)
        }
        s := &http.Server{}
        s.Serve(l)
        return
    }

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}