func init() {
	debug.FreeOSMemory()
	debug.SetGCPercent(0)
}

func maxArea(height []int) int {
    max, start, end := 0, 0, len(height) - 1
    for start < end {
        x := end - start
        y := 0
        if height[start] < height[end] {
            y = height[start]
            start ++
        } else {
            y = height[end]
            end --
        }

        tmp := x * y
        if tmp > max {
            max = tmp
        }
    }
    return max
}
