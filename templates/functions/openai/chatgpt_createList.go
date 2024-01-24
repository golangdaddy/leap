package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/golangdaddy/leap/sdk/cloudfunc"
	"github.com/kr/pretty"
	"github.com/richardboase/npgpublic/models"
	"github.com/sashabaranov/go-openai"
)

func (app *App) chatgpt_createList{{uppercase .Object.Name}}(w http.ResponseWriter, r *http.Request, parent *Internals, collection string) {

	m := map[string]interface{}{}
	if err := cloudfunc.ParseJSON(r, &m); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return
	}

	prompt, ok := models.AssertKeyValueSTRING(w, m, "prompt")
	if !ok {
		return
	}

	fmt.Println("prompt with parent", parent.ID, prompt)

	prompt = fmt.Sprintf(`
This JSON represents the schema of an item in a database table:


MY PROMPT: %s

REPLY ONLY WITH A JSON ENCODED ARRAY OF THE GENERATED OBJECTS.
`,
		prompt,
	)

	println(prompt)

	resp, err := app.ChatGPT().CreateChatCompletion(
		app.Context(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		err = fmt.Errorf("ChatCompletion error: %v\n", err)
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	reply := resp.Choices[0].Message.Content
	log.Println("reply >>", reply)

	newResults := []interface{}{}
	if err := json.Unmarshal([]byte(reply), &newResults); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

	for _, result := range newResults {
		item := result.(map[string]interface{})
		
		{{lowercase .Object.Name}} := 

		parent.New

		if updateInfo, err := parent.Firestore(app.App).Collection(collection).Doc(docID).Update(
			app.Context(),
			updates,
		); err != nil {
			cloudfunc.HttpError(w, err, http.StatusInternalServerError)
			return
		} else {
			log.Println(updateInfo)
		}
	}

	if err := cloudfunc.ServeJSON(w, newResults); err != nil {
		cloudfunc.HttpError(w, err, http.StatusInternalServerError)
		return
	}

}
