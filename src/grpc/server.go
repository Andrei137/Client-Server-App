package main

import (
    "os"
    "fmt"
    "net"
    "context"
    "strings"
    "strconv"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    utils "client_server/src/utils"
    "google.golang.org/grpc/metadata"
    "google.golang.org/protobuf/proto"
    pb "client_server/src/grpc/services"
    "google.golang.org/protobuf/encoding/protojson"
)

type server struct {
    pb.UnimplementedServicesServer
    logClient    bool
    serverLogger *zap.SugaredLogger
    clientLogger *zap.SugaredLogger
}

func prettifyProtoMessage(m proto.Message) string {
    marshalled, err := protojson.Marshal(m)
    if err != nil {
        return fmt.Sprintf("%v", m)
    }
    return strings.Trim(strings.Split(string(marshalled), ":")[1], "}")
}

func (s *server) MixWords(ctx context.Context, in *pb.StringSliceReq) (*pb.StringSliceRes, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.MixWords(input.GetStrings())
        return &pb.StringSliceRes{Strings: result}
    }

    response, err := handleRequest(ctx, "mix_words", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.StringSliceRes), nil
}

func (s *server) CountPerfectSquares(ctx context.Context, in *pb.StringSliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.CountPerfectSquares(input.GetStrings())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "count_perfect_squares", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) SumReversed(ctx context.Context, in *pb.Int32SliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.Int32SliceReq) interface{} {
        result := utils.SumReversed(input.GetNumbers())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "sum_reversed", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) AverageWithinRange(ctx context.Context, in *pb.Int32SliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.Int32SliceReq) interface{} {
        result := utils.AverageWithinRange(input.GetNumbers())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "average_within_range", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) ConvertBinary(ctx context.Context, in *pb.StringSliceReq) (*pb.Int32SliceRes, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.ConvertBinary(input.GetStrings())
        return &pb.Int32SliceRes{Numbers: result}
    }

    response, err := handleRequest(ctx, "convert_binary", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32SliceRes), nil
}

func (s *server) CaesarCipher(ctx context.Context, in *pb.CaesarCipherReq) (*pb.StringRes, error) {
    handler := func(input *pb.CaesarCipherReq) interface{} {
        result := utils.CaesarCipher(input.GetText(), input.GetDirection(), input.GetShift())
        return &pb.StringRes{Value: result}
    }

    response, err := handleRequest(ctx, "caesar_cipher", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.StringRes), nil
}

func (s *server) DecodeText(ctx context.Context, in *pb.StringReq) (*pb.StringRes, error) {
    handler := func(input *pb.StringReq) interface{} {
        result := utils.DecodeText(input.GetValue())
        return &pb.StringRes{Value: result}
    }

    response, err := handleRequest(ctx, "decode_text", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.StringRes), nil
}

func (s *server) CountDigitsPrime(ctx context.Context, in *pb.Int32SliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.Int32SliceReq) interface{} {
        result := utils.CountDigitsPrime(input.GetNumbers())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "count_digits_prime", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) CountEvenVowels(ctx context.Context, in *pb.StringSliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.CountEvenVowels(input.GetStrings())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "count_even_vowels", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) GCD(ctx context.Context, in *pb.StringSliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.GCD(input.GetStrings())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "gcd", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) SumPermuted(ctx context.Context, in *pb.SumPermutedReq) (*pb.Int32Res, error) {
    handler := func(input *pb.SumPermutedReq) interface{} {
        result := utils.SumPermuted(input.GetNumbers(), int(input.GetShift()))
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "sum_permuted", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) SumDuplicatedFirst(ctx context.Context, in *pb.Int32SliceReq) (*pb.Int32Res, error) {
    handler := func(input *pb.Int32SliceReq) interface{} {
        result := utils.SumDuplicatedFirst(input.GetNumbers())
        return &pb.Int32Res{Value: result}
    }

    response, err := handleRequest(ctx, "sum_duplicated_first", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.Int32Res), nil
}

func (s *server) FilterComplexNumbers(ctx context.Context, in *pb.Int32SliceReq) (*pb.DoubleSliceRes, error) {
    handler := func(input *pb.Int32SliceReq) interface{} {
        result := utils.FilterComplexNumbers(input.GetNumbers())
        return &pb.DoubleSliceRes{Numbers: result}
    }

    response, err := handleRequest(ctx, "filter_complex_numbers", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.DoubleSliceRes), nil
}

func (s *server) FilterValidPasswords(ctx context.Context, in *pb.StringSliceReq) (*pb.StringSliceRes, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.FilterValidPasswords(input.GetStrings())
        return &pb.StringSliceRes{Strings: result}
    }

    response, err := handleRequest(ctx, "filter_valid_passwords", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.StringSliceRes), nil
}

func (s *server) GenerateRandomPasswords(ctx context.Context, in *pb.StringSliceReq) (*pb.StringSliceRes, error) {
    handler := func(input *pb.StringSliceReq) interface{} {
        result := utils.GenerateRandomPasswords(input.GetStrings())
        return &pb.StringSliceRes{Strings: result}
    }

    response, err := handleRequest(ctx, "generate_random_passwords", in, handler, s.logClient, s.serverLogger, s.clientLogger)
    if err != nil {
        s.serverLogger.Error(err)
        return nil, err
    }
    return response.(*pb.StringSliceRes), nil
}

func handleRequest[T proto.Message](
    ctx context.Context,
    method string,
    input T,
    handler func(T) interface{},
    logClient bool,
    serverLogger *zap.SugaredLogger,
    clientLogger *zap.SugaredLogger,
) (interface{}, error) {
    params := prettifyProtoMessage(input)
    if strings.Count(params, ",") >= utils.MAX_PARAMS {
        return nil, fmt.Errorf("Limita de %d parametrii a fost depasita", utils.MAX_PARAMS)
    }

    if logClient {
        headers, _ := metadata.FromIncomingContext(ctx)
        name := strings.Split(headers.Get("user-agent")[0], " ")[0]
        clientLogger.Infof("Client %s a facut request la %s cu parametrii %s", name, method, params)
    }
    serverLogger.Infof("Server a primit request-ul %s cu parametrii %s", method, params)
    serverLogger.Info("Server proceseaza datele")
    response := handler(input)
    serverLogger.Infof("Server trimite %v catre client", prettifyProtoMessage(response.(proto.Message)))
    return response, nil
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

    lis, err := net.Listen("tcp", ":" + utils.GRPC_PORT)
    if err != nil {
        serverLogger.Error(err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterServicesServer(grpcServer, &server{
        logClient   : logClient,
        serverLogger: serverLogger,
        clientLogger: clientLogger,
    })
    serverLogger.Info("Server pornit pe portul ", utils.GRPC_PORT)
    if err := grpcServer.Serve(lis); err != nil {
        serverLogger.Error(err)
    }
}