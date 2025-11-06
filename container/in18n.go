package gdcontainer

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

type II18nProvider interface {
	Localize(lang string) *i18n.Localizer
	Translate(lang, category, messageId string, data map[string]interface{}) string
}

type i18nProvider struct {
	bundle      *i18n.Bundle
	defaultLang string
}

func NewI18nProvider(zap *zap.SugaredLogger, dirPath string) (II18nProvider, func(), error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read locale directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".json" {
			continue
		}

		filePath := filepath.Join(dirPath, file.Name())
		_, err = bundle.LoadMessageFile(filePath)
		if err != nil {
			return nil, nil, fmt.Errorf("failed loading message file %s: %w", filePath, err)
		}
	}

	provider := &i18nProvider{
		bundle:      bundle,
		defaultLang: DefaultLanguage,
	}

	// Optional cleanup (no open resources here, but consistent with other providers)
	cleanup := func() {
		zap.Info("closing i18n provider resources")
	}

	zap.Infof("loaded i18n translations from: %s", dirPath)
	return provider, cleanup, nil
}

func (p *i18nProvider) Localize(lang string) *i18n.Localizer {
	if lang == "" {
		lang = p.defaultLang
	}
	return i18n.NewLocalizer(p.bundle, lang)
}

func (p *i18nProvider) Translate(lang, category, messageId string, data map[string]interface{}) string {
	localizer := p.Localize(lang)
	fullID := fmt.Sprintf("%s.%s", category, messageId)

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    fullID,
		TemplateData: data,
	})
	if err != nil {
		return fmt.Sprintf("Missing translation: %s", fullID)
	}
	return msg
}
