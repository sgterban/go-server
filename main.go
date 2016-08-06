package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"encoding/xml"
	"strings"
	"strconv"
	"html/template"
)


func main() {
	http.HandleFunc("/", renderHTML)
	http.HandleFunc("/json", renderJson)
	http.HandleFunc("/xml", renderXml)	
	http.HandleFunc("/csv", renderCsv)	
	http.HandleFunc("/xls", renderXls)

	fmt.Println("Starting server...")	
	log.Fatal(http.ListenAndServe("0.0.0.0:7777", nil))
}


type Response struct {
	Msg string
}

func renderJson(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Error - only GET supported.", 400)
	}
	resp := Response{"hello"}
	
	x, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(x)

}

func renderXml(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Error - only GET supported", 400)
	}
	resp := Response{"hello"}
	
	x, err := xml.MarshalIndent(resp, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}

type InvoiceItem struct {
	Description string
	Price string
	Quantity int
}

func renderHTML(w http.ResponseWriter, req *http.Request) {
	temp, err := template.New("index").Parse(`<html><head><title>Terban.com</title></head><body><h1>Welcome.</h1></body></html>`)
	if err != nil {
		http.Error(w, "Error processing template", 500)
	}
	if err := temp.Execute(w, nil); err != nil {
		http.Error(w, "Error executing template", 500)
	}
}

func renderCsv(w http.ResponseWriter, req *http.Request) {
	invoice_id := req.FormValue("invoice")
	if invoice_id == "" {
		http.Error(w, "Error - no invoice or invoice id found", 400)
		return
	}
	headers := []string{"Description", "Price", "Quantity"}
	
	data_rows := []string{strings.Join(headers,",")}

	invoice := []InvoiceItem{}
	
	for _, item := range invoice {
		data_rows = append(data_rows, strings.Join([]string{item.Description, item.Price, strconv.Itoa(item.Quantity)}, ",")) 
	}	

	x := []byte(strings.Join(data_rows, "\n"))
	
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=test.csv")
	w.Write(x)
}

func renderXls(w http.ResponseWriter, req *http.Request) {
	invoice_id := req.FormValue("invoice")
	if invoice_id == "" {
		http.Error(w, "Error - no invoice or invoice id found", 400)
		return
	}
	data_rows := []row{}
	
	header_desc := data{SSType:"String", Value:"Description"}
	header_price := data{SSType:"String", Value:"Price"}
	header_quant := data{SSType:"String", Value:"Quantity"}
	header := []cell{cell{Data:header_desc}, cell{Data:header_price}, cell{Data:header_quant}}
	data_rows = append(data_rows, row{Cells:header})

	invoice := []InvoiceItem{}

	for _, item := range invoice {
		desc := data{SSType:"String",Value:item.Description}
		price := data{SSType:"String",Value:item.Price}
		quant := data{SSType:"Number",Value:strconv.Itoa(item.Quantity)}
		
		data_rows = append(data_rows, row{Cells:[]cell{cell{Data:desc}, cell{Data:price}, cell{Data:quant}}}) 
	}	
	table := table{Rows:data_rows}
	worksheet := worksheet{SSName:"Sheet1", Table:table}

	workbook := makeWorkbook(worksheet)
	
	x, err := xml.MarshalIndent(workbook, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	x = []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<?mso-application prodig="Excel.Sheet"?>` + "\n" + string(x))
	
	
	w.Header().Set("Content-Type", "application/vnd.ms-excel")
	w.Header().Set("Content-Disposition", "attachment; filename=test.xls")
	w.Write(x)
}
