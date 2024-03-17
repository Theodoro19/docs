# Abstract Factory

Imaginemos que temos uma montadora de automóveis.

Em nossa montadora é possível montar duas linhas de produto, Carros e Caminhões.

Para tornar isso possível e escalável em termos de código, vamos definir uma interface de "Vehicle" responsável por criar "Car" e "Truck".

```go
type Vehicle interface {
    WheelCount() int
    NumberOfDoors() int
    MaximumSpeed() int
}
```

Como é possível notar acima, os métodos `WheelCount`, `NumberOfDoor`, e `MaximumSpeed` são as funções que desejamos.

Agora vamos criar uma fábrica abstrata para grupos familiares de veículos.
Cada nova fábrica de grupo familiar deverá implementar essa classe.

```go
type VehicleFactory interface {
    NewVehicle(int) (Vehicle, error)
}
```

A principal ideia da Montadora, é a criação de Carros e Caminhões, estruturas que serão definidas abaixo. Cada uma das estruturas tem definições básicas porém também possuem propriedades específicas. Por conta disso, vamos criar sub famílias.

```go
type Car interface {
    Vehicle
    HasElectricEngine() bool
}
```

```go
type Truck interface {
    Vehicle
    NumberOfAxles() int
}
```

Agora vamos criar a fábrica onde será criado os veículos que são membros do grupo sub familiar de carros.

```go
import (
    "errors"
    "fmt"
)

const (
    HATCH_CAR = 1
    SEDAN_CAR = 2
)

type CarFactory struct{}

func (f *CarFactory) NewVehicle(carType int) (Vehicle, error){
    switch carType {
    case HATCH_CAR:
        return new(Hatch), nil
    case SEDAN_CAR:
        return new(Sedan), nil
    default:
        return nil, errors.New(fmt.Sprintf("Car type not found:%d", carType))
    }
}
```

Agora faremos o mesmo para os caminhões.

```go
import (
    "errors"
    "fmt"
)

const (
    SIMPLE_MECHANICAL_HORSE = 1
    ROAD_TRAIN = 2
)

type TruckFactory struct{}

func (f *TruckFactory) NewVehicle(truckType int) (Vehicle, error){
    switch truckType {
    case SIMPLE_MECHANICAL_HORSE:
        return new(SimpleMechanicalHorse), nil
    case ROAD_TRAIN:
        return new(RoadTrain),nil
    default:
        return nil, errors.New(fmt.Sprintf("Truck type not found:%d", truckType))
    }
}
```

Agora iremos criar os grupos da família carro que são os carros Hatch e Sedan.

```go
type Hatch struct{
    wheels int
    doors int
    speed int
    isElectric bool
}

func (car Hatch) WheelCount() int {
    return car.wheels
}

func (car Hatch) NumberOfDoors() int {
    return car.doors
}

func (car Hatch) MaximumSpeed() int {
    return car.speed
}

func (car Hatch) HasElectricEngine() bool {
    return car.isElectric
}
```

```go
type Sedan struct{
    wheels int
    doors int
    speed int
    isElectric bool
}

func (car Sedan) WheelCount() int {
    return car.wheels
}

func (car Sedan) NumberOfDoors() int {
    return car.doors
}

func (car Sedan) MaximumSpeed() int {
    return car.speed
}

func (car Sedan) HasElectricEngine() bool {
    return car.isElectric
}
```

Faremos o mesmo para o grupo da família caminhão que serão o Rodotrem e o Cavalo Mecânico Simples.

```go
type Road struct{
    wheels int
    doors int
    speed int
    axles bool
}

func (truck Road) WheelCount() int {
    return truck.wheels
}

func (truck Road) NumberOfDoors() int {
    return truck.doors
}

func (truck Road) MaximumSpeed() int {
    return truck.speed
}

func (truck Road) NumberOfAxles() int {
    return truck.axles
}
```

```go
type Mechanical struct{
    wheels int
    doors int
    speed int
    axles bool
}

func (truck Mechanical) WheelCount() int {
    return truck.wheels
}

func (truck Mechanical) NumberOfDoors() int {
    return truck.doors
}

func (truck Mechanical) MaximumSpeed() int {
    return truck.speed
}

func (truck Mechanical) NumberOfAxles() int {
    return truck.axles
}
```

Para finalizar, criaremos a classe main onde iremos ter uma função que cria fábricas de sub família da família Veículo.
O objetivo é dividir a complexidade da criação de famílias de objetos em pequenos blocos para facilitar a leitura e manutenção.
Basicamente delega criações de objetos de sub família para fábricas de sub família.

```go
import (
    "errors"
    "fmt"
)

const (
    TRUCK = 1
    CAR = 2
)

func CreateVehicleFactory(vehicleType int) (VehicleFactory, error) {
    switch vehicleType {
    case TRUCK:
        return new(TruckFactory), nil
    case CAR:
        return new(CarFactory), nil
    default:
        return nil, errors.New(fmt.Sprintf("Vehicle type not found:%d", truckType))
    }
}
```

Nesse link é possível encontrar o código de exemplo.