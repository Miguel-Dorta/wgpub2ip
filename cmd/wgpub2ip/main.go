package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	Version string
	pubKey  string
	network string
)

func main() {

}

func parseArgs() {
	var help bool
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "-help", false, "")
	flag.StringVar(&pubKey, "k", "", "")
	flag.StringVar(&pubKey, "-pubkey", "", "")
	flag.Parse()
	if help {
		printHelp()
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) != 1 {
		printHelp()
		os.Exit(1)
	}
	network = args[0]

	if pubKey == "" {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err) // TODO make this more descriptive
		}
		pubKey = string(data)
	}
}

func printHelp() {
	_, _ = fmt.Fprintf(os.Stderr,
		`USAGE:
  %s [options] [network]

DESCRIPTION:
  Calculate an IP given a WireGuard public key and network range.

OPTIONS:
  -h, --help      Print this help and exit
  -k, --pubkey    Key to generate the IP from

EXAMPLES:
  %s -k 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=' 192.0.2.0/24
  cat pubkey.txt | %s 2001:db8::/32

AUTHOR:
  Miguel Dorta <contact@migueldorta.com>

COPYRIGHT:
  Copyright © 2024 Miguel Dorta
  Licensed under the EUPL v1.2 <https://eupl.eu/1.2/en/>

VERSION:
  wgpub2ip %s

`, os.Args[0], os.Args[0], os.Args[0], Version)
}
