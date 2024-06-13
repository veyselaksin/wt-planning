package i18n

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"wt-planning/helpers"
)

type localizer struct {
	MessageId    string
	TemplateData map[string]string
	PluralCount  int
}

type localizerBuilder struct {
	localizer *localizer
}

func (b *localizerBuilder) WithTemplateData(templateData map[string]string) *localizerBuilder {
	b.localizer.TemplateData = templateData
	return b
}

func (b *localizerBuilder) WithPluralCount(pluralCount int) *localizerBuilder {
	b.localizer.PluralCount = pluralCount
	return b
}

func (b *localizerBuilder) Build(loc *i18n.Localizer) string {

	var pluralCount int
	if b.localizer.PluralCount == 0 && b.localizer.TemplateData != nil {
		pluralCount = 1
	} else {
		pluralCount = b.localizer.PluralCount
	}

	message := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    b.localizer.MessageId,
		TemplateData: b.localizer.TemplateData,
		PluralCount:  pluralCount,
	})

	return message
}

func (b *localizerBuilder) BuildWithContext(c *fiber.Ctx) string {

	var pluralCount int
	if b.localizer.PluralCount == 0 && b.localizer.TemplateData != nil {
		pluralCount = 1
	} else {
		pluralCount = b.localizer.PluralCount
	}

	lang := helpers.GetLanguage(c)
	loc := i18n.NewLocalizer(bundle, lang)

	message := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    b.localizer.MessageId,
		TemplateData: b.localizer.TemplateData,
		PluralCount:  pluralCount,
	})

	return message
}

func (b *localizerBuilder) BuildWithLanguage(lang string) string {

	var pluralCount int
	if b.localizer.PluralCount == 0 && b.localizer.TemplateData != nil {
		pluralCount = 1
	} else {
		pluralCount = b.localizer.PluralCount
	}

	loc := i18n.NewLocalizer(bundle, lang)

	message := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID:    b.localizer.MessageId,
		TemplateData: b.localizer.TemplateData,
		PluralCount:  pluralCount,
	})

	return message
}

// CreateMessageBuilder is a helper function for creating message builder
func CreateMessageBuilder(messageId string) *localizerBuilder {

	return &localizerBuilder{
		localizer: &localizer{
			MessageId: messageId,
		},
	}
}

// CreateMsg is a helper function for creating message with context
func CreateMsg(ctx *fiber.Ctx, messageId string, templateData ...map[string]string) string {

	loc := i18n.NewLocalizer(bundle, helpers.GetLanguage(ctx))
	msg := loc.MustLocalize(&i18n.LocalizeConfig{
		MessageID: messageId,
	})

	if templateData != nil {
		msg = loc.MustLocalize(&i18n.LocalizeConfig{
			MessageID:    messageId,
			TemplateData: templateData[0],
		})
	}

	return msg
}
