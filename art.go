package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func loadImg(filePath string) (image.Image, error) { //Загрузка фотки
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	image, _, err := image.Decode(f) //В конце попробовать вернуть string - 2 знач
	return image, err
}

func grayscale(c color.Color) int { //F = 0.30*R + 0.59*G + 0.11*B
	// Получаем на вход цвет пикселя и преобразуем его в оттеенок серого
	r, g, b, _ := c.RGBA()                                         //В конце попробовать что нибудь с альфой
	return int(0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) //Возвращаем оттенок
}

func avgPixel(img image.Image, x, y, w, h int) int { //x y - точки начала области, w на h пикселей
	cnt, sum, max := 0, 0, img.Bounds().Max //(1218,1624)(me.png)
	for i := x; i < x+w && i < max.X; i++ { //Проход по всем пикселям в указанной области
		for j := y; j < y+h && j < max.Y; j++ {
			sum += grayscale(img.At(i, j)) //Значение преобразованного серого пикселя добавляем в sum
			cnt++                          //++ каждый обработанный пиксель
		}
	}
	return sum / cnt //среднее значение пикселей в области = сумма значений пикселей / количество пикселей
}

func main() {
	var phName string
	fmt.Print("input image name/path: ")
	fmt.Scan(&phName)
	image, err := loadImg(phName)
	check(err)

	//ramp := "@%#*+=-:. "
	ramp := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,. "
	max := image.Bounds().Max            //(1218,1624)(me.png) - получаем доступ по max.X, max.Y
	scaleX, scaleY := 5, 10              //10, 5 - шаг между пикселями, 6,4 - + для me.png, 10,12 ok для roma.jpg, 5,4 - pres.jpg
	for y := 0; y < max.Y; y += scaleX { //(406, 541) / 16 = 25, 34
		for x := 0; x < max.X; x += scaleX {
			c := avgPixel(image, x, y, scaleX, scaleY)
			fmt.Print(string(ramp[len(ramp)*c/65536]))
		}
		fmt.Println()
	}

}
