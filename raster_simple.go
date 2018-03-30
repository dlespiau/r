package main

import (
	"image"
	"image/color"
)

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int32) int32 {
	return min(a, min(b, c))
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func max3(a, b, c int32) int32 {
	return max(a, max(b, c))
}

func orient2D(a, b, c point) int32 {
	return (b.X-a.X)*(a.Y-c.Y) - (a.Y-b.Y)*(c.X-a.X)
}

func drawTriangle(image *image.RGBA, v0, v1, v2 point) {
	red := color.RGBA{200, 0, 0, 255}

	// Compute AABB.
	minX := min3(v0.X, v1.X, v2.X)
	minY := min3(v0.Y, v1.Y, v2.Y)
	maxX := max3(v0.X, v1.X, v2.X)
	maxY := max3(v0.Y, v1.Y, v2.Y)

	// Clip against image extends.
	minX = max(minX, 0)
	minY = max(minY, 0)
	maxX = min(maxX, width-1)
	maxY = min(maxY, height-1)

	p := point{}
	for p.Y = minY; p.Y <= maxY; p.Y++ {
		for p.X = minX; p.X < maxX; p.X++ {
			// Determine barycentric coordinates
			w0 := orient2D(v1, v2, p)
			w1 := orient2D(v2, v0, p)
			w2 := orient2D(v0, v1, p)

			// if p is on or inside the triangle
			if w0 >= 0 && w1 >= 0 && w2 >= 0 {
				image.SetRGBA((int)(p.X), (int)(p.Y), red)
			}
		}
	}
}
