func init() {
	debug.FreeOSMemory()
	debug.SetGCPercent(0)
}

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0"{
        return "0"
    }

    n1, n2, tmp := []byte(num1), []byte(num2), make([]int, len(num1) + len(num2))

    for i := 0; i < len(n1); i++ {
        for j := 0; j < len(n2); j++ {
            tmp[i+j+1] += int(n1[i] - '0') * int(n2[j] - '0')
        }
    }

    for i := len(tmp) - 1; i > 0; i-- {
        tmp[i-1] += tmp[i] / 10
        tmp[i] = tmp[i] % 10
    }

    start := 0
	if tmp[0] == 0 {
		start = 1
	}
	res := make([]byte, len(tmp) - start)
	for i := start; i < len(tmp); i++ {
		res[i - start] = '0' + byte(tmp[i])
	}

	return string(res)
}


