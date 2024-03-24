# GO Variables

Em GO, os nomes das variáveis só podem começar com uma letra ou underline e ele é **Case Sensitive** então letras maiúsculas e minúsculas são tratadas de formas diferentes.

### Como declarar uma variável

Em GO, as variáveis podem ser declaradas das seguintes maneiras:

**Declaração simples sem um valor inicial**

```go
var <nome_da_variavel> <tipo>
```

Exemplo:

```go
package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)
}
```

Resultado:

```text
0
```

**Declaração simple com um valor inicial**

```go
var <nome_da_variavel> <tipo> = <valor>
```

Exemplo:

```go
package main

import "fmt"

func main() {
	var x int = 19
	fmt.Println(x)
}
```

Resultado:

```text
19
```

**Multiplas declarações sem valor inicial**

```go
var <nome_da_variavel1>, <nome_da_variavel2>, ..., <nome_da_variavelN> <tipo>
```

Exemplo:

```go
package main

import "fmt"

func main() {
	var x, y int
	fmt.Println(x)
    fmt.Println(y)
}
```

Resultado:

```text
0
0
```

**Multiplas declaraçoes com valor inicial**

```go
var <nome_da_variavel1>, <nome_da_variavel2>, ..., <nome_da_variavelN> <tipo> = <valor1>, <valor2>, ..., <valorN>
```

Exemplo:

```go
package main

import "fmt"

func main() {
	var x, y int = 19, 45
	fmt.Println(x)
    fmt.Println(y)
}
```

Resultado:

```text
19
45
```

**Declaração de diferentes tipos**

```go
var (
    <nome_da_variavel1> <tipo>
    <nome_da_varialve2> <tipo> = <valor>
)
```

Exemplo:

```go
package main

import "fmt"

func main() {
	var (
        x int
        y int = 19
        z string = "Raphael"
    )

    fmt.Println(x)
    fmt.Println(y)
    fmt.Println(z)
}
```

Resultado:

```text
0
19
Raphael
```

**Declaração sem tipo**

```go
var <nome_da_variavel> = <valor>
```

Exemplo:

```go
package main

import "fmt"

func main() {
    var t = 123      //O tipo de interferência será int
    var u = "circle" //O tipo de interferência será string
    var v = 5.6      //O tipo de interferência será float64
    var w = true     //O tipo de interferência será bool
    var x = 'a'      //O tipo de interferência será rune
    var y = 3 + 5i   //O tipo de interferência será complex128
    var z = sample{name: "test"}  //O tipo de interferência será main.Sample

    fmt.Printf("Type: %T Value: %v\n", t, t)
    fmt.Printf("Type: %T Value: %v\n", u, u)
    fmt.Printf("Type: %T Value: %v\n", v, v)
    fmt.Printf("Type: %T Value: %v\n", w, w)
    fmt.Printf("Type: %T Value: %v\n", x, x)
    fmt.Printf("Type: %T Value: %v\n", y, y)
    fmt.Printf("Type: %T Value: %v\n", z, z)
}

type sample struct {
    name string
}
```

Resultado:

```text
Type: int Value: 123
Type: string Value: circle
Type: float64 Value: 5.6
Type: bool Value: true
Type: int32 Value: 97
Type: complex128 Value: (3+5i)
Type: main.sample Value: {test}
```

**Declaração curta**

Existe uma outra maneira de declarar variáveis que é usando o operador `:=`.
Quando usamos esse operado o Go omite o tipo e a palavra-chave `var`.

Exemplo:

```go
package main

import "fmt"

func main() {
	t := 123                  //O tipo de interferência será int
	u := "circle"             //O tipo de interferência será string
	v := 5.6                  //O tipo de interferência será float64
	w := true                 //O tipo de interferência será bool
	x := 'a'                  //O tipo de interferência será rune
	y := 3 + 5i               //O tipo de interferência será complex128
	z := sample{name: "test"} //O tipo de interferência será main.Sample

	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

type sample struct {
	name string
}
```

Exemplo:

```text
Type: int Value: 123
Type: string Value: circle
Type: float64 Value: 5.6
Type: bool Value: true
Type: int32 Value: 97
Type: complex128 Value: (3+5i)
Type: main.sample Value: {test}
```

> [!NOTE]
> Esse operado só é permitido dentro de uma função.
> Uma variável declarada uma vez com := não pode ser redeclarada usando :=
> Ele pode ser usado para múltiplas variáveis em uma única linha.
