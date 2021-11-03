package main

func main() {

	type Subsubsubid struct {
		Value int "json:value"
	}
	type Subsubid struct {
		Subsubsubid "json:subsubsubid"
	}
	type Subid struct {
		Subsubid "json:subsubid"
	}
	type Id struct {
		Subid "json:subid"
	}
	type Product struct {
		Name      string "json:name"
		StartData string "json:startDate"
	}
	type DataElement struct {
		Product "json:product"
		Active  bool     "json:active"
		Price   float64  "json:price"
		Stock   uint     "json:stock"
		Colors  []string "json:color,omitempty"
	}
	type Data struct {
		Data []DataElement "json:data"
	}
	type Json struct {
		Title string "json:title"
		Id    "json:id"
		Data  "json:data"
	}

}
