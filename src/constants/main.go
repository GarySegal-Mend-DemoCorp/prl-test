package constants

type AuthorizationContextKey string

var (
	Name                                              = "Parallels Desktop DevOps Service"
	ExecutableName                                    = "prldevops"
	ServiceDefaultDirectory                           = "/etc/prl-devops-service"
	AUTHORIZATION_CONTEXT_KEY AuthorizationContextKey = "AUTHORIZATION_CONTEXT"
)

type RequestIdKey string

var REQUEST_ID_KEY RequestIdKey = "REQUEST_ID"

const (
	SqlNoRows = "sql: no rows in result set"
)

const (
	DELETE_REMOTE_MANIFEST_QUERY = "clean_remote"
)

const (
	LATEST_TAG = "latest"
)

const (
	DEFAULT_API_PREFIX                      = "/api"
	DEFAULT_API_PORT                        = "80"
	DEFAULT_API_TLS_PORT                    = "443"
	CURRENT_USER_ENV_VAR                    = "PD_CURRENT_USER"
	DEFAULT_TOKEN_DURATION_MINUTES          = 60
	DEFAULT_CATALOG_CACHE_FOLDER            = "./catalog_cache"
	DEFAULT_ORCHESTRATOR_PULL_FREQUENCY_SEC = 30
	SOURCE_ENV_VAR                          = "DEVOPS_SOURCE"
	LOCAL_ORCHESTRATOR_DESCRIPTION          = "Local Orchestrator"
	DEFAULT_SYSTEM_RESERVED_CPU             = 1
	DEFAULT_SYSTEM_RESERVED_MEMORY          = 2048
	DEFAULT_SYSTEM_RESERVED_DISK            = 20000

	API_MODE          = "api"
	CLI_MODE          = "cli"
	ORCHESTRATOR_MODE = "orchestrator"
	CATALOG_MODE      = "catalog"
)

const (
	API_PORT_ENV_VAR                            = "API_PORT"
	API_PREFIX_ENV_VAR                          = "API_PREFIX"
	HMAC_SECRET_ENV_VAR                         = "HMAC_SECRET"
	LOG_LEVEL_ENV_VAR                           = "LOG_LEVEL"
	ENCRYPTION_SECURITY_KEY_ENV_VAR             = "ENCRYPTION_PRIVATE_KEY"
	TLS_ENABLED_ENV_VAR                         = "TLS_ENABLED"
	TLS_PORT_ENV_VAR                            = "TLS_PORT"
	TLS_CERTIFICATE_ENV_VAR                     = "TLS_CERTIFICATE"
	TLS_PRIVATE_KEY_ENV_VAR                     = "TLS_PRIVATE_KEY"
	ROOT_PASSWORD_ENV_VAR                       = "ROOT_PASSWORD"
	DISABLE_CATALOG_CACHING_ENV_VAR             = "DISABLE_CATALOG_CACHING"
	TOKEN_DURATION_MINUTES_ENV_VAR              = "TOKEN_DURATION_MINUTES"
	MODE_ENV_VAR                                = "MODE"
	USE_ORCHESTRATOR_RESOURCES_ENV_VAR          = "USE_ORCHESTRATOR_RESOURCES"
	ORCHESTRATOR_PULL_FREQUENCY_SECONDS_ENV_VAR = "ORCHESTRATOR_PULL_FREQUENCY_SECONDS"
	DATABASE_FOLDER_ENV_VAR                     = "DATABASE_FOLDER"
	DATABASE_NUMBER_BACKUP_FILES_ENV_VAR        = "DATABASE_NUMBER_BACKUP_FILES"
	DATABASE_BACKUP_INTERVAL_ENV_VAR            = "DATABASE_BACKUP_INTERVAL_MINUTES"
	DATABASE_SAVE_INTERVAL_ENV_VAR              = "DATABASE_SAVE_INTERVAL_MINUTES"
	CATALOG_CACHE_FOLDER_ENV_VAR                = "CATALOG_CACHE_FOLDER"
	CORS_ALLOWED_HEADERS_ENV_VAR                = "CORS_ALLOWED_HEADERS"
	CORS_ALLOWED_METHODS_ENV_VAR                = "CORS_ALLOWED_METHODS"
	CORS_ALLOWED_ORIGINS_ENV_VAR                = "CORS_ALLOWED_ORIGINS"
	ENABLE_PACKER_PLUGIN_ENV_VAR                = "ENABLE_PACKER_PLUGIN"
	ENABLE_VAGRANT_PLUGIN_ENV_VAR               = "ENABLE_VAGRANT_PLUGIN"
	ENABLE_CORS_ENV_VAR                         = "ENABLE_CORS"
	VIRTUAL_MACHINES_FOLDER_ENV_VAR             = "VIRTUAL_MACHINES_FOLDER"
	EXECUTE_COMMAND_TIMEOUT_ENV_VAR             = "EXECUTE_COMMAND_TIMEOUT"
	PARALLELS_DESKTOP_REFRESH_INTERVAL_ENV_VAR  = "PARALLELS_DESKTOP_REFRESH_INTERVAL"
	SYSTEM_RESERVED_CPU_ENV_VAR                 = "SYSTEM_RESERVED_CPU"
	SYSTEM_RESERVED_MEMORY_ENV_VAR              = "SYSTEM_RESERVED_MEMORY"
	SYSTEM_RESERVED_DISK_ENV_VAR                = "SYSTEM_RESERVED_DISK"
)

