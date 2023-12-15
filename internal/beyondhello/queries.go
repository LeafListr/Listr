package beyondhello

import (
	"fmt"
)

func MenuQuery(storeId int) string {
	return fmt.Sprintf(`{"query":"","filters":"store_id :  %d","facets":["*"]}`, storeId)
}
