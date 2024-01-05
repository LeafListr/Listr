package beyondhello

import (
	"fmt"
)

func menuQuery(storeId int) string {
	return fmt.Sprintf(`{"query":"","filters":"store_id :  %d","hitsPerPage":5000}`, storeId)
}

func menuQueryWithCategory(storeId int, category string) string {
	return fmt.Sprintf(`{"query":"","filters":"store_id :  %d AND kind : %s","hitsPerPage":5000}`, storeId, category)
}

func facetQuery(storeId int, facet string) string {
	return fmt.Sprintf(`{"query":"","filters":"store_id : %d","hitsPerPage":0, "facets": ["%s"]}`, storeId, facet)
}
