package responses

type ProductBusinessPartnerRelationCollection struct {
	D struct {
		Results []struct {
			Metadata struct {
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ObjectID                                      string `json:"ObjectID"`
			ID                                            string `json:"ID"`
			Description                                   string `json:"Description"`
			TypeCode                                      string `json:"TypeCode"`
			CategoryCode                                  string `json:"CategoryCode"`
			LifeCycleStatusCode                           string `json:"LifeCycleStatusCode"`
			AutoProposalIndicator                         bool   `json:"AutoProposalIndicator"`
			ValidForAllBusinessPartnerIndicator           bool   `json:"ValidForAllBusinessPartnerIndicator"`
			ValidFromDate                                 string `json:"ValidFromDate"`
			ValidToDate                                   string `json:"ValidToDate"`
			EntityLastChangedOn                           string `json:"EntityLastChangedOn"`
			ETag                                          string `json:"ETag"`
			ProductBusinessPartnerRelationBusinessPartner struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationBusinessPartner"`
			ProductBusinessPartnerRelationProcessingType struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationProcessingType"`
			ProductBusinessPartnerRelationProduct struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationProduct"`
			ProductBusinessPartnerRelationProductCategory struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationProductCategory"`
			ProductBusinessPartnerRelationRequiredProduct struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationRequiredProduct"`
			ProductBusinessPartnerRelationRequiredProductCategory struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationRequiredProductCategory"`
			ProductBusinessPartnerRelationSalesArea struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationSalesArea"`
			ProductBusinessPartnerRelationSalesTerritory struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationSalesTerritory"`
			ProductBusinessPartnerRelationTargetGroup struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationTargetGroup"`
			ProductBusinessPartnerRelationTeam struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"ProductBusinessPartnerRelationTeam"`
		} `json:"results"`
	} `json:"d"`
}
