# go-server
Basic golang server for HTML, JSON, XML, CSV, and XLS 

To use:
```
go get github.com/sgterban/go-server
go install github.com/sgterban/go-server
$GOPATH/go-server
```

Server will run on **localhost:7777** at the following endpoints:

**/** = **HTML** Template Rendering<br/>
/**json** = **JSON** GET Response<br/>
/**xml** = **XML** GET Response<br/>
/**csv** = **CSV** File Response (**Required** Form value **invoice** e.g. **?invoice=1**)<br/>
/**xls** = **XLS** File Response (**Required** Form value **invoice** e.g. **?invoice=1**)<br/>


<h2>HTML Template</h2>
*http://localhost:7777/*
```html
<html>
  <head>
    <title>Terban.com</title>
  </head>
  <body>
    <h1>Welcome.</h1>
  </body>
</html>
```

<h2>JSON Response</h2>
*http://localhost:7777/json*
```js
{"Msg":"hello"}
```

<h2>XML Response</h2>
*http://localhost:7777/xml*
```xml
<Reponse>
  <Msg>hello</Msg>
</Response>
```

<h2>CSV File Download</h2>
*http://localhost:7777/csv*
```
Description,Price,Quantity
```

<h2>XLS File Download</h2>
*http://localhost:7777/xls*
```xml
<?xml version="1.0" encoding="UTF-8"?>
<?mso-application prodig="Excel.Sheet"?>
<Workbook xmlns="urn:schemas-microsoft-com:office:spreadsheet" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:x="urn:schemas-microsoft-com:office:excel" xmlns:ss="urn:schemas-microsoft-com:office:spreadsheet" xmlns:html="http://www.w3.org/TR/REC-html40">
 <Styles>
  <Style ss:ID="Default" ss:Name="Normal">
   <Alignment ss:Vertical="Bottom"></Alignment>
   <Borders></Borders>
   <Font ss:FontName="Calibri" x:Family="Swiss" ss:Size="11" ss:Color="#000000"></Font>
   <Interior></Interior>
   <NumberFormat></NumberFormat>
   <Protection></Protection>
  </Style>
 </Styles>
 <Worksheet ss:Name="Sheet1">
  <Table>
   <Row>
    <Cell>
     <Data ss:Type="String">Description</Data>
    </Cell>
    <Cell>
     <Data ss:Type="String">Price</Data>
    </Cell>
    <Cell>
     <Data ss:Type="String">Quantity</Data>
    </Cell>
   </Row>
  </Table>
 </Worksheet>
</Workbook>
```
