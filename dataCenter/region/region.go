package region

type Region struct {
	Region    string `json:region`
	ISOCode   string `json:isocode`
	Infected  int    `json:infected`
	Recovered int    `json:recovered`
	Deceased  int    `json:deceased`
	Country   string `json:country`
}

type Total struct {
	Infected  int    `json:infected`
	Recovered int    `json:recovered`
	Deceased  int    `json:deceased`
	Country   string `json:country`
}
