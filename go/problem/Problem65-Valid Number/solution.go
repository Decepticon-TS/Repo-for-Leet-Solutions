func isNumber(s string) bool {
    i, n := 0, len(s)

    // Skip leading whitespace
    for i < n && s[i] == ' ' {
        i++
    }

    // Skip leading sign
    if i < n && (s[i] == '+' || s[i] == '-') {
        i++
    }

    isNumeric := false

    // Integer part
    for i < n && s[i] >= '0' && s[i] <= '9' {
        i++
        isNumeric = true
    }

    // Decimal part
    if i < n && s[i] == '.' {
        i++
        for i < n && s[i] >= '0' && s[i] <= '9' {
            i++
            isNumeric = true
        }
    }

    // Exponent part
    if isNumeric && i < n && (s[i] == 'e' || s[i] == 'E') {
        i++
        if i < n && (s[i] == '+' || s[i] == '-') {
            i++
        }

        isNumeric = false
        for i < n && s[i] >= '0' && s[i] <= '9' {
            i++
            isNumeric = true
        }
    }

    // Skip trailing whitespace
    for i < n && s[i] == ' ' {
        i++
    }

    return isNumeric && i == n
}
