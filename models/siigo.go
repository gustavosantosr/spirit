package models

/*PAgination modelo de productos de terminados*/
type PAgination struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
	Total    int64 `json:"total_results"`
}

/*IdType modelo de productos de terminados*/
type IdType struct {
	Code     string `json:"code"`
	Name string `json:"name"`
}
/*FiscalResponsibilities modelo de productos de terminados*/
type Fiscalresponsibilities struct {
	Code     string `json:"code"`
	Name string `json:"name"`
}
/*ADress modelo de productos de terminados*/
type ADress struct {
	Address     string `json:"address"`
	City CIty `json:"city"`
	PostalCode string `json:"postal_code"` 
}
/*CIty modelo de productos de terminados*/
type CIty struct {
	CountryCode     string `json:"country_code"`
	CountryName     string `json:"country_name"`
	StateCode     string  `json:"state_code"`
	StateName     string `json:"state_name"`
	CityCode     string `json:"city_code"`
	CityName     string `json:"city_name"`
}
/*PHones modelo de productos de terminados*/
type PHones struct {
	Indicative     string `json:"indicative"`
	Number     string `json:"number"`
}
/*COntacts modelo de productos de terminados*/
type COntacts struct {
	FirstName     string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email     string `json:"email"`
}

/*ResultSiigo modelo de productos de terminados*/
type ResultSiigo struct {
	ID  string `json:"id"`
	Type string  `json:"type"`
	PersonType    string `json:"person_type"`
	IDType   IdType `json:"id_type"`
	Identification    string `json:"identification"`
	BranchOffice    int16 `json:"branch_office"`
	Name    []string `json:"name"`
	Active    bool `json:"active"`
	Vatresponsible    bool `json:"vat_responsible"`
	FiscalResponsibilities    []Fiscalresponsibilities `json:"fiscal_responsibilities"`
	Adress ADress `json:"address"`
	Phones []PHones `json:"phones"`
	Contacts []COntacts `json:"contacts"`
	
}
/*ResultProducto modelo de productos de terminados*/
type ResultProducto struct {
	ID  string `json:"id"`
	Code string  `json:"code"`
	Name    string `json:"name"`
	Active    bool `json:"active"`
	}
/*TerceroSiigo modelo de productos de terminados*/
type TerceroSiigo struct {
	Pagination     PAgination `json:"pagination"`
	Results     []ResultSiigo `json:"results"`
	
}
/*ProductoSiigo modelo de productos de terminados*/
type ProductoSiigo struct {
	Pagination     PAgination `json:"pagination"`
	Results     []ResultProducto `json:"results"`
	
}
/*MessageSiigo modelo de Messages de terminados*/
type MessageSiigo struct {
	Message     string 
	Error     string 
	
}



