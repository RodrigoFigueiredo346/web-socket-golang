package errors

// Códigos de erro relacionados a credenciais
const (
	NoError                = 0     // sucesso
	InvalidCredentials     = -101  // Credenciais inválidas
	InvalidEmailFormat     = -1010 // Formato de email inválido
	PasswordTooShort       = -1011 // Senha muito curta
	PasswordMissingSpecial = -1012 // Falta caractere especial na senha
	PasswordMissingNumber  = -1013 // Falta número na senha
	PasswordMissingUpper   = -1014 // Falta letra maiúscula na senha
	PasswordMissingLower   = -1015 // Falta letra minúscula na senha
)

// Códigos de erro relacionados a usuário
const (
	UserLocked     = -102 // Usuário bloqueado
	UserNotFound   = -109 // Usuário não encontrado
	UserInactive   = -503 // Usuário inativo
	DuplicateEntry = -504 // Entrada duplicada
)

// Códigos de erro relacionados a sessão
const (
	InvalidToken    = -103 // Token inválido
	SessionNotFound = -104 // Sessão não encontrada
)

// Códigos de erro relacionados a painel
const (
	PanelNotFound      = -105 // Painel não encontrado
	PanelAlreadyExists = -108 // Painel já existe
)

// Códigos de erro gerais
const (
	PermissionDenied        = -106 // Permissão negada
	InvalidParameters       = -107 // Parâmetros inválidos
	EmptyParameters         = -108 // Parâmetros vazios
	InternalServerError     = -500 // Erro interno do servidor
	InvalidRequestFormat    = -501 // Formato de requisição inválido
	RequiredFieldsMissing   = -502 // Campos obrigatórios ausentes
	DatabaseConnectionError = -505 // Erro de conexão com o banco de dados
	UnauthorizedAction      = -506 // Ação não autorizada
)
