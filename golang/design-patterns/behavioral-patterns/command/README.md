# Command

O padrão de projeto `Command` que sugere encapsularmos a solicitação como um objeto independente. O objeto criado possui todas as informações sobre a solicitação e assim pode executá-la de forma independente.

Os componentes básicos usados nesse padrão são:

- **Receptor** - É a classe que contém a lógica do negócio. O objeto de comando apenas atrasa suas solicitações ao receptor.

- **Comando** - Ela incorpora o receptor e vincula uma ação específica do receptor.

- **Invocador** - Ela incorpora o comando chamando o método de execução do comando.

- **Cliente** - Ela é quem cria o comando com o receptor apropriado ignorando o receptor para o construtor do comando e associa o comando resultando a um invocador.

## Exemplo

Vamos imaginar uma televisão, ela pode ser ligada através de um botão do controle remoto ou pelo botão na própria televisão.

Os dois possuem fazem a mesma coisa que é ligar a televisão.
Para ligar a televisão, podemos implementar o objeto de comando `ON` com o receptor sendo a `TV`. Quando o método `execute()` é chamado no objeto de comando `ON`, ele chamará a função `tv.on()`.

Então, neste caso teremos:

- TV como **receptor**
- O objeto de comando ON que incorporado a TV é o **comando**
- O botão ON do controle remoto ou da TV é o **invocador**

## Estrutura de Código

[button.go](main/button.go) - Invocador
[command.go](main/command.go) - Interface de Comando
[onCommand.go](main/onCommand.go) - Comando Concreto
[offCommand.go](main/offCommand.go) - Comando Concreto
[device.go](main/device.go) - Interface Receptora
[tv.go](main/tv.go) - Receptor Concreto
[main.go](main/main.go) - Cliente

## Código completo

```go
package main

import "fmt"

type button struct {
    command command
}

func (b *button) press() {
    b.command.execute()
}

type command interface {
    execute()
}

type offCommand struct {
    device device
}

func (c *offCommand) execute() {
    c.device.off()
}

type onCommand struct {
    device device
}

func (c *onCommand) execute() {
    c.device.on()
}

type device interface {
    on()
    off()
}

type tv struct {
    isRunning bool
}

func (t *tv) on() {
    t.isRunning = true
    fmt.Println("Turning tv on")
}

func (t *tv) off() {
    t.isRunning = false
    fmt.Println("Turning tv off")
}

func main() {
    tv := &tv{}
    onCommand := &onCommand{
        device: tv,
    }
    offCommand := &offCommand{
        device: tv,
    }
    onButton := &button{
        command: onCommand,
    }
    onButton.press()
    offButton := &button{
        command: offCommand,
    }
    offButton.press()
}
```

## Resultado

```text
Turning tv on
Turning tv off
```