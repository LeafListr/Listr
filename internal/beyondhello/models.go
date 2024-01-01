package beyondhello

import (
	"time"
)

type hit struct {
	Activities                 []string      `json:"activities"`
	AggregateRating            float64       `json:"aggregate_rating"`
	Allergens                  []interface{} `json:"allergens"`
	AllowMultipleFlowerCount   bool          `json:"allow_multiple_flower_count"`
	Amount                     *string       `json:"amount"`
	ApplicableBrandSpecialIds  []interface{} `json:"applicable_brand_special_ids"`
	ApplicableBulkSpecialIds   []interface{} `json:"applicable_bulk_special_ids"`
	ApplicableBundleSpecialIds struct {
		Qualify  []interface{} `json:"qualify"`
		Discount []interface{} `json:"discount"`
	} `json:"applicable_bundle_special_ids"`
	ApplicableSpecialIds   []int    `json:"applicable_special_ids"`
	ApplicableSpecialTypes []string `json:"applicable_special_types"`
	AvailableForDelivery   bool     `json:"available_for_delivery"`
	AvailableForPickup     bool     `json:"available_for_pickup"`
	AvailableWeights       []string `json:"available_weights"`
	BestSellerRank         *int     `json:"best_seller_rank"`
	Brand                  string   `json:"brand"`
	BrandSpecialPrices     struct {
		HalfGram interface{} `json:"half_gram"`
		Gram     interface{} `json:"gram"`
	} `json:"brand_special_prices"`
	BrandSubtype                string        `json:"brand_subtype"`
	BucketPrice                 int           `json:"bucket_price"`
	BusinessLicenses            []interface{} `json:"business_licenses"`
	Cannabinoids                []interface{} `json:"cannabinoids"`
	Category                    string        `json:"category"`
	Collections                 []interface{} `json:"collections"`
	CustomProductSubtype        string        `json:"custom_product_subtype"`
	CustomProductType           string        `json:"custom_product_type"`
	Description                 string        `json:"description"`
	DiscountedPriceEighthOunce  string        `json:"discounted_price_eighth_ounce"`
	DiscountedPriceGram         string        `json:"discounted_price_gram"`
	DiscountedPriceHalfGram     string        `json:"discounted_price_half_gram"`
	DiscountedPriceHalfOunce    string        `json:"discounted_price_half_ounce"`
	DiscountedPriceOunce        string        `json:"discounted_price_ounce"`
	DiscountedPriceQuarterOunce string        `json:"discounted_price_quarter_ounce"`
	DiscountedPriceTwoGram      string        `json:"discounted_price_two_gram"`
	DiscountedPriceEach         string        `json:"discounted_price_each"`
	Dosage                      interface{}   `json:"dosage"`
	Effects                     []interface{} `json:"effects"`
	Feelings                    []string      `json:"feelings"`
	Flavors                     []interface{} `json:"flavors"`
	Geoloc                      struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"_geoloc"`
	HasBrandDiscount   bool               `json:"has_brand_discount"`
	HasPhotos          bool               `json:"has_photos"`
	HasThcPotency      bool               `json:"has_thc_potency"`
	ImageUrls          []string           `json:"image_urls"`
	IndexedAt          time.Time          `json:"indexed_at"`
	Ingredients        []interface{}      `json:"ingredients"`
	InventoryPotencies []InventoryPotency `json:"inventory_potencies"`
	Kind               string             `json:"kind"`
	KindSubtype        string             `json:"kind_subtype"`
	LabResultUrls      []interface{}      `json:"lab_result_urls"`
	LabResults         []struct {
		PriceId    string `json:"price_id"`
		LabResults []struct {
			CompoundName string  `json:"compound_name"`
			Value        float64 `json:"value"`
			Unit         string  `json:"unit"`
			UnitId       string  `json:"unit_id"`
		} `json:"lab_results"`
	} `json:"lab_results"`
	MaxCartQuantity   int         `json:"max_cart_quantity"`
	Name              string      `json:"name"`
	NetWeightGrams    float64     `json:"net_weight_grams"`
	ObjectID          string      `json:"objectID"`
	OperatorStoreRank interface{} `json:"operator_store_rank"`
	PercentCbd        interface{} `json:"percent_cbd"`
	PercentCbda       interface{} `json:"percent_cbda"`
	PercentTac        interface{} `json:"percent_tac"`
	PercentThc        float64     `json:"percent_thc"`
	PercentThca       interface{} `json:"percent_thca"`
	Photos            []struct {
		Id       int `json:"id"`
		Position int `json:"position"`
		Urls     struct {
			Original   string `json:"original"`
			Small      string `json:"small"`
			Medium     string `json:"medium"`
			ExtraLarge string `json:"extraLarge"`
		} `json:"urls"`
	} `json:"photos"`
	PosProductLookup  struct{}    `json:"pos_product_lookup"`
	PriceEighthOunce  string      `json:"price_eighth_ounce"`
	PriceGram         string      `json:"price_gram"`
	PriceHalfGram     string      `json:"price_half_gram"`
	PriceHalfOunce    string      `json:"price_half_ounce"`
	PriceOunce        string      `json:"price_ounce"`
	PriceQuarterOunce string      `json:"price_quarter_ounce"`
	PriceTwoGram      string      `json:"price_two_gram"`
	PriceEach         string      `json:"price_each"`
	ProductBrandId    int         `json:"product_brand_id"`
	ProductId         int         `json:"product_id"`
	ProductPercentCbd interface{} `json:"product_percent_cbd"`
	ProductPercentThc interface{} `json:"product_percent_thc"`
	ProductPhotos     []struct {
		Id       string `json:"id"`
		Position int    `json:"position"`
		Urls     struct {
			Original   string `json:"original"`
			Small      string `json:"small"`
			Medium     string `json:"medium"`
			ExtraLarge string `json:"extraLarge"`
		} `json:"urls"`
	} `json:"product_photos"`
	QuantityUnits  *string `json:"quantity_units"`
	QuantityValue  float64 `json:"quantity_value"`
	Recommendation struct {
		Ordinal               int         `json:"ordinal"`
		CalibratedProbability interface{} `json:"calibrated_probability"`
	} `json:"recommendation"`
	ReviewCount             int           `json:"review_count"`
	RootSubtype             string        `json:"root_subtype"`
	RootTypes               []interface{} `json:"root_types"`
	RootsCustomRows         []interface{} `json:"roots_custom_rows"`
	SearchableSlug          string        `json:"searchable_slug"`
	SpecialAmount           string        `json:"special_amount"`
	SpecialId               int           `json:"special_id"`
	SpecialPriceEighthOunce interface{}   `json:"special_price_eighth_ounce"`
	SpecialPriceGram        *struct {
		SpecialId       int     `json:"special_id"`
		Price           string  `json:"price"`
		DiscountPrice   string  `json:"discount_price"`
		DiscountAmount  float64 `json:"discount_amount"`
		DiscountType    string  `json:"discount_type"`
		DiscountPercent string  `json:"discount_percent"`
	} `json:"special_price_gram"`
	SpecialPriceHalfGram *struct {
		SpecialId       int     `json:"special_id"`
		Price           string  `json:"price"`
		DiscountPrice   string  `json:"discount_price"`
		DiscountAmount  float64 `json:"discount_amount"`
		DiscountType    string  `json:"discount_type"`
		DiscountPercent string  `json:"discount_percent"`
	} `json:"special_price_half_gram"`
	SpecialPriceHalfOunce    interface{}   `json:"special_price_half_ounce"`
	SpecialPriceOunce        interface{}   `json:"special_price_ounce"`
	SpecialPriceQuarterOunce interface{}   `json:"special_price_quarter_ounce"`
	SpecialTitle             string        `json:"special_title"`
	StoreNotes               string        `json:"store_notes"`
	StoreSpecificProduct     bool          `json:"store_specific_product"`
	StoreTypes               []string      `json:"store_types"`
	Strain                   interface{}   `json:"strain"`
	Terpenes                 []interface{} `json:"terpenes"`
	Type                     string        `json:"type"`
	UniqueSlug               string        `json:"unique_slug"`
	UrlSlug                  string        `json:"url_slug"`
}

type LocationResponse []*struct {
	Id    int    `json:"id"`
	Link  string `json:"link"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
}

