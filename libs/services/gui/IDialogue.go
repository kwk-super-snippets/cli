package gui

type IDialogues interface {
    Modal(templateName string, data interface{}) *DialogueResponse
    Field(templateName string, data interface{}) *DialogueResponse
    MultiChoice(templateName string, header interface{}, options interface{}) *DialogueResponse
}