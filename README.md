# Servidor de Mensagens em Go

Este é um exemplo de implementação em Go de um servidor que recebe várias mensagens de vários clientes simultaneamente e exibe as mensagens na saída padrão (tela) do servidor.

## Funcionalidades

O servidor é capaz de:

- Receber múltiplas conexões de clientes (usando threads/goroutines).
- Exibir as mensagens recebidas de cada cliente na tela do servidor.

O cliente é capaz de:

- Conectar-se ao servidor.
- Enviar quantas mensagens quiser.
- Encerrar a conexão e o programa cliente digitando a mensagem "sair".

## Requisitos

- Go (versão 1.16 ou superior)

## Executando o Servidor

1. Abra um terminal e navegue até o diretório do projeto.
2. Execute o seguinte comando para compilar e executar o servidor:

   ```shell
   go run servidor.go
   ```

O servidor iniciará e ficará aguardando conexões na porta 8080.

## Executando o Cliente

1. Abra outro terminal e navegue até o diretório do projeto.
2. Execute o seguinte comando para compilar e executar o cliente:

   ```shell
   go run cliente.go
   ```

   O cliente se conectará ao servidor na porta 8080 do localhost.

   Digite uma mensagem e pressione Enter para enviá-la ao servidor.

   Para encerrar a conexão e o programa cliente, digite a mensagem "sair" e pressione Enter.

## Exemplo de Uso

1. Execute o servidor em um terminal.
2. Execute dois ou mais clientes em terminais diferentes.
3. Digite mensagens nos clientes e veja as mensagens sendo exibidas no console do servidor.

## Contribuição

Contribuições são bem-vindas! Se você encontrar algum problema, tiver alguma sugestão ou quiser adicionar novos recursos, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
```

Você pode salvar esse conteúdo em um arquivo chamado `README.md` na raiz do seu projeto. Lembre-se de incluir o arquivo `LICENSE` caso queira utilizar a licença MIT mencionada no `README.md`.