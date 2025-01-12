func init() {
	debug.FreeOSMemory()
	debug.SetGCPercent(0)
}

func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    n1, n2, tmp := []byte(num1), []byte(num2), make([]int, len(num1)+len(num2))

    for i := len(n1) - 1; i >= 0; i-- {
        for j := len(n2) - 1; j >= 0; j-- {
            mul := int(n1[i]-'0') * int(n2[j]-'0')
            sum := mul + tmp[i+j+1]
            tmp[i+j+1] = sum % 10
            tmp[i+j] += sum / 10
        }
    }

    if tmp[0] == 0 {
        tmp = tmp[1:]
    }

    res := make([]byte, len(tmp))
    for i := 0; i < len(tmp); i++ {
        res[i] = '0' + byte(tmp[i])
    }

    return string(res)
}

