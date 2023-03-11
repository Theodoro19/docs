# Design Patterns

No cotidiano de um desenvolvedor de software são enfrentados diversos problemas computacionais, como precisar atualizar um sistema ou desenvolver uma aplicação.Porém, assim como na matemática, física ou química, nas quais foram desenvolvidas teorias para se resolver determinadas questões, a programação passou pelo mesmo processo.

Por isso, assim como um físico ou física não precisa se preocupar em desenvolver uma teoria para a queda livre, pois Galileu já cuidou disso séculos atrás, quem programa também não precisa se preocupar em resolver certos problemas, pois outras pessoas no passado já trabalharam nisso.

O nome que damos para essas fórmulas é design patterns (Padrões de projeto), que no final, se resumem a soluções típicas para problemas comuns em um projeto de software. Elas podem ser aplicadas independentemente da linguagem. Mas antes de mergulharmos de cabeça nesse tema precisamos definir alguns detalhes.

## O que é?

Os Padrōes de Projeto são soluções generalistas para problemas recorrentes durante o desenvolvimento de um software. Não se trata de um framework ou um código pronto, mas de uma definição de alto nível de como um problema comum pode ser solucionado.

Os padrões de projeto (design patterns) são como plantas pré-projetadas de uma construção, que você pode alterar para se adequar melhor na resolução de um problema recorrente em seu código. O que diferencia os padrões de projeto das funções e bibliotecas é que você não pode simplesmente copiá-los direto para seu programa, já que eles não são um pedaço de código, mas sim um conceito que serve como uma solução.

Assim, para se implementar um padrão de projeto você deve seguir o conceito dos padrões escolhidos (dentre todos os existentes) e ajustá-lo ao problema que deseja resolver. Isso dependerá das características do projeto. Dessa forma, fazemos uma implementação que se encaixe exatamente à demanda da nossa aplicação.

## Origem

Em 1978 os arquitetos Christopher Alexander, Sara Ishikawa e Murray Silverstein escreveram um livro chamado “A Pattern Language: Towns, Buildings, Construction” que foi publicado em português com o nome “Uma Linguagem de Padrões”. Neste livro os autores catalogaram 253 tipos de problemas (ou desafios de projeto) e analisaram o que está por trás de cada situação, descrevendo-as na sua essência e propondo uma solução padrão.

Em 1987 durante a segunda edição da OOPSLA (Object-Oriented Programming, Systems, Languages, and Applications) o engenheiro de software Kent Back, que posteriormente foi um dos criadores das metodologias Extreme Programming e Test Driven Development (TDD), junto com Ward Cunningham apresentaram uma palestra intitulada “Using Pattern Languages for Object-Oriented Programs” (Utilizando a linguagem dos padrões para programação orientada a objetos, em tradução livre). Nesta palestra eles propuseram cinco padrões de projetos no campo da ciência da computação.

Mas esses conceitos ficaram realmente conhecidos em 1994, quando os engenheiros de software Erich Gamma, Richard Helm, Ralph Johnson e John Vlissides escreveram o livro “Design Patterns: Elements of Reusable Object-Oriented Software” com o objetivo de catalogar problemas comuns aos projetos de desenvolvimento de software e as formas de resolver esses problemas. Os autores catalogaram 23 padrões que utilizaram ao longo de suas carreiras. Este livro teve mais 500.000 exemplares vendidos e foi publicado em 13 idiomas. No Brasil foi publicado com o nome “Padrões de Projeto – Soluções Reutilizáveis de Software Orientado a Objetos”. Os autores do livro ficaram conhecidos como Gang of Four (Gangue dos quatro) ou “GoF”. Depois disso muitos outros livros surgiram, alguns criticando alguns desses padrões, e outros divulgando novos padrões.

Desde então, Design Patterns tem sido um tema bastante estudado por programadores e arquitetos de software pelo mundo todo.

## Benefícios

Design patterns são modelos que já foram utilizados e testados anteriormente, portanto podem representar um bom ganho de produtividade para os desenvolvedores.

Seu uso também contribui para a organização e manutenção de projetos, já que esses padrões se baseiam em baixo acoplamento entre as classes e padronização do código.

