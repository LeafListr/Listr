// Code generated by counterfeiter. DO NOT EDIT.
package repositoryfakes

import (
	"sync"

	"github.com/Linkinlog/LeafListr/internal/models"
	"github.com/Linkinlog/LeafListr/internal/repository"
)

type FakeRepository struct {
	GetCannabinoidsStub        func(string) ([]*models.Cannabinoid, error)
	getCannabinoidsMutex       sync.RWMutex
	getCannabinoidsArgsForCall []struct {
		arg1 string
	}
	getCannabinoidsReturns struct {
		result1 []*models.Cannabinoid
		result2 error
	}
	getCannabinoidsReturnsOnCall map[int]struct {
		result1 []*models.Cannabinoid
		result2 error
	}
	GetCategoriesStub        func(string) ([]*models.Category, error)
	getCategoriesMutex       sync.RWMutex
	getCategoriesArgsForCall []struct {
		arg1 string
	}
	getCategoriesReturns struct {
		result1 []*models.Category
		result2 error
	}
	getCategoriesReturnsOnCall map[int]struct {
		result1 []*models.Category
		result2 error
	}
	GetOffersStub        func(string) ([]*models.Offer, error)
	getOffersMutex       sync.RWMutex
	getOffersArgsForCall []struct {
		arg1 string
	}
	getOffersReturns struct {
		result1 []*models.Offer
		result2 error
	}
	getOffersReturnsOnCall map[int]struct {
		result1 []*models.Offer
		result2 error
	}
	GetProductStub        func(string, string) (*models.Product, error)
	getProductMutex       sync.RWMutex
	getProductArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getProductReturns struct {
		result1 *models.Product
		result2 error
	}
	getProductReturnsOnCall map[int]struct {
		result1 *models.Product
		result2 error
	}
	GetProductsStub        func(string) ([]*models.Product, error)
	getProductsMutex       sync.RWMutex
	getProductsArgsForCall []struct {
		arg1 string
	}
	getProductsReturns struct {
		result1 []*models.Product
		result2 error
	}
	getProductsReturnsOnCall map[int]struct {
		result1 []*models.Product
		result2 error
	}
	GetProductsForCategoryStub        func(string, models.Category) ([]*models.Product, error)
	getProductsForCategoryMutex       sync.RWMutex
	getProductsForCategoryArgsForCall []struct {
		arg1 string
		arg2 models.Category
	}
	getProductsForCategoryReturns struct {
		result1 []*models.Product
		result2 error
	}
	getProductsForCategoryReturnsOnCall map[int]struct {
		result1 []*models.Product
		result2 error
	}
	GetTerpenesStub        func(string) ([]*models.Terpene, error)
	getTerpenesMutex       sync.RWMutex
	getTerpenesArgsForCall []struct {
		arg1 string
	}
	getTerpenesReturns struct {
		result1 []*models.Terpene
		result2 error
	}
	getTerpenesReturnsOnCall map[int]struct {
		result1 []*models.Terpene
		result2 error
	}
	LocationStub        func(string) (*models.Location, error)
	locationMutex       sync.RWMutex
	locationArgsForCall []struct {
		arg1 string
	}
	locationReturns struct {
		result1 *models.Location
		result2 error
	}
	locationReturnsOnCall map[int]struct {
		result1 *models.Location
		result2 error
	}
	LocationsStub        func(float64, float64) ([]*models.Location, error)
	locationsMutex       sync.RWMutex
	locationsArgsForCall []struct {
		arg1 float64
		arg2 float64
	}
	locationsReturns struct {
		result1 []*models.Location
		result2 error
	}
	locationsReturnsOnCall map[int]struct {
		result1 []*models.Location
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) GetCannabinoids(arg1 string) ([]*models.Cannabinoid, error) {
	fake.getCannabinoidsMutex.Lock()
	ret, specificReturn := fake.getCannabinoidsReturnsOnCall[len(fake.getCannabinoidsArgsForCall)]
	fake.getCannabinoidsArgsForCall = append(fake.getCannabinoidsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetCannabinoidsStub
	fakeReturns := fake.getCannabinoidsReturns
	fake.recordInvocation("GetCannabinoids", []interface{}{arg1})
	fake.getCannabinoidsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetCannabinoidsCallCount() int {
	fake.getCannabinoidsMutex.RLock()
	defer fake.getCannabinoidsMutex.RUnlock()
	return len(fake.getCannabinoidsArgsForCall)
}

func (fake *FakeRepository) GetCannabinoidsCalls(stub func(string) ([]*models.Cannabinoid, error)) {
	fake.getCannabinoidsMutex.Lock()
	defer fake.getCannabinoidsMutex.Unlock()
	fake.GetCannabinoidsStub = stub
}

func (fake *FakeRepository) GetCannabinoidsArgsForCall(i int) string {
	fake.getCannabinoidsMutex.RLock()
	defer fake.getCannabinoidsMutex.RUnlock()
	argsForCall := fake.getCannabinoidsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) GetCannabinoidsReturns(result1 []*models.Cannabinoid, result2 error) {
	fake.getCannabinoidsMutex.Lock()
	defer fake.getCannabinoidsMutex.Unlock()
	fake.GetCannabinoidsStub = nil
	fake.getCannabinoidsReturns = struct {
		result1 []*models.Cannabinoid
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetCannabinoidsReturnsOnCall(i int, result1 []*models.Cannabinoid, result2 error) {
	fake.getCannabinoidsMutex.Lock()
	defer fake.getCannabinoidsMutex.Unlock()
	fake.GetCannabinoidsStub = nil
	if fake.getCannabinoidsReturnsOnCall == nil {
		fake.getCannabinoidsReturnsOnCall = make(map[int]struct {
			result1 []*models.Cannabinoid
			result2 error
		})
	}
	fake.getCannabinoidsReturnsOnCall[i] = struct {
		result1 []*models.Cannabinoid
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetCategories(arg1 string) ([]*models.Category, error) {
	fake.getCategoriesMutex.Lock()
	ret, specificReturn := fake.getCategoriesReturnsOnCall[len(fake.getCategoriesArgsForCall)]
	fake.getCategoriesArgsForCall = append(fake.getCategoriesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetCategoriesStub
	fakeReturns := fake.getCategoriesReturns
	fake.recordInvocation("GetCategories", []interface{}{arg1})
	fake.getCategoriesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetCategoriesCallCount() int {
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	return len(fake.getCategoriesArgsForCall)
}

func (fake *FakeRepository) GetCategoriesCalls(stub func(string) ([]*models.Category, error)) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = stub
}

func (fake *FakeRepository) GetCategoriesArgsForCall(i int) string {
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	argsForCall := fake.getCategoriesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) GetCategoriesReturns(result1 []*models.Category, result2 error) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = nil
	fake.getCategoriesReturns = struct {
		result1 []*models.Category
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetCategoriesReturnsOnCall(i int, result1 []*models.Category, result2 error) {
	fake.getCategoriesMutex.Lock()
	defer fake.getCategoriesMutex.Unlock()
	fake.GetCategoriesStub = nil
	if fake.getCategoriesReturnsOnCall == nil {
		fake.getCategoriesReturnsOnCall = make(map[int]struct {
			result1 []*models.Category
			result2 error
		})
	}
	fake.getCategoriesReturnsOnCall[i] = struct {
		result1 []*models.Category
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetOffers(arg1 string) ([]*models.Offer, error) {
	fake.getOffersMutex.Lock()
	ret, specificReturn := fake.getOffersReturnsOnCall[len(fake.getOffersArgsForCall)]
	fake.getOffersArgsForCall = append(fake.getOffersArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetOffersStub
	fakeReturns := fake.getOffersReturns
	fake.recordInvocation("GetOffers", []interface{}{arg1})
	fake.getOffersMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetOffersCallCount() int {
	fake.getOffersMutex.RLock()
	defer fake.getOffersMutex.RUnlock()
	return len(fake.getOffersArgsForCall)
}

func (fake *FakeRepository) GetOffersCalls(stub func(string) ([]*models.Offer, error)) {
	fake.getOffersMutex.Lock()
	defer fake.getOffersMutex.Unlock()
	fake.GetOffersStub = stub
}

func (fake *FakeRepository) GetOffersArgsForCall(i int) string {
	fake.getOffersMutex.RLock()
	defer fake.getOffersMutex.RUnlock()
	argsForCall := fake.getOffersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) GetOffersReturns(result1 []*models.Offer, result2 error) {
	fake.getOffersMutex.Lock()
	defer fake.getOffersMutex.Unlock()
	fake.GetOffersStub = nil
	fake.getOffersReturns = struct {
		result1 []*models.Offer
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetOffersReturnsOnCall(i int, result1 []*models.Offer, result2 error) {
	fake.getOffersMutex.Lock()
	defer fake.getOffersMutex.Unlock()
	fake.GetOffersStub = nil
	if fake.getOffersReturnsOnCall == nil {
		fake.getOffersReturnsOnCall = make(map[int]struct {
			result1 []*models.Offer
			result2 error
		})
	}
	fake.getOffersReturnsOnCall[i] = struct {
		result1 []*models.Offer
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProduct(arg1 string, arg2 string) (*models.Product, error) {
	fake.getProductMutex.Lock()
	ret, specificReturn := fake.getProductReturnsOnCall[len(fake.getProductArgsForCall)]
	fake.getProductArgsForCall = append(fake.getProductArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetProductStub
	fakeReturns := fake.getProductReturns
	fake.recordInvocation("GetProduct", []interface{}{arg1, arg2})
	fake.getProductMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetProductCallCount() int {
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	return len(fake.getProductArgsForCall)
}

func (fake *FakeRepository) GetProductCalls(stub func(string, string) (*models.Product, error)) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = stub
}

func (fake *FakeRepository) GetProductArgsForCall(i int) (string, string) {
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	argsForCall := fake.getProductArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetProductReturns(result1 *models.Product, result2 error) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = nil
	fake.getProductReturns = struct {
		result1 *models.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProductReturnsOnCall(i int, result1 *models.Product, result2 error) {
	fake.getProductMutex.Lock()
	defer fake.getProductMutex.Unlock()
	fake.GetProductStub = nil
	if fake.getProductReturnsOnCall == nil {
		fake.getProductReturnsOnCall = make(map[int]struct {
			result1 *models.Product
			result2 error
		})
	}
	fake.getProductReturnsOnCall[i] = struct {
		result1 *models.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProducts(arg1 string) ([]*models.Product, error) {
	fake.getProductsMutex.Lock()
	ret, specificReturn := fake.getProductsReturnsOnCall[len(fake.getProductsArgsForCall)]
	fake.getProductsArgsForCall = append(fake.getProductsArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetProductsStub
	fakeReturns := fake.getProductsReturns
	fake.recordInvocation("GetProducts", []interface{}{arg1})
	fake.getProductsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetProductsCallCount() int {
	fake.getProductsMutex.RLock()
	defer fake.getProductsMutex.RUnlock()
	return len(fake.getProductsArgsForCall)
}

func (fake *FakeRepository) GetProductsCalls(stub func(string) ([]*models.Product, error)) {
	fake.getProductsMutex.Lock()
	defer fake.getProductsMutex.Unlock()
	fake.GetProductsStub = stub
}

func (fake *FakeRepository) GetProductsArgsForCall(i int) string {
	fake.getProductsMutex.RLock()
	defer fake.getProductsMutex.RUnlock()
	argsForCall := fake.getProductsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) GetProductsReturns(result1 []*models.Product, result2 error) {
	fake.getProductsMutex.Lock()
	defer fake.getProductsMutex.Unlock()
	fake.GetProductsStub = nil
	fake.getProductsReturns = struct {
		result1 []*models.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProductsReturnsOnCall(i int, result1 []*models.Product, result2 error) {
	fake.getProductsMutex.Lock()
	defer fake.getProductsMutex.Unlock()
	fake.GetProductsStub = nil
	if fake.getProductsReturnsOnCall == nil {
		fake.getProductsReturnsOnCall = make(map[int]struct {
			result1 []*models.Product
			result2 error
		})
	}
	fake.getProductsReturnsOnCall[i] = struct {
		result1 []*models.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProductsForCategory(arg1 string, arg2 models.Category) ([]*models.Product, error) {
	fake.getProductsForCategoryMutex.Lock()
	ret, specificReturn := fake.getProductsForCategoryReturnsOnCall[len(fake.getProductsForCategoryArgsForCall)]
	fake.getProductsForCategoryArgsForCall = append(fake.getProductsForCategoryArgsForCall, struct {
		arg1 string
		arg2 models.Category
	}{arg1, arg2})
	stub := fake.GetProductsForCategoryStub
	fakeReturns := fake.getProductsForCategoryReturns
	fake.recordInvocation("GetProductsForCategory", []interface{}{arg1, arg2})
	fake.getProductsForCategoryMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetProductsForCategoryCallCount() int {
	fake.getProductsForCategoryMutex.RLock()
	defer fake.getProductsForCategoryMutex.RUnlock()
	return len(fake.getProductsForCategoryArgsForCall)
}

func (fake *FakeRepository) GetProductsForCategoryCalls(stub func(string, models.Category) ([]*models.Product, error)) {
	fake.getProductsForCategoryMutex.Lock()
	defer fake.getProductsForCategoryMutex.Unlock()
	fake.GetProductsForCategoryStub = stub
}

func (fake *FakeRepository) GetProductsForCategoryArgsForCall(i int) (string, models.Category) {
	fake.getProductsForCategoryMutex.RLock()
	defer fake.getProductsForCategoryMutex.RUnlock()
	argsForCall := fake.getProductsForCategoryArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) GetProductsForCategoryReturns(result1 []*models.Product, result2 error) {
	fake.getProductsForCategoryMutex.Lock()
	defer fake.getProductsForCategoryMutex.Unlock()
	fake.GetProductsForCategoryStub = nil
	fake.getProductsForCategoryReturns = struct {
		result1 []*models.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetProductsForCategoryReturnsOnCall(i int, result1 []*models.Product, result2 error) {
	fake.getProductsForCategoryMutex.Lock()
	defer fake.getProductsForCategoryMutex.Unlock()
	fake.GetProductsForCategoryStub = nil
	if fake.getProductsForCategoryReturnsOnCall == nil {
		fake.getProductsForCategoryReturnsOnCall = make(map[int]struct {
			result1 []*models.Product
			result2 error
		})
	}
	fake.getProductsForCategoryReturnsOnCall[i] = struct {
		result1 []*models.Product
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetTerpenes(arg1 string) ([]*models.Terpene, error) {
	fake.getTerpenesMutex.Lock()
	ret, specificReturn := fake.getTerpenesReturnsOnCall[len(fake.getTerpenesArgsForCall)]
	fake.getTerpenesArgsForCall = append(fake.getTerpenesArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetTerpenesStub
	fakeReturns := fake.getTerpenesReturns
	fake.recordInvocation("GetTerpenes", []interface{}{arg1})
	fake.getTerpenesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) GetTerpenesCallCount() int {
	fake.getTerpenesMutex.RLock()
	defer fake.getTerpenesMutex.RUnlock()
	return len(fake.getTerpenesArgsForCall)
}

func (fake *FakeRepository) GetTerpenesCalls(stub func(string) ([]*models.Terpene, error)) {
	fake.getTerpenesMutex.Lock()
	defer fake.getTerpenesMutex.Unlock()
	fake.GetTerpenesStub = stub
}

func (fake *FakeRepository) GetTerpenesArgsForCall(i int) string {
	fake.getTerpenesMutex.RLock()
	defer fake.getTerpenesMutex.RUnlock()
	argsForCall := fake.getTerpenesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) GetTerpenesReturns(result1 []*models.Terpene, result2 error) {
	fake.getTerpenesMutex.Lock()
	defer fake.getTerpenesMutex.Unlock()
	fake.GetTerpenesStub = nil
	fake.getTerpenesReturns = struct {
		result1 []*models.Terpene
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetTerpenesReturnsOnCall(i int, result1 []*models.Terpene, result2 error) {
	fake.getTerpenesMutex.Lock()
	defer fake.getTerpenesMutex.Unlock()
	fake.GetTerpenesStub = nil
	if fake.getTerpenesReturnsOnCall == nil {
		fake.getTerpenesReturnsOnCall = make(map[int]struct {
			result1 []*models.Terpene
			result2 error
		})
	}
	fake.getTerpenesReturnsOnCall[i] = struct {
		result1 []*models.Terpene
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Location(arg1 string) (*models.Location, error) {
	fake.locationMutex.Lock()
	ret, specificReturn := fake.locationReturnsOnCall[len(fake.locationArgsForCall)]
	fake.locationArgsForCall = append(fake.locationArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LocationStub
	fakeReturns := fake.locationReturns
	fake.recordInvocation("Location", []interface{}{arg1})
	fake.locationMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) LocationCallCount() int {
	fake.locationMutex.RLock()
	defer fake.locationMutex.RUnlock()
	return len(fake.locationArgsForCall)
}

func (fake *FakeRepository) LocationCalls(stub func(string) (*models.Location, error)) {
	fake.locationMutex.Lock()
	defer fake.locationMutex.Unlock()
	fake.LocationStub = stub
}

func (fake *FakeRepository) LocationArgsForCall(i int) string {
	fake.locationMutex.RLock()
	defer fake.locationMutex.RUnlock()
	argsForCall := fake.locationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) LocationReturns(result1 *models.Location, result2 error) {
	fake.locationMutex.Lock()
	defer fake.locationMutex.Unlock()
	fake.LocationStub = nil
	fake.locationReturns = struct {
		result1 *models.Location
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) LocationReturnsOnCall(i int, result1 *models.Location, result2 error) {
	fake.locationMutex.Lock()
	defer fake.locationMutex.Unlock()
	fake.LocationStub = nil
	if fake.locationReturnsOnCall == nil {
		fake.locationReturnsOnCall = make(map[int]struct {
			result1 *models.Location
			result2 error
		})
	}
	fake.locationReturnsOnCall[i] = struct {
		result1 *models.Location
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Locations(arg1 float64, arg2 float64) ([]*models.Location, error) {
	fake.locationsMutex.Lock()
	ret, specificReturn := fake.locationsReturnsOnCall[len(fake.locationsArgsForCall)]
	fake.locationsArgsForCall = append(fake.locationsArgsForCall, struct {
		arg1 float64
		arg2 float64
	}{arg1, arg2})
	stub := fake.LocationsStub
	fakeReturns := fake.locationsReturns
	fake.recordInvocation("Locations", []interface{}{arg1, arg2})
	fake.locationsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) LocationsCallCount() int {
	fake.locationsMutex.RLock()
	defer fake.locationsMutex.RUnlock()
	return len(fake.locationsArgsForCall)
}

func (fake *FakeRepository) LocationsCalls(stub func(float64, float64) ([]*models.Location, error)) {
	fake.locationsMutex.Lock()
	defer fake.locationsMutex.Unlock()
	fake.LocationsStub = stub
}

func (fake *FakeRepository) LocationsArgsForCall(i int) (float64, float64) {
	fake.locationsMutex.RLock()
	defer fake.locationsMutex.RUnlock()
	argsForCall := fake.locationsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRepository) LocationsReturns(result1 []*models.Location, result2 error) {
	fake.locationsMutex.Lock()
	defer fake.locationsMutex.Unlock()
	fake.LocationsStub = nil
	fake.locationsReturns = struct {
		result1 []*models.Location
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) LocationsReturnsOnCall(i int, result1 []*models.Location, result2 error) {
	fake.locationsMutex.Lock()
	defer fake.locationsMutex.Unlock()
	fake.LocationsStub = nil
	if fake.locationsReturnsOnCall == nil {
		fake.locationsReturnsOnCall = make(map[int]struct {
			result1 []*models.Location
			result2 error
		})
	}
	fake.locationsReturnsOnCall[i] = struct {
		result1 []*models.Location
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getCannabinoidsMutex.RLock()
	defer fake.getCannabinoidsMutex.RUnlock()
	fake.getCategoriesMutex.RLock()
	defer fake.getCategoriesMutex.RUnlock()
	fake.getOffersMutex.RLock()
	defer fake.getOffersMutex.RUnlock()
	fake.getProductMutex.RLock()
	defer fake.getProductMutex.RUnlock()
	fake.getProductsMutex.RLock()
	defer fake.getProductsMutex.RUnlock()
	fake.getProductsForCategoryMutex.RLock()
	defer fake.getProductsForCategoryMutex.RUnlock()
	fake.getTerpenesMutex.RLock()
	defer fake.getTerpenesMutex.RUnlock()
	fake.locationMutex.RLock()
	defer fake.locationMutex.RUnlock()
	fake.locationsMutex.RLock()
	defer fake.locationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ repository.Repository = new(FakeRepository)
