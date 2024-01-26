# bip39
Command line utility written in Go for generating and parsing bip-39 seed phrases. This utility can be used for the following:
1. Parse existing 12 - 24 word mneumonic to produce the wallet private/public key
1. Generate new mneumonic along with corresponding wallet private/public key
1. Generate batches of mneumonics along with corresponding wallet private/public keys

# Installation
This utility can be installed with Go or by downloading the zip file that corresponds to your distribution (e.g., Mac M1: darwin-arm64.zip, Windows 64-bit: windows-amd64.zip).

## Go Installation
To install this utility with Go, simply run the following:
```bash
go install github.com/jixer/bip39
```

Assuming you have the `$GOPATH` referenced in your core `$PATH`, you should then be able to execute from command line using `bip39` (see [Usage](#usage) instruction).

## Release Installation
Current release version is 1.0.0. Download the zip file that corresponds to your machine and extract to a place where you can run from a command line.

Steps:
1. Navigate to the [v1.0.0 Release](https://github.com/jixer/bip39/releases/tag/1.0.0) area in this github repo
1. In Assets area, download the zip file that corresponds to your machine
1. Extract to a place where you can run from a command line
1. Profit

# Usage
This utility includes help for the command via the `-h` or `--help` flag, so this section will not duplicativley cover that.

Runing the command without any flags will generate a new mneumonic along with it's corresponding priv/pub key. This will be printed to console as well as saved into a file named `<random-UUID>.json` (<random-UUID> will be replaced at runtime with a randomly generated UUID).

Simple example:
```bash
❯ bip39
Name: 07dc6453-6ddf-4f96-897a-c8837b9de04d
Mnemonic: payment oyster onion achieve memory surprise since crater ecology stereo acquire fit tooth kind invest discover person useful baby tiny pulp limb wisdom infant
Seed: 6134626634353666323465333738666633623062363839383131623130303564653734323438653066306639623831646665336132386464393762393463333866323439343932326331373663353032356132323438626362653632636538643234343639353666356533643533346139363636313431346132626466313832
Master private key: 38613732653238636238616536386662636261646662393936633235663430396439613332633763633563623865333862646334656464373834326263386230
Master public key: 303330393138356630643137643135656439636139313233343863623161303731346335653837656561396666643363366332373063643537613336303664616137

❯ cat 07dc6453-6ddf-4f96-897a-c8837b9de04d.json
{
  "name":"07dc6453-6ddf-4f96-897a-c8837b9de04d",
  "mnemonic":"payment oyster onion achieve memory surprise since crater ecology stereo acquire fit tooth kind invest discover person useful baby tiny pulp limb wisdom infant",
  "seed":"a4bf456f24e378ff3b0b689811b1005de74248e0f0f9b81dfe3a28dd97b94c38f2494922c176c5025a2248bcbe62ce8d2446956f5e3d534a96661414a2bdf182",
  "masterKey":"8a72e28cb8ae68fbcbadfb996c25f409d9a32c7cc5cb8e38bdc4edd7842bc8b0",
  "publicKey":"0309185f0d17d15ed9ca912348cb1a0714c5e87eea9ffd3c6c270cd57a3606daa7"
}
```
> NOTE: The output above was cleared and discarded after generating. These values do not reflect a real wallet in use. Please do not attempt to use these for any reason. If you transfer any funds to this public key or attempt to use the wallet, you run the risk of losing those funds.

For all other usages (batch, from existing mneumonic), please refer to the built-in help documentation via `bip39 --help`:
```bash
❯ bip39 --help
usage: bip39 [-h|--help] [-n|--name "<value>"] [-c|--disable-console-output]
             [-f|--disable-file-output] [-o|--output-folder "<value>"]
             [-m|--mnemonic "<value>"] [-b|--batch-mode <integer>]

             Command line utility written in Go for generating and parsing
             bip-39 seed phrases

Arguments:

  -h  --help                    Print help information
  -n  --name                    Override default UUID name with an explicit
                                name
  -c  --disable-console-output  Disable console output
  -f  --disable-file-output     Disable file output
  -o  --output-folder           Output folder for file output. Default: .
  -m  --mnemonic                Provide a mnemonic instead of letting command
                                generate one
  -b  --batch-mode              Generate a batch of mnemonics. Default: 0
```

# Credits
Special thanks to github user [@tyler-smith](https://github.com/tyler-smith) for doing most of the legwork here with the underlying bip39 and bip32 libraries:
- [tyler-smith/go-bip39](https://github.com/tyler-smith/go-bip39)
- [tyler-smith/go-bip32](https://github.com/tyler-smith/go-bip32)