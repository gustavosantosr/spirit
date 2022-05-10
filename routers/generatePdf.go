package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gustavosantosr/spirit/bd"

	gofpdf "github.com/jung-kurt/gofpdf"
)

/*GeneratePdf Leo las Statess */
func GeneratePdf(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("IDDespacho")) < 1 {
		http.Error(w, "debe enviar el IDDespacho", http.StatusBadRequest)
		return
	}
	IDDespacho, err := strconv.Atoi(r.URL.Query().Get("IDDespacho"))
	if err != nil {
		http.Error(w, "Error debe enviar IDDespacho mayor a 0", http.StatusBadRequest)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")

	
	encabezado := func() {
		RespDespacho, err := bd.GetDespachosbyID(IDDespacho)
		if err != nil {
			http.Error(w, "Error el leer el despacho por id", http.StatusBadRequest)
			return
		}
		// pdf.Image("./assets/logo.png", 120, 7, 15, 0, false, "", 0, "http://www.honesolutionst.com")
		
		pdf.AddUTF8Font("dejavu", "", "./font/DejaVuSansCondensed.ttf")
		pdf.AddUTF8Font("dejavu", "B", "./font/DejaVuSansCondensed-Bold.ttf")
		pdf.AddUTF8Font("dejavu", "I", "./font/DejaVuSansCondensed-Oblique.ttf")
		pdf.AddUTF8Font("dejavu", "BI", "./font/DejaVuSansCondensed-BoldOblique.ttf")
		pdf.SetFont("dejavu", "", 15)
		_, lineHt := pdf.GetFontSize()
		htmlStr := ` `
		html := pdf.HTMLBasicNew()
		html.Write(lineHt, htmlStr)
		pdf.SetLeftMargin(80)
		pdf.SetFontSize(13)
			pdf.Ln(-1)
			pdf.CellFormat(45, 6, "TEXTILES SPIRIT S.A.S.: ", "0", 0, "", false, 0, "")
			pdf.Ln(-1)
			pdf.CellFormat(45, 6, "NIT 901208883-1", "0", 0, "", false, 0, "")
			pdf.Ln(-1)
			pdf.CellFormat(45, 6, "CALLE 78 A # 63 - 64", "0", 0, "", false, 0, "")
			pdf.SetLeftMargin(10)
			pdf.SetFontSize(11)
		
		pdf.Ln(10)
		
			pdf.CellFormat(100, 6, "Cliente: "+fmt.Sprintf("%v", RespDespacho[0].Tercero.Name), "0", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, fmt.Sprintf("%v",""), "0", 0, "", false, 0, "")
			pdf.CellFormat(100, 6, fmt.Sprintf("%v",""), "0", 0, "", false, 0, "")
			pdf.Ln(-1)
			pdf.CellFormat(60, 6, "Departamento: "+fmt.Sprintf("%v", RespDespacho[0].Tercero.State), "0", 0, "", false, 0, "")
			pdf.CellFormat(60, 6, "Ciudad: "+fmt.Sprintf("%v", RespDespacho[0].Tercero.City), "0", 0, "", false, 0, "")
			pdf.CellFormat(120, 6, "Dirección: "+fmt.Sprintf("%v", RespDespacho[0].Tercero.Adress), "0", 0, "", false, 0, "")
			pdf.Ln(-1)
			pdf.CellFormat(100, 6, "Teléfono " + fmt.Sprintf("%v",  RespDespacho[0].Tercero.Phone), "0", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, fmt.Sprintf("%v",""), "0", 0, "", false, 0, "")
			pdf.CellFormat(100, 6, "Fecha " + fmt.Sprintf("%v",  RespDespacho[0].FechaDespacho), "0", 0, "", false, 0, "")
			pdf.SetLeftMargin(10)
			pdf.Ln(-1)
			pdf.Ln(-1)
		
	}
	basicTable := func() {
		RespSalida, err := bd.GetSalidasbyDespacho(IDDespacho)
		if err != nil {
			http.Error(w, "Error el leer el despacho", http.StatusBadRequest)
			return
		}
		
		pdf.SetFillColor(224, 235, 255)
		
		fill := false
		pdf.CellFormat(10, 6, "No", "1", 0, "", fill, 0, "")
			pdf.CellFormat(50, 6, "Código", "1", 0, "", fill, 0, "")
			pdf.CellFormat(20, 6, "Metros", "1", 0, "", fill, 0, "")
			pdf.CellFormat(80, 6, "Producto", "1", 0, "", fill, 0, "")
			pdf.CellFormat(30, 6, "Referencia", "1", 0, "", fill, 0, "")
		pdf.Ln(-1)
		x:=1
		 m := 0.0
		for _, c := range RespSalida {
			
			
			pdf.CellFormat(10, 6, fmt.Sprintf("%v", x), "1", 0, "", fill, 0, "")
			pdf.CellFormat(50, 6, fmt.Sprintf("%v", c.Pieza.Codigo), "1", 0, "", fill, 0, "")
			pdf.CellFormat(20, 6, fmt.Sprintf("%.2f", c.Pieza.MetrosRecibidos), "1", 0, "", fill, 0, "")
			pdf.CellFormat(80, 6, c.Pieza.Producto.Name, "1", 0, "", fill, 0, "")
			pdf.CellFormat(30, 6, c.Pieza.Producto.Code, "1", 0, "", fill, 0, "")
			pdf.Ln(-1)
			m=m+float64(c.Pieza.MetrosRecibidos)
			x++
		}
		x=x-1
		pdf.Ln(-1)
		pdf.CellFormat(45, 6, "Total Piezas: "+fmt.Sprintf("%v",x), "0", 0, "", false, 0, "")
		pdf.Ln(-1)
			pdf.CellFormat(40, 6, "Total Metros: " +fmt.Sprintf("%.2f", m), "0", 0, "", false, 0, "")
			pdf.Ln(-1)
			pdf.Ln(-1)
			pdf.Ln(-1)
			pdf.Ln(-1)
			pdf.Ln(-1)
			pdf.CellFormat(45, 6, "___________________", "0", 0, "", false, 0, "")
			pdf.CellFormat(45, 6, "", "0", 0, "", false, 0, "")
			pdf.CellFormat(45, 6, "___________________", "0", 0, "", false, 0, "")
			pdf.Ln(-1)
			pdf.CellFormat(45, 6, "FIRMA RECIBIDO", "0", 0, "", false, 0, "")
			pdf.CellFormat(45, 6, "", "0", 0, "", false, 0, "")
			pdf.CellFormat(45, 6, "      FIRMA", "0", 0, "", false, 0, "")
		
	}
	// Better table
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()
	encabezado()
	basicTable()
	err = pdf.OutputFileAndClose("./test/despacho"+fmt.Sprintf("%v",IDDespacho)+".pdf")
	if err != nil {
		fmt.Println("Error reading rows: " + err.Error())
	}
	f, err := os.Open("./test/despacho"+fmt.Sprintf("%v",IDDespacho)+".pdf")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	//Stream to response
	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	//Set header
	//w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-type", "application/pdf")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	//sjson.NewEncoder(w).Encode(true)
}

func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}
