# Chain of Responsibility

O padrão de projeto `Chain of Responsibility` permite criar uma cadeia de manipuladores de solicitações. 

Para cada solicitação recebida, ela é passada pela cadeia e por cada manipulador:

1. Processa a solicitação ou ignora o processamento.
2. Decide se deve passar a solicitação para o próximo manipulador na cadeia, ou não.

### Quando usar

O padrão é aplicável quando há vários candidatos para processar a mesma solicitação.
Quando não desejamos que o cliente escolha o receptor, pois vários objetos podem lidar com a solicitação.

Além disso, desejamos dissociar o cliente dos receptores. O cliente só precisa conhecer o primeiro elemento da cadeia.

## Exemplo

Suponhamos que somos diretores de um hospital e precisamos controlar todos os seus departamentos, como:

- Recepção
- Doutores
- Salas de Atendimento
- Caixa

Sempre que chega um novo paciente, primeiramente ele vai para a recepção se cadastrar.
Depois é encaminhado para um médico para fazer a triagem e em seguida, é encaminhado para uma sala de atendimento.
Para finalizar o atendimento, ele precisa passar no caixa para pagar o valor de sua consulta.

## Código Completo

```go
package main

import "fmt"

type department interface {
    execute(*patient)
    setNext(department)
}

type reception struct {
    next department
}

func (r *reception) execute(p *patient) {
    if p.registrationDone {
        fmt.Println("Patient registration already done")
        r.next.execute(p)
        return
    }
    fmt.Println("Reception registering patient")
    p.registrationDone = true
    r.next.execute(p)
}

func (r *reception) setNext(next department) {
    r.next = next
}

type doctor struct {
    next department
}

func (d *doctor) execute(p *patient) {
    if p.doctorCheckUpDone {
        fmt.Println("Doctor checkup already done")
        d.next.execute(p)
        return
    }
    fmt.Println("Doctor checking patient")
    p.doctorCheckUpDone = true
    d.next.execute(p)
}

func (d *doctor) setNext(next department) {
    d.next = next
}

type medical struct {
    next department
}

func (m *medical) execute(p *patient) {
    if p.medicineDone {
        fmt.Println("Medicine already given to patient")
        m.next.execute(p)
        return
    }
    fmt.Println("Medical giving medicine to patient")
    p.medicineDone = true
    m.next.execute(p)
}

func (m *medical) setNext(next department) {
    m.next = next
}

type cashier struct {
    next department
}

func (c *cashier) execute(p *patient) {
    if p.paymentDone {
        fmt.Println("Payment Done")
    }
    fmt.Println("Cashier getting money from patient patient")
}

func (c *cashier) setNext(next department) {
    c.next = next
}

type patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
    paymentDone       bool
}

func main() {
    cashier := &cashier{}
   
    //Set next for medical department
    medical := &medical{}
    medical.setNext(cashier)
   
    //Set next for doctor department
    doctor := &doctor{}
    doctor.setNext(medical)
   
    //Set next for reception department
    reception := &reception{}
    reception.setNext(doctor)
   
    patient := &patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}
```

## Resultado

```text
Reception registering patient
Doctor checking patient
Medical giving medicine to patient
Cashier getting money from patient patient
```
