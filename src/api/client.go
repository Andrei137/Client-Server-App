package main

import (
    "io"
    "os"
    "fmt"
    "sync"
    "strings"
    "strconv"
    "net/url"
    "net/http"
    "go.uber.org/zap"
    utils "client_server/src/utils"
)

func fetch(method string, params string, name string) (string, error) {
    var reqURL string = fmt.Sprintf(
        "http://localhost:%s/%s?params=%s&client=%s",
        utils.API_PORT, method, url.QueryEscape(params), name,
    )

    resp, err := http.Get(reqURL)
    if err != nil {
        return "", fmt.Errorf("Eroare la request-ul pentru metoda %s: %v", method, err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("Eroare la citirea raspunsului pentru metoda %s: %v", method, err)
    }

    var result string = string(body)
    if strings.Contains(result, "{\"result\":") {
        return result[10:len(result) - 2], nil
    }
    return "", fmt.Errorf("%s", strings.Trim(result, "\n"))
}

func runClient(name string, logger *zap.SugaredLogger, wg *sync.WaitGroup) {
    defer wg.Done()

    logger.Infof("Client %s conectat", name)

    queries := map[string]string{
        "mix_words"                 : "casa,masa,trei,tanc,4321",
        "count_perfect_squares"     : "abd4g5,1sdf6fd,fd2fdsf5",
        "sum_reversed"              : "12,13,14",
        "average_within_range"      : "2,10,11,39,32,80,84 | 1,2,3,4,5,6,7,8,9,10,11",
        "convert_binary"            : "2dasdas,12,dasdas,1010,101",
        "caesar_cipher"             : "abcdef,left,3 | hello,world,right,1",
        "decode_text"               : "1G11o1L",
        "count_digits_prime"        : "23,17,15,3,8",
        "count_even_vowels"         : "mama,iris,bunica,ala",
        "gcd"                       : "24,16,8,aaa,bbb",
        "sum_permuted"              : "1234,3456,4567,2",
        "sum_duplicated_first"      : "23,43,26,74 | hello,world",
        "filter_complex_numbers"    : "3,10,3,4,5,6,5,12",
        "filter_valid_passwords"    : "Ceva12!@,asa212,dasdas",
        "generate_random_passwords" : "a,b,e,3,!,A",
        "unknown_endpoint"          : "hello,world",
    }

    for method, params := range queries {
        var paramsList []string
        if strings.Contains(params, " | ") {
            paramsList = strings.Split(params, " | ")
        } else {
            paramsList = []string{params}
        }

        for _, paramsItem := range paramsList {
            logger.Infof("Client %s a facut request la %s cu parametrii %s", name, method, paramsItem)
            resp, err := fetch(method, paramsItem, name)
            if err != nil {
                logger.Errorf("Client %s a primit eroarea %v", name, err)
                continue
            }

            logger.Infof("Client %s a primit raspunsul %s", name, resp)
        }
    }
}

func main() {
    err := utils.LoadConfig()
    if err != nil {
        fmt.Println(err)
        return
    }

    var (
        numClients int
        wg         sync.WaitGroup
        logger     *zap.SugaredLogger
    )
    if len(os.Args) < 2 {
        numClients = 1
    } else {
        numClients, err = strconv.Atoi(os.Args[1])
        if err != nil || numClients <= 0 {
            fmt.Println("Numarul de clienti trebuie sa fie pozitiv!")
            return
        }
    }

    logger = utils.InitLogger(false).Sugar()
    defer logger.Sync()

    for i := 0; i < numClients; i++ {
        wg.Add(1)
        go runClient(utils.GetRandomName(), logger, &wg)
    }

    wg.Wait()
}