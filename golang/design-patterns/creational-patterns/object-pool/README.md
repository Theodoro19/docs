# Object Pool

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


### Código Completo

```go
package main

import (
    "fmt"
    "log"
    "strconv"
    "sync"
)

type iPoolObject interface {
    getID() string //This is any id which can be used to compare two different pool objects
}

type pool struct {
    idle   []iPoolObject
    active []iPoolObject
    capacity int
    mulock   *sync.Mutex
}

//InitPool Initialize the pool
func initPool(poolObjects []iPoolObject) (*pool, error) {
    if len(poolObjects) == 0 {
        return nil, fmt.Errorf("Cannot craete a pool of 0 length")
    }
    active := make([]iPoolObject, 0)
    pool := &pool{
        idle:     poolObjects,
        active:   active,
        capacity: len(poolObjects),
        mulock:   new(sync.Mutex),
    }
    return pool, nil
}

func (p *pool) loan() (iPoolObject, error) {
    p.mulock.Lock()
    defer p.mulock.Unlock()
    if len(p.idle) == 0 {
        return nil, fmt.Errorf("No pool object free. Please request after sometime")
    }
    obj := p.idle[0]
    p.idle = p.idle[1:]
    p.active = append(p.active, obj)
    fmt.Printf("Loan Pool Object with ID: %s\n", obj.getID())
    return obj, nil
}

func (p *pool) receive(target iPoolObject) error {
    p.mulock.Lock()
    defer p.mulock.Unlock()
    err := p.remove(target)
    if err != nil {
        return err
    }
    p.idle = append(p.idle, target)
    fmt.Printf("Return Pool Object with ID: %s\n", target.getID())
    return nil
}

func (p *pool) remove(target iPoolObject) error {
    currentActiveLength := len(p.active)
    for i, obj := range p.active {
        if obj.getID() == target.getID() {
            p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
            p.active = p.active[:currentActiveLength-1]
            return nil
        }
    }
    return fmt.Errorf("Targe pool object doesn't belong to the pool")
}

type connection struct {
    id string
}

func (c *connection) getID() string {
    return c.id
}

func main() {
    connections := make([]iPoolObject, 0)
    for i := 0; i < 3; i++ {
        c := &connection{id: strconv.Itoa(i)}
        connections = append(connections, c)
    }
    pool, err := initPool(connections)
    if err != nil {
        log.Fatalf("Init Pool Error: %s", err)
    }
    conn1, err := pool.loan()
    if err != nil {
        log.Fatalf("Pool Loan Error: %s", err)
    }
    conn2, err := pool.loan()
    if err != nil {
        log.Fatalf("Pool Loan Error: %s", err)
    }
    pool.receive(conn1)
    pool.receive(conn2)
}
```

### Resultado

```text
Loan Pool Object with ID: 0
Loan Pool Object with ID: 1
Return Pool Object with ID: 0
Return Pool Object with ID: 1
```
