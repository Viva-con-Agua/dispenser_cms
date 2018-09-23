package main

import (
   "encoding/json"
   b64 "encoding/base64"
  // "io"
   "io/ioutil"
   "fmt"
   //"os"
   "net/http"
   "bytes"
)

var dispenserUrl = "http://172.2.100.4:9000/dispenser/template/main/widget"

type MetaData struct {
   MsName string `json:"msName"`
   TemplateName string `json:"templateName"`
}
type TemplateData struct {
   Title string `json:"title"`
   Content string `json:"content"`
   Head string  `json:"head"`
}
type NavigationData struct {
   NavigationName string `json:"navigationName"`
   Active string `json:"active"`
}
type Template struct {
   MetaData MetaData `json:"metaData"`
   NavigationData NavigationData `json:"navigationData"`
   TemplateData TemplateData `json:"templateData"`
}

func loadFile(path string)[]byte {
  if file, err := ioutil.ReadFile(path); err != nil {
    fmt.Println(err)
    return []byte{0}
  }else{
    return file
  }
}

func createJson()(string) {
   
   //Create Metadata
   metaData := MetaData{"Arise", "simple"}
   
   //Create NavigationData
   navigationData:= NavigationData {"Drops", ""}
   
   /* Create TemplateData
      - load content and head from file
      - convert to Base64
   */
   //Load content file and head file
   content := loadFile("public/content.html")
   head := loadFile("public/head.html")
   content64 := b64.StdEncoding.EncodeToString(content)
   head64 := b64.StdEncoding.EncodeToString(head)
   templateData:= TemplateData {"Arise", content64, head64}
   template := Template {metaData, navigationData, templateData}
   templateJson := new(bytes.Buffer)
   json.NewEncoder(templateJson).Encode(template)
   res, err := http.Post(dispenserUrl, "application/json; charset=utf-8", templateJson)
   if err != nil {
      fmt.Println(err)
      return string(loadFile("public/index.html"))
   }else{
      buf := new(bytes.Buffer)
      buf.ReadFrom(res.Body)
      return buf.String()
   }
}

