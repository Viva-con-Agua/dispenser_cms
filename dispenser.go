package main

   
   type MetaData struct {
      msName string
      templateName string 
   }
   type TemplateData struct {
      title string 
      content string 
      head string 
   }
   type NavigationData struct {
      navigationName string
      active string 
   }
   
   type Template struct {
      metaData MetaData 
      navigationData NavigationData 
      templateData TemplateData
   }

   func create_json() {
      
   
   }
