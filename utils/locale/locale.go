package locale

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/PrinceNorin/bakanovels/config"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/i18n"
)

func InitLocale(c config.I18nConfig) error {
	defaultLangPath := path.Join(c.Directory, fmt.Sprintf("%s.json", c.DefaultLanguage))
	err := i18n.LoadTranslationFile(defaultLangPath)
	if err != nil {
		return fmt.Errorf("failed to load default translation file '%s':'%v'", defaultLangPath, err)
	}

	paths, err := filepath.Glob(path.Join(c.Directory, "*.json"))
	if err != nil {
		return fmt.Errorf("failed to get translation files: %v", err)
	}

	for _, file := range paths {
		err := i18n.LoadTranslationFile(file)
		if err != nil {
			return fmt.Errorf("failed to load translation file '%s':'%v'", file, err)
		}
	}

	return nil
}

func GetTfunc(c *gin.Context) (i18n.TranslateFunc, error) {
	lang := c.Query("lang")
	if lang == "" {
		lang = config.Get().I18n.DefaultLanguage
	}

	return i18n.Tfunc(lang)
}

func GetTfuncWithMissing(c *gin.Context) (i18n.TranslateFunc, error) {
	T, err := GetTfunc(c)
	if err != nil {
		return nil, err
	}

	return func(tranID string, args ...interface{}) string {
		msg := T(tranID, args...)
		if msg == tranID {
			return T("translation.missing", map[string]interface{}{"Key": tranID})
		}
		return msg
	}, nil
}
