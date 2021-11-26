dsieve
------
Take a single domain or read input file and extract unique parent domains. 
Filter results by domain level. 

 Supports any format of url, with or without protocol, port, path, parameters.


## Quickstart

```
> go run dsieve.go -i foo.bar.baz.tld 
bar.baz.tld
baz.tld
tld

> go run dsieve.go -i http://foo.bar.baz.tld/a/b?q=1 -f 2
baz.tld

> go run dsieve.go -i http://foo.bar.baz.tld:443 -f 2:
bar.baz.tld:443
baz.tld:443
```

## Usage
```
Usage of dsieve.go
  -f string
        Filter domain level. Use python slice notation to select range. 
        Example input: foo.bar.baz.tld 
          -f 3      bar.baz.tld 
          -f 3:     bar.baz.tld, foo.bar.baz.tld
          -f 2:4    baz.tld, bar.baz.tld
          -f :3     tld, baz.tld
  -i string
        Input url or domain
  -if string
        Input file path, one url/domain per line.
  -o string
        Output file path, optional

```


