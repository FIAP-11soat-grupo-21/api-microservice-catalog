# API Microservice Catalog

## Estrutura do Projeto

Este repositório contém dois principais diretórios:

- **infra/**: scripts e módulos de infraestrutura como código (IaC) para provisionamento de recursos na AWS, incluindo banco de dados, rede, ECS, ALB, entre outros.
- **microservice/**: código-fonte do microsserviço de catálogo, responsável pelo gerenciamento de produtos e categorias.

---

## infra/

A pasta `infra` utiliza Terraform para provisionar toda a infraestrutura necessária na AWS. Entre os recursos provisionados estão:

- **Banco de Dados (RDS/Postgres):**
  - O banco de dados é criado via módulo e seu endpoint é passado como variável de ambiente para o container do microsserviço no ECS.
  - O RDS garante alta disponibilidade, backups automáticos e segurança de dados.
- **ECS (Elastic Container Service):**
  - O microsserviço é executado em containers gerenciados pelo ECS, garantindo escalabilidade e fácil deploy.
- **ALB (Application Load Balancer):**
  - Gerencia o tráfego HTTP/HTTPS para o serviço, incluindo health checks.
- **Variáveis de ambiente:**
  - O endpoint do banco de dados, usuário, senha e outras configurações são passadas via variáveis de ambiente para o container.

---

## microservice/

O diretório `microservice` contém o código do microsserviço de catálogo, desenvolvido em Go, seguindo arquitetura limpa (clean architecture).

### Banco de Dados
- Utiliza Postgres, tanto localmente (via Docker) quanto na AWS (RDS).
- As configurações de conexão (host, porta, usuário, senha, nome do banco) são lidas de variáveis de ambiente, permitindo fácil portabilidade entre ambientes.

### Modelagem das Tabelas

#### Categoria
- `id` (varchar(36), PK)
- `name` varchar(100)
- `active` bool

#### Produto
- `id` (varchar(36), PK)
- `category_id` (varchar(36), FK para Categoria)
- `name` (varchar(100))
- `description` (text)
- `price` (numeric)
- `active` (bool)
- `created_at` (timestamptz)

#### Imagens do Produto
- `id` (varchar(36), PK)
- `product_id` (varchar(36), FK para Produto)
- `file_name` (text)
- `url` (text)
- `is_default` (bool)
- `created_at` (timestamptz)

## Diagrama de Entidade-Relacionamento (Mermaid)

```mermaid
erDiagram
  categories {
    id varchar(36) PK
    name varchar(100)
    active bool
  }
  products {
    id varchar(36) PK
    category_id varchar(36) FK 
    name varchar(100)
    description text
    price numeric
    active bool
    created_at timestamptz
  }
  product_images  {
    id varchar(36) PK
    product_id varchar(36) FK
    file_name text
    url text
    price numeric
    is_default bool
    created_at timestamptz
  }
  categories ||--o{ products : "possui"
  products ||--o{ product_images : "tem"
```

### Justificativa para Modelagem Relacional
- **Relacionamento entre produtos e categorias:** Cada produto pertence a uma categoria, o que exige integridade referencial e facilita consultas relacionais (ex: listar todos os produtos de uma categoria).
- **Transações e consistência:** O modelo relacional permite transações ACID, garantindo consistência dos dados em operações críticas.
- **Facilidade de manutenção:** A modelagem relacional facilita alterações futuras, como adição de novos relacionamentos ou entidades.
- **Validação de integridade:** O uso de chaves estrangeiras impede a existência de produtos sem categoria válida.

### Outros Pontos
- O microsserviço está preparado para rodar tanto localmente quanto na AWS, bastando ajustar as variáveis de ambiente.
- O código segue boas práticas de separação de camadas, facilitando manutenção e testes.

---

## Rotas Disponíveis

### Categorias
| Método | Rota                        | Descrição                    |
|--------|-----------------------------|------------------------------|
| POST   | /v1/categories              | Cadastrar nova categoria     |
| GET    | /v1/categories              | Listar todas as categorias   |
| GET    | /v1/categories/:id          | Buscar categoria por ID      |
| PUT    | /v1/categories/:id          | Atualizar categoria          |
| DELETE | /v1/categories/:id          | Remover categoria            |

### Produtos
| Método | Rota                                         | Descrição                        |
|--------|----------------------------------------------|----------------------------------|
| POST   | /v1/products                                | Cadastrar novo produto           |
| GET    | /v1/products                                | Listar todos os produtos         |
| GET    | /v1/products?category_id={id}               | Listar produtos por categoria    |
| GET    | /v1/products/:id                            | Buscar produto por ID            |
| PUT    | /v1/products/:id                            | Atualizar produto                |
| DELETE | /v1/products/:id                            | Remover produto                  |
| PATCH  | /v1/products/:id/images                     | Adicionar imagem ao produto      |
| DELETE | /v1/products/:id/images/:image_file_name     | Remover imagem do produto        |

### Outros
| Método | Rota                 | Descrição                  |
|--------|----------------------|----------------------------|
| POST   | /v1/uploads          | Upload direto de arquivo   |
| GET    | /health              | Health check               |
| GET    | /v1/health           | Health check v1            |
| GET    | /swagger/index.html  | Documentação Swagger       |

---

## Checklist de Testes das Rotas da API

## Categorias
| Rota                                      | Método | Testado? | Observações                       |
|-------------------------------------------|--------|----------|-----------------------------------|
| /v1/categories                           | POST   | [x]      | Cadastrar nova categoria          |
| /v1/categories                           | GET    | [x]      | Listar todas as categorias        |
| /v1/categories/:id                       | GET    | [x]      | Buscar categoria por ID           |
| /v1/categories/:id                       | PUT    | [x]      | Atualizar categoria               |
| /v1/categories/:id                       | DELETE | [x]      | Remover categoria                 |

## Produtos
| Rota                                      | Método | Testado? | Observações                       |
|-------------------------------------------|--------|----------|-----------------------------------|
| /v1/products                             | POST   | [ ]      | Cadastrar novo produto            |
| /v1/products                             | GET    | [ ]      | Listar todos os produtos          |
| /v1/products?category_id={id}            | GET    | [ ]      | Listar produtos por categoria     |
| /v1/products/:id                         | GET    | [ ]      | Buscar produto por ID             |
| /v1/products/:id                         | PUT    | [ ]      | Atualizar produto                 |
| /v1/products/:id                         | DELETE | [ ]      | Remover produto                   |
| /v1/products/:id/images                  | PATCH  | [ ]      | Adicionar imagem ao produto       |
| /v1/products/:id/images/:image_file_name | DELETE | [ ]      | Remover imagem do produto         |

## Outros
| Rota                                      | Método | Testado? | Observações                       |
|-------------------------------------------|--------|----------|-----------------------------------|
| /v1/uploads                              | POST   | [ ]      | Testar upload direto              |
| /health                                  | GET    | [ ]      |                                   |
| /v1/health                               | GET    | [ ]      |                                   |
| /swagger/index.html                      | GET    | [ ]      | Verificar documentação            |

---

**Observações Gerais:**
- Teste com dados válidos e inválidos.
- Verifique respostas de erro (400, 404, 500).
- Teste upload e remoção de imagens.
- Teste filtro por categoria.
- Anote problemas, melhorias ou comportamentos inesperados.

---

## Como rodar localmente

1. Suba os serviços necessários com Docker Compose:
   ```sh
   docker compose up --build
   ```
2. Exporte as variáveis de ambiente conforme o exemplo do `.env`.
3. Execute o microsserviço normalmente.

---

## Imagem Default para Produtos (Minio)

Se estiver rodando localmente com Minio, é necessário subir manualmente a imagem default do produto para o bucket. Essa imagem é usada quando um produto não possui imagem cadastrada.

1. Certifique-se de que o serviço Minio está rodando.
2. Acesse o painel do Minio (geralmente em http://localhost:9000) ou utilize o client Minio.
3. Utilize o usuário e senha definido no seu arquivo .env nas variáveis MINIO_ROOT_USER e MINIO_ROOT_PASSWORD para acessar o Minio
2. Crie o bucket com o nome definido no seu arquivo .env na variável **AWS_S3_BUCKET_NAME**

3. Faça upload do arquivo que está na pasta `uploads/default_product_image.webp` para o bucket que foi criado anteriormente.

Assim, sempre que um produto não tiver imagem, o sistema irá retornar a imagem default.

---

## Observações
- Para produção, garanta que as variáveis de ambiente estejam corretas e nunca sobrescreva o host do banco para `localhost` no código.
- O seed de dados deve ser executado apenas em ambiente controlado, pois insere dados iniciais no banco.
