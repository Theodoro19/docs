# SOLID

## Introdução

`SOLID` é um acrônimo para cinco postulados de design de código em Programação Orientada a Objeto, utilizados para facilitar a compreensão, o desenvolvimento e a manutenção de software.

> [!IMPORTANT]
> Estes princípios são fundamentais na programação orientada a objetos e podem ser aplicados em qualquer linguagem que adote este paradigma.

Os cinco princípios são:

- **Single Responsibility Principle**
- **Open/Closed Principal**
- **Liskov Substitution Principle**
- **Interface Segregation Principle**
- **Dependency Inversion Principle**

Embora Go não ser uma linguagem orientada a objetos, ela ainda incorpora os conceitos orientados a objetos e pode se beneficiar dos princípios `SOLID`.
Aplicar `SOLID` em nossos projetos Go pode nos levar a um melhor design de software, uma melhor qualidade de código e a colaboração entre os membros da equipe são facilitados.

## O que é significa cada um?

### Princípio da Responsabilidade Única

O princípio `Single Responsibility` nos diz que, cada classe dentro de um projeto deve se especializar em um único assunto e possuir uma única responsabilidade, tendo uma única tarefa ou ação para executar.

#### Exemplo

Suponhamos que temos um sistema de registro simples. Em vez de ter uma única função que lide com a formatação e gravação de logs em saídas diferentes, iremos separar essas responsabilidades em funções separadas.

```go
type LogFormatter interface {
	Format(message string) string
}
type LogWriter interface {
	Write(message string) error
}

func LogMessage(formatter LogFormatter, writer LogWriter, message string) error {
	formattedMessage := formatter.Format(message)
	return writer.Write(formattedMessage)
}
```

### Princípio Aberto/Fechado

O princípio `Open/Closed` nos diz que, ao criar um objeto ou entidade, devemos prepará-lo para que possa ser implementado por outro futuramente. Assim, não será necessário modificar o objeto pai.

#### Exemplo

Suponhamos que temos uma loja de aplicativos e precisamo aplicar desconto para cada tipo de produto. Ao usarmos uma interface conseguimos adicionar novos descontos.

```go
type Product struct {
	Price float64
}
type Discount interface {
	Calculate(product Product) float64
}

func ApplyDiscount(discount Discount, product Product) float64 {
	return product.Price - discount.Calculate(product)
}
```

### Princípio de Substituição de Liskov

O princípio `Liskov Substitution` nos diz que, se uma função necessita de uma classe que está sendo implementada por outras, podemos utilizar tanto a classe base quanto as classes que são derivadas desta primeira. Isso porque todas seguem os mesmos padrões, alterando apenas as implementações da original.

#### Exemplo

Vamos pegar as formas geométricas, tanto o Retângulo quanto o Quadrado podem ser usados alternadamente.

```go
type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	SideLength float64
}

func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}
```

### Princípio da Segregação de Interface

O princípio `Interface Segregation` nos diz que, devemos criar interfaces mais específicas para nossos objetos, ao invés de uma classe mais genérica para todos do mesmo tipo.

#### Exemplo

Suponhamos que temos um sistema de armazenamento de arquivos com vários provedores de armazenamento.

```go
type FileReader interface {
    ReadFile(path string) ([]byte, error)
}

type FileWriter interface {
    WriteFile(path string, data []byte) error
}

type FileStorage struct {

}

func (fs *FileStorage) ReadFile(path string) ([]byte, error) {
    
}


func (fs *FileStorage) WriteFile(path string, data []byte) error {
    
}
```

### Princípio da Inversão de Dependência

O princípio `Dependency Inversion` nos diz que, ao desacoplarmos nossas classes de bibliotecas específicas e fazer com que outras ferramentas possam ser utilizadas no lugar desta primeira. Assim, as classes utilizarão abstrações de interfaces ao invés de outras classes ou de instâncias de objetos.

#### Exemplo

Suponhamos que temos um sistema de notificação que envia mensagens por sms ou e-mail.

