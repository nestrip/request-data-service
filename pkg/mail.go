package pkg

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/nestrip/request-data-service/ent"
	"os"
)

var MG *mailgun.MailgunImpl

//InitMailGun loads the stuff from the environment variables
func InitMailGun() {
	MG = mailgun.NewMailgun(os.Getenv("MAILGUN_DOMAIN"), os.Getenv("MAILGUN_API_KEY"))
	MG.SetAPIBase(mailgun.APIBaseEU)
}

func SendDataRequestComplete(user *ent.User, request *ent.DataRequest, c context.Context) {
	message := MG.NewMessage(fmt.Sprintf("nest.rip <%s>", os.Getenv("MAILGUN_SENDER")), "Password changed",
		fmt.Sprintf(
			"Hello %s "+
				"\n\n"+
				"Your data request has completed, and you are able to download it."+
				"\n\n"+
				"Please click the link to download your data package: \n"+
				"https://cdn.nest.rip/data-requests/%s"+
				"\n\n"+
				"This link will expire after 30 days, so make sure to keep a copy!",
			user.Username, request.CdnName),
		user.Email)

	_, _, err := MG.Send(c, message)

	if err != nil {
		fmt.Printf("Could not send data request email to %s with error: %s\n", user.Username, err.Error())
	}
}
