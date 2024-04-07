# Single Responsibility Principle

## Introdução

O princípio de Responsabilidade Única `(SRP)` afirma que uma classe ou módulo deve ter apenas um motivo para mudar, o que significa que deve ter apenas uma responsabilidade.
Este princípio promove melhor organização, capacidade de manutenção e reutilização de código.

## Explicação

Em Go, o conceito `SRP` pode ser aplicado tanto a pacotes quanto a funções. Embora Go, não seja uma linguagem orientada a objetos no sentido tradicional, ela ainda promove bons princípios de design de software, incluindo `SRP`.

### Packages
Idealmente, um pacote em Go deve focar em uma única responsabilidade. Isso ajuda a criar limites e responsabilidades bem definidas para o nosso código, facilitando o gerenciamento e a compreensão.
Por exemplo, um pacote que manipula operações de arquivo não deve incluir funções relacionadas às comunicações de rede.

### Functions
As funções em Go também devem aderir ao SRP, concentrando-se em uma única tarefa. Uma função com múltiplas responsabilidades pode tornar-se difícil de compreender, manter e testar. Ao dividir funções complexas em funções menores e mais focadas, você pode criar uma base de código mais modular e de fácil manutenção.

### Exemplos

Para compreendermos um pouco mais vamos ver dois exemplos abaixo, um que ilustra a boa e a má adesão ao `SRP`.

#### Violando SRP

Suponhamos que tenhamos uma função que lê dados de um arquivo .csv, faz o seu processamento e depois envia os resultados para uma API.

```go
package main

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strings"
)

func readProcessAndSendData(filename, apiUrl string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		processedData := strings.ToUpper(strings.Join(record, ","))
		_, err = http.Post(apiUrl, "text/plain", strings.NewReader(processedData))
		if err != nil {
			return err
		}
	}
	return nil
}
```

Essa função viola o `SRP` porque tem múltiplas responsabilidades.

#### Seguindo SRP

Para aderirmos ao `SRP`, podemos refatorar o código acima em 3 funções cada uma com sua responsabilidade.

```go
package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"strings"
)

func readCSVData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}
func processCSVData(data [][]string) []string {
	var processedData []string
	for _, record := range data {
		processedData = append(processedData, strings.ToUpper(strings.Join(record, ",")))
	}
	return processedData
}
func sendDataToAPI(apiUrl string, data []string) error {
	for _, record := range data {
		_, err := http.Post(apiUrl, "text/plain", strings.NewReader(record))
		if err != nil {
			return err
		}
	}
	return nil
}
```

Agora cada função tem uma única responsabilidade e nosso código é mais modular, sustentável e testável.

## Vantagens e Desvantagens

### Vantagens

- **Mais fácil de entender**
O código que segue o `SRP` é mais simples de compreender, pois cada função ou pacote tem uma responsabilidade única e bem definida.

- **Mais fácil de manter**
Quando cada componente do seu código tem uma única responsabilidade, as modificações e correções de bugs tornam-se mais gerenciáveis. Isso permite isolar as alterações e minimizar seu impacto em outras partes da base de código.

- **Maior capacidade de reutilização**
Funções ou pacotes com uma única responsabilidade tem maior probabilidade de serem reutilizáveis em diferentes partes do seu aplicativo ou até mesmo em diferentes projetos.

- **Testabilidade aprimorada**
Ao limitar o escopo de cada função ou pacote, nós podemos escrever testes mais focados e completos, resultando em maior qualidade de código e menos bugs.

### Desvantagens

- **Fragmentação excessiva**
Seguir o `SRP` às vezes pode levar a um número excessivo de pequenas funções ou pacotes, o que pode tornar o código mais difícil de navegar e entender.

- **Maior complexidade**
Embora o `SRP` possa tornar funções ou pacotes individuais mais simples, ele pode introduzir complexidade na estrutura geral do seu aplicativo se não for gerenciado com cuidado.

- **Possíveis compensações de desempenho**
Em alguns casos, aderir ao `SRP` pode exigir a divisão de uma única função em múltiplas funções menores, o que poderia introduzir alguma sobrecarga de desempenho devido a chamadas de função.
