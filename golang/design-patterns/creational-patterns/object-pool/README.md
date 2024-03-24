# Object Pool

### Introdução

O `Object Pool` é um padrão de projeto no qual um pool de objetos são inicializados e criados antecipadamente e mantidos em um pool.
Conforme, e quando necessário, um cliente pode solicitar um objeto do pool, usá-lo e devolvê-lo ao pool.
O objeto pool nunca é destruído.

### Quando usar

1. Quando o custo para criar o objeto da classe é alto e o número desses objetos que serão necessários em um determinado momento não é muito.

Vamos pegar o exemplo das conexões de um banco de dados.

Cada criação de objeto de conexão tem um custo alto, pois há chamadas de rede envolvidas e também por vez, mais do que uma determinada conexão pode ser necessária.

Esse padrão de projeto é perfeitamente adequado para esses casos.

2. Quando o objeto de pool é imutável.

Pegando o mesmo exemplo acima das conexões de um banco de dados.

Uma conexão de banco de dados é um objeto imutável. Quase nenhuma de suas propriedades precisam ser alteradas.

3. Por motivos de desempenho. Isso aumentará significativamente o desempenho de nossa aplicação, uma vez que o pool já está criado.

