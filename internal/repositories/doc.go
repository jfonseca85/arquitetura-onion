//Driven adapter ( Fluxo de Saída do Core )
//Todos os Driven adapter serão colocados em pacotes dentro do diretório ./internal/repositories.
//São aqueles que esperam que o Core seja aquele que aciona a comunicação.
//Nesse caso, é o Core que precisa de algo que o ator fornece, portanto, ele envia uma solicitação ao ator e invoca uma ação específica sobre ela.
//Por exemplo, se o núcleo precisa salvar dados em um banco de dados MySQL,
//O núcleo aciona a comunicação para executar uma consulta INSERT no cliente MySQL.
package repositories
