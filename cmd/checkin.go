package cmd

import (
	"log"
	"os"
	"time"

	"github.com/buildoric/acheckin/pkg/config"
	httprequest "github.com/buildoric/acheckin/pkg/http-request"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var checkinCmd = &cobra.Command{
	Use: "checkin",
	Run: func(cmd *cobra.Command, args []string) {
		flag := rootCmd.PersistentFlags().Lookup("config")
		configPath := flag.Value.String()
		if configPath == "" {
			log.Fatal("Required a config file")
			return
		}

		yamlFile, err := os.ReadFile(configPath)

		if err != nil {
			log.Fatalln("Read file error:", err)
		}
		var c *config.Config

		err = yaml.Unmarshal(yamlFile, &c)

		if err != nil {
			log.Fatalln("Unmarshal yaml error:", err)
		}

		createRequest := httprequest.CreateRequest{
			Config: c,
		}

		user := createRequest.Login()

		if user != nil && user.Token != "" {

			createRequest.Checkin(user.Token)

			timeData, _ := createRequest.CanCheckIn(user.Token, user.Data.UserObjId)

			timeCheckin, _ := time.Parse("2006-01-02 15:04:05", timeData.Data.TimeKeepingMonth.TimeCheckIn)

			timeCheckout := timeCheckin.Add(time.Hour*9 + time.Minute*30)
			cronSpec := timeCheckout.Format("04 15 * * *")
			c := cron.New()
			_, err := c.AddFunc(cronSpec, func() {
				createRequest.Checkin(user.Token)
			})
			if err != nil {
				log.Fatalf("Error scheduling cron job: %v", err)
			}

			log.Println("Starting cron scheduler...")
			c.Start()
			select {}
		} else {
			log.Fatalln("Auth failed.")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkinCmd)
}
