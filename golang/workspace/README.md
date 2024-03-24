# GO Workspace

Nessa seção iremos compreender algumas variáveis muito importantes do `GO ENV`:

### **GOROOT**

Quando ouvirmos falar de `GOROOT` devemos automaticamente entender que é aonde está localizado o `Go SDK`.

Para configurarmos localmente, podemos seguir o seguinte tutorial:
-- INSERIR LINK DO TUTORIAL --

### **GOPATH**

> [!NOTE]
> Tópico apenas para conhecermos mais sobre o GOPATH.
> Utilizamos à partir da versão 1.22.

**Antes da versão 1.13**

A variável env `GOPATH` é usada para resolver instruções go imports e informa qual a localização do nosso espaço de trabalho (É a raiz do nosso espaço de trabalho).

> [!IMPORTANT]
> Uma dependência só pode ser importada se estiver dentro do GOPATH.

As versões anteriores a 1.13 exigiam que todo o nosso código-fonte estivesse dentro do `GOPATH`.

O que conseguimos compreender é que basicamente, o `GOPATH` é usado para:

- **Resolver importações**

    Uma dependência não poderá ser importada, se não estiver dentro do `GOPATH`. Portanto, ainda é necessário que seu código esteja dentro da `GOPATH`

- **Pacotes instalados**

    Os pacotes são instalados no diretório
    ```bash
    $GOPATH/pkg/$GOOS_$GOARCH
    ```

- **Binários**

    O aplicativo binário compilado era armazenado em:
    ```bash
    $GOPATH/bin
    ```

    A pasta bin, contém os seguintes sub pastas:

    - **src**

        Localização do arquivo de origem. Ao instalarmos qualquer dependência usando o `go get`, todos os arquivos de pacote são armazenados nesse local.
        É nele que contém os arquivos `.go`.
    
    - **pkg**

        Ele contém os pacotes compilados do diretório `src/`. Eles são usados no momento do link para criar os binários executáveis que ficam no diretório `bin/`.
        Era uma boa ideia compilar o pacote apenas uma vez e usá-lo para criar diferentes binários executáveis.

    - **bin**

        Localização dos binários executáveis.

**Depois da versão 1.13**

Na versão 1.13, a linguagem introduziu um novo recurso para gerenciamento de dependência chamado **módulos**.

Com esse novo recurso, não é mais necessário colocar todo o código dentro de um espaço de trabalho ou dentro do diretório `$GOPATH/src`. Agora é possível criarmos o diretório em qualquer lugar e colocar o nosso programa Go lá.

> [!NOTE]
> O comportamento legado do GOPATH ainda é valido para as novas versões.

Os programas podem ser executados de duas maneiras:

- **Usando módulos**

    Usando módulos, o `GOPATH` será usado para armazenar o nosso código-fonte no diretório:

    ```bash
    $GOPATH/pkg/mod
    ```

    E o binário compilado ainda será armazenado em:

    ```bash
    $GOPATH/bin
    ```

- **Não usando módulos**

    Não usando módulos o comportamento é o mesmo das versões anteriores a 1.13.

Para configurarmos localmente, podemos seguir o seguinte tutorial:
-- INSERIR LINK DO TUTORIAL --

### **GOBIN**

A variável env `GOBIN` nos informará o diretório onde o binário da aplicação será inserido após construir o pacote principal.
Por padrão ele é armazenado em:

```bash
$GOPATH/bin
```

Para configurarmos localmente, podemos seguir o seguinte tutorial:
-- INSERIR LINK DO TUTORIAL --
