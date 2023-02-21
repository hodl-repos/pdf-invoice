package localize

import (
	"context"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	go2 "github.com/adam-hanna/arrayOperations"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// contextKey is a private string type to prevent collisions in the context map.
type contextKey string

// loggerKey points to the value in the context where the logger is stored
const loggerKey = contextKey("localizer")

type LocalizeService struct {
	bundle *i18n.Bundle
	langs  *[]language.Tag
}

type LocalizeClient struct {
	service   *LocalizeService
	localizer *i18n.Localizer
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

func ContextWithLocalizeService(ctx context.Context, service *LocalizeService) context.Context {
	return context.WithValue(ctx, loggerKey, service)
}

func FromContext(ctx context.Context) *LocalizeService {
	if logger, ok := ctx.Value(loggerKey).(*LocalizeService); ok {
		return logger
	}

	panic("must register localizer before using")
}

func (service *LocalizeService) CreateLocalizer(preferedLangs ...string) *i18n.Localizer {
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

func (service *LocalizeService) CreateClient(preferedLangs ...string) *LocalizeClient {
	localizer := service.CreateLocalizer(preferedLangs...)

	return &LocalizeClient{
		service:   service,
		localizer: localizer,
	}
}
