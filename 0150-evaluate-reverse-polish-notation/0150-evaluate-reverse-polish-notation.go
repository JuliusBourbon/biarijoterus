import(
    "strconv" // for converting string to int
)
func evalRPN(tokens []string) int {
    temps := make([]int, 0, len(tokens))
    count := 0

    for _, token := range tokens{
        switch token {
            case "*":
                TempsLastIndex := len(temps)-1
                // do arithmetic operation on the last two numbers in temps[]
                count = temps[TempsLastIndex-1] * temps[TempsLastIndex]
                // POP last two numbers in temps[]
                temps = temps[:len(temps)-2]
                // PUSH count
                temps = append(temps, count)
            case "/":
                TempsLastIndex := len(temps)-1
                count = temps[TempsLastIndex-1] / temps[TempsLastIndex]
                temps = temps[:len(temps)-2]
                temps = append(temps, count)
            case "+":
                TempsLastIndex := len(temps)-1
                count = temps[TempsLastIndex-1] + temps[TempsLastIndex]
                temps = temps[:len(temps)-2]
                temps = append(temps, count)
            case "-":
                TempsLastIndex := len(temps)-1
                count = temps[TempsLastIndex-1] - temps[TempsLastIndex]
                temps = temps[:len(temps)-2]
                temps = append(temps, count)
            default: // operand
                num, _ := strconv.Atoi(token)
                temps = append(temps, num)
        }
    }
    return temps[0]
}