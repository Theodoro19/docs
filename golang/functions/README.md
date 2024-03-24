# GO Functions

### Introdução

Em Go, uma função é como se fosse o método na linguagem Java.
As funções são variáveis de primeira ordem. Eles podem ser passados como qualquer outra variável.

> [!IMPORTANT]
> O nome da função é Case Sensitive.
> Uma função que começa com a letra maiúscula será exportada para fora do seu pacote e poderá ser chamada fora de seu pacote.
> Uma função que começa com a letra minúscula não será exportada e só ficará visível dentro de seu pacote.

```go
func <nome_da_funcao>(<parametros>) <retorno>{
    // corpo
}
```

Uma função em Go, tem as seguintes características:

- **func** é a palavra-chave
- deve possuir um nome
- pode ter nenhum ou vários parâmetros separados por vírgula
- pode ter nenhum ou vários retornos separados por vírgula
- corpo funcional
- pode retornar vários valores

Exemplo:

```go
package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3))
}

func sum(a int, b int) int {
	return a + b
}
```

Resultado:

```text
5
```

### Parâmetros

Como mencionamos, podemos ter nenhum ou mais parâmetros.
Quando tivermos parâmetros do mesmo tipo, e consecutivos, podemos declarar apenas uma única vez.

```go
func sum(a, b int)
```

### Valores de Retorno

Como mencionamos, podemos ter nenhum ou mais retornos.
Por exemplo, se tivermos uma função que queremos que retorne o valor da soma e média.

```go
func sum_avg(a, b int) (int, int)
```

Por convenção, retornamos error como último argumento.

```go
func sum(a, b int) (int, error)
```

Para coletarmos múltiplos retornos das funções, podemos escrever desta forma:

```go
result, err := sum(2, 3)
```

#### Valores de Retorno Nomeados

Em Go as funções podem ter um nome para os retornos.
Quando demos um nome para o retorno, ele não precisa ser inicializado dentro da função.

```go
func sum(a, b int) (result int)
```

### Uso das Funções

Em Go, temos 3 tipos:

- Uso genérico
- Função como tipo
- Função como valor

A diferença entre uma função como tipo e uma função como valor é que no tipo usamos apenas a assinatura da função, enquanto na valor é usada a assinatura junto com o corpo.

#### Uso genérico

Exemplo:

```go
package main

import "fmt"

func main() {
    result := sum(2, 3)
    fmt.Println(result)
}

func sum(a, b int) int {
    return a + b
}
```

Resultado:

```text
5
```

#### Funções como Tipo

Em Go uma função também pode ser um tipo.

Duas funções serão do mesmo tipo se:

- Eles têm o mesmo número de argumentos e cada argumento é do mesmo tipo.
- Eles têm o mesmo número de valores de retorno e cada valor de retorno é do mesmo tipo.

A função como tipo é útil em:

- No caso de funções de ordem superior. O argumento e o tipo de retorno são especificados usando o tipo da função.
- No caso de definir interfaces, pois apenas o tipo da função é especificado.

Segue um exemplo de uma função como tipo na interface:

```go
area() int
getType() string
```

Exemplo:
```go
package main
import "fmt"
func main() {
    var shapes []shape
    s := &square{side: 2}
    shapes = append(shapes, s)
    r := &rectangle{length: 2, breath: 3}
    shapes = append(shapes, r)
    for _, shape := range shapes {
        fmt.Printf("Type: %s, Area %d\n", shape.getType(), shape.area())
    }
}
type shape interface {
    area() int
    getType() string
}
type rectangle struct {
    length int
    breath int
}
func (r *rectangle) area() int {
    return r.length * r.breath
}
func (r *rectangle) getType() string {
    return "rectangle"
}
type square struct {
    side int
}
func (s *square) area() int {
    return s.side * s.side
}
func (s *square) getType() string {
    return "square"
}
```

```text
Type: square, Area 4
Type: rectangle, Area 6
```

Uma função como tipo definida pelo desenvolvedor pode ser declarada usando a palavra-chave tipo.

Exemplo:

```go
package main

import "fmt"

func main() {
    areaF := getAreaFunc()
    print(3, 4, areaF)
}

type area func(int, int) int

func print(x, y int, a area) {
    fmt.Printf("Area is: %d\n", a(x, y))
}

func getAreaFunc() area {
    return func(x, y int) int {
        return x * y
    }
}
```

Resultado:

```text
12
```

#### Funções como Valor

Em Go uma função também pode ser um valor.

Ela também pode ser chamada de **Função Anônima** porque uma função não tem nome e pode ser atribuída a uma variável e transmitida.

> [!IMPORTANT]
> Elas geralmente são criadas para uso de curto prazo ou alguma funcionalidade limitada.

Exemplo:

```go
package main

import "fmt"

var max = func(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

func main() {
    res := max(2, 3)
    fmt.Println(res)
}
```

Resultado:

```text
3
```
