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

func productQuery(storeId, productId int) string {
	return fmt.Sprintf(`{"filters":"store_id = %d AND product_id = %d","hitsPerPage":1}`, storeId, productId)
}
