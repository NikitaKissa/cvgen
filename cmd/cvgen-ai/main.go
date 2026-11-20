package main

import (
	"context"
	"fmt"
	"time"

	"github.com/NikitaKissa/cvgen/internal/ai"
	aigenerate "github.com/NikitaKissa/cvgen/internal/app/ai-generate"
	"github.com/NikitaKissa/cvgen/internal/llm"
	"github.com/NikitaKissa/cvgen/internal/llm/llama"
)

func main() {
	ctx := context.Background()
	gpt := llama.New()
	defer gpt.Close()

	fmt.Println("Loading model...")
	err := gpt.LoadModel(ctx, "/home/nikita/downloads/chrome/Qwen3.5-4B-Q4_K_M.gguf")
	if err != nil {
		fmt.Println("Error loading model:", err)
		return
	}
	ctx = llm.WithContext(ctx, gpt)
	fmt.Println("Model loaded.")

	start := time.Now()

	err = aigenerate.Run(
		ctx,
		"resume",
		ai.Vacancy{
			Position: "Frontend dev",
			Company:  "INTERIA",
			Description: `Twój zakres obowiązków
			- tworzenie i rozwijanie warstwy frontowej naszych aplikacji,
			- wsparcie w procesie wdrożenia nowych wersji aplikacji na systemy produkcyjne,
			- dbałość o wysoką jakość kodu,
			- tworzenie testów jednostkowych oraz funkcjonalnych.
			Nasze wymagania
			- bardzo dobra znajomość JavaScript/ES6/Type Script – minimum 1 rok doświadczenia komercyjnego,
			- dobra znajomość React 17+, HTML5 i CSS3,
			- wiedza na temat dobrych praktyk i wzorców projektowych.
			Mile widziane
			- znajomość Node.js,
			- umiejętność korzystania z platform PAAS np. OpenShift (OKD) / Kubernetes,
			- doświadczenie w realizacji projektów w metodykach zwinnych,
			- zainteresowanie wykorzystaniem AI i automatyzacji w codziennej pracy developera.`,
		},
		ai.Options{
			Language: "Polish",
		},
		aigenerate.Options{
			OutputPath: "/tmp/testcv.html",
		},
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Time taken:", time.Since(start))
}
