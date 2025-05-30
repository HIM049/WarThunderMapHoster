# WarThunderMapHoster
A simple WarThunder map host, used to limit who can download your map.

## Feature:
- Visit control with password.
- Multi map host.
- Map manager with upload and delete.

## Roadmap ~~(may added)~~
- [x] multiple map
- [ ] Anything more? Submit your issue.

## Guide to use
1. Download the server executable program or build the source code by self.
2. Edit the config.toml with you own idea.
3. Fill the secret key in config.toml (required). 
4. Edit tmpl files in `./tamplates` if you want to customize your website.
5. Start server

## Notice
### Secret Key
You can use any string you want as the key without limit. But we recommend you use the key fully random (generate with tools), and longer than 256bit (32byte).

The command below can help you generate a random key with 32byte long.

```shell
openssl rand -base64 64
```

### Public IP address required

You should run this server with a public server but not your local pc. The game cannot visit the map with private IP address.
