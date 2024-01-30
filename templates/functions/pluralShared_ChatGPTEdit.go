package main

import (
	"fmt"
	"log"
	"errors"
	"encoding/json"

	"github.com/sashabaranov/go-openai"
)

func (app *App) {{lowercase .Object.Name}}ChatGPTEdit(user *User, parent *Internals, object *{{uppercase .Object.Name}}, prompt string) error {

	fmt.Println("prompt with parent", parent.ID, prompt)

	objectBytes, err := app.MarshalJSON(object)
	if err != nil {
		return err
	}

	system := fmt.Sprintf(`ATTENTION! YOUR ENTIRE RESPONSE TO THIS PROMPT NEEDS TO BE A VALID JSON...

Here is the object we need to edit:
%s

The purpose of the object is to represent: {{.Object.Context}}

RULES:
1: GENERATE DATA FOR REQUIRED FIELDS
2: UNLESS SPECIFICALLY TOLD NOT TO, GENERATE ALL FIELDS... DON'T BE LAZY.
3: OMIT ANY NON-REQUIRED FIELDS WHEN NO DATA FOR THE FIELD IS GENERATED.
4: DON'T INCLUDE FIELDS WITH EMPTY STRINGS.
5: RESPECT ANY VALIDATION INFORMATION SPECIFIED FOR FIELDS, SUCH AS MIN AND MAX LENGTHS.
6: REPLY TO THE USER PROMPT ONLY WITH THE JSON ENCODED MUTATED OBJECT
`,
		string(objectBytes),
	)

	println(system)
	println(prompt)


	resp, err := app.ChatGPT().CreateChatCompletion(
		app.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo1106,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		err = fmt.Errorf("ChatCompletion error: %v\n", err)
		return err
	}

	reply := resp.Choices[0].Message.Content
	log.Println("reply >>", reply)

	newResults := []interface{}{}
	if err := json.Unmarshal([]byte(reply), &newResults); err != nil {
		return err
	}

	for _, r := range newResults {
		result, ok := r.(map[string]interface{})
		if !ok {
			return errors.New("array item is not a map")
		}
		// remove any empty fields
		for k, v := range result {
			vv, ok := v.(string)
			if ok && vv == "" {
				delete(result, k)
			}
		}
		object := New{{uppercase .Object.Name}}(parent, Fields{{uppercase .Object.Name}}{})
		if err := object.ValidateObject(result); err != nil {
			return err
		}
		if err := app.CreateDocument{{uppercase .Object.Name}}(parent, object); err != nil {
			return err
		}
		app.SendMessageToUser(user, &Message{Type: "async-create", Body: object})
	}

	return nil
}