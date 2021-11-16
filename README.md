Elite Planet Name Generator
===========================

[![Go Reference](https://pkg.go.dev/badge/github.com/fherr/elite.svg)](https://pkg.go.dev/github.com/fherr/elite)

Name Generator based on the Ian Bell's "text Elite" (see link below).



Example
-----

```
package main

import (
	"fmt"
	"github.com/fherr/elite"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", elite.GetRandomName())
	}
}

-----
output:

requgele
ceinriar
beatti
orbe
mamaxea
leenso
esbea
enlais
leraarar
onusatma
```

Links
-----

* Ian Bell's Text Elite Page (<http://www.iancgbell.clara.net/elite/text/index.htm>): A C implementation of the classic Elite trading system.
* [richcarl/txtelite](https://github.com/richcarl/txtelite): An Erlang version of Ian Bell's "text Elite". Also contains an updated version of the original C code.
* [jabbalaci/Elite](https://github.com/jabbalaci/Elite/blob/master/pyplanets.py) Elite Planet Name Generator wrtitten in python 