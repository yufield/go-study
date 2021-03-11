package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(regexp.QuoteMeta("[foo]"))                                    //\[foo\]
	fmt.Println(regexp.Match("\\d[a-z]{3}", []byte("1abc")))                  //true <nil>
	fmt.Println(regexp.MatchString("\\d[a-z]{3}", "1abc"))                    //true <nil>
	fmt.Println(regexp.MatchString("\\d[a-z]{3}", "123"))                     //false <nil>
	fmt.Println(regexp.MatchReader("\\d[a-z]{3}", strings.NewReader("1abc"))) //true <nil>
	compile, _ := regexp.Compile("a\\d[a-z]{3}")
	fmt.Println(compile)                 //a\d[a-z]{3}
	fmt.Println(compile.LiteralPrefix()) //a false

}
