# Password Generator

This library generates easy to remember (one time) 8 character long passwords. It does so by making up passwords with a mix of consonants and vowels pair and then appending random digits in the end.

## What it doesn't do
  - Generate cryptographically strong random passwords
  - Toast bread

### Demo
* [Here](https://radiant-shard-237505.appspot.com/)
* [Demo Code](https://github.com/sn123/passwordgen-demo)

### Sample Passwords Generated

```
kimo1270
gima6384
bowa1861
bepo6623
sadi5777
yata5342
reye2346
jivi0682
```

### Usage
```golang
package main

import "fmt"
import "./passwordgen"

func main() {
	num := 0
	for num < 8 {
		fmt.Println(passwordgen.Generate())
		num++
	}
}
```
### TODO
* Add better unit tests
* Support for stop words
* Make it configurable

### License
MIT
