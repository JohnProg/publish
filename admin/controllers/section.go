package controllers

import (
    "fmt"
    "net/http"
    _ "github.com/lib/pq"
    //"encoding/json"
)

type SectionController struct {
    Controller
}

func (this *SectionController) Get(w http.ResponseWriter, r *http.Request) {
    // this.Data["Username"] = "astaxie"
    // this.Data["Email"] = "astaxie@gmail.com"
    
    
    // jsonStr := {
    //     {"name":"Content","cssclass":"traycontent","alias":"content"},
    //     {"name":"Media","cssclass":"traymedia","alias":"media"},
    //     {"name":"Settings","cssclass":"traysettings","alias":"settings"},
    //     {"name":"Developer","cssclass":"traydeveloper","alias":"developer"},
    //     {"name":"Users","cssclass":"trayuser","alias":"users"},
    //     {"name":"Members","cssclass":"traymember","alias":"member"},
    //     {"name":"Tea Commerce","cssclass":"icon-shopping-basket-alt-2","alias":"teacommerce"}
    // }
    //b := []byte(`{"name":"Content","cssclass":"traycontent","alias":"content"}`)
    
    //this.Data["All"] = b
    // this.TplNames = "index.tpl"

    // return "return from controller"
    str := r.URL.Query().Get(":id")
    fmt.Fprintf(w,"lolollol oloolbol %s: ", str)


}