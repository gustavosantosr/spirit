package models

/*Usuario modelo de productos de terminados*/
type Usuario struct {
	IDUsuario  int64
	Nombre     string
	Email      string
	Contrasena string
	Tipo       int32
	Activo     int16
}

/*TokenSiigo modelo de productos de terminados*/
type TokenSiigo struct {
	Access  string `json:"access_token"`
	Expires int32  `json:"expires_in"`
	Type    string `json:"token_type"`
	Scope   string  `json:"scope"`
}
