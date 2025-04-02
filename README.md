Aplicação Web em Go
Este projeto é uma aplicação web desenvolvida do zero utilizando a linguagem Go, seguindo as principais convenções da linguagem e boas práticas de desenvolvimento.

Funcionalidades
Servidor Web: Inicialização e configuração de um servidor web em Go para atender requisições HTTP.

Manipulação de Dados com Structs: Criação e utilização de structs para modelar produtos e outras entidades.

Integração com Banco de Dados: Conexão com PostgreSQL para armazenamento, consulta, atualização e deleção de registros.

Modularização e Partials: Organização do código em módulos e utilização de partials para evitar duplicação de código nas views.

Gerenciamento de Rotas: Configuração centralizada das rotas da aplicação no arquivo routes.go.

Conteúdo do Projeto
O desenvolvimento foi dividido em módulos que abordam:

Servidor e Manipulação de Structs
Configuração do ambiente e organização dos pacotes.

Inicialização do servidor e definição de rotas básicas.

Criação de structs para modelagem dos produtos e outras entidades.

Conexão com Banco de Dados
Instalação e configuração do PostgreSQL.

Implementação da conexão entre a aplicação Go e o banco de dados.

Operações de leitura e exibição dos dados armazenados.

Refatoração e Modularização
Divisão do código em módulos para melhorar a organização e a manutenção.

Criação de páginas para cadastro de novos produtos, com formulários e validações.

Implementação de partials para reutilização de componentes na interface.

Deleção e Atualização de Registros
Funcionalidades para deletar registros, com feedback utilizando JavaScript.

Tela e lógica para atualização e edição dos produtos, com pré-carregamento dos dados.

Gerenciamento das rotas responsáveis pelas operações de edição, atualização e deleção.

Tecnologias Utilizadas
Go (Golang)

PostgreSQL

Bootstrap (para estilização)

HTML, CSS e JavaScript (para a interface)

Como Executar o Projeto
Pré-requisitos:

Go instalado na máquina

PostgreSQL instalado e configurado

Git para clonar o repositório

Clonando o Repositório:

bash
Copiar
Editar
git clone https://github.com/GuiCMartini/GuiLoja.git
cd GuiLoja
Configuração:
Configure as variáveis de ambiente necessárias para a conexão com o banco de dados (por exemplo, utilizando um arquivo .env ou definindo as variáveis no sistema).

Executando a Aplicação:
No diretório raiz do projeto, execute:

bash
Copiar
Editar
go run main.go
A aplicação ficará disponível em http://localhost:8080 (ou conforme a porta configurada).

Contribuições
Contribuições são bem-vindas! Caso deseje sugerir melhorias ou corrigir algum problema, sinta-se à vontade para fazer um fork do projeto, criar sua branch com as alterações e enviar um pull request.

Licença
Este projeto está licenciado sob a MIT License.

Desenvolvido com dedicação e seguindo as melhores práticas de Go.
