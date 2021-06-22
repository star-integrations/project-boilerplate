// Package props is a scaffold file for props of controllers
package props

import (
	"github.com/star-integrations/project-boilerplate/back/pkg/config"
)

// ControllerProps is passed from Bootstrap() to all controllers
type ControllerProps struct {
	// DB, config, etc...
	Config *config.Config
}
