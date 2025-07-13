package docker

type InsertDto struct {
	Service struct {
		Name  string  `json:"name"`
		Value Service `json:"value"`
	} `json:"service"`
	Network struct {
		Name  string  `json:"name"`
		Value Network `json:"value"`
	} `json:"network"`
	Volume struct {
		Name  string `json:"name"`
		Value Volume `json:"value"`
	} `json:"volume"`
}

type DeleteDto struct {
	Name string `json:"name"`
}
