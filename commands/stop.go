package commands

import (
	"context"

	"github.com/pavlo/slack-time/models"
)

//Stop - handles the '/timer stop` command received from Slack
type Stop struct {
}

// cases:
// 1. Successfully stopped a timer
// 2. No currently ticking timer existed
// 3. Any other errors

// Handle - SlackCustomCommandHandler interface
func (c *Stop) Handle(ctx context.Context, slackCommand models.SlackCustomCommand) *ResponseToSlack {

	// dataService := data.CreateDataService()
	// db := utils.GetDBTransactionFromContext(ctx)
	// team, user, project := dataService.CreateTeamAndUserAndProject(db, slackCommand)

	// log.Printf("%v %v %v", team, user, project)
	return nil
}
