package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-product-list-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToProductBusinessPartnerRelationCollection(raw []byte, l *logger.Logger) ([]ProductBusinessPartnerRelationCollection, error) {
	pm := &responses.ProductBusinessPartnerRelationCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ProductBusinessPartnerRelationCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	productBusinessPartnerRelationCollection := make([]ProductBusinessPartnerRelationCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		productBusinessPartnerRelationCollection = append(productBusinessPartnerRelationCollection, ProductBusinessPartnerRelationCollection{
			ObjectID:                            data.ObjectID,
			ID:                                  data.ID,
			Description:                         data.Description,
			TypeCode:                            data.TypeCode,
			CategoryCode:                        data.CategoryCode,
			LifeCycleStatusCode:                 data.LifeCycleStatusCode,
			AutoProposalIndicator:               data.AutoProposalIndicator,
			ValidForAllBusinessPartnerIndicator: data.ValidForAllBusinessPartnerIndicator,
			ValidFromDate:                       data.ValidFromDate,
			ValidToDate:                         data.ValidToDate,
			EntityLastChangedOn:                 data.EntityLastChangedOn,
			ETag:                                data.ETag,
		})
	}

	return productBusinessPartnerRelationCollection, nil
}
