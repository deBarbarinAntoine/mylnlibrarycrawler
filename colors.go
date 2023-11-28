package webscraping

import (
	"strconv"
)

type Color struct {
	Name string
	R    int
	G    int
	B    int
}

var (
	//Red colors
	Lightsalmon = Color{"Lightsalmon", 255, 160, 122}
	Salmon      = Color{"Salmon", 250, 128, 114}
	Crimson     = Color{"Crimson", 220, 20, 60}
	Red         = Color{"Red", 255, 0, 0}
	//All Reds
	Reds = []Color{Lightsalmon, Salmon, Crimson, Red}

	//Orange colors
	Coral      = Color{"Coral", 255, 127, 80}
	Tomato     = Color{"Tomato", 255, 99, 71}
	Orangered  = Color{"Orangered", 255, 69, 0}
	Orange     = Color{"Orange", 255, 165, 0}
	Darkorange = Color{"Darkorange", 255, 140, 0}
	//All Oranges
	Oranges = []Color{Coral, Tomato, Orangered, Orange, Darkorange}

	//Yellow colors
	Palegoldenrod = Color{"Palegoldenrod", 238, 232, 170}
	Moccasin      = Color{"Moccasin", 255, 228, 181}
	Darkkhaki     = Color{"Darkkhaki", 189, 183, 107}
	Gold          = Color{"Gold", 255, 215, 0}
	Yellow        = Color{"Yellow", 255, 255, 0}
	//All Yellows
	Yellows = []Color{Palegoldenrod, Moccasin, Darkkhaki, Gold, Yellow}

	//Green colors
	Lawngreen      = Color{"Lawngreen", 124, 252, 0}
	Limegreen      = Color{"Limegreen", 50, 205, 50}
	Forestgreen    = Color{"Forestgreen", 34, 139, 34}
	Darkgreen      = Color{"Darkgreen", 0, 100, 0}
	Springgreen    = Color{"Springgreen", 0, 255, 127}
	Lightgreen     = Color{"Lightgreen", 144, 238, 144}
	Mediumseagreen = Color{"Mediumseagreen", 60, 179, 113}
	Olive          = Color{"Olive", 128, 128, 0}
	Darkolivegreen = Color{"Darkolivegreen", 85, 107, 47}
	Olivedrab      = Color{"Olivedrab", 107, 142, 35}
	//All Greens
	Greens = []Color{Lawngreen, Limegreen, Forestgreen, Darkgreen, Springgreen, Lightgreen, Mediumseagreen, Olive, Darkolivegreen, Olivedrab}

	//Cyan colors
	Cyan       = Color{"Cyan", 0, 255, 255}
	Aquamarine = Color{"Aquamarine", 127, 255, 212}
	Turquoise  = Color{"Turquoise", 64, 224, 208}
	Teal       = Color{"Teal", 0, 128, 128}
	//All Cyans
	Cyans = []Color{Cyan, Aquamarine, Turquoise, Teal}

	//Blue colors
	Deepskyblue    = Color{"Deepskyblue", 0, 191, 255}
	Cornflowerblue = Color{"Cornflowerblue", 100, 149, 237}
	Royalblue      = Color{"Royalblue", 65, 105, 225}
	Blue           = Color{"Blue", 0, 0, 255}
	//All Blues
	Blues = []Color{Deepskyblue, Cornflowerblue, Royalblue, Blue}

	//Purple colors
	Plum         = Color{"Plum", 221, 160, 221}
	Magenta      = Color{"Magenta", 255, 0, 255}
	Mediumpurple = Color{"Mediumpurple", 147, 112, 219}
	Darkorchid   = Color{"Darkorchid", 153, 50, 204}
	Darkmagenta  = Color{"Darkmagenta", 139, 0, 139}
	Purple       = Color{"Purple", 128, 0, 128}
	//All Purples
	Purples = []Color{Plum, Magenta, Mediumpurple, Darkorchid, Darkmagenta, Purple}

	//Pink colors
	Pink     = Color{"Pink", 255, 192, 203}
	Hotpink  = Color{"Hotpink", 255, 105, 180}
	Deeppink = Color{"Deeppink", 255, 20, 147}
	//All Pinks
	Pinks = []Color{Pink, Hotpink, Deeppink}

	//White colors
	White     = Color{"White", 255, 255, 255}
	Beige     = Color{"Beige", 245, 245, 220}
	Mistyrose = Color{"Mistyrose", 255, 228, 225}
	//All Whites
	Whites = []Color{White, Beige, Mistyrose}

	//Gray colors
	Darkgray  = Color{"Darkgray", 169, 169, 169}
	Slategray = Color{"Slategray", 112, 128, 144}
	//All Grays
	Grays = []Color{Darkgray, Slategray}

	//Brown colors
	Tan         = Color{"Tan", 210, 180, 140}
	Peru        = Color{"Peru", 205, 133, 63}
	Chocolate   = Color{"Chocolate", 210, 105, 30}
	Sienna      = Color{"Sienna", 160, 82, 45}
	Brown       = Color{"Brown", 165, 42, 42}
	Saddlebrown = Color{"Saddlebrown", 139, 69, 19}
	//All Browns
	Browns = []Color{Tan, Peru, Chocolate, Sienna, Brown, Saddlebrown}
)

const CLEARCOLOR = "\033[0m"

// colorCode generates a string in the adequate format for the terminal to change the color using the RGB components present in the Color struct.
func ColorCode(color Color) string {
	return "\033[38;2;" + strconv.Itoa(color.R) + ";" + strconv.Itoa(color.G) + ";" + strconv.Itoa(color.B) + "m"
}
