package sap_api_output_formatter

type ProductList struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	ProductListCode string `json:"product_list_code"`
	Deleted         bool   `json:"deleted"`
}

type ProductBusinessPartnerRelationCollection struct {
	ObjectID                            string `json:"ObjectID"`
	ID                                  string `json:"ID"`
	Description                         string `json:"Description"`
	TypeCode                            string `json:"TypeCode"`
	CategoryCode                        string `json:"CategoryCode"`
	LifeCycleStatusCode                 string `json:"LifeCycleStatusCode"`
	AutoProposalIndicator               bool   `json:"AutoProposalIndicator"`
	ValidForAllBusinessPartnerIndicator bool   `json:"ValidForAllBusinessPartnerIndicator"`
	ValidFromDate                       string `json:"ValidFromDate"`
	ValidToDate                         string `json:"ValidToDate"`
	EntityLastChangedOn                 string `json:"EntityLastChangedOn"`
	ETag                                string `json:"ETag"`
}
