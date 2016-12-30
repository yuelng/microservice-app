package utils

import (
	"log"

	"github.com/smartwalle/sendcloud"
)

func SendTemplateMail(invokeName, subject string, to []map[string]string) (ok bool, result string) {
	from, err := ConfigGetString("sendcloud", "from")
	FailOnError("get sender failed: %s", err)

	fromName, err := ConfigGetString("sendcloud", "from_name")
	FailOnError("get sender name failed: %s", err)

	apiKey, err := ConfigGetString("sendcloud", "api_key")
	FailOnError("get access_key failed: %s", err)

	apiUser, err := ConfigGetString("sendcloud", "api_user")
	FailOnError("get access_user failed: %s", err)

	sendcloud.UpdateApiInfo(apiUser, apiKey)

	ok, err, result = sendcloud.SendTemplateMail(invokeName, from, fromName, "", subject, to, nil)
	if !ok {
		log.Println(result, err)
	}
	return
}
