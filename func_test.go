package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func TestGeminiApi(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-and-image input (multimodal), use the gemini-pro-vision model
	model := client.GenerativeModel("gemini-pro-vision")

	pathToImage1 := "images/abhi.jpeg"
	imgData1, err := os.ReadFile(pathToImage1)
	if err != nil {
		log.Fatal(err)
	}

	pathToImage2 := "images/announcement.jpg"
	imgData2, err := os.ReadFile(pathToImage2)
	if err != nil {
		log.Fatal(err)
	}

	prompt := []genai.Part{
		genai.ImageData("jpeg", imgData1),
		genai.ImageData("jpeg", imgData2),
		genai.Text("tell me story using these images?"),
	}
	resp, err := model.GenerateContent(ctx, prompt...)

	log.Printf("%+v", resp.Candidates[0].Content.Parts[0])
	if err != nil {
		log.Fatal(err)
	}
}
