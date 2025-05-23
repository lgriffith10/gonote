package translations

import (
	"embed"
	"encoding/json"
	"log"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	//go:embed active.en.json
	fs     embed.FS
	bundle *i18n.Bundle
)

var localizer *i18n.Localizer

func LoadTranslations() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	_, err := bundle.LoadMessageFileFS(fs, "active.en.json")
	if err != nil {
		log.Fatalf("Erreur chargement fichier : %v", err)
	}

	localizer = i18n.NewLocalizer(bundle, "en")

	// TODO: Add unit testing for translations
	/*
		 	tests := []string{"WelcomeUser", "emailAlreadyExists", "nonexistentID"}

			for _, id := range tests {
				msg, err := localizer.Localize(&i18n.LocalizeConfig{
					MessageID: id,
					TemplateData: map[string]interface{}{
						"User": "Alice",
					},
				})
				if err != nil {
					fmt.Printf("Erreur pour %q : %v\n", id, err)
					continue
				}
				fmt.Printf("Message pour %q : %s\n", id, msg)
			}
	*/
}

func GetTranslation(id string, args map[string]any) string {
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    id,
		TemplateData: args,
	})

	if err != nil {
		return id
	}

	return msg
}
