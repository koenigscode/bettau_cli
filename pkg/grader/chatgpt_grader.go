package grader

import (
	"context"
	"encoding/json"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

type ChatGPTGrader struct {
	Client *openai.Client
}

func NewChatGPTGrader(token string) ChatGPTGrader {
	return ChatGPTGrader{Client: openai.NewClient(token)}
}

func (g *ChatGPTGrader) Grade(q GradeQuery) GradeResult {
	ctx := context.Background()
	params := jsonschema.Definition{
		Type: jsonschema.Object,
		Properties: map[string]jsonschema.Definition{
			"correct": {
				Type:        jsonschema.Boolean,
				Description: "The translation is correct and without grammar mistakes.",
			},
			"solution": {
				Type:        jsonschema.String,
				Description: "The correct translation.",
			},
			"feedback": {
				Type:        jsonschema.String,
				Description: "Explain to the user why his translation is wrong.",
			},
		},
		Required: []string{"correct", "solution"},
	}
	functionDefinition := openai.FunctionDefinition{
		Name:        "grade",
		Description: "grade the user's solution",
		Parameters:  params,
	}
	tool := openai.Tool{
		Type:     openai.ToolTypeFunction,
		Function: functionDefinition,
	}

	dialogue := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: "You are a language learning app, telling the user if their translation is correct and provide them with solution and feedback. The user speaks English and learns Swedish. Always pay attention to proper capitalization and punctuation."},
		{Role: openai.ChatMessageRoleSystem, Content: fmt.Sprintf("The user translated %s to %s", q.Question, q.Input)},
	}

	resp, err := g.Client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: dialogue,
			Tools:    []openai.Tool{tool},
		},
	)
	if err != nil {
		panic(err)
	}

	var gradeResult GradeResult
	jsonData := resp.Choices[0].Message.ToolCalls[0].Function.Arguments

	err = json.Unmarshal([]byte(jsonData), &gradeResult)
	if err != nil {
		panic(err)
	}

	return gradeResult

}
