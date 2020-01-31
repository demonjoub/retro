package schema

type Category struct {
	Id   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Menu struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	CategoryId  int    `json:"category_id" xml:"category_id"`
	Recommend   int    `json:"recommend" xml:"recommend"`
	SpicyRank   int    `json:"spicy_rank" xml:"spicy_rank"`
	Price       int    `json:"price" xml:"price"`
	ImageAvatar string `json:"image_avatar" xml:"image_avatar"`
	Images      string `json:"images" xml:"images"`
}

type Error struct {
	Error string `json:"error" xml:"error"`
}

type Response struct {
	Data    interface{} `json:"data" xml:"data"`
	Message string      `json:"message" xml:"message"`
}