type ProductResponse struct {
	Hits        []hit `json:"hits"`
	NbHits      int   `json:"nbHits"`
	Page        int   `json:"page"`
	NbPages     int   `json:"nbPages"`
	HitsPerPage int   `json:"hitsPerPage"`
	Facets      struct {
		Kind struct {
			Flower   int `json:"flower"`
			Vape     int `json:"vape"`
			Extract  int `json:"extract"`
			Edible   int `json:"edible"`
			Tincture int `json:"tincture"`
			Gear     int `json:"gear"`
			Topical  int `json:"topical"`
		} `json:"kind"`
		Brand struct {
			OrganicRemedies   int `json:"Organic Remedies"`
			SeCHe             int `json:"SeCHe"`
			Shine             int `json:"&Shine"`
			Lab               int `json:"Lab"`
			GLeaf             int `json:"gLeaf"`
			DoubleBear        int `json:"Double Bear"`
			MaitriGenetics    int `json:"Maitri Genetics"`
			Revel             int `json:"Revel"`
			Terrapin          int `json:"Terrapin"`
			GageCannabisCo    int `json:"Gage Cannabis Co."`
			FarmaceuticalRx   int `json:"FarmaceuticalRx"`
			Verano            int `json:"Verano"`
			Savvy             int `json:"Savvy"`
			Bank              int `json:"Bank"`
			NaturalSelections int `json:"Natural Selections"`
			PrimeWellness     int `json:"Prime Wellness"`
			Grassroots        int `json:"Grassroots"`
			Ilera             int `json:"Ilera"`
			Crops             int `json:"Crops"`
			ResoluteRemedies  int `json:"Resolute Remedies"`
			TheWoodsReserve   int `json:"The Woods Reserve"`
			VytalOptions      int `json:"Vytal Options"`
			AgriKind          int `json:"Agri-Kind"`
			KindTree          int `json:"Kind Tree"`
			WonderWellnessCo  int `json:"Wonder Wellness Co."`
			BeyondHello       int `json:"Beyond Hello"`
			DrSolomonS        int `json:"Dr. Solomons"`
			Insa              int `json:"Insa"`
			Nira              int `json:"Nira+"`
			StashLogix        int `json:"Stash Logix"`
			Tasteology        int `json:"Tasteology"`
			Avexia            int `json:"Avexia"`
			Cookies           int `json:"Cookies"`
			PAX               int `json:"PAX"`
			Parea             int `json:"Parea"`
			Puffco            int `json:"Puffco"`
			SteelCityGreats   int `json:"Steel City Greats"`
			Strane            int `json:"Strane"`
			Valhalla          int `json:"Valhalla"`
			Yocan             int `json:"Yocan"`
			Curaleaf          int `json:"Curaleaf"`
			FloraCal          int `json:"FloraCal"`
			GPEN              int `json:"GPEN"`
			GarciaHandPicked  int `json:"Garcia Hand Picked"`
			GoodGreen         int `json:"Good Green"`
			HigherStandards   int `json:"Higher Standards"`
			Integra           int `json:"Integra"`
			KYND              int `json:"KYND"`
			Lookah            int `json:"Lookah"`
			Moxie             int `json:"Moxie"`
			Ooze              int `json:"Ooze"`
			PennHealthGroup   int `json:"Penn Health Group"`
			TripleSeven       int `json:"Triple Seven"`
			WholePlants       int `json:"Whole Plants"`
			Hijinks           int `json:"hijinks"`
		} `json:"brand"`
		Category struct {
			Hybrid int `json:"hybrid"`
			Indica int `json:"indica"`
			Sativa int `json:"sativa"`
			Cbd    int `json:"cbd"`
		} `json:"category"`
		Feelings struct {
			Relaxed   int `json:"Relaxed"`
			Blissful  int `json:"Blissful"`
			PainFree  int `json:"Pain free"`
			Creative  int `json:"Creative"`
			Sleepy    int `json:"Sleepy"`
			Energetic int `json:"Energetic"`
			Hungry    int `json:"Hungry"`
			NotHigh   int `json:"Not high"`
		} `json:"feelings"`
		StoreId struct {
			Field1 int `json:"2935"`
		} `json:"store_id"`
		Activities struct {
			EaseMyMind      int `json:"Ease my mind"`
			GetRelief       int `json:"Get relief"`
			GetSomeSleep    int `json:"Get some sleep"`
			StimulateMyMind int `json:"Stimulate my mind"`
			HangWithFriends int `json:"Hang with friends"`
			GetActive       int `json:"Get Active"`
			GetIntimate     int `json:"Get intimate"`
		} `json:"activities"`
		ProductId struct{} `json:"product_id"`
		RootTypes struct {
			Flower                   int `json:"flower"`
			Vape                     int `json:"vape"`
			VapeCartridges           int `json:"vape:Cartridges"`
			WhatSNew2935             int `json:"Whats New?-2935"`
			HighTHC2935              int `json:"High THC-2935"`
			Sale                     int `json:"sale"`
			Extract                  int `json:"extract"`
			Edible                   int `json:"edible"`
			EdibleCandies            int `json:"edible:Candies"`
			FlowerSmalls             int `json:"flower:Smalls"`
			Tincture                 int `json:"tincture"`
			FlowerPremium            int `json:"flower:Premium"`
			BestSelling              int `json:"best_selling"`
			Gear                     int `json:"gear"`
			FlowerShake              int `json:"flower:Shake"`
			VapeDisposables          int `json:"vape:Disposables"`
			ExtractRickSimpsonOilRSO int `json:"extract:Rick Simpson Oil (RSO)"`
			TinctureSublinguals      int `json:"tincture:Sublinguals"`
			ExtractWaxes             int `json:"extract:Waxes"`
			GearVaporizers           int `json:"gear:Vaporizers"`
			ExtractLiveResins        int `json:"extract:Live Resins"`
			Topical                  int `json:"topical"`
			TinctureTinctures        int `json:"tincture:Tinctures"`
			FlowerGroundFlower       int `json:"flower:Ground Flower"`
			GearAccessories          int `json:"gear:Accessories"`
			EdibleCapsules           int `json:"edible:Capsules"`
			FlowerSelect             int `json:"flower:Select"`
			TopicalPatches           int `json:"topical:Patches"`
			ExtractIsolates          int `json:"extract:Isolates"`
			GearGrinders             int `json:"gear:Grinders"`
			TopicalBalms             int `json:"topical:Balms"`
			ExtractRosins            int `json:"extract:Rosins"`
			DeepDiscounts2935        int `json:"Deep Discounts-2935"`
			ExtractDiamondsCrystals  int `json:"extract:Diamonds/Crystals"`
			ExtractSauces            int `json:"extract:Sauces"`
			GearParaphernalia        int `json:"gear:Paraphernalia"`
			TinctureSprays           int `json:"tincture:Sprays"`
			TopicalCreams            int `json:"topical:Creams"`
			TopicalLotions           int `json:"topical:Lotions"`
			TopicalOther             int `json:"topical:Other"`
		} `json:"root_types"`
		Collections struct {
			GiveDanks3     int `json:"Give Danks:3"`
			ForAll5        int `json:"420 For All:5"`
			EndlessSummer1 int `json:"Endless Summer:1"`
			Holiday20236   int `json:"Holiday 2023:6"`
		} `json:"collections"`
		PercentCbd struct{} `json:"percent_cbd"`
		PercentThc struct{} `json:"percent_thc"`
		StoreTypes struct {
			Medical int `json:"medical"`
		} `json:"store_types"`
		UniqueSlug  struct{} `json:"unique_slug"`
		BucketPrice struct{} `json:"bucket_price"`
		PercentThca struct {
			Field1 int `json:"20.818"`
		} `json:"percent_thca"`
		ReviewCount  struct{} `json:"review_count"`
		BrandSubtype struct {
			Cartridge                     int `json:"Cartridge"`
			Flower                        int `json:"Flower"`
			PremiumFlower                 int `json:"Premium Flower"`
			Smalls                        int `json:"Smalls"`
			Tincture                      int `json:"Tincture"`
			Troches                       int `json:"Troches"`
			FineGrind                     int `json:"Fine Grind"`
			DistillateCartridge           int `json:"Distillate Cartridge"`
			MysticSpiritCartridge         int `json:"Mystic Spirit Cartridge"`
			LiquidLiveResinCartridge      int `json:"Liquid Live Resin Cartridge"`
			AllDay                        int `json:"All Day"`
			LiveResinCartridge            int `json:"Live Resin Cartridge"`
			RickSimpsonOilRSO             int `json:"Rick Simpson Oil (RSO)"`
			Sugar                         int `json:"Sugar"`
			VapeCartridge                 int `json:"Vape Cartridge"`
			HometownGrown                 int `json:"Hometown Grown"`
			Disposable                    int `json:"Disposable"`
			LiveSugar                     int `json:"Live Sugar"`
			FineFlower                    int `json:"Fine Flower"`
			SweetTroches                  int `json:"Sweet Troches"`
			Capsules                      int `json:"Capsules"`
			DisposablePen                 int `json:"Disposable Pen"`
			PortableVaporizer             int `json:"Portable Vaporizer"`
			Select                        int `json:"Select"`
			THCASandInfusedFlower         int `json:"THCA Sand Infused Flower"`
			TransdermalPatch              int `json:"Transdermal Patch"`
			ExtraStrengthTincture         int `json:"Extra Strength Tincture"`
			FullSpectrumDisposableVapePen int `json:"Full Spectrum Disposable Vape Pen"`
			GroundFlower                  int `json:"Ground Flower"`
			GroundFlowerSand              int `json:"Ground Flower + Sand"`
			LiveSauce                     int `json:"Live Sauce"`
			Premium                       int `json:"Premium"`
			StorageBag                    int `json:"Storage Bag"`
			Balm                          int `json:"Balm"`
			FullSpectrumVapeCartridge     int `json:"Full Spectrum Vape Cartridge"`
			Grinder                       int `json:"Grinder"`
			HerbGrinder                   int `json:"Herb Grinder"`
			LiveResinBeachSand            int `json:"Live Resin Beach Sand"`
			LiveRosin                     int `json:"Live Rosin"`
			LiveRosinCartridge            int `json:"Live Rosin Cartridge"`
			LiveSauceCartridge            int `json:"Live Sauce Cartridge"`
			NanoEmulsionHydroTincture     int `json:"Nano-Emulsion Hydro Tincture"`
			RSOSweetTroches               int `json:"RSO Sweet Troches"`
			RSOSyringe                    int `json:"RSO Syringe"`
			ResinCartridge                int `json:"Resin Cartridge"`
			SmallFlower                   int `json:"Small Flower"`
			Badder                        int `json:"Badder"`
			Blades                        int `json:"Blades"`
			BoxMod                        int `json:"Box Mod"`
			Bud                           int `json:"Bud"`
			DryHerbVaporizer              int `json:"Dry Herb Vaporizer"`
			EssenceFlower                 int `json:"Essence Flower"`
			ExtraStrengthSalve            int `json:"Extra Strength Salve"`
			FSEDisposablePen              int `json:"FSE Disposable Pen"`
			FlavoredElixir                int `json:"Flavored Elixir"`
			FullSpectrumDisposable        int `json:"Full Spectrum Disposable"`
			GrowerSSelect                 int `json:"Growers Select"`
			HTECartridge                  int `json:"HTE Cartridge"`
			HumidityPack                  int `json:"Humidity Pack"`
			InfusedLotion                 int `json:"Infused Lotion"`
			LiquidSugarCartridge          int `json:"Liquid Sugar Cartridge"`
			LiveDiamondsSauce             int `json:"Live Diamonds & Sauce"`
			LiveRickSimpsonOilRSO         int `json:"Live Rick Simpson Oil (RSO)"`
			NanoSpray                     int `json:"NanoSpray"`
			PenBattery                    int `json:"Pen Battery"`
			RollingTray                   int `json:"Rolling Tray"`
			Sand                          int `json:"Sand"`
			SeahorseVaporizer             int `json:"Seahorse Vaporizer"`
			SoftLozenges                  int `json:"Soft Lozenges"`
			Topical                       int `json:"Topical"`
			TopicalCream                  int `json:"Topical Cream"`
			Tray                          int `json:"Tray"`
			Vaporizer                     int `json:"Vaporizer"`
			Wax                           int `json:"Wax"`
			WaxPenVaporizerCoil           int `json:"Wax Pen Vaporizer Coil"`
			Wipes                         int `json:"Wipes"`
			IKrusher                      int `json:"iKrusher"`
		} `json:"brand_subtype"`
		AggregateRating struct{} `json:"aggregate_rating"`
		AtVisibleStore  struct {
			False int `json:"false"`
		} `json:"at_visible_store"`
		ProductBrandId   struct{} `json:"product_brand_id"`
		AvailableWeights struct {
			EighthOunce  int `json:"eighth ounce"`
			HalfGram     int `json:"half gram"`
			Gram         int `json:"gram"`
			QuarterOunce int `json:"quarter ounce"`
			HalfOunce    int `json:"half ounce"`
			Ounce        int `json:"ounce"`
		} `json:"available_weights"`
		HasBrandDiscount struct {
			False int `json:"false"`
		} `json:"has_brand_discount"`
		AvailableForPickup struct {
			True int `json:"true"`
		} `json:"available_for_pickup"`
		ApplicableSpecialIds struct {
			Field1 int `json:"452252"`
			Field2 int `json:"2512536"`
			Field3 int `json:"2725311"`
		} `json:"applicable_special_ids"`
		AvailableForDelivery struct {
			False int `json:"false"`
		} `json:"available_for_delivery"`
		StoreSpecificProduct struct {
			False int `json:"false"`
		} `json:"store_specific_product"`
		ApplicableSpecialTypes struct {
			Product int `json:"product"`
			Bundle  int `json:"bundle"`
		} `json:"applicable_special_types"`
		ApplicableBundleSpecialIdsQualify struct {
			Field1 int `json:"1920020"`
			Field2 int `json:"1920222"`
		} `json:"applicable_bundle_special_ids.qualify"`
		ApplicableBundleSpecialIdsDiscount struct {
			Field1 int `json:"1920020"`
			Field2 int `json:"1920222"`
		} `json:"applicable_bundle_special_ids.discount"`
	} `json:"facets"`
	FacetsStats struct {
		StoreId struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"store_id"`
		ProductId struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"product_id"`
		PercentCbd struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
			Avg float64 `json:"avg"`
			Sum float64 `json:"sum"`
		} `json:"percent_cbd"`
		PercentThc struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
			Avg float64 `json:"avg"`
			Sum float64 `json:"sum"`
		} `json:"percent_thc"`
		BucketPrice struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"bucket_price"`
		PercentThca struct {
			Min float64 `json:"min"`
			Max float64 `json:"max"`
			Avg float64 `json:"avg"`
			Sum float64 `json:"sum"`
		} `json:"percent_thca"`
		ReviewCount struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"review_count"`
		AggregateRating struct {
			Min int     `json:"min"`
			Max int     `json:"max"`
			Avg float64 `json:"avg"`
			Sum float64 `json:"sum"`
		} `json:"aggregate_rating"`
		ProductBrandId struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"product_brand_id"`
		ApplicableSpecialIds struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"applicable_special_ids"`
		ApplicableBundleSpecialIdsQualify struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"applicable_bundle_special_ids.qualify"`
		ApplicableBundleSpecialIdsDiscount struct {
			Min int `json:"min"`
			Max int `json:"max"`
			Avg int `json:"avg"`
			Sum int `json:"sum"`
		} `json:"applicable_bundle_special_ids.discount"`
	} `json:"facets_stats"`
	ExhaustiveFacetsCount bool `json:"exhaustiveFacetsCount"`
	ExhaustiveNbHits      bool `json:"exhaustiveNbHits"`
	ExhaustiveTypo        bool `json:"exhaustiveTypo"`
	Exhaustive            struct {
		FacetsCount bool `json:"facetsCount"`
		NbHits      bool `json:"nbHits"`
		Typo        bool `json:"typo"`
	} `json:"exhaustive"`
	Query               string `json:"query"`
	Params              string `json:"params"`
	ProcessingTimeMS    int    `json:"processingTimeMS"`
	ProcessingTimingsMS struct {
		Request struct {
			RoundTrip int `json:"roundTrip"`
		} `json:"_request"`
		AfterFetch struct {
			DedupFacets int `json:"dedupFacets"`
			Format      struct {
				Total int `json:"total"`
			} `json:"format"`
			Merge struct {
				Total int `json:"total"`
			} `json:"merge"`
			Total int `json:"total"`
		} `json:"afterFetch"`
		Total int `json:"total"`
	} `json:"processingTimingsMS"`
	ServerTimeMS int `json:"serverTimeMS"`
}

type InventoryPotency struct {
	PriceId    string  `json:"price_id"`
	TacPotency float64 `json:"tac_potency"`
	ThcPotency float64 `json:"thc_potency"`
	CbdPotency float64 `json:"cbd_potency"`
}
