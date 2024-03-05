package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/google/uuid"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type bip39Result struct {
	Name      string `json:"name"`
	Mnemonic  string `json:"mnemonic"`
	Seed      string `json:"seed"`
	MasterKey string `json:"masterKey"`
	PublicKey string `json:"publicKey"`
}

func generateMnemonic() string {
	// Generate a mnemonic for memorization or user-friendly seeds
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create entropy: %s", err)
		panic(errMsg)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to generate mnemonic: %s", err)
		panic(errMsg)
	}

	return mnemonic
}

func parseSeed(mnemonic string) (seed []byte, masterKey *bip32.Key, publicKey *bip32.Key) {
	// Generate seed from mnemonic
	seed = bip39.NewSeed(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to generate master key: %s", err)
		panic(errMsg)
	}

	publicKey = masterKey.PublicKey()
	return
}

func generate(name string, mnemonic string) bip39Result {
	if name == "" {
		name = uuid.NewString()
	}

	if mnemonic == "" {
		mnemonic = generateMnemonic()
	}

	seed, masterKey, publicKey := parseSeed(mnemonic)
	return bip39Result{
		Name:      name,
		Mnemonic:  mnemonic,
		Seed:      fmt.Sprintf("%x", seed),
		MasterKey: fmt.Sprintf("%x", masterKey.Key),
		PublicKey: fmt.Sprintf("%x", publicKey.Key),
	}
}

func generateBatch(count int) []bip39Result {
	results := make([]bip39Result, count)
	for i := 0; i < count; i++ {
		results[i] = generate("", "")
	}
	return results
}

func printToConsole(result bip39Result) {
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Mnemonic: %s\n", result.Mnemonic)
	fmt.Printf("Seed: %s\n", result.Seed)
	fmt.Printf("Seed (hex): %x\n", result.Seed)
	fmt.Printf("Master private key: %s\n", result.MasterKey)
	fmt.Printf("Master private key (hex): %x\n", result.MasterKey)
	fmt.Printf("Master public key: %s\n", result.PublicKey)
	fmt.Printf("Master public key (hex): %x\n", result.PublicKey)
}

func saveToFile(result bip39Result, outputFolder string) {
	// Parse as JSON string
	resultJson, err := json.Marshal(result)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to generate JSON: %s", err)
		panic(errMsg)
	}

	// Check output folder exists
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		if err := os.Mkdir(outputFolder, 0755); err != nil {
			errMsg := fmt.Sprintf("Failed to create output folder: %s", err)
			panic(errMsg)
		}
	}

	// Create file
	fileName := outputFolder + "/" + result.Name + ".json"
	file, err := os.Create(fileName)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to create file: %s", err)
		panic(errMsg)
	}
	defer file.Close()

	// Write JSON to file
	if _, err := file.Write(resultJson); err != nil {
		errMsg := fmt.Sprintf("Failed to write to file: %s", err)
		panic(errMsg)
	}
}

func main() {
	parser := argparse.NewParser("bip39", "Command line utility written in Go for generating and parsing bip-39 seed phrases")

	// Create flags
	help := parser.Flag("h", "help", &argparse.Options{Required: false, Help: "Show help information for this command"})
	name := parser.String("n", "name", &argparse.Options{Required: false, Help: "Override default UUID name with an explicit name"})
	disableConsole := parser.Flag("c", "disable-console-output", &argparse.Options{Required: false, Help: "Disable console output"})
	disableFile := parser.Flag("f", "disable-file-output", &argparse.Options{Required: false, Help: "Disable file output"})
	outputFolder := parser.String("o", "output-folder", &argparse.Options{Required: false, Help: "Output folder for file output", Default: "."})
	mneumonic := parser.String("m", "mnemonic", &argparse.Options{Required: false, Help: "Provide a mnemonic instead of letting command generate one"})
	batchMode := parser.Int("b", "batch-mode", &argparse.Options{Required: false, Help: "Generate a batch of mnemonics", Default: 0})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// In case of error print error and print usage
		// This can also be done by passing -h or --help flags
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if *help {
		fmt.Print(parser.Usage(nil))
		os.Exit(0)
	}

	// Generate a mnemonic for memorization or user-friendly seeds
	var results []bip39Result
	if *batchMode > 0 {
		results = generateBatch(*batchMode)
	} else {
		result := generate(*name, *mneumonic)
		results = []bip39Result{result}
	}

	for _, result := range results {
		// Print to console unless console output disabled
		if !*disableConsole {
			printToConsole(result)
		}

		// Save to file unless file output disabled
		if !*disableFile {
			saveToFile(result, *outputFolder)
		}
	}
}
