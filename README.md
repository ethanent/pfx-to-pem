# pfx-to-pem
> Convert PKCS12 to PEM keys

## Install

```sh
go get github.com/ethanent/pfxpem
```


## Usage

```
Usage of pfxpem:
  -password string
        Password for PFX file
  -source string
        Source PFX file
```

The CLI will output each encoded PEM Block to stdout.

Each PEM Block will be prefixed by "PEM Block [*number*]:"
