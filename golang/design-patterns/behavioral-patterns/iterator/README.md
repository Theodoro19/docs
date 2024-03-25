# Iterator

Nesse padrão de projeto, `Iterator`, a estrutura da coleção fornece um iterador que permite percorrer cada elemento da estrutura da coleção em sequência, sem expor sua implementação subjacente.

Os componentes básicos usados nesse padrão são:

- **Interface do Iterador** - Esta interface fornece as operações básicas.

- **Interface de Coleção** - Essa interface representa a coleção que precisa ser percorrida.

- **Iterador Concreto** - É a implementação concreta da Interface do Iterador.

- **Coleção Concreta** - É a implementação concreta da Interface da Coleção.

A ideia principal por trás desse padrão é expor a lógica da iteração de uma estrutura coleção em um objeto diferente (que implemente a interface do iterador). Esse iterador fornece um método genérico de iteração sobre uma coleção independente de seu tipo.

## Estrutura de Código

[collection.go](main/collection.go) - Coleção
[iterator.go](main/iterator.go) - Iterador
[main.go](main/main.go) - Cliente
[userCollection.go](main/userCollection.go) - Coleção Concreta
[userIterator.go](main/userIterator.go) - Iterador Concreto

## Código Completo

```go
package main

import "fmt"

type collection interface {
    createIterator() iterator
}

type userCollection struct {
    users []*user
}

func (u *userCollection) createIterator() iterator {
    return &userIterator{
        users: u.users,
    }
}

type iterator interface {
    hasNext() bool
    getNext() *user
}

type userIterator struct {
    index int
    users []*user
}

func (u *userIterator) hasNext() bool {
    if u.index < len(u.users) {
        return true
    }
    return false
}

func (u *userIterator) getNext() *user {
    if u.hasNext() {
        user := u.users[u.index]
        u.index++
        return user
    }
    return nil
}

type user struct {
    name string
    age  int
}

func main() {
    user1 := &user{
        name: "a",
        age:  30,
    }
    user2 := &user{
        name: "b",
        age:  20,
    }
    userCollection := &userCollection{
        users: []*user{user1, user2},
    }
    iterator := userCollection.createIterator()
    for iterator.hasNext() {
        user := iterator.getNext()
        fmt.Printf("User is %+v\n", user)
    }
}
```

## Resultado

```text
User is &{name:a age:30}
User is &{name:b age:20}
```
