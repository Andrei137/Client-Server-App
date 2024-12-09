package main

import (
    "os"
    "fmt"
    "sync"
    "time"
    "context"
    "strings"
    "strconv"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    utils "client_server/src/utils"
    pb "client_server/src/grpc/services"
    "google.golang.org/grpc/credentials/insecure"
)

func snakeCaseMethod(method string) string {
    if method == "GCD" {
        return "gcd"
    }
    var result string = ""
    for i, char := range method {
        if i > 0 && 'A' <= char && char <= 'Z' {
            result += "_"
        }
        result += strings.ToLower(string(char))
    }
    return result
}

func parseParams(paramsStr string, method string) (interface{}, error) {
    switch method {
        case "DecodeText":
            return &pb.StringReq{Value: paramsStr}, nil
        case "MixWords",
             "CountPerfectSquares",
             "ConvertBinary",
             "CountEvenVowels",
             "GCD",
             "FilterValidPasswords",
             "GenerateRandomPasswords":
            return &pb.StringSliceReq{Strings: strings.Split(paramsStr, ",")}, nil
        case "SumReversed",
             "AverageWithinRange",
             "CountDigitsPrime",
             "SumDuplicatedFirst",
             "FilterComplexNumbers":
            params := strings.Split(paramsStr, ",")
            numbers := make([]int32, 0, len(params))
            for _, param := range params {
                number, err := strconv.Atoi(param)
                if err != nil {
                    return nil, fmt.Errorf("%s nu este un numar valid", param)
                }
                numbers = append(numbers, int32(number))
            }
            return &pb.Int32SliceReq{Numbers: numbers}, nil
        case "SumPermuted":
            params := strings.Split(paramsStr, ",")
            numbers := make([]int32, 0, len(params) - 1)
            for i, param := range params {
                if i == len(params) - 1 {
                    number, err := strconv.Atoi(param)
                    if err != nil {
                        return nil, fmt.Errorf("%s nu este un numar valid", param)
                    }
                    return &pb.SumPermutedReq{Numbers: numbers, Shift: int32(number)}, nil
                }
                number, err := strconv.Atoi(param)
                if err != nil {
                    return nil, fmt.Errorf("%s nu este un numar valid", param)
                }
                numbers = append(numbers, int32(number))
            }
            return nil, fmt.Errorf("Numarul de parametrii (%d) nu corespunde cu numarul necesar (2)", len(params))
        case "CaesarCipher":
            params := strings.Split(paramsStr, ",")
            if len(params) != 3 {
                return nil, fmt.Errorf("Numarul de parametrii (%d) nu corespunde cu numarul necesar (3)", len(params))
            }
            shift, err := strconv.Atoi(params[2])
            if err != nil {
                return nil, fmt.Errorf("%s nu este un numar valid", params[2])
            }
            return &pb.CaesarCipherReq{Text: params[0], Direction: params[1], Shift: int32(shift)}, nil
        default:
            return nil, fmt.Errorf("Metoda nu exista")
    }
}

func callGRPC(client pb.ServicesClient, method string, paramsStr string) (string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    params, err := parseParams(paramsStr, method)
    if err != nil {
        return "", err
    }
    switch method {
        case "MixWords":
            response, err := client.MixWords(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetStrings()), nil
        case "CountPerfectSquares":
            response, err := client.CountPerfectSquares(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "SumReversed":
            response, err := client.SumReversed(ctx, params.(*pb.Int32SliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "AverageWithinRange":
            response, err := client.AverageWithinRange(ctx, params.(*pb.Int32SliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "ConvertBinary":
            response, err := client.ConvertBinary(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetNumbers()), nil
        case "CaesarCipher":
            response, err := client.CaesarCipher(ctx, params.(*pb.CaesarCipherReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "DecodeText":
            response, err := client.DecodeText(ctx, params.(*pb.StringReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "CountDigitsPrime":
            response, err := client.CountDigitsPrime(ctx, params.(*pb.Int32SliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "CountEvenVowels":
            response, err := client.CountEvenVowels(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "GCD":
            response, err := client.GCD(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "SumPermuted":
            response, err := client.SumPermuted(ctx, params.(*pb.SumPermutedReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "SumDuplicatedFirst":
            response, err := client.SumDuplicatedFirst(ctx, params.(*pb.Int32SliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetValue()), nil
        case "FilterComplexNumbers":
            response, err := client.FilterComplexNumbers(ctx, params.(*pb.Int32SliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetNumbers()), nil
        case "FilterValidPasswords":
            response, err := client.FilterValidPasswords(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetStrings()), nil
        case "GenerateRandomPasswords":
            response, err := client.GenerateRandomPasswords(ctx, params.(*pb.StringSliceReq))
            if err != nil {
                return "", err
            }
            return fmt.Sprint(response.GetStrings()), nil
        default:
            return "", fmt.Errorf("Metoda nu exista")
    }
}

func runClient(name string, logger *zap.SugaredLogger, wg *sync.WaitGroup) {
    defer wg.Done()

    conn, err := grpc.NewClient(
        "localhost:" + utils.GRPC_PORT,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithUserAgent(name),
    )
    if err != nil {
        logger.Errorf("Client %s nu s-a putut conecta la server", name)
    }
    logger.Infof("Client %s conectat", name)
    defer conn.Close()

    queries := map[string]string{
        "MixWords"                : "casa,masa,trei,tanc,4321",
        "CountPerfectSquares"     : "abd4g5,1sdf6fd,fd2fdsf5",
        "SumReversed"             : "12,13,14",
        "AverageWithinRange"      : "2,10,11,39,32,80,84 | 1,2,3,4,5,6,7,8,9,10,11",
        "ConvertBinary"           : "2dasdas,12,dasdas,1010,101",
        "CaesarCipher"            : "abcdef,left,3 | hello,world,right,1",
        "DecodeText"              : "1G11o1L",
        "CountDigitsPrime"        : "23,17,15,3,8",
        "CountEvenVowels"         : "mama,iris,bunica,ala",
        "GCD"                     : "24,16,8,aaa,bbb",
        "SumPermuted"             : "1234,3456,4567,2",
        "SumDuplicatedFirst"      : "23,43,26,74 | hello,world",
        "FilterComplexNumbers"    : "3,10,3,4,5,6,5,12",
        "FilterValidPasswords"    : "Ceva12!@,asa212,dasdas",
        "GenerateRandomPasswords" : "a,b,e,3,!,A",
        "UnknownEndpoint"         : "hello,world",
    }

    for method, params := range queries {
        var paramsList []string
        if strings.Contains(params, " | ") {
            paramsList = strings.Split(params, " | ")
        } else {
            paramsList = []string{params}
        }

        for _, paramsItem := range paramsList {
            logger.Infof("Client %s a facut request la %s cu parametrii %s", name, snakeCaseMethod(method), paramsItem)
            response, err := callGRPC(pb.NewServicesClient(conn), method, paramsItem)
            if err != nil {
                if strings.Contains(err.Error(), "rpc error") {
                    logger.Errorf("Client %s a primit eroarea %s", name, strings.Split(err.Error(), "desc = ")[1])
                } else {
                    logger.Errorf("Client %s a primit eroarea %v", name, err)
                }
                continue
            }
            logger.Infof("Client %s a primit raspunsul %s", name, response)
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