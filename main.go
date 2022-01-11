package main

import (
	"fmt"

	"os"

	"github.com/ananchev/homeconnect-proxy/internal/logger"
	"github.com/ananchev/homeconnect-proxy/internal/proxy"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config is the application configuration structure
type Config struct {
	OAuth struct {
		ClientID     string `env:"CLIENT_ID" env-description:"Home Connect application client ID"`
		ClientSecret string `env:"CLIENT_SECRET" env-description:"Home Connect application client secret"`
		ClientScopes string `env:"CLIENT_SCOPES" env-description:"Home Connect application authorization scopes"`
	}

	Server struct {
		Host string `env:"HOST" env-description:"Server host" env-default:"localhost"`
		Port string `env:"PORT" env-description:"Server port" env-default:"8088"`
	}
}

func main() {
	var cfg Config

	// read configuration from environment variables
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// file, _ := json.MarshalIndent(cfg.Commands, "", "    ")
	// _ = ioutil.WriteFile("commands.json", file, 0644)

	logger.Info("Starting the Home Connect client proxy ...")
	proxy.Run(cfg.Server.Port, cfg.OAuth.ClientID, cfg.OAuth.ClientSecret, cfg.OAuth.ClientScopes)

}