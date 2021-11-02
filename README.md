
# dicdh

Dynamic IP Cloudflare DNS Handler (dicdh) is a very simple program that updates the IP of the specified
Cloudflare DNS A-records for a given Zone.

## Common use case

The ideal use of this software is to install it and configure it to run at startup. Thus, everytime
the machine boots, the DNS will be updated.

This is useful for people who don't have a static IP and want some persistent identifier that doesn't
change over time. For example, a simple homelab. This way, the same DNS,
for example, `homelab1.example.com` always points to your server, instead of using dynamic IPs that change over time.
## Installation

The easiest way to install it is cloning the repository and running the bash scripts in
`build` directory.

```bash
  sudo bash ./build/install.sh
```

This will automatically copy the binary to `/usr/bin/` and create a daemon that will run
at startup.

### Configuration file

You will need to create a file to configure the Cloudflare API key, the DNS' to update, and more.

Create the following file and fill your data.

`/etc/dicdh/dicdh.json`
```json
{
    "key": "<api key from cloudflare>",
    "zone": "<zone id of the cloudflare domain>",
    "dns": [
        {
            "name": "<full dns>",
            "proxy": "<use cloudflare proxy or not>"
        }
    ]
}
```
Example configuration file:
```json
{
    "key": "aaaaaaaaaaa111111111",
    "zone": "aaaaaaaaaaa11111111",
    "dns": [
        {
            "name": "test.example.com",
            "proxy": true
        },
        {
            "name": "test2.example.com",
            "proxy": false
        }
    ]
}
```
## How it works

### Cloudflare API

It uses the Cloudflare API to update the DNS records. Thus, a valid API key with
permissions to edit DNS records is needed.

Authenticating with email is not supported as it is considered unsecure and a bad practice.

### Public Ip APIs

It uses a list of well-known public APIs to retrieve the client public ip.
