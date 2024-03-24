# GO Constants

### Introdução

Uma constante é qualquer coisa que não muda o seu valor de origem.

> [!IMPORTANT]
> Em Go, uma constante pode ser dos seguintes tipos: "String", "Numérico", "Booleano" ou "Caracteres"

Para declararmos uma constante podemos usar a palavra-chave **`const`**.
E uma característica é que não podemos atribuir o valor depois, como é feito com as variávei.

### Como declarar uma constante

**Declaração junto com a especificação de tipo**

```go
const <nome_da_constante> <tipo> = <valor>
```

Exemplo:

```go
package main

import "fmt"

func main() {
	const x int = 19
	fmt.Println(x)
}
```

Resultado:

```text
19
```

**Declaração sem a especificação de tipo**

```go
const <nome_da_variavel> = <valor>
```

Exemplo:

```go
package main

import "fmt"

func main() {
	const x = 19
	fmt.Println(x)
}
```

Resultado:

```text
19
```

**Múltiplas declarações**

```go
const (
    <nome_da_variavel1> = <valor1>
    <nome_da_variavel2> = <valor2>
)
```

Exemplo:

```go
package main

import "fmt"

func main() {
	const (
		x = 19
		y = "Raphael"
	)
	fmt.Println(x)
	fmt.Println(y)
}
```

Resultado:

```text
19
Raphael
```

### Pontos Importantes

Uma constante não pode ser reassinada depois de sua declaração

Exemplo:

```go
package main

import "fmt"

func main() {
	const x = 19
	x = 45
	fmt.Println(x)
}
```

Resultado:

```text
./prog.go:7:2: cannot assign to x (neither addressable nor a map index expression)
```

O valor da constante deve ser conhecido em tempo de compilação, portanto, não é possível atribuirmos a chamada de uma função avaliada em tempo de execução.

Exemplo:

```go
package main

import "fmt"

const y = 19

func main() {
	const x = getValue()
	fmt.Println(x)
}

func getValue() int {
	return 45
}
```

Resultado:

```text
./prog.go:8:12: getValue() (value of type int) is not constant
```
