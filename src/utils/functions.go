package client_server

import (
    "math"
    "sort"
    "regexp"
    "slices"
    "unicode"
    "strings"
    "strconv"
    "math/rand"
    "math/cmplx"
)

// Helpers
func isPerfectSquare(num uint32) bool {
    var sqrt float64 = math.Sqrt(float64(num))
    return sqrt == math.Floor(sqrt)
}

func isPrime(num int32) bool {
    if num == 2 {
        return true
    } else if num < 2 || num % 2 == 0 {
        return false
    }
    for i := int32(3); i * i <= num; i += 2 {
        if num % i == 0 {
            return false
        }
    }
    return true
}

func divisors(num int32) []int32 {
    var divs []int32
    for i := int32(1); i * i <= num; i++ {
        if num % i == 0 {
            divs = append(divs, i)
            if i * i != num {
                divs = append(divs, num / i)
            }
        }
    }
    return divs
}

func reverseNum(num int32) int32 {
    var negative bool = false
    if num < 0 {
        negative = true
        num = -num
    }
    var reversed int32 = 0
    for num > 0 {
        reversed = reversed * 10 + num % 10
        num /= 10
    }
    if negative {
        reversed = -reversed
    }
    return reversed
}

func permuteNum(num int32, steps int) int32 {
    var negative bool = false
    if num < 0 {
        negative = true
        num = -num
    }
    var digits string = strconv.Itoa(int(num))
    var n int = len(digits)
    steps = steps % n
    var permuted string = digits[n-steps:] + digits[:n-steps]
    result, _ := strconv.Atoi(permuted)
    if negative {
        result = -result
    }
    return int32(result)
}

func duplicateFirstDigit (num int32) int32 {
    var negative bool = false
    if num < 0 {
        negative = true
        num = -num
    }
    var result, p int32 = 0, 1
    for num > 0 {
        result = num % 10 * p + result
        p *= 10
        if num < 10 {
            result = num * p + result
        }
        num /= 10
    }
    if (negative) {
        result = -result
    }
    return result
}

func sumDigits(num int32) int32 {
    if num < 0 {
        num = -num
    }
    var sum int32 = 0
    for num > 0 {
        sum += num % 10
        num /= 10
    }
    return sum
}

func nrDigits(num int32) int32 {
    if num < 0 {
        num = -num
    } else if num == 0 {
        return 1
    }
    var count int32 = 0
    for num > 0 {
        count++
        num /= 10
    }
    return count
}

func extractNum(word string) (uint32, bool) {
    var num uint32 = 0
    var found bool = false
    for _, c := range word {
        if c >= '0' && c <= '9' {
            num = num * 10 + uint32(c - '0')
            found = true
        }
    }
    return num, found
}

func convertBinary(bits string) int32 {
    var num int32 = 0
    for _, bit := range bits {
        num = num * 2 + int32(bit - '0')
    }
    return num
}

func intersection(first, second []int32) []int32 {
    var result []int32 = []int32{}
    var bucket map[int32]bool = make(map[int32]bool)
    for _, i := range first {
        for _, j := range second {
            if i == j && !bucket[i] {
                result = append(result, i)
                bucket[i] = true
            }
        }
    }
    return result
}

// Main functions
func MixWords(words []string) []string {
    if len(words) == 0 {
        return nil
    }

    var result []string = make([]string, len(words[0]))
    for i := range result {
        for _, word := range words {
            result[i] += string(word[i])
        }
    }

    return result
}

func CountPerfectSquares(words []string) int32 {
    var count int32 = 0
    for _, word := range words {
        num, found := extractNum(word)
        if found && isPerfectSquare(num) {
            count++
        }
    }
    return count
}

func SumReversed(nums []int32) int32 {
    var sum int32 = 0
    for _, num := range nums {
        sum += reverseNum(num)
    }
    return sum
}

func AverageWithinRange(nums []int32) int32 {
    if len(nums) < 2 {
        return 0
    }

    var lowBound, highBound, sum, count int32 = nums[0], nums[1], 0, 0
    for _, num := range nums[2:] {
        var sumDigitsNum int32 = sumDigits(num)
        if sumDigitsNum >= lowBound && sumDigitsNum <= highBound {
            sum += num
            count++
        }
    }
    if count == 0 {
        return 0
    }
    return sum / count
}

func ConvertBinary(words []string) []int32 {
    var result []int32
    hasOnlyBinary := regexp.MustCompile("^[01]+$")
    for _, word := range words {
        if hasOnlyBinary.MatchString(word) {
            result = append(result, convertBinary(word))
        }
    }
    return result
}

