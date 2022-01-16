## Website Lookup CLI

Clone this github repo:

```bash 
git clone https://github.com/zunedz/go-cli.git
cd go-cli
```
Make sure Docker is already running, then run:

```bash
make run-cli
```
Inside the terminal, run `cli` to list out cli commands and its functionalities.

## Description
This CLI provides query for IPs, CNAMEs, MX records and Name Servers

## Usage
```bash
[global options] command [command options] [arguments...]
```
### With global options:
```bash
 --host value  (default: "http://youtube.com")
 --help, -h    show help
```
### Commands:
-   ns       will retrieve the name servers
-   cname    will lookup the CNAME for a given host
-   mx       will lookup the mail exchange records for a given host
-   ip       will lookup the ip address for a given host
-   help, h  Shows a list of commands or help for one command






