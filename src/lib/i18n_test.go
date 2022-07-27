package lib_test

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/arcadia-lib/src/lib"
)

var _ = Describe("i18n", func() {

	Context("go-i18n", func() {
		When("using map of any", func() {
			It("🧪 should: translate", func() {
				notSupportedMsg := &i18n.Message{
					ID:    "language-not-supported.arcadia-lib",
					Other: "language '{{.Language}}' not supported",
				}

				localised := lib.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: notSupportedMsg,
					TemplateData:   map[string]any{"Language": "foo"},
				})
				expected := "language 'foo' not supported"
				Expect(localised).To(Equal(expected))
			})
		})

		// When("using template", func() {
		// 	It("🧪 should: translate", func() {
		// 		violationMsg := &i18n.Message{
		// 			ID:    "ov-failed-out-of-range",
		// 			Other: "({{.Flag}}): option validation failed, '{{.Value}}', out of range: [{{.Lo}}]..[{{.Hi}}]",
		// 		}

		// 		localised := lib.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
		// 			DefaultMessage: violationMsg,
		// 			TemplateData: l10n.WithinOptValidationTemplData{
		// 				OutOfRangeOV: l10n.OutOfRangeOV{Flag: "Strike", Value: 999, Lo: 1, Hi: 99},
		// 			},
		// 		})
		// 		expected := "(Strike): option validation failed, '999', out of range: [1]..[99]"
		// 		Expect(localised).To(Equal(expected))
		// 	})
		// })
	})
})
