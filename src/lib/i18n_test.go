package lib_test

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"

	"github.com/snivilised/arcadia-lib/src/lib"
	"github.com/snivilised/arcadia-lib/src/lib/internal/l10n"
)

var _ = Describe("i18n", Ordered, func() {
	var es, us, expected string

	BeforeAll(func() {
		es = language.Spanish.String()
		us = language.AmericanEnglish.String()
		expected = fmt.Sprintf("language '%v' not supported", es)
	})

	Context("UseTag", func() {
		When("given: tag is supported", func() {
			It("🧪 should: not return error", func() {
				us := language.AmericanEnglish
				Expect(lib.UseTag(us)).Error().To(BeNil())
				Expect(lib.GetLanguageInfo().Current).To(Equal(us))
			})

			It("🧪 should: localise in requested non default language", func() {
				lib.UseTag(language.AmericanEnglish)
				data := l10n.LanguageNotSupportedTemplData{
					Language: us,
				}

				_, tag, _ := lib.GetLocaliser().LocalizeWithTag(&i18n.LocalizeConfig{
					DefaultMessage: data.Message(),
					TemplateData:   data,
				})
				Expect(tag.String()).To(Equal(language.AmericanEnglish.String()))
			})
		})

		When("given: tag is NOT supported", func() {
			It("🧪 should: return error", func() {
				Expect(lib.UseTag(language.Spanish)).Error().ToNot(BeNil())
			})
		})
	})

	Context("go-i18n", func() {
		When("using map of any", func() {
			It("🧪 should: translate", func() {
				notSupportedMsg := &i18n.Message{
					ID:    "language-not-supported.arcadia-lib",
					Other: "language '{{.Language}}' not supported",
				}

				localised := lib.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: notSupportedMsg,
					TemplateData:   map[string]any{"Language": es},
				})

				Expect(localised).To(Equal(expected))
			})
		})

		When("using template", func() {
			It("🧪 should: translate", func() {
				localised := lib.GetLanguageNotSupportedErrorMessage(language.Spanish)

				Expect(localised).To(Equal(expected))
			})
		})
	})
})
