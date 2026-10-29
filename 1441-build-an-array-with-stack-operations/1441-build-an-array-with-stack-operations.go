func buildArray(target []int, n int) []string {
    stack := make([]int, 0, n)
    output := make([]string, 0, n)
    sameValue := false

    for i := 0; i < n; i++ {
        stack = append(stack, i+1) // PUSH
        output = append(output, "Push")


        // if len(stack) == len(target) {
        for j := 0; j < len(stack); j++ {
            // in case first index is not 1
            if stack[0] != target[0] {
            }

            if stack[j] != target[j] {
                if j == 0 {
                    stack = stack[:0] // POP
                    sameValue = false
                    break
                } else {
                    stack = stack[:len(stack)-1] // POP
                    sameValue = false
                    break
                }
            } else {
                sameValue = true
            }
        }

        if sameValue == false {
            output = append(output, "Pop")
        }
        // }

        if len(stack) == len(target) && sameValue == true {
            break
        }
    }

    return output
}