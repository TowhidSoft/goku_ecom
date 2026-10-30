package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageURL"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productId int) *Product {
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}

	return nil
}

func Update(product Product) {
	for idx, p := range productList { // [a, b, c]
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product

	for _, p := range productList { // [a, b, c]
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}

	productList = tempList
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red. I love orange",
		Price:       100,
		ImgUrl:      "https://www.dole.com/sites/default/files/styles/1536w1152h-webp-80/public/media/2025-01/oranges.png.webp?itok=Gy4-VdDs-LRUygKxu-qoEwM6XP",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I hate apple",
		Price:       100,
		ImgUrl:      "https://daganghalal.blob.core.windows.net/28749/Product/1000x1000__apple1-1657088917775.jpg",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is boring. I feel bored eating banana",
		Price:       5,
		ImgUrl:      "https://daganghalal.blob.core.windows.net/28749/Product/1000x1000__apple1-1657088917775.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
}
