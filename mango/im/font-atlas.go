package im

// This might need to be moved out of the `im` package

type FontAtlasItem struct {
	char  string
	x     int
	y     int
	width int

	texturePositions [4]float32
}

const FONT_SIZE = 32

var chars = [][]string{
	{"X", "Y", "Z", "[", "\\", "]", "^", "_"},
	{"P", "Q", "R", "S", "T", "U", "V", "W"},
	{"H", "I", "J", "K", "L", "M", "N", "O"},
	{"@", "A", "B", "C", "D", "E", "F", "G"},
	{"8", "9", ":", ";", "<", "=", ">", "?"},
	{"0", "1", "2", "3", "4", "5", "6", "7"},
	{"(", ")", "*", "+", ",", "-", ".", "/"},
	{" ", "!", "\"", "#", "$", "%", "&", "'"},
}

var _atlas map[string]*FontAtlasItem

func InitFontAtlas() {

	_atlas = make(map[string]*FontAtlasItem)

	xTicker := 0
	yTicker := 0

	for row := 0; row < len(chars); row++ {

		for col := 0; col < len(chars[0]); col++ {

			item := new(FontAtlasItem)
			item.char = chars[row][col]

			item.x = xTicker
			item.y = yTicker

			item.width = FONT_SIZE

			item.texturePositions = [4]float32{
				float32(item.x) / float32(len(chars)*FONT_SIZE),     // X
				float32(item.y) / float32(len(chars)*FONT_SIZE),     // Y
				float32(item.width) / float32(len(chars)*FONT_SIZE), // WIDTH
				float32(FONT_SIZE) / float32(len(chars)*FONT_SIZE),  // HEIGHT
			}

			item.texturePositions[2] += item.texturePositions[0]
			item.texturePositions[3] += item.texturePositions[1]

			_atlas[item.char] = item

			xTicker += FONT_SIZE
		}

		xTicker = 0
		yTicker += FONT_SIZE

	}

}