const (
	TEST_COMMAND                  = "test"
	API_COMMAND                   = "api"
	REVERSE_PROXY_COMMAND         = "reverse-proxy"
	GENERATE_SECURITY_KEY_COMMAND = "gen-rsa"
	INSTALL_SERVICE_COMMAND       = "install"
	UNINSTALL_SERVICE_COMMAND     = "uninstall"
	VERSION_COMMAND               = "version"
	HELP_COMMAND                  = "help"
	CATALOG_COMMAND               = "catalog"
	CATALOG_PUSH_COMMAND          = "push"
	CATALOG_PULL_COMMAND          = "pull"
	UPDATE_ROOT_PASSWORD_COMMAND  = "update-root-pass"
	DELETE_COMMAND                = "delete"
	START_COMMAND                 = "start"
	STOP_COMMAND                  = "stop"
	EXEC_COMMAND                  = "exec"
	CLONE_COMMAND                 = "clone"

	TEST_FLAG                       = "test"
	TEST_CATALOG_PROVIDERS_FLAG     = "catalog-providers"
	API_PORT_FLAG                   = "port"
	UPDATE_ROOT_PASSWORD_FLAG       = "update-root-pass"
	FILE_FLAG                       = "file"
	RSA_KEY_SIZE                    = "rsa-key-size"
	MODE_FLAG                       = "mode"
	HELP_FLAG                       = "help"
	PASSWORD_FLAG                   = "password"
	USE_ORCHESTRATOR_RESOURCES_FLAG = "use-orchestrator-resources"
	CONFIG_FILE_FLAG                = "config"
)

const (
	PD_FILE_FROM_FLAG             = "from"
	PD_FILE_VERSION_FLAG          = "version"
	PD_FILE_ARCHITECTURE_FLAG     = "architecture"
	PD_FILE_LOCAL_PATH_FLAG       = "local-path"
	PD_FILE_ROLE_FLAG             = "role"
	PD_FILE_CLAIM_FLAG            = "claim"
	PD_FILE_PROVIDER_FLAG         = "provider"
	PD_FILE_DO_FLAG               = "do"
	PD_FILE_RUN_FLAG              = "run"
	PD_FILE_USERNAME_FLAG         = "username"
	PD_FILE_PASSWORD_FLAG         = "password"
	PD_FILE_API_KEY_FLAG          = "api-key"
	PD_FILE_CATALOG_ID_FLAG       = "catalog-id"
	PD_FILE_INSECURE_FLAG         = "insecure"
	PD_FILE_OUTPUT_FLAG           = "output"
	PD_FILE_DESTINATION_FLAG      = "destination"
	PD_FILE_DESCRIPTION_FLAG      = "description"
	PD_FILE_OWNER_FLAG            = "owner"
	PD_FILE_START_AFTER_PULL_FLAG = "start-after-pull"
	PD_FILE_MACHINE_NAME_FLAG     = "machine-name"
	PD_FILE_TAG_FLAG              = "tag"
)

