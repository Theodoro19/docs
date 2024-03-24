# GO Loops

### Introdução

Em Go, temos duas formas de fazermos repetições:

- `for loop`
- `for-range loop`

> [!IMPORTANT]
> A estrutura de repetição `while` está faltando em Go, porém pode ser estruturada usando a repetição `for loop`.

### For Loop

A estrutura de repetição `for loop` é composta de três partes:

- **Parte de inicialização**
    É executada primeiro antes da primeira iteração.
    Pode ser qualquer instrução como uma declaração curta, função ou atribuição.
    Se tiver a declaração de uma variável, o escopo dessa variável será limitado ao laço.
    É opcional.
- **Parte de condição**
    É executada antes de cada iteração.
    Se a iteração for falsa, o laço será encerrado.
    Se a iteração for verdadeira, o laço continuará a iterar.
- **Parte posterior**
    É executada após cada iteração.
    Se a condição é verdadeira, o laço continua.
    Se a condição é falsa, o laço termina.
    Pode ser qualquer instrução, geralmente contém a lógica de incremento.
    Não pode conter inicialização.
    É opcional.

```go
for <init_part>; <condition_part>; <post_part> {

}
```

**Loop Simples**

Exemplo:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

Resultado:

```text
0
1
2
3
4
```

**Com apenas uma condição**

Exemplo:

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

Resultado:

```text
0
1
2
3
4
```

**Laço infinito**

Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    i := 0
    for {
        fmt.Println(i)
        i++
        time.Sleep(time.Second * 1)
    }
}
```

Resultado:

```text
0
1
2
3
4
5
.
.
```

**Usando o Break**

O `break` ajuda a sair do laço. Nada é executado após ele ser executado.

Exemplo:

```go
package main

import (
    "fmt"
)

func main() {
    i := 0
    for {
        fmt.Println(i)
        i++
        if i >= 5 {
            break
        }
    }
}
```

Resultado:

```text
0
1
2
3
4
```

**Usando o Continue**

O `continue` ajuda a pular a interação atual do laço. Nenhuma instrução após o `continue` é executada e a execução atinge o início novamente na próxima iteração.

Exemplo:

```go
package main

import "fmt"

func main() {
    for i := 1; i < 10; i++ {
        if i%3 == 0 {
            continue
        }
        fmt.Println(i)
    }
}
```

Resultado:

```text
1
2
4
5
7
8
10
```

**Laço Aninhado**

Exemplo:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        fmt.Printf("Outer loop iteration %d\n", i)
        for j := 0; j < 2; j++ {
            fmt.Printf("i= %d j=%d\n", i, j)
        }
    }
}
```

Resultado:

```text
Outer loop iteration 0
i= 0 j=0
i= 0 j=1
Outer loop iteration 1
i= 1 j=0
i= 1 j=1
Outer loop iteration 2
i= 2 j=0
i= 2 j=1
```

**Chamando uma função**

Exemplo:

```go
package main

import "fmt"

func main() {
    i := 1

    for test(); i < 3; i++ {
        fmt.Println(i)
    }

    for i = 2; i < 3; i++ {
        fmt.Println(i)
    }
}

func test() {
    fmt.Println("In test function")
}
```

Resultado:

```text
In test function
1
2
2
```

**While**

Como o Go não possui a palavra-chave `while` podemos usar a palavra-chave `for` para simularmos o while.

> [!NOTE]
> Para termos o `for` com o mesmo comportamento do `while`, precisamos ignorar a `init_part` e `post_part`.

Exemplo:

```go
package main

import "fmt"

func main() {
    i := 1
    for i <= 5 {
        fmt.Println(i)
        i++
    }
}
```

Resultado:

```text
1
2
3
4
5
```

### For-Range Loop

Ele é usado para iterar sobre diferentes coleções de estruturas de dados. Como:

- Strings
- Array
- Slice
- Maps
- Channel

**array/slice**

```go
for index, value := range array/slice 
```

Exemplo:

```go
package main

import "fmt"

func main() {
    letters := []string{"a", "b", "c"}

    fmt.Println("Both Index and Value")
    for i, letter := range letters {
        fmt.Printf("Index: %d Value:%s\n", i, letter)
    }

    fmt.Println("\nOnly value")
    for _, letter := range letters {
        fmt.Printf("Value: %s\n", letter)
    }

    fmt.Println("\nOnly Index")
    for i := range letters {
        fmt.Printf("Index: %d\n", i)
    }

    fmt.Println("\nWithout Index and Value")
    i := 0
    for range letters {
        fmt.Printf("Index: %d Value: %s\n", i, letters[i])
        i++
    }
}

```

Resultado:

```text
Both Index and Value
Index: 0 Value:a
Index: 1 Value:b
Index: 2 Value:c

Only value
Value: a
Value: b
Value: c

Only Index
Index: 0
Index: 1
Index: 2

Without Index and Value
Index: 0 Value: a
Index: 1 Value: b
Index: 2 Value: c

Program exited.
```

**string**

Uma string em Go é uma sequência de bites.
Seu unicode padrão é UTF-8. Em UTF-8, os 128 caracteres ASCII são correspondentes a um bite. Todos os outros são entre um e quatro bites.

```go
for index, character := range string
```

Exemplo:

```go
package main

import "fmt"

func main() {
	sample := "a£c"

	fmt.Printf("Length is %d\n", len(sample))

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%c\n", sample[i])
	}

	fmt.Println("\nBoth Index and Value")
	for i, letter := range sample {
		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))
	}

	fmt.Println("\nOnly value")
	for _, letter := range sample {
		fmt.Printf("Value:%s\n", string(letter))
	}

	fmt.Println("\nOnly Index")
	for i := range sample {
		fmt.Printf("Start Index: %d\n", i)
	}
}
```

Resultado:

```text
Length is 4
a
Â
£
c
Both Index and Value
Start Index: 0 Value:a
Start Index: 1 Value:£
Start Index: 3 Value:c

Only value
Value:a
Value:£
Value:c

Only Index
Start Index: 0
Start Index: 1
Start Index: 3
```

**map**

```go
for key, value := range map
```

Exemplo:

```go
package main

import "fmt"

func main() {
    sample := map[string]string{
        "a": "x",
        "b": "y",
    }

    fmt.Println("Both Key and Value")
    for k, v := range sample {
        fmt.Printf("key :%s value: %s\n", k, v)
    }

    fmt.Println("\nOnly keys")
    for k := range sample {
        fmt.Printf("key :%s\n", k)
    }

    fmt.Println("\nOnly values")
    for _, v := range sample {
        fmt.Printf("value :%s\n", v)
    }
}
```

Resultado:

```text
Both Key and Value
key :a value: x
key :b value: y

Only keys
key :a
key :b

Only values
value :x
value :y
```

**channel**

```go
for value := range channel
```

Exemplo:

```go
package main

import "fmt"

func main() {
    ch := make(chan string)
    go pushToChannel(ch)
    for val := range ch {
        fmt.Println(val)
    }
}
func pushToChannel(ch chan<- string) {
    ch <- "a"
    ch <- "b"
    ch <- "c"
    close(ch)
}
```

Resultado:

```text
a
b
c
```
s