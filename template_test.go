package goaiml

import "testing"

func TestTemplate_SetTag(t *testing.T) {
	aiml := NewAIMLInterpreter()
	v, err := aiml.ProcessSetTag(`this is a <set name="key">value</set>`)

	if err != nil {
		t.Error("Error to parser:", err)
	}

	if v != "this is a value" {
		t.Error("Result not match:", v)
	}

	m, ok := aiml.Memory["key"]

	if !ok {
		t.Error("Key not found at memory")
	}

	if m != "value" {
		t.Error("Result different of value")
	}
}

func TestTemplate_GetTag(t *testing.T) {
	aiml := NewAIMLInterpreter()
	_, err := aiml.ProcessSetTag(`this is a <set name="key">value</set>`)

	if err != nil {
		t.Error("Error to parser:", err)
	}

	v, errG := aiml.ProcessGetTag(`get the <get name="key"/>`)

	if errG != nil {
		t.Error("Error to parser:", errG)
	}

	if v != "get the value" {
		t.Error("Result different of value:", v)
	}
}

func TestTemplate_BotTag(t *testing.T) {
	aiml := NewAIMLInterpreter()
	v, err := aiml.ProcessBotTag(`the bot name is <bot name="name"/>`)

	if err != nil {
		t.Error("Error to parser:", err)
	}

	if v != "the bot name is "+BOT_NAME {
		t.Error("Result different of value:", v)
	}
}

func TestTemplate_StarTag(t *testing.T) {
	aiml := NewAIMLInterpreter()
	v := aiml.ProcessStarTag(`WHATS APP <star/> JOW <star/>`, []string{"", "MY", "FRIEND"})

	if v != "WHATS APP MY JOW FRIEND" {
		t.Error("Result different of value:", v)
	}
}

func TestTemplate_Think(t *testing.T) {
	aiml := NewAIMLInterpreter()

	starContent := []string{}
	v, err := aiml.ProcessThinkTag(`<think><set name="goodquestion">thequestion</set></think>JOW`, starContent)

	if err != nil {
		t.Error("Error to parser:", err)
	}

	_, ok := aiml.Memory["goodquestion"]

	if !ok {
		t.Error("Memory not setted")
	}

	if v != "JOW" {
		t.Error("Result different of value:", v)
	}
}

func TestTemplate_Random(t *testing.T) {
	aiml := NewAIMLInterpreter()

	xml := `<random>
		<li>opt 1</li>
		<li>opt 2</li>
	</random>`

	starContent := []string{}
	v, err := aiml.ProcessRandomTag(xml, starContent)

	if err != nil {
		t.Error("Error to parser:", err)
	}

	if v != "opt 1" && v != "opt 2" {
		t.Error("Result different of value:", v)
	}
}