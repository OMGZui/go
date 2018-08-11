package main

import "fmt"

const (
	WHITE  = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box //a slice of boxes

// 体积
func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}

// 设置颜色
func (b *Box) SetColor(c Color) {
	(*b).color = c
}

// 最大块颜色和体积
func (bl BoxList) Biggest() (Color, float64 ) {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k, v
}


// 涂黑
func (bl BoxList) PaintItBlack() {
	for i := range bl {
		(&bl[i]).SetColor(BLACK)
	}
}

// 实现了String()，可以用interface{}表示接受任意的参数类型
func (c Color) String() interface{} {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

	maxColor, maxVolume := boxes.Biggest()

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is", maxColor.String())
	fmt.Printf("The biggest color is %s, volume is %0.2f\n", maxColor.String(), maxVolume)

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].color.String())

	maxColor2, _ := boxes.Biggest()
	fmt.Println("Obviously, now, the biggest one is", maxColor2.String())
}
