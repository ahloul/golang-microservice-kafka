package helpers

type Pagination struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

type Link struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Action struct {
	Id          string `jsonapi:"primary,action" json:"-"`
	EndpointUrl string `jsonapi:"attr,endpoint_url" json:"endpoint_url"`
	Key         string `jsonapi:"attr,key" json:"key"`
	Label       string `jsonapi:"attr,label" json:"label"`
	Method      string `jsonapi:"attr,method" json:"method"`
}
