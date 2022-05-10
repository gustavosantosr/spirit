package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gustavosantosr/spirit/middlew"
	"github.com/gustavosantosr/spirit/routers"
	"github.com/rs/cors"
)

/*Manejadores manejador de urls*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/getip", middlew.ChequeoBD(routers.GetIp)).Methods("GET")
	router.HandleFunc("/getdocumentotipo", middlew.ChequeoBD(routers.GetDocumentoTipo)).Methods("GET")
	router.HandleFunc("/insertdocumentotipo", middlew.ChequeoBD(routers.InsertDocumentoTipo)).Methods("POST")
	router.HandleFunc("/updatedocumentotipo", middlew.ChequeoBD(routers.UpdateDocumentoTipo)).Methods("PUT")
	/*EndPoints Estados Despacho*/
	router.HandleFunc("/getdespachoestado", middlew.ChequeoBD(routers.GetDespachoEstado)).Methods("GET")
	router.HandleFunc("/insertdespachoestado", middlew.ChequeoBD(routers.InsertDespachoEstado)).Methods("POST")
	router.HandleFunc("/updatedespachoestado", middlew.ChequeoBD(routers.UpdateDespachoEstado)).Methods("PUT")
	/*EndPoints Piezas*/

	router.HandleFunc("/getpieza", middlew.ChequeoBD(routers.GetPieza)).Methods("GET")
	router.HandleFunc("/getallpieza", middlew.ChequeoBD(routers.GetAllPieza)).Methods("GET")
	router.HandleFunc("/getpiezabycode", middlew.ChequeoBD(routers.GetPiezabyCode)).Methods("GET")
	router.HandleFunc("/insertpieza", middlew.ChequeoBD(routers.InsertPieza)).Methods("POST")
	router.HandleFunc("/insertcorte", middlew.ChequeoBD(routers.InsertCorte)).Methods("POST")
	router.HandleFunc("/updatepieza", middlew.ChequeoBD(routers.UpdatePieza)).Methods("PUT")
	router.HandleFunc("/getpiezasbycliente", middlew.ChequeoBD(routers.GetPiezabyCliente)).Methods("GET")
	
	/*EndPoints Despachos*/

	router.HandleFunc("/getdespacho", middlew.ChequeoBD(routers.GetDespacho)).Methods("GET")
	router.HandleFunc("/getdespachoactivo", middlew.ChequeoBD(routers.GetDespachoActivo)).Methods("GET")
	router.HandleFunc("/insertdespacho", middlew.ChequeoBD(routers.InsertDespacho)).Methods("POST")
	router.HandleFunc("/updatedespacho", middlew.ChequeoBD(routers.UpdateDespacho)).Methods("PUT")
	/*EndPoints Salidas*/

	router.HandleFunc("/getsalida", middlew.ChequeoBD(routers.GetSalida)).Methods("GET")
	router.HandleFunc("/getsalidabyiddespacho", middlew.ChequeoBD(routers.GetSalidabyIDDespacho)).Methods("GET")
	router.HandleFunc("/insertsalida", middlew.ChequeoBD(routers.InsertSalida)).Methods("POST")
	router.HandleFunc("/updatesalida", middlew.ChequeoBD(routers.UpdateSalida)).Methods("PUT")
	router.HandleFunc("/deletesalida", middlew.ChequeoBD(routers.DeleteSalida)).Methods("POST")
	/*EndPoints Usuarios*/
	router.HandleFunc("/getauth", middlew.ChequeoBD(routers.GetAuth)).Methods("POST")
	router.HandleFunc("/getusuario", middlew.ChequeoBD(routers.GetUsuario)).Methods("GET")
	router.HandleFunc("/insertusuario", middlew.ChequeoBD(routers.InsertUsuario)).Methods("POST")
	router.HandleFunc("/updateusuario", middlew.ChequeoBD(routers.UpdateUsuario)).Methods("PUT")
	/*EndPoints Cliente*/
	router.HandleFunc("/insertcliente", middlew.ChequeoBD(routers.InsertCliente)).Methods("POST")
	router.HandleFunc("/getcliente", middlew.ChequeoBD(routers.GetCliente)).Methods("GET")
	router.HandleFunc("/updatecliente", middlew.ChequeoBD(routers.UpdateCliente)).Methods("PUT")
	/*EndPoints Diseños*/
	router.HandleFunc("/insertserviciotipo", middlew.ChequeoBD(routers.InsertServicioTipo)).Methods("POST")
	router.HandleFunc("/getserviciotipo", middlew.ChequeoBD(routers.GetServicioTipo)).Methods("GET")
	router.HandleFunc("/updateserviciotipo", middlew.ChequeoBD(routers.UpdateServicioTipo)).Methods("PUT")
	//EndPoint Generate PDF
	router.HandleFunc("/generatePdf", middlew.ChequeoBD(routers.GeneratePdf)).Methods("GET")
	/*EndPoints Chat
	router.HandleFunc("/insertchat", middlew.ChequeoBD(routers.InsertChat)).Methods("PUT")
	router.HandleFunc("/sendemail", middlew.ChequeoBD(routers.SendEmail)).Methods("PUT")*/
	/*EndPoints Diseños*/
	/*EndPoints Diseños*/
	router.HandleFunc("/insertproducto", middlew.ChequeoBD(routers.InsertProducto)).Methods("POST")
	router.HandleFunc("/getproductos", middlew.ChequeoBD(routers.GetProducto)).Methods("GET")
	router.HandleFunc("/updateproductos", middlew.ChequeoBD(routers.UpdateProducto)).Methods("PUT")

	router.HandleFunc("/getterceros", middlew.ChequeoBD(routers.GetTercero)).Methods("GET")
	router.HandleFunc("/getproveedores", middlew.ChequeoBD(routers.GetProveedor)).Methods("GET")
	router.HandleFunc("/updateterceros", middlew.ChequeoBD(routers.UpdateTercero)).Methods("PUT")

	router.HandleFunc("/getsiigo", middlew.ChequeoBD(routers.GetSiigoProduct)).Methods("GET")
	router.HandleFunc("/getsiigotercero", middlew.ChequeoBD(routers.GetSiigoTercer)).Methods("GET")
	router.HandleFunc("/insertencuesta", middlew.ChequeoBD(routers.InsertEncuesta)).Methods("PUT")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
