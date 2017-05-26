package v3

import (
	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/v3/shared"
)

//go:generate counterfeiter . V3SetDropletActor

type V3SetDropletActor interface {
	SetAppDroplet(appName string, dropletGUID string) (v3action.Warnings, error)
}

type V3SetDropletCommand struct {
	usage       interface{} `usage:"CF_NAME v3-set-droplet --name [name] --droplet-guid [guid]"`
	AppName     string      `short:"n" long:"name" description:"The desired application name" required:"true"`
	DropletGUID string      `long:"droplet-guid" description:"The guid of the droplet to stage" required:"true"`

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       V3SetDropletActor
}

func (cmd *V3SetDropletCommand) Setup(config command.Config, ui command.UI) error {
	// cmd.UI = ui
	// cmd.Config = config
	// cmd.SharedActor = sharedaction.NewActor()

	// ccClient, uaaClient, err := shared.NewClients(config, ui, true)
	// if err != nil {
	// 	return err
	// }
	// cmd.Actor = v3action.NewActor(ccClient, config)

	// dopplerURL, err := hackDopplerURLFromUAA(ccClient.UAA())
	// if err != nil {
	// 	return err
	// }
	// cmd.NOAAClient = shared.NewNOAAClient(dopplerURL, config, uaaClient, ui)

	return nil
}

func (cmd V3SetDropletCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(cmd.Config, true, true)
	if err != nil {
		return shared.HandleError(err)
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Setting app {{.AppName}} to droplet {{.DropletGUID}} in org {{.OrgName}} / space {{.SpaceName}} as {{.Username}}...", map[string]interface{}{
		"AppName":     cmd.AppName,
		"DropletGUID": cmd.DropletGUID,
		"OrgName":     cmd.Config.TargetedOrganization().Name,
		"SpaceName":   cmd.Config.TargetedSpace().Name,
		"Username":    user.Name,
	})

	warnings, err := cmd.Actor.SetAppDroplet(cmd.AppName, cmd.DropletGUID)
	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return shared.HandleError(err)
	}
	cmd.UI.DisplayOK()

	return nil
}
