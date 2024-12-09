package client_server

func MixWordsWrapper(input interface{}) interface{} {
    var words []string = input.([]string)
    return MixWords(words)
}

func CountPerfectSquaresWrapper(input interface{}) interface{} {
    var words []string = input.([]string)
    return CountPerfectSquares(words)
}

func SumReversedWrapper(input interface{}) interface{} {
    var nums []int32 = input.([]int32)
    return SumReversed(nums)
}

func AverageWithinRangeWrapper(input interface{}) interface{} {
    var nums []int32 = input.([]int32)
    return AverageWithinRange(nums)
}

func ConvertBinaryWrapper(input interface{}) interface{} {
    var words []string = input.([]string)
    return ConvertBinary(words)
}

func CaesarCipherWrapper(input interface{}) interface{} {
    var (
        args      []interface{} = input.([]interface{})
        text      string        = args[0].(string)
        direction string        = args[1].(string)
        shift     int32         = args[2].(int32)
    )
    return CaesarCipher(text, direction, shift)
}

func DecodeTextWrapper(input interface{}) interface{} {
    var encoded string = input.(string)
    return DecodeText(encoded)
}

func CountDigitsPrimeWrapper(input interface{}) interface{} {
    var nums []int32 = input.([]int32)
    return CountDigitsPrime(nums)
}

func CountEvenVowelsWrapper(input interface{}) interface{} {
    var words []string = input.([]string)
    return CountEvenVowels(words)
}

func GCDWrapper(input interface{}) interface{} {
    var words []string = input.([]string)
    return GCD(words)
}

func SumPermutedWrapper(input interface{}) interface{} {
    var (
        args  []interface{} = input.([]interface{})
        nums  []int32       = args[0].([]int32)
        shift int           = args[1].(int)
    )
    return SumPermuted(nums, shift)
}

func SumDuplicatedFirstWrapper(input interface{}) interface{} {
    var nums []int32 = input.([]int32)
    return SumDuplicatedFirst(nums)
}

func FilterComplexNumbersWrapper(input interface{}) interface{} {
    var nums []int32 = input.([]int32)
    return FilterComplexNumbers(nums)
}

func FilterValidPasswordsWrapper(input interface{}) interface{} {
    var passwords []string = input.([]string)
    return FilterValidPasswords(passwords)
}

func GenerateRandomPasswordsWrapper(input interface{}) interface{} {
    var characters []string = input.([]string)
    return GenerateRandomPasswords(characters)
}