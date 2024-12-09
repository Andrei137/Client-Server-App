package client_server

import (
    "testing"
    "regexp"
)

func equalSlice[T comparable](a, b []T) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

// Cerinta 1
func TestMixWords(t *testing.T) {
    tests := []struct {
        words    []string
        expected []string
    }{
        {
            words:    []string{"casa", "masa", "trei", "tanc", "4321"},
            expected: []string{"cmtt4", "aara3", "ssen2", "aaic1"},
        },
        {
            words:    []string{"hello", "world"},
            expected: []string{"hw", "eo", "lr", "ll", "od"},
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := MixWords(test.words)
            if !equalSlice(actual, test.expected) {
                t.Errorf("For input %v, expected %v, but got %v", test.words, test.expected, actual)
            }
        })
    }
}

// Cerinta 2
func TestCountPerfectSquares(t *testing.T) {
    tests := []struct {
        words    []string
        expected int32
    }{
        {
            words:    []string{"abd4g5", "1sdf6fd", "fd2fdsf5"},
            expected: 2,
        },
        {
            words:    []string{"hello", "world"},
            expected: 0,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := CountPerfectSquares(test.words)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.words, test.expected, actual)
            }
        })
    }
}

// Cerinta 3
func TestSumReversed(t *testing.T) {
    tests := []struct {
        nums     []int32
        expected int32
    }{
        {
            nums:     []int32{12, 13, 14},
            expected: 93,
        },
        {
            nums:     []int32{123, 456, -123},
            expected: 654,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := SumReversed(test.nums)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, actual)
            }
        })
    }
}

// Cerinta 4
func TestAverageWithinRange(t *testing.T) {
    tests := []struct {
        nums     []int32
        expected int32
    }{
        {
            nums:     []int32{2, 10, 11, 39, 32, 80, 84},
            expected: 41,
        },
        {
            nums:     []int32{6, 6, -123},
            expected: -123,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := AverageWithinRange(test.nums)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, actual)
            }
        })
    }
}

// Cerinta 5
func TestConvertBinary(t *testing.T) {
    tests := []struct {
        words    []string
        expected []int32
    }{
        {
            words:    []string{"2dasdas", "12", "dasdas", "1010", "101"},
            expected: []int32{10, 5},
        },
        {
            words:    []string{"1001", "1000", "0111", "110", "101"},
            expected: []int32{9, 8, 7, 6, 5},
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := ConvertBinary(test.words)
            if !equalSlice(actual, test.expected) {
                t.Errorf("For input %v, expected %v, but got %v", test.words, test.expected, actual)
            }
        })
    }
}

// Cerinta 6
func TestCaesarCipher(t *testing.T) {
    tests := []struct {
        text      string
        direction string
        shift     int32
        expected  string
    }{
        {
            text:      "abcdef",
            direction: "RIGHT",
            shift:     1,
            expected:  "bcdefg",
        },
        {
            text:      "abcdef hello world",
            direction: "left",
            shift:     3,
            expected:  "xyzabc ebiil tloia",
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := CaesarCipher(test.text, test.direction, test.shift)
            if actual != test.expected {
                t.Errorf("For input %v, %v, %v, expected %v, but got %v", test.text, test.direction, test.shift, test.expected, actual)
            }
        })
    }
}

// Cerinta 7
func TestDecodeText(t *testing.T) {
    tests := []struct {
        encoded  string
        expected string
    }{
        {
            encoded:  "1G11o1L",
            expected: "GoooooooooooL",
        },
        {
            encoded:  "10a2b3c",
            expected: "aaaaaaaaaabbccc",
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := DecodeText(test.encoded)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.encoded, test.expected, actual)
            }
        })
    }
}

// Cerinta 8
func TestCountDigitsPrime(t *testing.T) {
    tests := []struct {
        nums     []int32
        expected int32
    }{
        {
            nums:     []int32{23, 17, 15, 3, 8},
            expected: 5,
        },
        {
            nums:     []int32{10, 11, 12, 13, 14, 15},
            expected: 4,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := CountDigitsPrime(test.nums)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, actual)
            }
        })
    }
}

