package apispace

type BaseBody struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type BirthdayFlower struct {
	BaseBody BaseBody             `json:",inline"`
	Data     []BirthdayFlowerData `json:"data"`
}

type BirthdayFlowerData struct {
	BirthdayFlower        string `json:"birthday_flower"`
	BirthdayFlowerContent string `json:"birthday_flower_content"`
	FlowerLng             string `json:"flower_lng"`
	FlowerLngContent      string `json:"flower_lng_content"`
	Birthstone            string `json:"birthstone"`
	BirthstoneContent     string `json:"birthstone_content"`
	Moon                  int    `json:"moon"`
	Day                   int    `json:"day"`
}
