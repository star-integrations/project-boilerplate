package environ

import (
	"os"
	"strings"

	"github.com/star-integrations/project-boilerplate/back/pkg/definitions"
)

var (
	projectID      = getProjectID()
	serviceName    = os.Getenv(definitions.EnvKeyServiceName)
	serviceVersion = os.Getenv(definitions.EnvKeyServiceVersion)
)

func getProjectID() string {
	id, ok := os.LookupEnv(definitions.EnvKeyProjectID)
	if ok {
		return id
	}

	id, ok = os.LookupEnv(definitions.EnvKeyGoogleCloudProject)
	if ok {
		return id
	}

	return ""
}

// IsLocal - determine if it is a local environment
func IsLocal() bool {
	return projectID == ""
}

// IsDev - determine if it is a dev environment
func IsDev() bool {
	return strings.HasSuffix(projectID, "-dev")
}

// IsQA - determine if it is a qa environment
func IsQA() bool {
	return strings.HasSuffix(projectID, "-qa")
}

// IsStaging - determine if it is a staging environment
func IsStaging() bool {
	return strings.HasSuffix(projectID, "-stg")
}

// IsProd - determine if it is a production environment
func IsProd() bool {
	return strings.HasSuffix(projectID, "-prod")
}

// GetProjectID - get the project id
func GetProjectID() string {
	id := projectID
	if id == "" {
		return "localProject"
	}
	return id
}

// GetServiceName - get the service name
func GetServiceName() string {
	if serviceName == "" {
		return "localService"
	}
	return serviceName
}

// GetServiceVersion - get the service version
func GetServiceVersion() string {
	if serviceVersion == "" {
		return "1.0"
	}
	return serviceVersion
}
