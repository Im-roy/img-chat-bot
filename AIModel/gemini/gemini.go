package gemini

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"img-chat-bot/constants"
	"img-chat-bot/model"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type GeminiAI struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func (g *GeminiAI) GetClient() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	g.client = client
	g.SetGenerativeModel("gemini-pro-vision")
}

func (g *GeminiAI) Close() {
	g.client.Close()
}

func (g *GeminiAI) SetGenerativeModel(model string) {
	g.model = g.client.GenerativeModel(model)
}

func (g *GeminiAI) GenerateContent(ctx context.Context, promptMessage string, additionalImageData []model.PromptImageModel) (string, error) {

	g.GetClient()
	if len(additionalImageData) > 0 {
		g.SetGenerativeModel(constants.GEMINI_PRO_VISION)
	} else {
		g.SetGenerativeModel(constants.GEMINI_PRO)
	}

	if g.client == nil || g.model == nil {
		return "", fmt.Errorf("error occured")
	}
	defer g.Close()
	prompt := g.buildPrompt(promptMessage, additionalImageData)
	resp, err := g.model.GenerateContent(ctx, prompt...)
	if err != nil {
		return "", err
	}
	return g.formatResponse(resp), nil
}

func (g *GeminiAI) buildPrompt(promptMessage string, additionalImageData []model.PromptImageModel) []genai.Part {
	prompt := []genai.Part{
		genai.Text(promptMessage),
	}
	for _, data := range additionalImageData {
		prompt = append(prompt, genai.ImageData(data.ExtensionName, data.Data))
	}
	return prompt
}

func (g *GeminiAI) formatResponse(resp *genai.GenerateContentResponse) string {
	var formattedContent strings.Builder
	if resp != nil && resp.Candidates != nil {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					formattedContent.WriteString(fmt.Sprintf("%v", part))
				}
			}
		}
	}

	return formattedContent.String()
}
