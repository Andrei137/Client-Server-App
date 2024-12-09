package client_server

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
    "math/rand"
    "encoding/json"
)

var (
    API_PORT      string
    GRPC_PORT     string
    MAX_PARAMS    int
    COLORED_LOGS  bool
    names         []string
    usedNames     map[string]bool
)

type Config struct {
    API_PORT      int  `json:"API_PORT"`
    GRPC_PORT     int  `json:"GRPC_PORT"`
    MAX_PARAMS    int  `json:"MAX_PARAMS"`
    COLORED_LOGS  bool `json:"COLORED_LOGS"`
}

func loadNames(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        name := strings.TrimSpace(scanner.Text())
        if name != "" {
            names = append(names, name)
        }
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

func GetRandomName() string {
    if len(usedNames) == len(names) {
        usedNames = make(map[string]bool)
    }

    for {
        var name string = names[rand.Intn(len(names))]

        if !usedNames[name] {
            usedNames[name] = true
            return name
        }
    }
}

func LoadConfig() error {
    if API_PORT != "" && GRPC_PORT != "" {
        return nil
    }

    var (
        configPath   string = "../config/"
        settingsFile string = configPath + "settings.json"
        namesFile    string = configPath + "names.txt"
        config       Config
    )

    file, err := os.Open(settingsFile)
    if err != nil {
        return fmt.Errorf("Eroare la citirea fisierului de configurare: %v", err)
    }
    defer file.Close()

    err = json.NewDecoder(file).Decode(&config)
    if err != nil {
        return fmt.Errorf("Eroare la parsarea fisierului de configurare: %v", err)
    }

    API_PORT     = strconv.Itoa(config.API_PORT)
    GRPC_PORT    = strconv.Itoa(config.GRPC_PORT)
    MAX_PARAMS   = config.MAX_PARAMS
    COLORED_LOGS = config.COLORED_LOGS

    err = loadNames(namesFile)
    if err != nil {
        return fmt.Errorf("Eroare la citirea fisierului de nume: %v", err)
    }

    usedNames = make(map[string]bool)
    return nil
}
