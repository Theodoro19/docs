# GO Routines

## Introdução

As `goroutines` podem ser consideradas uma thread leve que possui uma execução independente separada e que pode ser executada simultaneamente com outras outras `goroutines`. É totalmente gerenciado pelo tempo de execução.

Por conta do Go ser uma linguagem concorrente, cada `goroutine` é uma execução independente e é ela que nos ajuda a alcançar a simultaneidade nessa linguagem.

## Como usar

Para usar devemos utilizar a palavra-chave **`go`** antes da chamada de uma função ou método.

```go
go test()
```

Como no exemplo acima, a função `test()` agora será executada na `goroutine`.

> [!NOTE]
> Não é a função ou método que determina se é uma goroutine.
> Se a função ou método for chamado com a palavra-chave go ele será executado em uma goroutine.

Agora vamos ver a diferença de uma função rodando normalmente e uma função rodando em goroutine:

- Uma função normal

```go
test1
start()
test2
```

Para o cenário acima, teremos:

1. O `test1` será executado
2. A função `start()` será chamada
3. Assim que finalizar a função, o `test2` será executado

- Uma função em goroutine

```go
test1
go start()
```

Para o cenário acima, teremos:

1. O `test1`será executado
2. A função `start()` será chamada como `goroutine` que será executada de forma **assíncrona**
3. O `test2` será executado imediatamente. 

Quando executamos uma função em goroutine ela é executada simultaneamente com a execução do resto do programa.

Basicamente, ao chamar uma função em goroutine, a chamada retornará imediatamente e a execução continuará a partir da próxima linha enquanto a goroutine está sendo executada simultaneamente em segundo plano.

> [!IMPORTANT]
> Qualquer valor da goroutine será ignorado.

Para entendermos mais do motivo de ser ignorado, vamos ver o exemplo abaixo:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go start()
    fmt.Println("Started")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished")
}

func start() {
    fmt.Println("In Goroutine")
}
```

Resultado:

```text
Started
In Goroutine
Finished
```

No programa acima, iniciará uma goroutine que executará a função `start()`. O programa irá imprimir a mensagem "Started" e note que o imprimimos a mensagem e depois nossa goroutine é iniciada.
Isso ilustra o ponto em que mencionamos que depois que uma goroutine é iniciada, a chamada continuará para a próxima linha.
Colocamos um timeout para que a goroutine seja agendada antes que a goroutine principal exista.
Depois disso, a goroutine é iniciada e imprime a mensagem "In Goroutine".
E pra finalizar o nosso programa, é impresso a mensagem "Finished".

Agora se removermos o timeout, acontece o seguinte:

```go
package main
import (
    "fmt"
)
func main() {
    go start()
    fmt.Println("Started")
    fmt.Println("Finished")
}
func start() {
    fmt.Println("In Goroutine")
}
```

Resultado:

```text
Started
Finished
```

Vejam que a mensagem "In Goroutine" não foi impressa e isso significa que nossa goroutine não foi executada.
Isso aconteceu porque o nosso programa foi encerrado antes que a goroutine pudesse ser agendada.

Para entendermos um pouco mais, vamos trazer para a discussão a `goroutine principal`.

## Main Goroutine

A função `main()` do pacote main é a goroutine principal.

Essa goroutine pode então iniciar outras goroutines e assim por diante.

Goroutines não possuem pais e filhos. Quando uma goroutine é iniciada, ela é executada com todas as outras goroutines em execução. Cada goroutine sai apenas quando a sua função retorna (a única exceção é quando todas as goroutines saem quando a goroutine principal sai).

Segue abaixo um exemplo de que as goroutines não possuem pais e filhos:

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    go start()
    fmt.Println("Started")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished")
}

func start() {
    go start2()
    fmt.Println("In Goroutine")
}
func start2() {
    fmt.Println("In Goroutine2")
}
```

Resultado:

```text
Started
In Goroutine
In Goroutine2
Finished
```

> [!CAUTION]
> O timeout foi apenas para exemplo, é recomendado não utilizar em Produção.

## Múltiplas Goroutines

Exemplo:

```go
package main

import (
    "fmt"
    "time"
)

func execute(id int) {
    fmt.Printf("id: %d\n", id)
}

func main() {
    fmt.Println("Started")
    for i := 0; i < 10; i++ {
        go execute(i)
    }
    time.Sleep(time.Second * 2)
    fmt.Println("Finished")
}
```

Resultado:

```text
Started
id: 3
id: 1
id: 0
id: 6
id: 4
id: 5
id: 7
id: 8
id: 2
id: 9
Finished

Program exited.
```

Cada vez que executarmos o programa acima, obteremos resultados diferentes, pois como as goroutines são executadas em paralelo, não conseguimos saber qual será executado primeiro.

## Agendamento de Goroutine

Assim que nosso programa é iniciado, o tempo de execução iniciará threads do sistema operacional equivalente ao número de CPUs lógicas utilizáveis pelo processo atual.

Existe uma CPU lógica por núcleo virtual, onde um núcleo virtual significa:

```go
nucleo_virtual = x * numero_nucleo_fisico
```

Onde "x" é o número de threads de hardware por núcleo.

Existe uma função em Go, que permite descobrirmos o número de processadores lógicos disponíveis.

```go
package main
import (
    "fmt"
    "runtime"
)
func main() {
    fmt.Println(runtime.NumCPU())
}
```

Resultado:

```text
8
```

> [!NOTE]
> O teste acima foi feito em [go dev](https://go.dev/play/).

O programa Go iniciará threads do sistema operacional igual ao número de CPUs lógicas disponíveis para ele ou para a saída da função `runtime.NumCPU()`.
Essas threads serão gerenciadas pelo sistema operacional e o agendamento dessas threads nos núcleos da CPU é de responsabilidade apenas do sistema operacional.

Em Go seu tempo de execução tem seu próprio agendador que irá multiplexar as goroutines nas threads de nível do sistema operacional no tempo de execução do programa. Então, cada goroutine está sendo executado em uma thread do sistema operacional atribuído a uma CPU lógica.

Existem duas filas que podemos utilizar para gerenciar as nossas goroutines.

### Fila de execução local

Quando dentro do tempo de execução do programa, cada thread do sistema operacional possui uma fila associada a ele, é chamado de fila de execução local `LRQ`.

Ele contém todas as goroutines que serão executadas no contexto dessa thread.
O tempo de execução do programa fará o agendamento e a troca do contexto das goroutines pertencentes a um LRQ específico para a thread de nível de sistema operacional correspondente que possui esse LRQ.

### Fila de execução global

Ele contém todas as goroutines que não foram movidas para nenhum LRQ de qualquer thread do sistema operacional.
O agendador atribuirá uma goroutine desta fila à fila de execução local de qualquer thread do sistema operacional.

INCLUIR DIAGRAMA

## Agendador Cooperativo

## Vantagens sobre Threads

## Goroutines anônimas