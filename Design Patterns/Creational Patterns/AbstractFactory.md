# Abstract Factory

É um padrão que resolve o problema na criação de famílias inteiras de produtos sem especificar suas classes concretas.

O Abstract Factory define uma Interface para criar os produtos distintos e deixa a criação do produto real para as classes de fabricas concretas. Cada tipo de fabrica corresponde a uma determinada variedade de produto.

O código do cliente chama os métodos de criação de um objeto de fábrica em vez de criar produtos diretamente com uma chamada de construtor (novo operador). Como uma fábrica corresponde a uma única variante de produto, todos os seus produtos serão compatíveis.

O código do cliente funciona com fábricas e produtos apenas por meio de suas interfaces abstratas. Isso permite que o código do cliente funcione com quaisquer variantes de produto criadas pelo objeto fábrica.
