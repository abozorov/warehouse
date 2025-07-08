package models

type Product struct {
	ID      int     `json:"id"`
	PostProduct
}

type PostProduct struct {
	Article string  `json:"article"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
}

func PostProductToProduct(postP PostProduct) Product {
	return Product{
		PostProduct: postP,
	}
}