package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/kr/pretty"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/api/iterator"
)

func (app *App) {{lowercase .Object.Name}}ChatGPTModify(user *User, parent *{{uppercase .Object.Name}}, prompt string) error {

	println("grabbing results for prompt:", parent.Meta.Name, prompt)

	idList := []string{}
	var list []map[string]interface{}
	iter := parent.Meta.Firestore(app.App).Collection(parent.Meta.Class).OrderBy("Meta.Created", firestore.Asc).Documents(app.Context())
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err)
			break
		}
		m := map[string]interface{}{}
		if err := doc.DataTo(&m); err != nil {
			log.Println(err)
			continue
		}

		idList = append(idList, m["Meta"].(map[string]interface{})["ID"].(string))
		// prune metadata
		delete(m, "Meta")

		cleaned := map[string]interface{}{
			"_id": len(list),
		}
		for k, v := range m["fields"].(map[string]interface{}) {
			cleaned[k] = v
		}

		pretty.Println(cleaned)

		list = append(list, cleaned)
	}

	fmt.Println("prompt with parent", parent.Meta.ID, prompt)

	b, err := json.Marshal(list)
	if err != nil {
		return err
	}

	prompt = fmt.Sprintf(`ATTENTION! YOUR ENTIRE RESPONSE TO THIS PROMPT NEEDS TO BE VALID JSON...

	This JSON represents the current state of items in a database table:
%s

MY MUTATION PROMPT: %s

REPLY ONLY WITH A JSON ENCODED ARRAY OF THE END RESULT
`,
		string(b),
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
		return fmt.Errorf("ChatCompletion error: %v\n", err)
	}

	reply := resp.Choices[0].Message.Content
	log.Println("reply >>", reply)

	newResults := []interface{}{}
	if err := json.Unmarshal([]byte(reply), &newResults); err != nil {
		return err
	}

	for n, result := range newResults {
		updates := []firestore.Update{
			{
				Path:  "Meta.Modified",
				Value: app.TimeNow().Unix(),
			},
			{
				Path:  "Meta.Context.Order",
				Value: n,
			},
		}
		data := result.(map[string]interface{})
		id, ok := data["_id"].(float64)
		if !ok {
			return fmt.Errorf("no _id present: %d", n)
		}
		if n != int(id) {
			log.Println("order doesn't match:", n, id)
		}
		delete(data, "_id")

		for field, value := range data {
			updates = append(updates, firestore.Update{
				Path:  "fields." + strings.ToLower(field),
				Value: value,
			})
		}
		docID := idList[int(id)]
		log.Println("updating firestore doc:", docID)
		pretty.Println(updates)

		if updateInfo, err := parent.Meta.Firestore(app.App).Collection(parent.Meta.Class).Doc(docID).Update(
			app.Context(),
			updates,
		); err != nil {
			return err
		} else {
			log.Println(updateInfo)
		}
	}

	return nil
}
