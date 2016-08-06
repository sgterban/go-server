package main 

import "encoding/xml"

type workbook struct {
        XMLName xml.Name `xml:"Workbook"`
        Xmlns string `xml:"xmlns,attr"`
        Xmlnso string `xml:"xmlns:o,attr"`
        Xmlnsx string `xml:"xmlns:x,attr"`
        Xmlnss string `xml:"xmlns:ss,attr"`
        Xmlnsh string `xml:"xmlns:html,attr"`
        Styles styles
        Worksheet worksheet
}

type documentproperties struct {
        XMLName xml.Name `xml:"DocumentProperties"`
        Xmlns string `xml:"xmlns,attr"`
        Version string
}

type documentsettings struct {
        Xmlns string `xml:"xmlns,attr"`
        AllowPNG string
}

type excelworkbook struct {
        XMLName xml.Name `xml:"ExcelWorkbook"`
        Xmlns string `xml:"xmlns,attr"`
        WindowHeight int
        WindowWidth int
        WindowTopX int
        WindowTopY int
        ProtectStructure string
        ProtectWindows string
}

type styles struct {
        XMLName xml.Name `xml:"Styles"`
        Style style
}

type style struct {
        Ssid string `xml:"ss:ID,attr"`
        Ssname string `xml:"ss:Name,attr"`
        Alignment alignment
        Borders string
        Font font
        Interior string
        NumberFormat string
        Protection string
}


type alignment struct {
        Ssvert string `xml:"ss:Vertical,attr"`
}


type font struct {
        Ssfont string `xml:"ss:FontName,attr"`
        Xfamily string `xml:"x:Family,attr"`
        Sssize string `xml:"ss:Size,attr"`
        Sscolor string `xml:"ss:Color,attr"`
}

type worksheet struct {
        XMLName xml.Name `xml:"Worksheet"`
        SSName string `xml:"ss:Name,attr"`
        Table table
}

type table struct {
        XMLName xml.Name `xml:"Table"`
        Rows []row
}

type row struct {
        XMLName xml.Name `xml:"Row"`
        Cells []cell
}

type cell struct {
        XMLName xml.Name `xml:"Cell"`
        Data data
}


type data struct {
        XMLName xml.Name `xml:"Data"`
        SSType string `xml:"ss:Type,attr"`
        Value string `xml:",chardata"`
}

type worksheetoptions struct {
        Xmlns string `xml:"xlmns,attr"`
        Selected string
        ProtectObjects string
        ProtectScenarios string
}

func makeWorkbook(wksheet worksheet) workbook { 
        excel_font := font{Ssfont:"Calibri", Xfamily:"Swiss", Sssize:"11", Sscolor:"#000000"}
    
        excel_style := style{Ssid:"Default", Ssname:"Normal", Font:excel_font, Alignment:alignment{Ssvert:"Bottom"}}

	workbook := workbook{Xmlns:"urn:schemas-microsoft-com:office:spreadsheet", Xmlnso:"urn:schemas-microsoft-com:office:office", Xmlnsx:"urn:schemas-microsoft-com:office:excel",Xmlnss:"urn:schemas-microsoft-com:office:spreadsheet",Xmlnsh:"http://www.w3.org/TR/REC-html40", Styles:styles{Style:excel_style}, Worksheet:wksheet}
	
	return workbook
}
