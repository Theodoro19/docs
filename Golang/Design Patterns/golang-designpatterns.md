# Golang Design Patterns

### Introdução

Sabemos que o desenvolvimento de software pode ser uma tarefa difícil e uma linguagem que não temos tanta familiaridade, pode demandar mais tempo para o desenvolvimento.

A nossa ideia foi criar guias sobre os `Design Patterns` que podemos utilizar na linguagem `Golang`.

O que iremos demonstrar são alguns padrões populares discutidos pela `Gang of Four` que podem ser aplicados. Ao compreender esses padrões, podemos criar sistemas de softwares mais flexíveis, escaláveis e de fácil manutenção.

> [!NOTE]
> Todos os Padrões de Projeto mencionados podem ser encontrados na página do [Refactoring Guru](https://refactoring.guru/pt-br/design-patterns/go).

### O que é?

Os **Padrões de Projeto** são soluções típicas para problemas comuns em projeto de software.
Por exemplo, quando vamos construir uma casa pedimos para um Arquiteto construir uma planta.
Nesse momento, o Arquiteto vai consultar uns modelos de plantas de obra pré fabricadas e encaminhará para nós selecionarmos qual queremos. Escolhemos um modelo X porém queremos fazer alguma estilização.

Com os padrões de projeto funciona de forma parecida, quando vamos construir uma casa podemos copiar um projeto pré moldado e tudo bem, mas quando encontramos um padrão, não podemos apenas copiá-lo em nosso código. O padrão de projeto é um conceito que para resolvermos um problema e não uma função de uma lib que importamos para utilizar.

### Porque aprender?

Saber os padrões, mesmo que tenhamos anos de experiência e nunca os utilizados, é muito útil porque eles nos ensinam como resolver vários problemas usando os princípios de projeto orientado a objetos.

Alé do que, quando aprendemos sobre eles, podemos nos comunicar com outros desenvolvedores de forma mais eficiente. Quando algum desenvolvedor que ler o nosso código, e conhece sobre padrões, podemos apenas dizer *"Nesse projeto utilizamos a Cadeia de Responsabilidade"*. Assim que falarmos isso, o desenvolvedor que está lendo nosso código saberá a ideia e não precisaremos explicá-lo.

### Quais podemos utilizar?

1. **Creational Patterns**

    Os `Padrões Criacionais` fornecem maneiras de criar objetos e classes de maneira flexível, permitindo criar objetos sem conhecer suas classes específicas. Abaixo mostraremos alguns dos padrões mais populares em `Go`.

    1.1 **Abstract Factory**

    O padrão `Fábrica Abstrata` cria famílias de objetos relacionados sem especificar suas classes concretas. Ele fornece uma interface para criação de famílias de objetos relacionados, sem especificar suas classes concretas. Esse padrão é útil quando precisamos criar um grupo de objetos relacionados que funcionam juntos.

    1.2 **Builder**

    O padrão `Construtor` separa a construção de objetos de sua representação, permitindo criar objetos complexos passo a passo. Ele fornece uma maneira de criar objetos complexos sem expor sua estrutura interna. Esse padrão é útil quando precisamos criar objetos complexos com múltiplas dependências.

    1.3 **Factory Method**

    O padrão `Método de Fábrica` define uma interface para criação de objetos, permitindo que as subclasses decidam qual classe instanciar. Ele fornece uma interface para criação de objetos, mas delega às subclasses a responsabilidade de qual classe instanciar. Esse padrão é útil quando precisamos criar objetos sem conhecer sua classe específica.

    1.4 **Prototype**

    O padrão `Protótipo` cria novos objetos clonando os existentes. Ele fornece uma maneira de criar novos objetos copiando os existentes. Esse padrão é útil quando precisamos criar um grande número de objetos com propriedades semelhantes.

    1.5 **Singleton**

    O padrão `Solteiro` garante que uma classe tenha apenas uma instância e fornece um ponto global de acesso a ela. Ele fornece uma maneira de garantir que uma classe tenha apenas uma instância e fornece um ponto global de acesso a ela. Esse padrão é útil quando precisamos limitar o número de instâncias de uma classe.