const (
	USER_ROLE       = "USER"
	SUPER_USER_ROLE = "SUPER_USER"
	READ_ONLY_CLAIM = "READ_ONLY"

	LIST_CLAIM   = "LIST"
	CREATE_CLAIM = "CREATE"
	DELETE_CLAIM = "DELETE"
	UPDATE_CLAIM = "UPDATE"

	LIST_USER_CLAIM   = "LIST_USER"
	CREATE_USER_CLAIM = "CREATE_USER"
	DELETE_USER_CLAIM = "DELETE_USER"
	UPDATE_USER_CLAIM = "UPDATE_USER"

	LIST_API_KEY_CLAIM   = "LIST_API_KEY"
	CREATE_API_KEY_CLAIM = "CREATE_API_KEY"
	DELETE_API_KEY_CLAIM = "DELETE_API_KEY"
	UPDATE_API_KEY_CLAIM = "UPDATE_API_KEY"

	LIST_CLAIM_CLAIM   = "LIST_CLAIM"
	CREATE_CLAIM_CLAIM = "CREATE_CLAIM"
	DELETE_CLAIM_CLAIM = "DELETE_CLAIM"
	UPDATE_CLAIM_CLAIM = "UPDATE_CLAIM"

	LIST_ROLE_CLAIM   = "LIST_ROLE"
	CREATE_ROLE_CLAIM = "CREATE_ROLE"
	DELETE_ROLE_CLAIM = "DELETE_ROLE"
	UPDATE_ROLE_CLAIM = "UPDATE_ROLE"

	LIST_VM_CLAIM            = "LIST_VM"
	CREATE_VM_CLAIM          = "CREATE_VM"
	DELETE_VM_CLAIM          = "DELETE_VM"
	UPDATE_VM_STATES_CLAIM   = "UPDATE_VM_STATES"
	UPDATE_VM_CLAIM          = "UPDATE_VM"
	EXECUTE_COMMAND_VM_CLAIM = "EXECUTE_COMMAND_VM"

	LIST_PACKER_TEMPLATE_CLAIM   = "LIST_PACKER_TEMPLATE"
	CREATE_PACKER_TEMPLATE_CLAIM = "CREATE_PACKER_TEMPLATE"
	DELETE_PACKER_TEMPLATE_CLAIM = "DELETE_PACKER_TEMPLATE"
	UPDATE_PACKER_TEMPLATE_CLAIM = "UPDATE_PACKER_TEMPLATE"

	LIST_CATALOG_MANIFEST_CLAIM   = "LIST_CATALOG_MANIFEST"
	CREATE_CATALOG_MANIFEST_CLAIM = "CREATE_CATALOG_MANIFEST"
	DELETE_CATALOG_MANIFEST_CLAIM = "DELETE_CATALOG_MANIFEST"
	UPDATE_CATALOG_MANIFEST_CLAIM = "UPDATE_CATALOG_MANIFEST"
	PULL_CATALOG_MANIFEST_CLAIM   = "PULL_CATALOG_MANIFEST"
	PUSH_CATALOG_MANIFEST_CLAIM   = "PUSH_CATALOG_MANIFEST"
	IMPORT_CATALOG_MANIFEST_CLAIM = "IMPORT_CATALOG_MANIFEST"
)

var AllSystemRoles = []string{
	USER_ROLE,
	SUPER_USER_ROLE,
}

