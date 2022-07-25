<h1 align="center">dsieve <a href="https://twitter.com/intent/tweet?text=dsieve%20-%20Filter%20and%20enrich%20a%20list%20of%20subdomains%20by%20level%20https%3A%2F%2Fgithub.com%2Ftrickest%2Fdsieve&hashtags=bugbounty,bugbountytips,infosec"><img src="https://img.shields.io/badge/Tweet--lightgrey?logo=twitter&style=social" alt="Tweet" height="20"/></a></h1>
<h3 align="center">Filter and enrich a list of subdomains by level</h3>

![dsieve](dsieve.png "dsieve")

Take a single domain or read an input file and extract unique parent domains, enrich subdomains, filter subdomains by level, or find out which subdomains have the most number of sub-subdomains (or sub-sub-subdomains or sub-sub-sub...). Dsieve supports any format of URL, with or without protocol, port, path, parameters.

# Installation
## Binary
Binaries are available in the [latest release](https://github.com/trickest/dsieve/releases/latest).

## Docker
```
docker run quay.io/trickest/dsieve
```

## From source
```
go install github.com/trickest/dsieve@latest
```

# Usage
```
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
  -top int
    	Only consider top X subdomains of a certain level and return all their subdomains
```

### Example
##### test.txt
```
a.foo.target.com
b.foo.target.com
c.foo.target.com
a.bar.target.com
b.bar.target.com
a.baz.target.com
```

```shell script
# All levels by default
$ dsieve -if test.txt
a.foo.target.com
foo.target.com
target.com
b.foo.target.com
c.foo.target.com
a.bar.target.com
bar.target.com
b.bar.target.com
a.baz.target.com
baz.target.com

# Level 2, the main domain
$ dsieve -if test.txt -f 2
target.com

# Level 3, one level above the main domain
$ dsieve -if test.txt -f 3
foo.target.com
bar.target.com
baz.target.com

# Levels 2 and above, main domain and all its subdomains
$ dsieve -if test.txt -f 2:
a.foo.target.com
foo.target.com
target.com
b.foo.target.com
c.foo.target.com
a.bar.target.com
bar.target.com
b.bar.target.com
a.baz.target.com
baz.target.com

# The top one level 3 subdomain with the highest number of sub-subdomains
$ dsieve -if test.txt -f 3 -top 1
foo.target.com

# The top two level 3 subdomain with the highest number of sub-subdomains
$ dsieve -if test.txt -f 3 -top 1
foo.target.com
bar.target.com
```

# Report Bugs / Feedback
We look forward to any feedback you want to share with us or if you're stuck with a problem you can contact us at [support@trickest.com](mailto:support@trickest.com). You can also create an [Issue](https://github.com/trickest/dsieve/issues/new) or pull request on the Github repository.

# Where does this fit in your methodology?
Dsieve is an integral part of many workflows in the Trickest store. Sign up on [trickest.com](https://trickest.com) to get access to these workflows or build your own from scratch!

[<img src="./banner.png" />](https://trickest-access.paperform.co/)
