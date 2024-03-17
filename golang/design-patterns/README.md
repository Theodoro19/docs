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

    Os `Padrões Criacionais` fornecem maneiras de criar objetos e classes de maneira flexível, permitindo criar objetos sem conhecer suas classes específicas. Abaixo mostraremos alguns dos padrões criacionais mais populares em `Go`:

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

2. **Structural Patterns**

    Os `Padrões Estruturais` fornecem maneiras de compor classes e objetos para formar estruturas maiores. Abaixo mostraremos os padrões estruturais mais populares em `Go`:

    2.1 **Adapter**

    O padrão `Adaptador`permite que classes com interfaces incompatíveis trabalhem juntas criando um objeto intermediário. Ele fornece uma maneira de conectar duas interfaces incompatíveis criando um objeto intermediário. Esse padrão é útil quando precisamos adapter uma interface existente para uma nova.

    2.2 **Bridge**

    O padrão `Ponte` separa a interface de um objeto de sua implementação, permitindo que variem de forma independente. Ele fornece uma maneira de separar a interface de um objeto de sua implementação, permitindo que elas variem de forma independente. Esse padrão é útil quando precisamos dissociar uma abstração de sua implementação.

    2.3 **Composite**

    O padrão `Composto` trata objetos e grupos de objetos de maneira uniforme, permitindo criar estruturas complexas a partir de estruturas simples. Ele fornece uma maneira de tratar objetos e grupos de objetos de maneira uniforme, permitindo criar estruturas complexas a partir de estruturas simples. Esse padrão é útil quando precisamos criar uma hierarquia de objetos.

    2.4 **Decorator**

    O padrão `Decorador` atribui responsabilidades adicionais a um objeto dinamicamente. Ele fornece uma maneira de adicionar novas funcionalidades a um objeto sem alterar sua interface. Esse padrão é útil quando precisamos adicionar novas funcionalidades a um objeto existente.

    2.5 **Facade**

    O padrão `Fachada` fornece uma interface unificada para um conjunto de interfaces em um subsistema. Ele fornece uma maneira de simplificar um subsistema complexo, fornecendo uma interface unificada para ele. Esse padrão é útil quando precisamos simplificar um sistema complexo.

3. **Behavioral Patterns**

    Os `Padrões Comportamentais` fornecem maneiras de coordenar objetos e classes para cumprir tarefas específicas. Abaixo mostraremos alguns dos padrões comportamentais mais populares em `Go`:

    3.1 **Chain of Responsibility**

    O padrão `Cadeia de Responsabilidade` permite passar solicitações entre uma cadeia de objetos até que um deles lide com a solicitação. Ele fornece uma maneira de passar solicitações entre uma cadeia de objetos até que um deles lide com a solicitação. Esse padrão é útil quando precisamos fornecer a mais de um objeto a oportunidade de lidar com uma solicitação.

    3.2 **Command**

    O padrão `Comando`encapsula uma solicitação como um objeto, permitindo parametrizar clientes com diferentes solicitações, enfileirar ou registrar solicitações e oferecer suporte a operações que podem ser revertidas. Ele fornece uma maneira de encapsular uma solicitação como um objeto, permitindo parametrizar clientes com diferentes solicitações, enfileirar ou registrar solicitações e oferecer suporte a operações que podem ser revertidas. Esse padrão é útil quando precisamos dissociar uma solicitação de sua implementação.

    3.3 **Interpreter**

    O padrão `Intérprete` define uma representação gramatical para uma linguagem e um intérprete que interpreta a linguagem. Ele fornece uma maneira de definir uma representação gramatical para um idioma e um intérprete que interpreta o idioma. Esse padrão é útil quando precisamos criar um intérprete de linguagem.

    3.4 **Iterator**

    O padrão `Iterador` fornece uma maneira de acessar sequencialmente os elementos de um objeto agregado sem expor sua representação subjacente. Ele fornece uma maneira de acessar sequencialmente os elementos de um objeto agregado sem expor sua representação subjacente. Esse padrão é útil quando precisamos acessar os elementos de uma coleção sem conhecer sua implementação específica.

    3.5 **Mediator**

    O padrão `Mediador` reduz o acoplamento entre classes, permitindo que elas se comuniquem indiretamente através de um objeto mediador. Ele fornece uma maneira de reduzir o acoplamento entre classes, permitindo que elas se comuniquem indiretamente através de um objeto mediador. Esse padrão é útil quando precisamos reduzir o acoplamento entre classes.

    3.6 **Memento**

    O padrão `Lembrança` fornece uma maneira de capturar e restaurar o estado interno de um objeto. Ele fornece uma maneira de capturar e restaurar o estado interno de um objeto. Esse padrão é útil quando precisamos salvar e restaurar o estado de um objeto.

    3.7 **Observer**

    O padrão `Observador` define uma dependência um-para-muitos entre objetos, de modo que quando um objeto muda de estado, todos os seus dependentes são notificados e atualizados automaticamente. Ele fornece uma maneira de definir uma dependência um-para-muitos entre objetos, de modo que quando um objeto muda de estado, todos os seus dependentes sejam notificados e atualizados automaticamente. Esse padrão é útil quando precisamos criar um sistema que reaja às mudanças no estado de um objeto.

    3.8 **State**

    O padrão `Estado` permite que um objeto altere seu comportamento quando seu estado interno muda. Ele fornece uma maneira de alterar o comportamento de um objeto quando seu estado interno muda. Esse padrão é útil quando precisamos criar um sistema que reaja às mudanças no estado de um objeto.

    3.9 **Strategy**

    O padrão `Estratégia` define uma família de algoritmos, encapsula cada um deles e os torna intercambiáveis. Ele fornece uma maneira de definir uma família de algoritmos, encapsular cada um deles e torná-los intercambiáveis. Esse padrão é útil quando precisamos selecionar um algoritmo em tempo de execução.

    3.10 **Template Method**

    O padrão `Método de Modelo` define o esqueleto de um algoritmo em um método, permitindo que subclasses forneçam suas próprias implementações para determinadas etapas. Ele fornece uma maneira de definir o esqueleto de um algoritmo em um método, permitindo que subclasses forneçam suas próprias implementações para determinadas etapas. Esse padrão é útil quando precisamos definir um algoritmo comum com variações.

    3.11 **Visitor**

    O padrão `Visitante` separa um algoritmo de uma estrutura de objeto movendo o algoritmo para um novo objeto. Ele fornece uma maneira de separar um algoritmo de uma estrutura de objeto, movendo o algoritmo para um novo objeto. Esse padrão é útil quando precisamos aplicar algoritmos diferentes a uma estrutura de objeto.