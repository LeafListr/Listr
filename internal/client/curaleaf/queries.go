package curaleaf

import "fmt"

func ProductQuery(menuId, productId, menuType string) string {
	return fmt.Sprintf(`
{
    "query": "fragment full on Product { brand { description id image { url } name slug } category { displayName key } descriptionHtml effects { displayName } id images { id url } labResults { cannabinoids { cannabinoid { description name } unit value } terpenes { terpene { description name } unitSymbol value } thc { formatted range } } name offers { description id title } strain { key displayName } subcategory { key displayName } variants { flowerEquivalent { unit value } id isSpecial option price quantity specialPrice } } query PDP($dispensaryUniqueId: ID!, $menuType: MenuType!, $productId: ID!) { product( dispensaryUniqueId: $dispensaryUniqueId menuType: $menuType productId: $productId ) { product { ...full } } }",
    "variables": {
        "dispensaryUniqueId": "%s",
        "menuType": "%s",
		"productId": "%s"
    }
}`, menuId, menuType, productId)
}

func AllProductQuery(menuId, menuType string) string {
	return fmt.Sprintf(`
{
    "query": "query PGP($dispensaryUniqueId: ID!, $menuType: MenuType!, $categoryType: [Category]) { dispensaryMenu(dispensaryUniqueId: $dispensaryUniqueId, menuType: $menuType, categories: $categoryType) { products { brand { description id image { url } name slug } cardDescription category { displayName key } id images { url } labResults { cannabinoids { cannabinoid { description name } unit value } terpenes { terpene { description name } unitSymbol value } thc { formatted range } } name offers { id title } strain { key displayName  } subcategory { key displayName } variants { id isSpecial option price quantity specialPrice } } } }",
    "variables": {
        "dispensaryUniqueId": "%s",
        "menuType": "%s"
    }
}`, menuId, menuType)
}

func AllProductForCategoryQuery(menuId, menuType, category string) string {
	return fmt.Sprintf(`
{
    "query": "query PGP($dispensaryUniqueId: ID!, $menuType: MenuType!, $categoryType: [Category]) { dispensaryMenu(dispensaryUniqueId: $dispensaryUniqueId, menuType: $menuType, categories: $categoryType) { products { brand { description id image { url } name slug } cardDescription category { displayName key } id images { url } labResults { cannabinoids { cannabinoid { description name } unit value } terpenes { terpene { description name } unitSymbol value } thc { formatted range } } name offers { id title } strain { key displayName  } subcategory { key displayName } variants { id isSpecial option price quantity specialPrice } } } }",
    "variables": {
        "dispensaryUniqueId": "%s",
        "menuType": "%s",
        "categoryType": "%s"
    }
}`, menuId, menuType, category)
}

func AllOffersQuery(menuId, menuType string) string {
	return fmt.Sprintf(`
{
    "query": "query PGP($dispensaryUniqueId: ID!, $menuType: MenuType!) { dispensaryMenu(dispensaryUniqueId: $dispensaryUniqueId, menuType: $menuType) { offers { id title } } }",
    "variables": {
        "dispensaryUniqueId": "%s",
        "menuType": "%s"
    }
}`, menuId, menuType)
}

func AllCategoriesQuery(menuId, menuType string) string {
	return fmt.Sprintf(`
{
    "query": "query PGP($dispensaryUniqueId: ID!, $menuType: MenuType!) { dispensaryMenu(dispensaryUniqueId: $dispensaryUniqueId, menuType: $menuType) { allFilters { categories { displayName key } } } }",
	"variables": {
        "dispensaryUniqueId": "%s",
        "menuType": "%s"
    }
}`, menuId, menuType)
}

func AllLocationsQuery(longitude, latitude float64) string {
	return fmt.Sprintf(`
{
    "query": "query Dispensaries($coordinates: CoordinatesInput, $forDelivery: Boolean) { dispensaries(coordinates: $coordinates, forDelivery: $forDelivery) { uniqueId name slug orderTypes menuTypes isOpened location { coordinates { latitude longitude } address city distance distanceFormatted state stateAbbreviation stateSlug zipCode } nextTime validForDelivery } }",
    "variables": {
        "coordinates": {
            "longitude": %f,
            "latitude": %f
        },
        "forDelivery": false
    }
}`, longitude, latitude)
}
