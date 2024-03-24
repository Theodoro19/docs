# GO Defer

### Introdução

Em Go, o `defer` é usado para limpar uma função.
Essa limpeza, geralmente, é feito no final de uma função e essas atividades são feitas por uma função diferente chamada `defer`.
Essa função diferente é chamada no final da função circundante.

```go
defer {nome_da_funcao_ou_metodo}
```

> [!NOTE]
> A execução da função é adiada até o momento em que a função circundante retorna.
> Será executada se a função envolvente terminar abruptamente.

### Função Defer Padrão

Por exemplo, suponhamos que temos um código para abrir um arquivo e escrever nele.

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    err := writeToTempFile("Some text")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
    file, err := os.Open("temp.txt")
    if err != nil {
        return err
    }
    n, err := file.WriteString("Some text")
    if err != nil {
        return err
    }
    fmt.Printf("Number of bytes written: %d", n)
    file.Close()
    return nil
}
```

A função acima `writeToTempFile` estamos tentando abrir um arquivo e tentamos gravar algum conteúdo nele. Depois de escrevermos o conteúdo do arquivo, o fechamos. É possível que durante a operação de gravação resulte um erro e a função retorne sem fechar o arquivo. A função `defer` ajuda a resolver esse problema. A função `defer` é sempre executada antes que a função circundante retorne.

Agora vamos reescrever o código acima, utilizando a função **defer**:

No código abaixo, adiamos o `file.Close()` após abrir o arquivo.
Isso garantirá que o fechamento do arquivo seja executado mesmo que a gravação no arquivo resulte em erro.
A função `defer` garantirá que o arquivo será fechado independentemente da quantidade de retornos na função.

```go
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    err := writeToTempFile("Some text")
    if err != nil {
        log.Fatalf(err.Error())
    }
    fmt.Printf("Write to file succesful")
}

func writeToTempFile(text string) error {
    file, err := os.Open("temp.txt")
    if err != nil {
        return err
    }
    defer file.Close()

    n, err := file.WriteString("Some text")
    if err != nil {
        return err
    }
    fmt.Printf("Number of bytes written: %d", n)
    return nil
}
```

### Função Defer Personalizada

Temos a possibilidade de chamar uma função personalizada em `defer`.

Exemplo:

```go
package main
import "fmt"
func main() {
    defer test()
    fmt.Println("Executed in main")
}
func test() {
    fmt.Println("In Defer")
}
```

Resultado:

```text
Executed in main
```

No exemplo acima, ha uma função personalizada defer chamada `test()`.
Notemos que a função `test()` foi chamada após tudo dentro da função `main()` ser executada e antes dela retornar.

### Múltiplas Funções Defer

Caso tenhamos várias funções `defer` dentro de uma função específica, todas as funções `defer` serão executadas em última ordem.

Como veremos no código abaixo, teremos 3 funções `defer` onde cada uma imprimirá um valor.
Vocês vão notar que a saída está invertida, pois quando temos várias funções de `defer` dentro de uma função, uma regra é seguida *Último a entrar, primeiro a sair*  e sendo assim, teremos o seguinte resultado:

Exemplo:

```go
package main
import "fmt"
func main() {
    i := 0
    i = 19
    defer fmt.Println(i)
    i = 30
    defer fmt.Println(i)
    i = 45
    defer fmt.Println(i)
}
```

Resultado:

```text
45
30
19
```

### Como funciona

Quando o compilador encontra uma instrução `defer`, ele a coloca em uma lista. Esta lista implementa, internamente, uma estrutura de dados de pilha.
Todas as instruções `defer` encontradas na mesma função são colocadas nesta lista.
Quando uma função circundante retorna, todas as funções na pilha, começando de cima para baixo, são executadas antes que a execução possa começar na função de chamada.

### Defer e Methods

A instrução `defer` também pode ser aplicadas para os métodos como são feitos com as funções.

### Defer e Panic

A instrução `defer` também será executada se ocorrer pânico em algum programa. Quando o `panic` é gerado em uma função, a execução dessa função é interrompida e qualquer função `defer` será executada.

Na verdade, a função `defer` de todas as chamadas de função na pilha também será executada até que todas as funções tenham retornado. Nesse momento, o programa será encerrado e imprimirá a mensagem de pânico.

Portanto, se uma função `defer` estiver presente, ela será executada e o controle será retornado à função chamadora, que executará novamente sua função `defer`, se presente, e a cadeia continuará até que o programa exista.

Exemplo:

```go
package main
import "fmt"
func main() {
    defer fmt.Println("Defer in main")
    panic("Panic with Defer")
    fmt.Println("After painc in f2")
}
```

Resultado:

```text
# [play]
./prog.go:8:2: unreachable code

Go vet failed.

Defer in main
panic: Panic with Defer

goroutine 1 [running]:
main.main()
	/tmp/sandbox934533563/prog.go:7 +0x68
```
