package curaleaf

func AllProductQuery() string {
	return `fragment grid on Product { brand { description id image { url __typename } name slug __typename } cardDescription category { displayName key __typename } id images { url __typename } labResults { thc { formatted range __typename } __typename } name offers { id title __typename } strain { key displayName __typename } subcategory { key displayName __typename } variants { id isSpecial option price quantity specialPrice __typename } __typename } query PGP($dispensaryUniqueId: ID!, $menuType: MenuType!) { dispensaryMenu(dispensaryUniqueId: $dispensaryUniqueId, menuType: $menuType) { offers { id title __typename } products { ...grid __typename } __typename } }`
}

func AllProductForCategoryQuery() string {
	return `fragment grid on Product { brand { description id image { url __typename } name slug __typename } cardDescription category { displayName key __typename } id images { url __typename } labResults { thc { formatted range __typename } __typename } name offers { id title __typename } strain { key displayName __typename } subcategory { key displayName __typename } variants { id isSpecial option price quantity specialPrice __typename } __typename } query PGP($dispensaryUniqueId: ID!, $menuType: MenuType!, $categoryType: [Category]) { dispensaryMenu(dispensaryUniqueId: $dispensaryUniqueId, menuType: $menuType, categories: $categoryType) { offers { id title __typename } products { ...grid __typename } __typename } }`
}
