package main

type color struct {
	ID      int
	Hex     string
	ValueTr string
	ValueEn string
}

var colorMap = map[int]color{
	1:  {1, "#000000", "Siyah", "Black"},
	2:  {2, "#FFFFFF", "Beyaz", "White"},
	3:  {3, "#808080", "Gri", "Gray"},
	4:  {4, "#C0C0C0", "Gümüş", "Silver"},
	5:  {5, "#FF0000", "Kırmızı", "Red"},
	6:  {6, "#800000", "Bordo", "Maroon"},
	7:  {7, "#FFA500", "Turuncu", "Orange"},
	8:  {8, "#FFD700", "Altın", "Gold"},
	9:  {9, "#FFFF00", "Sarı", "Yellow"},
	10: {10, "#F0E68C", "Haki", "Khaki"},
	11: {11, "#00FF00", "Yeşil", "Green"},
	12: {12, "#008000", "Koyu Yeşil", "Dark Green"},
	13: {13, "#0000FF", "Mavi", "Blue"},
	14: {14, "#000080", "Lacivert", "Navy Blue"},
	15: {15, "#800080", "Mor", "Purple"},
	16: {16, "#4B0082", "Çivit", "Indigo"},
	17: {17, "#FFC0CB", "Pembe", "Pink"},
	18: {18, "#FF69B4", "Sıcak Pembe", "Hot Pink"},
	19: {19, "#A52A2A", "Kahverengi", "Brown"},
	20: {20, "#D2691E", "Çikolata", "Chocolate"},
}

type VariationColor struct {
	VariationID int    `json:"variation_id"`
	ColorHex    string `json:"color_hex"`
	ColorID     int    `json:"color_id"`
	ProductID   int    `json:"product_id"`
	ColorName   string `json:"color_name"`
	IsProcessed int    `json:"is_processed"`
	IsMigrated  int    `json:"is_migrated"`
}

func colorAI(userPrompt string) (string, error) {
	var systemPrompt = colorSystemPrompt

	output, err := askAI(userPrompt, systemPrompt)
	if err != nil {
		return "", err
	}

	return output, nil
}