```go
type Notifier interface {
 Notify(message string) error
}
type EmailNotifier struct {
 
}
func (en *EmailNotifier) Notify(message string) error {
 
}
type SMSNotifier struct {
 
}
func (sn *SMSNotifier) Notify(message string) error {
 
}
func SendNotification(notifier Notifier, message string) error {
    return notifier.Notify(message)
}
```

## Vantagens e Desvantagens

### Vantagens

- **Manutenção**
Utilizando `SOLID` resultamos em um código mais sustentável, tornando mais fácil corrigir bugs, adicionar novas funcionalidades e refatorarmos o código.

- **Escalabilidade**
Os princípios `SOLID` promovem a criação de código flexível e modular, o que é crucial para dimensionar nossas aplicações e acomodarmos novos recursos.

- **Reutilização**
Utilizando `SOLID` podemos criar componentes modulares e desaclopados, aumentando a capacidade de reutilização do seu código em diferentes projetos ou partes de um projeto.

- **Testabilidade**
Os princípios `SOLID`, em especial o princípio de Inversão de Dependência, podem melhorar significativamente a testabilidade do seu código, facilitando a escrita e a manutenção de testes unitários.

- **Legibilidade**
Os princípios `SOLID` leva a um código mais organizado e legível, tornando mais fácil para nós entendermos e trabalharmos.

### Desvantagens

- **Excesso de Engenharia**
Uma desvantagem de aplicarmos os princípios `SOLID` é o risco de engenharia excessiva de código. É muito importante buscarmos o equilíbrio entro os princípios e mantermos o nosso código simples e direto.

- **Maior Complexidade**
Ao aderirmos aos princípios `SOLID` às vezes pode resultar em maior complexidade, especialmente se não tomarmos cuidado ao gerenciarmos abstrações e dependências. No entanto, isso muitas vezes pode ser mitigado com design e uma refatoração cuidadosa.

- **Curva de Aprendizado**
Quem não está familiarizado com os princípios `SOLID`, pode haver uma curva de aprendizado envolvida na compreensão e aplicação eficaz desses princípios.

Apesar das desvantagens, os benefícios de aplicarmos os princípios `SOLID` em Golang superam as desvantagens, resultando em melhor qualidade de software, capacidade de manutenção e escalabilidade.

## Dicas

Se ainda assim não conseguiram encontrar benefícios em utilizar os princípios `SOLID` segue abaixo algumas dicas para aplicarmos os princípios de forma eficaz.

- **Aproveitando as Interfaces**
As interfaces em Go são ferramentas poderosas e flexíveis que podem nos ajudar a implementar muitos dos princípios `SOLID`.
Devemos usar interfaces com sabedoria para definir o comportamento de seus tipos, abstrair dependências e criar código modular e reutilizável.

- **Abraçando a Composição**
Em Go somos incentivados a abraçar a composição ao invés da herança. Devemos aproveitar esta abordagem para criarmos códigos flexíveis e extensíveis que aderem ao princípio `Open/Closed` e ao princípio `Liskov Substitution`.

- **Apliquem Injeção de Dependência**
A injeção de dependência é uma técnica que pode nos ajudar a seguirmos o princípio `Dependency Inversion`.
Devemos usar injeção de dependência para fornecermos implementações concretas de interfaces e desacloparmos o nosso código, o tornando mais fácil de manter e testar.

- **Iterando e Refatorando**
Não devemos ter medo de revisitar e refatorar o nosso código a medida que obtemos uma compreensão mais profunda dos princípios, e benefícios, `SOLID`.

### Conclusão

Conseguimos compreender melhor os princípios `SOLID` e como eles podem ser aplicados para melhorar o design e a qualidade do software de nossos projetos.
Aderindo esses princípios, podemos criar códigos sustentáveis, escaláveis e robustos, mesmo que Go não seja uma linguagem orientado a objeto.

Embora os princípios `SOLID` sejam diretrizes valiosas para o design de software, eles não devem ser seguidos dogmaticamente. Encontrar um equilíbrio entre aderir aos princípios `SOLID` e mantermos um código simples e direto é a chave para a criação de um software eficaz e fácil de entender.
