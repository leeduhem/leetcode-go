package p149

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func maxPoints(points []Point) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	result := 0
	for i := 0; i < n; i++ {
		mem := map[Point]int {}
		overlap, max := 0, 0

		pi := points[i]
		for j := i+1; j < n; j++ {
			pj := points[j]
			dx := pj.X - pi.X
			dy := pj.Y - pi.Y
			if dx == 0 && dy == 0 {
				overlap++
				continue
			}

			d := gcd(dx, dy)
			if d != 0 {
				dx /= d
				dy /= d
			}

			key := Point{ dx, dy }
			if val, ok := mem[key]; ok {
				mem[key] = val + 1
			} else {
				mem[key] = 1
			}

			max = intMax(max, mem[key])
		}
		result = intMax(result, max + overlap + 1)
	}

	return result
}
