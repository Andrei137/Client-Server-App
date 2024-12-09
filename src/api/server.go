package main

import (
    "os"
    "fmt"
    "strconv"
    "strings"
    "net/http"
    "encoding/json"
    "go.uber.org/zap"
    utils "client_server/src/utils"
)

type ProcessFunc func(interface{}) interface{}

func convertToIntSlice(paramList []string, numsLength int) ([]int32, error) {
    nums := make([]int32, numsLength)
    for i, param := range paramList[:numsLength] {
        num, err := strconv.Atoi(param)
        if err != nil {
            return nil, fmt.Errorf("%s nu este un numar valid", param)
        }
        nums[i] = int32(num)
    }
    return nums, nil
}

func parseParams(method string, params string) (interface{}, error) {
    var paramList []string = strings.Split(params, ",")
    if len(paramList) > utils.MAX_PARAMS {
        return nil, fmt.Errorf("Numarul de parametrii (%d) depaseste limita permisa (%d)", len(paramList), utils.MAX_PARAMS)
    }

    switch method {
        // string
        case "decode_text":
            return params, nil
        // []string
        case "mix_words",
             "count_perfect_squares",
             "convert_binary",
             "count_even_vowels",
             "gcd",
             "filter_valid_passwords",
             "generate_random_passwords":
            return strings.Split(params, ","), nil
        // []int32
        case "sum_reversed",
             "average_within_range",
             "count_digits_prime",
             "sum_duplicated_first",
             "filter_complex_numbers":
            nums, err := convertToIntSlice(paramList, len(paramList))
            if err != nil {
                return nil, err
            }
            return nums, nil
        // ([]int32, int)
        case "sum_permuted":
            var numsLength int = len(paramList) - 1
            nums, err := convertToIntSlice(paramList, numsLength)
            if err != nil {
                return nil, err
            }
            shift, err := strconv.Atoi(paramList[numsLength])
            if err != nil {
                return nil, fmt.Errorf("%s nu este un numar valid", paramList[numsLength])
            }
            return []interface{}{nums, shift}, nil
        // (string, string, int32)
        case "caesar_cipher":
            if len(paramList) != 3 {
                return nil, fmt.Errorf("Numarul de parametrii (%d) nu corespunde cu numarul necesar (3)", len(paramList))
            }
            shift, err := strconv.Atoi(paramList[2])
            if err != nil {
                return nil, fmt.Errorf("%s nu este un numar valid", paramList[2])
            }
            return []interface{}{paramList[0], paramList[1], int32(shift)}, nil
        default:
            return nil, fmt.Errorf("Metoda %s nu are un format valid", method)
    }
}

func handler(w http.ResponseWriter, r *http.Request, serverLogger *zap.SugaredLogger, clientLogger *zap.SugaredLogger) {
    var functionMap = map[string]ProcessFunc{
        "mix_words"                 : utils.MixWordsWrapper,
        "count_perfect_squares"     : utils.CountPerfectSquaresWrapper,
        "sum_reversed"              : utils.SumReversedWrapper,
        "average_within_range"      : utils.AverageWithinRangeWrapper,
        "convert_binary"            : utils.ConvertBinaryWrapper,
        "caesar_cipher"             : utils.CaesarCipherWrapper,
        "decode_text"               : utils.DecodeTextWrapper,
        "count_digits_prime"        : utils.CountDigitsPrimeWrapper,
        "count_even_vowels"         : utils.CountEvenVowelsWrapper,
        "gcd"                       : utils.GCDWrapper,
        "sum_permuted"              : utils.SumPermutedWrapper,
        "sum_duplicated_first"      : utils.SumDuplicatedFirstWrapper,
        "filter_complex_numbers"    : utils.FilterComplexNumbersWrapper,
        "filter_valid_passwords"    : utils.FilterValidPasswordsWrapper,
        "generate_random_passwords" : utils.GenerateRandomPasswordsWrapper,
    }

    var (
        method string = strings.TrimPrefix(r.URL.Path, "/")
        params string = r.URL.Query().Get("params")
        client string = r.URL.Query().Get("client")
    )

    if clientLogger != nil {
        clientLogger.Infof("Client %s a facut request la %s cu parametrii %s", client, method, params)
    }
    serverLogger.Infof("Server a primit request-ul %s cu parametrii %s", method, params)

    function, exists := functionMap[method]
    if !exists {
        http.Error(w, "Metoda nu exista", http.StatusBadRequest)
        serverLogger.Errorf("Metoda %s nu exista", method)
        return
    }

    resultChan := make(chan interface{})
    errChan    := make(chan error)

    go func() {
        serverLogger.Info("Server proceseaza datele")
        input, err := parseParams(method, params)
        if err != nil {
            errChan <- err
            return
        }
        result := function(input)
        resultChan <- result
    }()

    select {
        case result := <-resultChan:
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(map[string]interface{}{
                "result": result,
            })
            serverLogger.Infof("Server trimite %v catre client", result)
            if clientLogger != nil {
                clientLogger.Infof("Client %s a primit raspunsul %v", client, result)
            }
        case err := <-errChan:
            http.Error(w, err.Error(), http.StatusBadRequest)
            serverLogger.Error(err)
    }
}

func main() {
    err := utils.LoadConfig()
    if err != nil {
        fmt.Println(err)
        return
    }

    var (
        logClient    bool
        serverLogger *zap.SugaredLogger
        clientLogger *zap.SugaredLogger
    )

    if len(os.Args) < 2 {
        logClient = false
    } else {
        logClient, err = strconv.ParseBool(os.Args[1])
        if err != nil {
            fmt.Println("Parametrul trebuie sa fie de tipul bool!")
            return
        }
    }

    serverLogger = utils.InitLogger(true).Sugar()
    defer serverLogger.Sync()
    if (logClient) {
        clientLogger = utils.InitLogger(false).Sugar()
        defer clientLogger.Sync()
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        handler(w, r, serverLogger, clientLogger)
    })
    serverLogger.Info("Server pornit pe portul ", utils.API_PORT)
    http.ListenAndServe(":" + utils.API_PORT, nil)
}
