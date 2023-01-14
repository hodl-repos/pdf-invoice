package document

import "fmt"

// translations hold all the languages and their translations
var translations = map[string]map[string]string{}

// LoadTranslations loads given translations to the global translations after
// all keys are checked for non-existence.
func LoadTranslation(newTranslation map[string]map[string]string) error {
	// check if all keys are new
	for lang, t := range newTranslation {
		if _, ok := translations[lang]; ok {
			for k := range t {
				if _, ok := translations[lang][k]; ok {
					return fmt.Errorf("%v %v already exists", lang, k)
				}
			}
		}
	}

	// add keys to translation map
	for lang, t := range newTranslation {
		if _, ok := translations[lang]; !ok {
			translations[lang] = map[string]string{}
		}
		for k, v := range t {
			translations[lang][k] = v
		}
	}

	return nil
}

func T(key, language string) string {
	if t, ok := translations[language]; ok {
		if v, ok := t[key]; ok {
			return v
		} else {
			panic(fmt.Sprintf("Translation for key %s not found in language %s", key, language))
		}
	} else {
		panic(fmt.Sprintf("Language %s not found", language))
	}
}

// clearTranslations is only used by tests to reset state
func clearTranslations() {
	translations = map[string]map[string]string{}
}