// Cerinta 9
func TestCountEvenVowels(t *testing.T) {
    tests := []struct {
        words    []string
        expected int32
    }{
        {
            words:    []string{"mama", "iris", "bunica", "ala"},
            expected: 2,
        },
        {
            words:    []string{"hello", "world", "are"},
            expected: 1,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := CountEvenVowels(test.words)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.words, test.expected, actual)
            }
        })
    }
}

// Cerinta 10
func TestGCD(t *testing.T) {
    tests := []struct {
        words    []string
        expected int32
    }{
        {
            words:    []string{"24", "16", "8", "aaa", "bbb"},
            expected: 8,
        },
        {
            words:    []string{"hello12", "world14", "are"},
            expected: 2,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := GCD(test.words)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.words, test.expected, actual)
            }
        })
    }
}

// Cerinta 11
func TestSumPermutedPrimes(t *testing.T) {
    tests := []struct {
        nums     []int32
        shift    int
        expected int32
    }{
        {
            nums:     []int32{1234, 3456, 4567},
            shift:    2,
            expected: 15791,
        },
        {
            nums:     []int32{10, 11, 12, -13},
            shift:    1,
            expected: 2,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := SumPermuted(test.nums, test.shift)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, actual)
            }
        })
    }
}

// Cerinta 12
func TestSumDuplicatedFirst(t *testing.T) {
    tests := []struct {
        nums     []int32
        expected int32
    }{
        {
            nums:     []int32{23, 43, 26, 74},
            expected: 1666,
        },
        {
            nums:     []int32{1, 2, 3, 4, 5},
            expected: 165,
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := SumDuplicatedFirst(test.nums)
            if actual != test.expected {
                t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, actual)
            }
        })
    }
}

// Cerinta 13
func TextFilterComplexNumbers(t *testing.T) {
    tests := []struct {
        nums     []int32
        expected []float64
    }{
        {
            nums:     []int32{3, 10, 3, 4, 5, 6},
            expected: []float64{},
        },
        {
            nums:     []int32{3, 10, 5, 12},
            expected: []float64{13},
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := FilterComplexNumbers(test.nums)
            if !equalSlice(actual, test.expected) {
                t.Errorf("For input %v, expected %v, but got %v", test.nums, test.expected, actual)
            }
        })
    }
}

// Cerinta 14
func TestFilterValidPasswords(t *testing.T) {
    tests := []struct {
        passwords []string
        expected  []string
    }{
        {
            passwords: []string{"Ceva12!@", "asa212", "dasdas"},
            expected:  []string{"Ceva12!@"},
        },
        {
            passwords: []string{"1234", "abcd", "1234_Abcd"},
            expected:  []string{"1234_Abcd"},
        },
    }

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            actual := FilterValidPasswords(test.passwords)
            if !equalSlice(actual, test.expected) {
                t.Errorf("For input %v, expected %v, but got %v", test.passwords, test.expected, actual)
            }
        })
    }
}

// Cerinta 15
func TestGenerateRandomPasswords(t *testing.T) {
    tests := []struct {
        characters []string
    }{
        {
            characters: []string{"a", "b", "e", "3", "!", "A"},
        },
        {
            characters: []string{"a", "1", "!", "B", "3", "@"},
        },
    }

    hasLower := regexp.MustCompile(`[a-z]`)
    hasUpper := regexp.MustCompile(`[A-Z]`)
    hasDigit := regexp.MustCompile(`\d`)
    hasSymbol := regexp.MustCompile(`[\W_]`)

    for _, test := range tests {
        t.Run("", func(t *testing.T) {
            result := GenerateRandomPasswords(test.characters)
            for _, password := range result {
                if !(hasLower.MatchString(password) &&
                     hasUpper.MatchString(password) &&
                     hasDigit.MatchString(password) &&
                     hasSymbol.MatchString(password)) {
                    t.Errorf("For input %v, expected valid password, but got %v", test.characters, password)
                }
            }
        })
    }
}