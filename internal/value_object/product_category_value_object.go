package valueobject

import "fmt"

const (
	PRODUCT_CATEGORY_CLOTHING = iota
	PRODUCT_CATEGORY_ACCESSORIS
	PRODUCT_CATEGORY_FOOTWEAR
	PRODUCT_CATEGORY_BEVERAGES
)

var productCategoryType = map[int]string{
	PRODUCT_CATEGORY_CLOTHING:   "Clothing",
	PRODUCT_CATEGORY_ACCESSORIS: "Accessories",
	PRODUCT_CATEGORY_FOOTWEAR:   "Footwear",
	PRODUCT_CATEGORY_BEVERAGES:  "Beverages",
}

type ProductCategoryType struct {
	types int
}

func CheckProductCategory(types interface{}) error {
	switch t := types.(type) {
	case int:
		if _, ok := productCategoryType[t]; ok {
			return nil
		}
	case string:
		for _, v := range productCategoryType {
			if v == t {
				return nil
			}
		}
	}

	return fmt.Errorf("invalid product category type '%v'", types)
}

func NewProductCategoryType(types interface{}) (*ProductCategoryType, error) {
	switch t := types.(type) {
	case int:
		if _, ok := productCategoryType[t]; ok {
			return &ProductCategoryType{
				types: t,
			}, nil
		}
	case string:
		for k, v := range productCategoryType {
			if v == t {
				return &ProductCategoryType{
					types: k,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("invalid product category type '%v'", types)
}

func (b *ProductCategoryType) GetTypesInt() int {
	return b.types
}

func (b *ProductCategoryType) GetTypesStr() string {
	return productCategoryType[b.types]
}
