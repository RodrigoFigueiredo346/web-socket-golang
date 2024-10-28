package errors

const (
	NoError = 0 // sucesso
)

// Códigos de erro relacionados a credenciais range -101
const (
	InvalidCredentials     = -1010 // Credenciais inválidas
	InvalidEmailFormat     = -1011 // Formato de email inválido
	PasswordTooShort       = -1012 // Senha muito curta
	PasswordMissingSpecial = -1013 // Falta caractere especial na senha
	PasswordMissingNumber  = -1014 // Falta número na senha
	PasswordMissingUpper   = -1015 // Falta letra maiúscula na senha
	PasswordMissingLower   = -1016 // Falta letra minúscula na senha
	PasswordNotMatch       = -1017 // Password incorreto
)

// Códigos de erro relacionados a usuário range -102
const (
	UserLocked     = -1020 // Usuário bloqueado
	UserNotFound   = -1021 // Usuário não encontrado
	UserInactive   = -1022 // Usuário inativo
	DuplicateEntry = -1023 // Entrada duplicada
)

// Códigos de erro relacionados a sessão range -103
const (
	InvalidToken    = -1030 // Token inválido
	SessionNotFound = -1031 // Sessão não encontrada
)

// Códigos de erro relacionados a painel range -104
const (
	PanelNotFound      = -1040 // Painel não encontrado
	PanelAlreadyExists = -1041 // Painel já existe
)

// Códigos de erro gerais range -50
const (
	PermissionDenied        = -501 // Permissão negada
	InvalidParameters       = -502 // Parâmetros inválidos
	EmptyParameters         = -503 // Parâmetros vazios
	InvalidMethod           = -504 // Método inválido
	InternalServerError     = -505 // Erro interno do servidor
	InvalidRequestFormat    = -506 // Formato de requisição inválido
	RequiredFieldsMissing   = -507 // Campos obrigatórios ausentes
	DatabaseConnectionError = -508 // Erro de conexão com o banco de dados
	UnauthorizedAction      = -509 // Ação não autorizada
)
