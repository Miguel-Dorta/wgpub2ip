package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"golang.org/x/crypto/blake2s"
	"io"
	"net"
	"os"
)

var (
	Version string
	pubKey  string
	network string
)

func main() {
	readArgs()
	_, ipnet, err := net.ParseCIDR(network)
	printFatalErr(err)

	binPubKey, err := base64.StdEncoding.DecodeString(pubKey)
	printFatalErr(err)
	hashsum := blake2s.Sum256(binPubKey)

	ip := net.IP(make([]byte, len(ipnet.IP)))
	for i := range ipnet.IP {
		ip[i] = (ipnet.IP[i] & ipnet.Mask[i]) | (hashsum[i+blake2s.Size-len(ipnet.IP)] & ^ipnet.Mask[i])
	}
	fmt.Println(ip)
}

func readArgs() {
	var help bool
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "")
	flag.StringVar(&pubKey, "k", "", "")
	flag.StringVar(&pubKey, "pubkey", "", "")
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
		printFatalErr(err)
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
  %s -k 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=' 2001:db8::/64
  cat pubkey.txt | %s 192.0.2.0/24

AUTHOR:
  Miguel Dorta <contact@migueldorta.com>

COPYRIGHT:
  Copyright © 2024 Miguel Dorta
  Licensed under the EUPL v1.2 <https://eupl.eu/1.2/en/>

VERSION:
  wgpub2ip %s

`, os.Args[0], os.Args[0], os.Args[0], Version)
}

func printFatalErr(err error) {
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
