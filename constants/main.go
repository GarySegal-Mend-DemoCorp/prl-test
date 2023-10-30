package constants

const (
	SqlNoRows = "sql: no rows in result set"
)

const (
	API_PREFIX = "/api/v1"
)

const (
	HMAC_SECRET_ENV_VAR       = "HMAC_SECRET"
	AUTHORIZATION_CONTEXT_KEY = "AUTHORIZATION_CONTEXT"
	TOKEN_DURATION_MINUTES    = 60
)

const (
	USER_ROLE       = "USER"
	SUPER_USER_ROLE = "SUPER_USER"
	READ_ONLY_CLAIM = "READ_ONLY"

	CREATE_CLAIM = "CREATE"
	DELETE_CLAIM = "DELETE"
	UPDATE_CLAIM = "UPDATE"
	LIST_CLAIM   = "LIST"

	LIST_USER   = "LIST_USER"
	ADD_USER    = "ADD_USER"
	DELETE_USER = "DELETE_USER"
	UPDATE_USER = "UPDATE_USER"

	CREATE_VM_CLAIM          = "CREATE_VM"
	DELETE_VM_CLAIM          = "DELETE_VM"
	LIST_VM_CLAIM            = "LIST_VM"
	UPDATE_VM_STATES_CLAIM   = "UPDATE_VM_STATES"
	UPDATE_VM_CLAIM          = "UPDATE_VM"
	EXECUTE_COMMAND_VM_CLAIM = "EXECUTE_COMMAND_VM"

	CREATE_TEMPLATE_CLAIM = "CREATE_TEMPLATE"
	DELETE_TEMPLATE_CLAIM = "DELETE_TEMPLATE"
	LIST_TEMPLATE_CLAIM   = "LIST_TEMPLATE"
	UPDATE_TEMPLATE_CLAIM = "UPDATE_TEMPLATE"

	CREATE_CATALOG_MANIFEST_CLAIM = "CREATE_CATALOG_MANIFEST"
	DELETE_CATALOG_MANIFEST_CLAIM = "DELETE_CATALOG_MANIFEST"
	LIST_CATALOG_MANIFEST_CLAIM   = "LIST_CATALOG_MANIFEST"
	UPDATE_CATALOG_MANIFEST_CLAIM = "UPDATE_CATALOG_MANIFEST"
	PULL_CATALOG_MANIFEST_CLAIM   = "PULL_CATALOG_MANIFEST"
	PUSH_CATALOG_MANIFEST_CLAIM   = "PUSH_CATALOG_MANIFEST"
)
