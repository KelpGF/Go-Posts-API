# Go-Posts-API

Go http server about Posts.

## Topics

- HTTP
- REST
- CRUD
- Chi
- Graceful shutdown
- Panic recovery
- Clean Code
- Unit tests
- Clean Architecture
- Design Patterns
- Docker

## Commands

### Development

É iniciado um container com o banco de dados PostgreSQL e outro com a aplicação Go.

O entrypoint do app é `tail -f /dev/null` para manter o container rodando. Dessa forma consigo acessar o container e rodar os comandos necessários.

```bash
sh scripts/start-dev.sh
```

```bash
sh scripts/stop-dev.sh
```

### Production

```bash
sh scripts/start.sh
```

```bash
sh scripts/stop.sh
```

## Considerações

- Utilizei o Gorm para facilitar a manipulação do banco de dados. Ele possui uma API muito boa, é bem documentado e também tem algumas funcionalidades padrões como Soft Delete e facilidade com Transactions.
- Utilizei o Chi para facilitar a criação de rotas. Ele é um roteador HTTP leve e rápido.
- Utilizei o Testify para facilitar a criação de testes. Ele é uma biblioteca que fornece algumas funções auxiliares para testes.
- Não utilizei o Viper para configurações, pois não achei necessário para esse projeto. Utilizei variáveis de ambiente para configurações.
- Normalmente crio uma camada entre o caso de uso e o handler http. Essa camada são as controller que possui assinaturas de funções únicas com tipos únicos HttpRequest e HttpResponse. Isso facilita a criação de testes unitários e também facilita caso seja necessário trocar o framework HTTP. Nesse projeto não criei essa camada, pois achei que não era necessário porém gostaria de informar essa prática.
- Também crio uma camada chamada DAO entre a Repository e a instância direta do banco de dados. Mas também não achei necessário para esse projeto. Crio essa camada pois em alguns projetos os objetos de domínio tem estruturas diferentes das tabelas do banco de dados. Nesse caso, a camada Repository é responsável por converter os objetos de domínio um formato melhor para o banco de dados (e também o contrário) e a camada DAO é responsável por acessar o banco de dados.
- Os mocks dos testes estão no mesmo arquivo dos testes. Normalmente eu crio uma pasta chamada mocks e coloco os mocks lá. Mas nesse projeto não achei necessário pois os mocks são utilizados apenas nos testes, exceto pelo mock de Post.
