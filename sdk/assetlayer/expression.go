package assetlayer

import (
	"encoding/json"
)

type Expression struct {
	SlotID         string `json:"slotId"`
	ExpressionName string `json:"expressionName"`
	ExpressionID   string `json:"expressionId"`
	ExpressionType struct {
		ExpressionTypeName   string `json:"expressionTypeName"`
		ExpressionAttributes []struct {
			ExpressionAttributeName string `json:"expressionAttributeName"`
			ExpressionAttributeID   string `json:"expressionAttributeId"`
		} `json:"expressionAttributes"`
		ExpressionTypeID string `json:"expressionTypeId"`
	} `json:"expressionType"`
	Description      string `json:"description,omitempty"`
	ExpressionTypeID string `json:"expressionTypeId"`
}

func (client *Client) NewExpression(slotID, name string) (string, error) {

	expression := &Expression{
		SlotID:           slotID,
		ExpressionTypeID: "6281963ab23c7bf548942139",
		ExpressionName:   name,
		Description:      "test description",
	}

	data, err := client.Try("POST", "/api/v1/slot/expressions/new", nil, expression)
	if err != nil {
		return "", err
	}
	m, err := assertMapStringInterface(data)
	if err != nil {
		return "", err
	}
	id, err := assertString(m["expressionId"])
	if err != nil {
		return "", err
	}
	return id, nil
}

type ExpressionType struct {
	ExpressionTypeID     string `json:"expressionTypeId"`
	ExpressionTypeName   string `json:"expressionTypeName"`
	ExpressionAttributes []struct {
		ExpressionAttributeName string `json:"expressionAttributeName"`
		ExpressionAttributeID   string `json:"expressionAttributeId"`
	} `json:"expressionAttributes"`
}

func (client *Client) GetExpressionTypes() ([]*ExpressionType, error) {

	data, err := client.Try(
		"GET",
		"/api/v1/slot/expressions/types",
		nil,
	)
	if err != nil {
		return nil, err
	}

	m, err := assertMapStringInterface(data)
	if err != nil {
		return nil, err
	}
	a, err := assertInterfaceArray(m["expressionTypes"])
	if err != nil {
		return nil, err
	}

	expressions := []*ExpressionType{}

	for _, item := range a {
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		expression := &ExpressionType{}
		if err := json.Unmarshal(b, expression); err != nil {
			return nil, err
		}
		expressions = append(expressions, expression)
	}

	return expressions, nil
}
