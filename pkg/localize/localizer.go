package localize

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type LocalizeService struct {
	bundle *i18n.Bundle
	langs  *[]language.Tag
}

type LocalizeClient struct {
	service   *LocalizeService
	localizer *i18n.Localizer
	printer   *message.Printer
}

// NewLogger creates a new logger with the given configuration.
func NewLocalizeService(config *Config) *LocalizeService {
	langs := strings.Split(config.LangKeys, ",")

	if len(langs) == 0 {
		fmt.Sprintln("only english will be avaliable")
	}

	defLang := language.MustParse("en")

	bundle := i18n.NewBundle(defLang)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("active.de.toml")

	langArray := make([]language.Tag, 0)

	langArray = append(langArray, defLang)

	for _, lang := range langs {
		if parsed, err := language.Parse(lang); err == nil {
			if _, ok := go2.FindOne(langArray, func(f language.Tag) bool { return f == parsed }); ok {
				continue
			}

			langArray = append(langArray, parsed)
		}
	}

	return &LocalizeService{
		bundle: bundle,
		langs:  &langArray,
	}
}

func (service *LocalizeService) createLocalizer(preferedLangs ...string) *i18n.Localizer {
	langs := make([]language.Tag, 0)

	for _, lang := range preferedLangs {
		if parsed, err := language.Parse(lang); err == nil {
			if _, ok := go2.FindOne(langs, func(f language.Tag) bool { return f == parsed }); ok {
				continue
			}

			langs = append(langs, parsed)
		}
	}

	if len(langs) == 0 {
		langs = *service.langs
	}

	resLangs := make([]string, 0)
	for _, lang := range langs {
		resLangs = append(resLangs, lang.String())
	}

	return i18n.NewLocalizer(service.bundle, resLangs...)
}

func (service *LocalizeService) CreateClient(locale string, preferedLangs ...string) *LocalizeClient {
	localizer := service.createLocalizer(preferedLangs...)

	bestLang := preferedLangs[0]

	return &LocalizeClient{
		service:   service,
		localizer: localizer,
		printer:   createPrinter(bestLang, locale),
	}
}

func createPrinter(bestLang, locale string) *message.Printer {
	lang := language.Make(bestLang + "-" + strings.ToUpper(locale))
	return message.NewPrinter(lang)
}

func (client *LocalizeClient) FFloat32(data float32) string {
	return client.printer.Sprintf("%.2f", data)
}

func (client *LocalizeClient) FFloat64(data float64) string {
	return client.printer.Sprintf("%.2f", data)
}

func (client *LocalizeClient) FInt(data int) string {
	return client.printer.Sprintf("%g", data)
}
