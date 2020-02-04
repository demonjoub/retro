package schema

type Category struct {
	Id   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

type Menu struct {
	Id         int      `json:"id" xml:"id"`
	Name       string   `json:"name" xml:"name"`
	CategoryId int      `json:"category_id" xml:"category_id"`
	Recommend  int      `json:"recommend" xml:"recommend"`
	SpicyRank  int      `json:"spicy_rank" xml:"spicy_rank"`
	Price      float32  `json:"price" xml:"price"`
	Path       []string `gorm:"-" json:"path" xml:"path"`
}

type MenuImage struct {
	Id     int    `gorm:"index"`
	MenuId int    `json:"id" xml:"id"`
	Image  string `gorm:"type:varchar(100);unique_index"`
}

type Error struct {
	Error string `json:"error" xml:"error"`
}

type Response struct {
	Code    int         `json:"code" xml:"code"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}
