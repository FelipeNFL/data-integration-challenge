# data-integration-challenge

Para executar a solução, é necessário ter instalado o Docker e Docker-Compose (não estará disponível no Makefile para não conflitar com versões previamente instaladas).

As APIs serão executadas nas portas 8080 e 8081, e as documentações na 8082 e 8083. Certifique-se que elas não estejam sendo usadas por outros processos.

O comando para iniciar as APIs é `make start` e para executar os testes `make check`. É importante ressaltar que os testes fazem manipulações no banco de dados. Caso você execute os testes, execute o start novamente para que os dados sejam resetados.

## Documentação das APIs

A documentação das APIs pode ser consultada em:

http://localhost:8082 - Integration API
http://localhost:8083 - Matching API