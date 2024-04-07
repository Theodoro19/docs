# Open/Closed Principle

## Introdução

O princípio Aberto/Fechado `(OCP)` fornece orientação sobre a criação de componentes de software que são abertos para extensão, mas fechados para modificação.
Seguindo este princípio, podemos desenvolver um software mais robusto e flexível que pode se adaptar facilmente a novo requisitos ou mudanças sem modificar o código existente.

## Explicação

Em Go, o conceito `OCP` pode ser aplicado tanto a pacotes quanto a funções. Embora Go não possua classes no sentido tradicional de orientação a objetos, ele possui estruturas e interfaces, que podem ser usadas para atingir o nível desejado de abstração e extensibilidade.

### Interfaces

As interfaces desempenham um papel crucial na aplicação do `OCP` em Go. Uma interface define um contrato ou um conjunto de métodos que um tipo deve implementar, permitindo que vários tipos sejam usados de forma intercambiável enquanto aderem ao mesmo contrato.
Ao usar Interfaces, podemos estender a funcionalidade do nosso código sem alterar a implementação existente.

### Composição

Outra forma de aplicar `OCP` em Go é usando Composição. A composição permite criar tipos complexos combinando tipos menores e mais focados, cada um responsável por uma funcionalidade específica. Ao compor tipos menores, você pode estender o comportamento do seu código sem modificar os componentes existentes.

### Exemplos

#### Violando OCP

Suponha que tenhamos uma função que calcula a área de diferentes formas geométricas.

```go
package main

import "math"

type Shape struct {
	Type   string
	Width  float64
	Height float64
	Radius float64
}

func CalculateArea(s Shape) float64 {
	if s.Type == "rectangle" {
		return s.Width * s.Height
	} else if s.Type == "circle" {
		return math.Pi * s.Radius * s.Radius
	}
	return 0
}
```

Essa implementação viola o princípio porque adicionar suporte para uma nova forma exigiria a modificação da função Área, potencialmente introduzindo bugs ou quebrando funcionalidades existentes.

#### Seguindo OCP com Interfaces

Para aplicarmos o princípio, vamos reescrever o código acima usando interfaces.

Primeiro, vamos criar uma interface para calcularmos a area

```go
type Shape interface {
 Area() float64
}
```

Agora, vamos criar estrutura para para cada tipo de forma e implementar o método área para cada um.

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}
```

Agora, o método área faz parte da implementação de cada forma, aderindo ao `OCP`. Para adiciona uma nova forma, basta criarmos uma nova estrutura e implementar a interface `Shape`, sem modificar nenhum código existente.

#### Seguindo OCP com Composição

Neste exemplo, temos um PaymentProcessor que processa pagamentos usando diferentes gateways de pagamento, como Stripe e PayPal. Em vez de codificar a lógica do gateway de pagamento dentro do PaymentProcessor, podemos aplicar OCP usando composição.

Primeiro, vamos definir uma interface PaymentGateway.

```go
type PaymentGateway interface {
 ProcessPayment(amount float64) error
}
```

Em seguida, criaremos estruturas para cada gateway de pagamento e implementaremos o método ProcessPayment para cada um deles.

```go
type StripeGateway struct {
 APIKey string
}
func (s StripeGateway) ProcessPayment(amount float64) error {

}
type PayPalGateway struct {
 ClientID string
 ClientSecret string
}
func (p PayPalGateway) ProcessPayment(amount float64) error {
 
}
```

Por fim, componha o PaymentProcessor usando a interface PaymentGateway.

```go
type PaymentProcessor struct {
 Gateway PaymentGateway
}
func (p PaymentProcessor) ProcessPayment(amount float64) error {
 return p.Gateway.ProcessPayment(amount)
}
```

Agora, o PaymentProcessor está aberto para extensão e fechado para modificação. Para adicionar um novo gateway de pagamento, basta criarmos uma nova estrutura implementando a interface PaymentGateway e passá-la para PaymentProcessor. Nenhuma modificação do código existente é necessária.

## Vantagens e Desvantagens

### Vantagens

- **Flexibilidade**
Usando `OCP` promove flexibilidade, permitindo estender seu código sem modificar a funcionalidade existente. Isto torna mais fácil acomodar novos requisitos ou alterações sem afetar a implementação atual.

- **Capacidade de Manutenção**
Seguindo o `OCP` resultamos em um código mais sustentável, pois é menos provável que introduzimos bugs ou quebre funcionalidades existentes ao fazer alterações ou adicionar novos recursos.

- **Testabilidade**
O código que segue o `OCP` é geralmente mais fácil de testar, pois cada componente pode ser testado isoladamente, tornando mais simples identificar e corrigir problemas.

- **Reutilização**
`OCP` nos incentiva a criar de componentes modulares com responsabilidades bem definidas, que podem ser facilmente reutilizados em diferentes partes da sua aplicação ou em diferentes projetos.

### Desvantagens

- **Aumento de complexidade**
A aplicação do `OCP` às vezes pode resultar em aumento da complexidade devido à introdução de abstrações adicionais e ao aumento do número de componentes.

- **Abstração prematura**
Existe o risco de criarmos abstrações desnecessárias ao tentar aderir ao `OCP`, o que pode levar a um excesso de engenharia ou a um design mais complicado do que o necessário.

- **Compensações de desempenho**
Em alguns casos, seguir o `OCP` pode nos levar a compensações de desempenho devido ao uso de interfaces ou camadas adicionais de abstração.
