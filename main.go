package main

const (
	width  = 1024
	height = 512
)

var triangle = [3]point{
	{124, 20}, {81, 210}, {196, 86},
}

func main() {
	image := NewImage(width, height)
	drawTriangle(image.image, triangle[0], triangle[1], triangle[2])
	image.WriteToFile("result.png")
}
