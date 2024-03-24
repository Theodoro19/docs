# GO Pointer

### Introdução

Ponteiro é uma variável que contém um endereço de memória de outra variável.

### Declaração do Ponteiro

```go
var x *T
```

### Inicialização do Ponteiro

Existem duas maneiras de se inicializar um ponteiro:

- Usando um novo operador
- Usando o operador E comercial (&)

**Usando um novo operador**

Exemplo:

```go
package main

import "fmt"

func main() {
	a := new(int)
	*a = 10
	fmt.Println(*a)
}
```

Resultado:

```text
10
```

> [!NOTE]
> O operador * pode ser usado para desreferenciar um ponteiro e isso significará que obteremos o valor no endereço armazenado no ponteiro.

**Usando &**

Usamos o `&` para o obtermos o endereço de uma variável.

Exemplo:

```go
package main

import "fmt"

func main() {
	a := 2
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	b = new(int)
	*b = 10
	fmt.Println(*b)
}
```

Resultado:

```text
0xc000012028
2
10
```
