package main

import (
	"encoding/json"
	"fmt"
	"github.com/openziti/zrok/environment"
	"github.com/openziti/zrok/sdk/golang/sdk"
	"github.com/openziti/zrok/tui"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"slices"
	"time"
)

func init() {
	rootCmd.AddCommand(newReserveCommand().cmd)
}

type reserveCommand struct {
	basicAuth          []string
	frontendSelection  []string
	backendMode        string
	jsonOutput         bool
	oauthProvider      string
	oauthEmailDomains  []string
	oauthCheckInterval time.Duration
	cmd                *cobra.Command
}

func newReserveCommand() *reserveCommand {
	cmd := &cobra.Command{
		Use:   "reserve <public|private> <target>",
		Short: "Create a reserved share",
		Args:  cobra.ExactArgs(2),
	}
	command := &reserveCommand{cmd: cmd}
	cmd.Flags().StringArrayVar(&command.frontendSelection, "frontends", []string{"public"}, "Selected frontends to use for the share")
	cmd.Flags().StringVarP(&command.backendMode, "backend-mode", "b", "proxy", "The backend mode (public|private: proxy, web, caddy, drive) (private: tcpTunnel, udpTunnel)")
	cmd.Flags().BoolVarP(&command.jsonOutput, "json-output", "j", false, "Emit JSON describing the created reserved share")
	cmd.Flags().StringArrayVar(&command.basicAuth, "basic-auth", []string{}, "Basic authentication users (<username:password>,...)")
	cmd.Flags().StringVar(&command.oauthProvider, "oauth-provider", "", "Enable OAuth provider [google, github]")
	cmd.Flags().StringArrayVar(&command.oauthEmailDomains, "oauth-email-domains", []string{}, "Allow only these email domains to authenticate via OAuth")
	cmd.Flags().DurationVar(&command.oauthCheckInterval, "oauth-check-interval", 3*time.Hour, "Maximum lifetime for OAuth authentication; reauthenticate after expiry")
	cmd.MarkFlagsMutuallyExclusive("basic-auth", "oauth-provider")

	cmd.Run = command.run
	return command
}

func (cmd *reserveCommand) run(_ *cobra.Command, args []string) {
	shareMode := sdk.ShareMode(args[0])
	privateOnlyModes := []string{"tcpTunnel", "udpTunnel"}
	if shareMode != sdk.PublicShareMode && shareMode != sdk.PrivateShareMode {
		tui.Error("invalid sharing mode; expecting 'public' or 'private'", nil)
	} else if shareMode == sdk.PublicShareMode && slices.Contains(privateOnlyModes, cmd.backendMode) {
		tui.Error(fmt.Sprintf("invalid sharing mode for a %s share: %s", sdk.PublicShareMode, cmd.backendMode), nil)
	}

	var target string
	switch cmd.backendMode {
	case "proxy":
		v, err := parseUrl(args[1])
		if err != nil {
			tui.Error("invalid target endpoint URL", err)
		}
		target = v

	case "web":
		target = args[1]

	case "tcpTunnel":
		target = args[1]

	case "udpTunnel":
		target = args[1]

	case "caddy":
		target = args[1]

	case "drive":
		target = args[1]

	default:
		tui.Error(fmt.Sprintf("invalid backend mode '%v'; expected {proxy, web, tcpTunnel, udpTunnel, caddy, drive}", cmd.backendMode), nil)
	}

	env, err := environment.LoadRoot()
	if err != nil {
		tui.Error("error loading environment", err)
	}

	if !env.IsEnabled() {
		tui.Error("unable to load environment; did you 'zrok enable'?", nil)
	}

	req := &sdk.ShareRequest{
		Reserved:    true,
		BackendMode: sdk.BackendMode(cmd.backendMode),
		ShareMode:   shareMode,
		BasicAuth:   cmd.basicAuth,
		Target:      target,
	}
	if shareMode == sdk.PublicShareMode {
		req.Frontends = cmd.frontendSelection
	}
	if cmd.oauthProvider != "" {
		if shareMode != sdk.PublicShareMode {
			tui.Error("--oauth-provider only supported for public shares", nil)
		}
		req.OauthProvider = cmd.oauthProvider
		req.OauthEmailDomains = cmd.oauthEmailDomains
		req.OauthAuthorizationCheckInterval = cmd.oauthCheckInterval
	}
	shr, err := sdk.CreateShare(env, req)
	if err != nil {
		tui.Error("unable to create share", err)
	}

	if !cmd.jsonOutput {
		logrus.Infof("your reserved share token is '%v'", shr.Token)
		for _, fpe := range shr.FrontendEndpoints {
			logrus.Infof("reserved frontend endpoint: %v", fpe)
		}
	} else {
		out, err := json.Marshal(shr)
		if err != nil {
			tui.Error("error emitting JSON", err)
		}
		fmt.Println(string(out))
	}
}
