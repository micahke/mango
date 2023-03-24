package im


// This might need to be moved out of the `im` package


type FontAtlasItem struct {

  char string
  x int
  y int
  width int
  
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

      _atlas[item.char] = item

      xTicker += FONT_SIZE
    }

    xTicker = 0;
    yTicker += FONT_SIZE

  }

}


