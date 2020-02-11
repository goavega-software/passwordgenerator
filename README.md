# Password Generator

This library generates easy to remember (one time) 8 character long passwords. It does so by making up passwords with a mix of consonants and vowels pair and then appending random digits in the end.

The password gen has few options that can be passed to it:
* LeetMode - replaces the first vowel with its 1337 equivalent
* AppendSymbol - adds a symbol in the end (making the generated password 9 chars long)

## What it doesn't do
  - Generate cryptographically strong random passwords
  - Toast bread

### Demo
* [Here](https://shh.goave.ga/)
* [Demo Code](https://github.com/sn123/passwordgen-demo)

### Sample Passwords Generated

```
j3doco67$
f3lole11$
x1rayi18#
f0qese46@
bepeha05
xirayi18#
```

### Usage
```golang
package main

import "fmt"
import "./passwordgen"

func main() {
	num := 0
	request := new(passwordgen.GenerateRequest)
	request.LeetMode = true
	for num < 8 {
		fmt.Println(passwordgen.Generate(*request))
		num++
	}
}
```
### TODO
* Add better unit tests
* Support for stop words

### License
MIT