Além disso, com a padronização dos termos, as discussões técnicas são facilitadas. É mais fácil falar o nome de um design pattern em vez de ter que explicar todo o seu comportamento.

No mercado de trabalho você irá encontrar diversos programadores experientes que não conhecem nenhum padrão de projeto, e que no entanto, já implementam alguns padrões sem nem mesmo saber. Portanto, você sairá na frente se souber os padrões a serem aplicados.

Além disso, a utilização deles claramente traz economia de tempo de trabalho, pois você reaproveita um caminho que já foi deixado anteriormente para ser seguido por outras pessoas desenvolvedoras. Um conselho que costumamos dar e receber na área de exatas é para evitarmos de reinventar a roda e ele consegue se encaixar bem nessa situação, pois não faz sentido gastar tempo analisando e criando quando já existe um “guia” pronto para ser seguido.

## E o SOLID? Onde se encaixa?

Frequentemente as pessoas que iniciam os estudos de design patterns se deparam com esse termo. Mas afinal, do que se trata? Tem realmente a ver com padrões de projeto?

O SOLID é uma sigla em inglês para cinco princípios de projeto que possuem o objetivo de fazer programas mais compreensíveis, flexíveis e sustentáveis.

Não faz parte dos design patterns, mas por serem boas práticas que garantem diversos benefícios é bem comum de aprender ambos os temas ao mesmo tempo.

Veja abaixo uma breve definição destes princípios:

“S” Single Responsibility Principle (Princípio de responsabilidade única): uma classe deve ter uma e apenas uma razão para mudar.

“O” Open-Closed Principle (Princípio aberto/fechado): objetos devem estar disponíveis para extensão, mas fechados para modificação.

“L” Liskov Substitution Principle (Princípio de substituição de Liskov): uma subclasse deve ser substituível por sua superclasse.

“I” Interface Segregation Principle (Princípio de segregação de interface): uma classe não deve ser obrigada a implementar métodos e interfaces que não serão utilizadas.

“D” Dependency Inversion Principle (Princípio de inversão de dependência): dependa de abstrações e não de implementações.

## Padrões de projeto: quais são?

Após surfar por alguns conceitos importantes, finalmente vamos dar um mergulho nos principais tópicos de padrões de projeto.

Os três principais padrões de projeto definidos pelo livro “Design Patterns: Elements of Reusable Object-Oriented Software” de 1994, escrito por GOF (Gang of Four: Rich Gamma, Richard Helm, Ralph Johnson e John Vlissides) são os padrões criacionais, estruturais e comportamentais. Esses padrões foram divididos e agrupados de acordo com a natureza do problema que eles solucionam.

## Tipos de padrões

1) Padrões Criacionais

    Estes padrões oferecem diversas alternativas de criação de objetos, o que aumenta a flexibilidade e a reutilização de código.

    - Factory Method
    - Abstract Factory
    - Builder
    - Prototype
    - Singleton

2) Padrões Estruturais

    Nos mostram como montar objetos e classes em estruturas maiores, sem perder a eficiência e flexibilidade.

    - Adapter
    - Bridge
    - Composite
    - Decorator
    - Facade
    - Business Delegate
    - Flyweight
    - Proxy

3) Padrões comportamentais

    Nos ajudam a trabalhar melhor com os algoritmos e com a delegação de responsabilidades entre os objetos.

    - Chain of Responsibility
    - Command
    - Interpreter
    - Iterator
    - Mediator
    - Memento
    - Observer
    - State
    - Strategy
    - Template Method
    - Visitor

4) Padrōes Arquiteturais

    - Interceptor
    - Model View Controller
    - Model View ViewMode
    - Model View Presenter
    - Arquitetura Multicamada (n-tier)
    - Specification
    - Publish-Subscribe
    - Inversion of Control
    - Dependency Injection
    - Intercepting Filter
    - Lazy Loading
    - Mock Object
    - Method Chaining
    - Unit of Work
    - Delegation
    - Forwarding

## Conclusão

Conhecer Design Patterns é algo de extrema importância no desenvolvimento de qualquer software. A utilização desses padrões nos ajuda a desenvolver de forma mais rápida frente a desafios semelhantes, fornece uma linguagem comum durante a documentação e discussões técnicas além de nos auxiliar a organizar o código fonte do software que estamos desenvolvendo.