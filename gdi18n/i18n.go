package gdi18n

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

const DefaultLanguage = "en"

func NewBundle(sugar *zap.SugaredLogger, dirPath string) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		sugar.Fatalf("failed to read locale directory: %s", err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(dirPath, file.Name())

		_, err = bundle.LoadMessageFile(filePath)
		if err != nil {
			sugar.Fatalf("failed loading message file: %s", err)
		}
	}

	return bundle
}

type Localized struct {
	Localize map[string]*i18n.Localizer
	lang     string
}

func NewLocalize(bundle *i18n.Bundle, lang string) *Localized {
	if lang == "" {
		lang = DefaultLanguage
	}

	localizes := map[string]*i18n.Localizer{
		lang: i18n.NewLocalizer(bundle, lang),
	}

	return &Localized{
		Localize: localizes,
		lang:     lang,
	}
}

func (l *Localized) Translate(category, messageId string, data map[string]interface{}) string {
	localize, exists := l.Localize[l.lang]
	if !exists {
		localize = l.Localize["en"]
	}

	messageId = fmt.Sprintf("%s.%s", category, messageId)
	msg, err := localize.Localize(&i18n.LocalizeConfig{
		MessageID:    messageId,
		TemplateData: data,
	})
	if err != nil {
		return fmt.Sprintf("Missing translation: %s", messageId)
	}

	return msg
}
