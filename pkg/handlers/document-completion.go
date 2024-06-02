package handlers

import (
	"github.com/opa-oz/emoji-lsp/pkg/mappers"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"

	_ "github.com/tliron/commonlog/simple"
)

func TextDocumentCompletion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var completionItems []protocol.CompletionItem

	for word, emoji := range mappers.EmojiMapper {
		emojiCopy := emoji // Create a copy of emoji
		completionItems = append(completionItems, protocol.CompletionItem{
			Label:      word + " - [Emoji]",
			Detail:     &emojiCopy,
			InsertText: &emojiCopy,
		})
	}

	return completionItems, nil
}