var AllSystemClaims = []string{
	READ_ONLY_CLAIM,
	CREATE_CLAIM,
	DELETE_CLAIM,
	UPDATE_CLAIM,
	LIST_CLAIM,
	LIST_USER_CLAIM,
	CREATE_USER_CLAIM,
	DELETE_USER_CLAIM,
	UPDATE_USER_CLAIM,
	LIST_API_KEY_CLAIM,
	CREATE_API_KEY_CLAIM,
	DELETE_API_KEY_CLAIM,
	UPDATE_API_KEY_CLAIM,
	LIST_CLAIM_CLAIM,
	CREATE_CLAIM_CLAIM,
	DELETE_CLAIM_CLAIM,
	UPDATE_CLAIM_CLAIM,
	LIST_ROLE_CLAIM,
	CREATE_ROLE_CLAIM,
	DELETE_ROLE_CLAIM,
	UPDATE_ROLE_CLAIM,
	CREATE_VM_CLAIM,
	DELETE_VM_CLAIM,
	LIST_VM_CLAIM,
	UPDATE_VM_STATES_CLAIM,
	UPDATE_VM_CLAIM,
	EXECUTE_COMMAND_VM_CLAIM,
	CREATE_PACKER_TEMPLATE_CLAIM,
	DELETE_PACKER_TEMPLATE_CLAIM,
	LIST_PACKER_TEMPLATE_CLAIM,
	UPDATE_PACKER_TEMPLATE_CLAIM,
	CREATE_CATALOG_MANIFEST_CLAIM,
	DELETE_CATALOG_MANIFEST_CLAIM,
	LIST_CATALOG_MANIFEST_CLAIM,
	UPDATE_CATALOG_MANIFEST_CLAIM,
	PULL_CATALOG_MANIFEST_CLAIM,
	PUSH_CATALOG_MANIFEST_CLAIM,
	IMPORT_CATALOG_MANIFEST_CLAIM,
}

var AllSuperUserClaims = []string{
	CREATE_CLAIM,
	DELETE_CLAIM,
	UPDATE_CLAIM,
	LIST_CLAIM,
	LIST_USER_CLAIM,
	CREATE_USER_CLAIM,
	DELETE_USER_CLAIM,
	UPDATE_USER_CLAIM,
	LIST_API_KEY_CLAIM,
	CREATE_API_KEY_CLAIM,
	DELETE_API_KEY_CLAIM,
	UPDATE_API_KEY_CLAIM,
	LIST_CLAIM_CLAIM,
	CREATE_CLAIM_CLAIM,
	DELETE_CLAIM_CLAIM,
	UPDATE_CLAIM_CLAIM,
	LIST_ROLE_CLAIM,
	CREATE_ROLE_CLAIM,
	DELETE_ROLE_CLAIM,
	UPDATE_ROLE_CLAIM,
	CREATE_VM_CLAIM,
	DELETE_VM_CLAIM,
	LIST_VM_CLAIM,
	UPDATE_VM_STATES_CLAIM,
	UPDATE_VM_CLAIM,
	EXECUTE_COMMAND_VM_CLAIM,
	CREATE_PACKER_TEMPLATE_CLAIM,
	DELETE_PACKER_TEMPLATE_CLAIM,
	LIST_PACKER_TEMPLATE_CLAIM,
	UPDATE_PACKER_TEMPLATE_CLAIM,
	CREATE_CATALOG_MANIFEST_CLAIM,
	DELETE_CATALOG_MANIFEST_CLAIM,
	LIST_CATALOG_MANIFEST_CLAIM,
	UPDATE_CATALOG_MANIFEST_CLAIM,
	PULL_CATALOG_MANIFEST_CLAIM,
	PUSH_CATALOG_MANIFEST_CLAIM,
	IMPORT_CATALOG_MANIFEST_CLAIM,
}

var DefaultRoles = []string{
	USER_ROLE,
}

var DefaultClaims = []string{
	READ_ONLY_CLAIM,
	LIST_CATALOG_MANIFEST_CLAIM,
	LIST_PACKER_TEMPLATE_CLAIM,
	LIST_VM_CLAIM,
	LIST_CLAIM,
}
