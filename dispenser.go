package main

import (
   "encoding/json"
   b64 "encoding/base64"
  // "io"
   "io/ioutil"
   "fmt"
)
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
   NavigationName string `json:"navigationData"`
   Active string `json:"active"`
}
type Template struct {
   MetaData MetaData `json:"metaData"`
   NavigationData NavigationData `json:"navigationData"`
   TemplateData TemplateData `json:"templateData"`
}
func create_json() {
   //Create Metadata
   metaData := MetaData{"Arise", "simple"}
   /*metaDataJson, _ := json.Marshal(metaData)
   fmt.Printf("%+v\n", string(metaDataJson))*/

   //Create NavigationData
   navigationData:= NavigationData {"Drops", ""}
   
   /* Create TemplateData
      - load content and head from file
      - convert to Base64
   */
   content, _ := ioutil.ReadFile("public/content.html")
   head, _ := ioutil.ReadFile("public/head.html")
   content64 := b64.StdEncoding.EncodeToString(content)
   head64 := b64.StdEncoding.EncodeToString(head)
   templateData:= TemplateData {"Arise", content64, head64}
   template := Template {metaData, navigationData, templateData}
   templateJson, _ := json.Marshal(template)
   fmt.Printf("%+v\n", string(templateJson))
  // response, _ := http.Post(dispenserURL, "application/json; charset=utf-8", templateJson)
}

