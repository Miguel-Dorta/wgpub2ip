# wgpub2ip

`wgpub2ip` is a command-line utility that allows you to calculate a
deterministic IP address given a WireGuard public key and network range.

## Usage

```
wgpub2ip [options] [network]
```

## Description

`wgpub2ip` calculates an IP address corresponding to a WireGuard public key within a specified network range.

## Options

- `-h, --help`: Print help information and exit.
- `-k, --pubkey`: Specify the public key to generate the IP address from.

## Examples

```sh
wgpub2ip -k 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=' 2001:db8::/64
```
This command calculates the IP address corresponding to the given public key within the specified IPv6 network range.

```sh
cat pubkey.txt | wgpub2ip 192.0.2.0/24
```
This command reads the public key from a file named `pubkey.txt` and calculates the IP address within the specified IPv4 network range.

## Author

Miguel Dorta
- Email: [contact@migueldorta.com](mailto:contact@migueldorta.com)
- Website: [migueldorta.com](https://migueldorta.com)

## Copyright

Copyright © 2024 Miguel Dorta  
Licensed under the [EUPL v1.2](https://eupl.eu/1.2/en/)