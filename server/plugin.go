package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/techbot/server/command"
	"github.com/techbot/server/config"
	"github.com/techbot/server/controller"
	"github.com/techbot/server/techbuzz"
	"github.com/techbot/server/util"

	algorithmia "github.com/algorithmiaio/algorithmia-go"
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
	"mvdan.cc/xurls"
)

type Plugin struct {
	plugin.MattermostPlugin
	running bool
	handler http.Handler
}

func (p *Plugin) OnActivate() error {
	config.Mattermost = p.API

	if err := p.setupStaticFileServer(); err != nil {
		p.API.LogError(err.Error())
		return err
	}

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	if err := p.RegisterCommands(); err != nil {
		config.Mattermost.LogError(err.Error())
		return err
	}

	p.Run()

	return nil
}

func (p *Plugin) setupStaticFileServer() error {
	return nil
}

func (p *Plugin) OnConfigurationChange() error {
	if config.Mattermost != nil {
		var configuration config.Configuration

		botID, err := p.setUpBot()
		if err != nil {
			return err
		}
		configuration.BotUserID = botID

		if err := config.Mattermost.LoadPluginConfiguration(&configuration); err != nil {
			config.Mattermost.LogError("Error in LoadPluginConfiguration: " + err.Error())
			return err
		}

		if err := configuration.ProcessConfiguration(); err != nil {
			config.Mattermost.LogError("Error in ProcessConfiguration: " + err.Error())
			return err
		}

		if err := configuration.IsValid(); err != nil {
			config.Mattermost.LogError("Error in Validating Configuration: " + err.Error())
			return err
		}

		config.SetConfig(&configuration)
	}
	return nil
}

func (p *Plugin) setUpBot() (string, error) {
	botID, err := p.Helpers.EnsureBot(&model.Bot{
		Username:    config.BotUsername,
		DisplayName: config.BotDisplayName,
		Description: "Bot for Tech Post.",
	})
	if err != nil {
		return "", err
	}

	return botID, nil
}

func (p *Plugin) RegisterCommands() error {
	if err := config.Mattermost.RegisterCommand(command.Master().Command); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split, argErr := util.SplitArgs(args.Command)
	if argErr != nil {
		return &model.CommandResponse{
			Type: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text: argErr.Error(),
		}, nil
	}

	function := split[0]
	var params []string

	if len(split) > 1 {
		params = split[1:]
	}

	if function != "/"+command.Master().Command.Trigger {
		return nil, &model.AppError{Message: "Unknown command: [" + function + "] encountered"}
	}

	context := p.prepareContext(args)
	if response, err := command.Master().Validate(params, context); response != nil {
		return response, err
	}

	// todo add error logs here
	return command.Master().Execute(params, context)
}

func (p *Plugin) prepareContext(args *model.CommandArgs) command.Context {
	return command.Context{
		CommandArgs: args,
		Props:       make(map[string]interface{}),
	}
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	conf := config.GetConfig()

	if err := conf.IsValid(); err != nil {
		p.API.LogError("This plugin is not configured: " + err.Error())
		http.Error(w, "This plugin is not configured.", http.StatusNotImplemented)
		return
	}

	path := r.URL.Path
	endpoint := controller.Endpoints[path]

	if endpoint == nil {
		p.handler.ServeHTTP(w, r)
	} else {
		endpoint.Execute(w, r)
	}
}

func (p *Plugin) Run() {
	if !p.running {
		p.running = true
		p.runner()
	}
}

func (p *Plugin) runner() {
	go func() {
		<-time.NewTimer(config.RunnerInterval).C
		techbuzz.SendPost()
		if !p.running {
			return
		}
		p.runner()
	}()
}

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	if post.ChannelId == config.GetConfig().TechBuzzChannel {
		rxRelaxed := xurls.Relaxed()
		URL := rxRelaxed.FindString(post.Message)
		if URL == "" {
			return
		}
		var client = algorithmia.NewClient(config.GetConfig().Apikey, "")
		algo, _ := client.Algo("tags/AutoTagURL/0.1.9?timeout=300") // timeout is optional
		flag := false
		for i := 0; i < 5; i++ {
			resp, err := algo.Pipe(URL)
			if err != nil {
				break
			}
			response := resp.(*algorithmia.AlgoResponse)
			fmt.Println(response.Result)
			for val, _ := range response.Result.(map[string]interface{}) {
				if techbuzz.TechTag[strings.ToLower(val)] {
					//frequency := int (fre.(float64))
					techbuzz.InsertData(strings.ToLower(val), post.Message)
					fmt.Println("From Api", val)
					flag = true
				}
			}
			if flag {
				break
			}
		}
		if flag == false {
			for val, _ := range techbuzz.TechTag {
				if strings.Contains(strings.ToLower(URL), val) {
					techbuzz.InsertData(strings.ToLower(val), post.Message)
					fmt.Println("From tag search", val)
					break
				}
			}
		} else {
			fmt.Println("Insertin in other tag:")
			techbuzz.InsertData("other", post.Message)
		}
	}
}

func main() {
	plugin.ClientMain(&Plugin{})
}
