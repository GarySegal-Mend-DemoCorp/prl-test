package startup

import (
	"Parallels/pd-api-service/constants"
	"Parallels/pd-api-service/controllers"
	"Parallels/pd-api-service/restapi"
	"Parallels/pd-api-service/service_provider"
)

var listener *restapi.HttpListener

func InitApi() *restapi.HttpListener {
	listener = restapi.GetHttpListener()
	listener.Options.ApiPrefix = constants.API_PREFIX
	listener.AddJsonContent().AddLogger().AddHealthCheck()
	listener.WithPublicUserRegistration()
	RegisterControllers()

	return listener
}

func ResetApi() {
	listener.WaitAndShutdown()
}

func RegisterControllers() {
	provider := service_provider.Get()
	listener.AddController(controllers.GetTokenController(), "/auth/token", "POST")
	listener.AddController(controllers.ValidateTokenController(), "/auth/token/validate", "POST")

	// Users Controller
	listener.AddAuthorizedControllerWithRoles(controllers.GetUsersController(), "/auth/users", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.CreateUserController(), "/auth/users", []string{constants.SUPER_USER_ROLE}, "POST")
	listener.AddAuthorizedControllerWithRoles(controllers.GetUserController(), "/auth/users/{id}", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.DeleteUserController(), "/auth/users/{id}", []string{constants.SUPER_USER_ROLE}, "DELETE")
	listener.AddAuthorizedControllerWithRoles(controllers.UpdateUserController(), "/auth/users/{id}", []string{constants.SUPER_USER_ROLE}, "PUT")
	listener.AddAuthorizedControllerWithRoles(controllers.GetUserRolesController(), "/auth/users/{id}/roles", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.AddRoleToUserController(), "/auth/users/{id}/roles", []string{constants.SUPER_USER_ROLE}, "POST")
	listener.AddAuthorizedControllerWithRoles(controllers.RemoveRoleFromUserController(), "/auth/users/{id}/roles/{role_id}", []string{constants.SUPER_USER_ROLE}, "DELETE")
	listener.AddAuthorizedControllerWithRoles(controllers.GetUserClaimsController(), "/auth/users/{id}/claims", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.AddClaimToUserController(), "/auth/users/{id}/claims", []string{constants.SUPER_USER_ROLE}, "POST")
	listener.AddAuthorizedControllerWithRoles(controllers.RemoveClaimFromUserController(), "/auth/users/{id}/claims/{role_id}", []string{constants.SUPER_USER_ROLE}, "DELETE")

	// ApiKey Controller
	listener.AddAuthorizedControllerWithRoles(controllers.GetApiKeysController(), "/auth/api_keys", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.CreateApiKeyController(), "/auth/api_keys", []string{constants.SUPER_USER_ROLE}, "POST")
	listener.AddAuthorizedControllerWithRoles(controllers.GetApiKeyByIdOrNameController(), "/auth/api_keys/{id}", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.DeleteApiKeyController(), "/auth/api_keys/{id}", []string{constants.SUPER_USER_ROLE}, "DELETE")
	listener.AddAuthorizedControllerWithRoles(controllers.RevokeApiKeyController(), "/auth/api_keys/{id}/revoke", []string{constants.SUPER_USER_ROLE}, "PUT")

	// Config Controller
	listener.AddAuthorizedControllerWithRoles(controllers.InstallToolsController(), "/config/tools/install", []string{constants.SUPER_USER_ROLE}, "POST")
	listener.AddAuthorizedControllerWithRoles(controllers.UninstallToolsController(), "/config/tools/uninstall", []string{constants.SUPER_USER_ROLE}, "POST")
	listener.AddAuthorizedControllerWithRoles(controllers.GetParallelsDesktopLicenseController(), "/config/tools/parallels_desktop/license", []string{constants.SUPER_USER_ROLE}, "GET")
	listener.AddAuthorizedControllerWithRoles(controllers.RestartController(), "/config/restart", []string{constants.SUPER_USER_ROLE}, "POST")

	// Virtual Machines Controller
	if provider.IsParallelsDesktopAvailable() {
		listener.AddAuthorizedControllerWithClaims(controllers.GetParallelsDesktopLicenseController(), "/parallels_desktop/key", []string{constants.LIST_VM_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.GetMachinesController(), "/machines", []string{constants.LIST_VM_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.CreateMachine(), "/machines", []string{constants.CREATE_VM_CLAIM}, "POST")
		listener.AddAuthorizedControllerWithClaims(controllers.RegisterMachineController(), "/machines/register", []string{constants.CREATE_VM_CLAIM}, "POST")
		listener.AddAuthorizedControllerWithClaims(controllers.DeleteMachineController(), "/machines/{id}", []string{constants.DELETE_VM_CLAIM}, "DELETE")
		listener.AddAuthorizedControllerWithClaims(controllers.GetMachineController(), "/machines/{id}", []string{constants.LIST_VM_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.StartMachineController(), "/machines/{id}/start", []string{constants.UPDATE_VM_STATES_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.StopMachineController(), "/machines/{id}/stop", []string{constants.UPDATE_VM_STATES_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.RestartMachineController(), "/machines/{id}/restart", []string{constants.UPDATE_VM_STATES_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.SuspendMachineController(), "/machines/{id}/suspend", []string{constants.UPDATE_VM_STATES_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.ResetMachineController(), "/machines/{id}/reset", []string{constants.UPDATE_VM_STATES_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.PauseMachineController(), "/machines/{id}/pause", []string{constants.UPDATE_VM_STATES_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.StatusMachineController(), "/machines/{id}/status", []string{constants.LIST_VM_CLAIM}, "GET")
		listener.AddAuthorizedControllerWithClaims(controllers.SetMachineController(), "/machines/{id}/set", []string{constants.UPDATE_VM_CLAIM}, "POST")
		listener.AddAuthorizedControllerWithClaims(controllers.ExecuteCommandOnMachineController(), "/machines/{id}/execute", []string{constants.EXECUTE_COMMAND_VM_CLAIM}, "POST")
		listener.AddAuthorizedControllerWithClaims(controllers.RenameMachineController(), "/machines/{id}/rename", []string{constants.UPDATE_VM_CLAIM}, "POST")
		listener.AddAuthorizedControllerWithClaims(controllers.UnregisterMachineController(), "/machines/{id}/unregister", []string{constants.UPDATE_VM_CLAIM}, "POST")
	}

	// Templates Controller
	listener.AddAuthorizedControllerWithClaims(controllers.GetVirtualMachinesTemplatesController(), "/templates/virtual_machines", []string{constants.LIST_TEMPLATE_CLAIM}, "GET")
	listener.AddAuthorizedControllerWithClaims(controllers.GetVirtualMachineTemplateController(), "/templates/virtual_machines/{name}", []string{constants.LIST_VM_CLAIM}, "GET")

	// Remote Machines Catalog Controller
	listener.AddAuthorizedControllerWithClaims(controllers.GetCatalogManifestsController(), "/catalog", []string{constants.LIST_CATALOG_MANIFEST_CLAIM}, "GET")
	listener.AddAuthorizedControllerWithClaims(controllers.GetCatalogManifestController(), "/catalog", []string{constants.CREATE_CATALOG_MANIFEST_CLAIM}, "POST")
	listener.AddAuthorizedControllerWithClaims(controllers.GetCatalogManifestController(), "/catalog/{id}", []string{constants.DELETE_CATALOG_MANIFEST_CLAIM}, "DELETE")
	listener.AddAuthorizedControllerWithClaims(controllers.GetCatalogManifestController(), "/catalog/{id}", []string{constants.LIST_CATALOG_MANIFEST_CLAIM}, "GET")
	listener.AddAuthorizedControllerWithClaims(controllers.PushCatalogManifestController(), "/catalog/push", []string{constants.PUSH_CATALOG_MANIFEST_CLAIM}, "POST")
	listener.AddAuthorizedControllerWithClaims(controllers.PullCatalogManifestController(), "/catalog/pull", []string{constants.PULL_CATALOG_MANIFEST_CLAIM}, "POST")
}
