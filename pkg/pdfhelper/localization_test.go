package pdfhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadTranslation(t *testing.T) {
	defer clearTranslations() // clear translations after test

	// load single language translation key
	var data = map[string]map[string]string{
		"en": {
			"VAT": "VAT",
		},
	}

	err := LoadTranslation(data)
	assert.NoError(t, err)

	// load another translation key to same language
	data = map[string]map[string]string{
		"en": {
			"AMOUNT_IN_EUR_EXCL_VAT": "Amount in EUR excl. VAT",
		},
	}

	err = LoadTranslation(data)
	assert.NoError(t, err)

	// loading same translation key should result in an error
	err = LoadTranslation(data)
	assert.Error(t, err)

	// load different language
	data = map[string]map[string]string{
		"de": {
			"AMOUNT_IN_EUR_EXCL_VAT": "Betrag in EUR zzgl. USt.",
			"VAT":                    "USt.",
		},
	}

	err = LoadTranslation(data)
	assert.NoError(t, err)
}

func TestT(t *testing.T) {
	defer clearTranslations() // clear translations after test

	// prepare translations
	var data = map[string]map[string]string{
		"en": {
			"VAT": "VAT",
		},
	}

	err := LoadTranslation(data)
	assert.NoError(t, err)

	// get key without panic
	assert.NotPanics(t, func() { T("VAT", "en") })

	// panics with wrong key
	assert.PanicsWithValue(t, "Translation for key WRONG_KEY not found in language en", func() { T("WRONG_KEY", "en") })

	// panics with wrong lang
	assert.PanicsWithValue(t, "Language at not found", func() { T("VAT", "at") })

	// panics with wrong lang and key, lang takes precedence
	assert.PanicsWithValue(t, "Language at not found", func() { T("VAT", "at") })

}
