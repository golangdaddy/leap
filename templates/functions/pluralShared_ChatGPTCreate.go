package main

import (
	"fmt"
	"log"
	"errors"
	"encoding/json"

	"github.com/sashabaranov/go-openai"
)

func (app *App) {{lowercase .Object.Name}}ChatGPTCreate(user *User, parent *{{uppercase .Object.Name}}, prompt string) error {

	fmt.Println("prompt with parent", parent.Meta.ID, prompt)

	system := `Your role is a helpful preprocessor that follows the prompt to create one or more JSON objects, ultimately outputting raw valid JSON array.

We want to create one or more of these data objects: 
// {{.Object.Context}}
{
{{range .Object.Fields}}
	// {{.Context}} {{if .Required}} (THIS FIELD IS REQUIRED){{end}}
	{{lowercase .Name}} ({{lowercase .Type}})
{{end}}
}

The response should be a raw JSON array with one or more objects, based on the user prompt: `

	println(prompt)

	resp, err := app.ChatGPT().CreateChatCompletion(
		app.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo1106,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: system+prompt,
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
	replyBytes := []byte(reply)
	if err := json.Unmarshal(replyBytes, &newResults); err != nil {
		newResult := map[string]interface{}{}
		if err := json.Unmarshal(replyBytes, &newResult); err != nil {
			return err
		}
		newResults = append(newResults, newResult)
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
		object := user.New{{uppercase .Object.Name}}(&parent.Meta, Fields{{uppercase .Object.Name}}{})
		if err := object.ValidateObject(result); err != nil {
			return err
		}
		if err := app.CreateDocument{{uppercase .Object.Name}}(&parent.Meta, object); err != nil {
			return err
		}
		app.SendMessageToUser(user, "create", object)
	}

	return nil
}