func CaesarCipher(text string, direction string, shift int32) string {
    var result string = ""
    for _, c := range text {
        if !unicode.IsLetter(c) {
            result += string(c)
            continue
        }
        if strings.ToLower(direction) == "left" {
            c = (c - 'a' - shift + 26) % 26 + 'a'
        } else
        {
            c = (c - 'a' + shift) % 26 + 'a'
        }
        result += string(c)
    }
    return result
}

func DecodeText(encoded string) string {
    var decoded string = ""
    var length int32 = 0
    for _, c := range encoded {
        if unicode.IsDigit(c) {
            length = length * 10 + int32(c - '0')
        } else {
            decoded += strings.Repeat(string(c), int(length))
            length = 0
        }
    }
    return decoded
}

func CountDigitsPrime(nums []int32) int32 {
    var count int32 = 0
    for _, num := range nums {
        if isPrime(num) {
            count += nrDigits(num)
        }
    }
    return count
}

func CountEvenVowels(words []string) int32 {
    var count int32 = 0
    var vowels string = "aeiouAEIOU"
    for _, word := range words {
        var vowelsCount int32 = 0
        for i, c := range word {
            if i % 2 == 0 && strings.ContainsRune(vowels, c) {
                vowelsCount++
            }
        }
        if vowelsCount > 0 && vowelsCount % 2 == 0 {
            count++
        }
    }
    return count
}

func GCD(words []string) int32 {
    var nums []int32
    for _, word := range words {
        num, found := extractNum(word)
        if found {
            nums = append(nums, int32(num))
        }
    }
    if len(nums) == 0 {
        return 0
    }
    var divs []int32 = divisors(nums[0])
    for _, num := range nums[1:] {
        divs = intersection(divs, divisors(num))
    }
    if len(divs) == 0 {
        return 0
    }
    return slices.Max(divs)
}

func SumPermuted(nums []int32, shift int) int32 {
    var sum int32 = 0
    for _, num := range nums {
        sum += permuteNum(num, shift)
    }
    return sum
}

func SumDuplicatedFirst(nums []int32) int32 {
    var sum int32 = 0
    for _, num := range nums {
        sum += duplicateFirstDigit(num)
    }
    return sum
}

func FilterComplexNumbers(nums []int32) []float64 {
    var lowBound, highBound int32 = nums[0], nums[1]
    var modules []float64
    for i := 2; i < len(nums); i += 2 {
        var module float64 = cmplx.Abs(complex(float64(nums[i]), float64(nums[i+1])))

        if module < float64(lowBound) || module > float64(highBound) {
            modules = append(modules, module)
        }
    }

    sort.Float64s(modules)
    return modules
}

func FilterValidPasswords(passwords []string) []string {
    hasLower := regexp.MustCompile(`[a-z]`)
    hasUpper := regexp.MustCompile(`[A-Z]`)
    hasDigit := regexp.MustCompile(`\d`)
    hasSymbol := regexp.MustCompile(`[\W_]`)
    var validPasswords []string
    for _, password := range passwords {
        if hasLower.MatchString(password) &&
           hasUpper.MatchString(password) &&
           hasDigit.MatchString(password) &&
           hasSymbol.MatchString(password) {
            validPasswords = append(validPasswords, password)
        }
    }
    return validPasswords
}

func GenerateRandomPasswords(characters []string) []string {
    hasLower := regexp.MustCompile(`[a-z]`)
    hasUpper := regexp.MustCompile(`[A-Z]`)
    hasDigit := regexp.MustCompile(`\d`)
    hasSymbol := regexp.MustCompile(`[\W_]`)
    var foundLower, foundUpper, foundDigit, foundSymbol bool = false, false, false, false

    for _, ch := range characters {
        if hasLower.MatchString(ch) {
            foundLower = true
        }
        if hasUpper.MatchString(ch) {
            foundUpper = true
        }
        if hasDigit.MatchString(ch) {
            foundDigit = true
        }
        if hasSymbol.MatchString(ch) {
            foundSymbol = true
        }
    }

    if !(foundLower && foundUpper && foundDigit && foundSymbol) {
        return []string{}
    }

    var passwords []string
    var numPasswords int = rand.Intn(10) + 1

    for i := 0; i < numPasswords; i++ {
        var passwordLength int = rand.Intn(11) + 5
        var password []byte = make([]byte, passwordLength)

        for j := 0; j < passwordLength; j++ {
            password[j] = characters[rand.Intn(len(characters))][0]
        }

        var passwordStr string = string(password)

        if !(hasLower.MatchString(passwordStr) &&
             hasUpper.MatchString(passwordStr) &&
             hasDigit.MatchString(passwordStr) &&
             hasSymbol.MatchString(passwordStr)) {
            i--
        } else {
            passwords = append(passwords, passwordStr)
        }
    }

    return passwords
}