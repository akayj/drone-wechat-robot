package main

type Message interface{}

type MessageType struct {
	Type string `json:"msgtype"`
}

type TextBody struct {
	Content string `json:"content"`
	// MentionedList       []string `json:"mentioned_list"`
	// MentionedMobileList []string `json:"mentioned_mobile_list"`
}

type MarkdownBody struct {
	Content string `json:"content"`
}

type TextMessage struct {
	MessageType
	Body TextBody `json:"text"`
}

type MarkdownMessage struct {
	MessageType
	Body MarkdownBody `json:"markdown"`
}
