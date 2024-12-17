package app

import (
	"log"
	"net/http"

	"github.com/wDRxxx/yandex_calculator_api/internal/models"
	"github.com/wDRxxx/yandex_calculator_api/internal/service"
	"github.com/wDRxxx/yandex_calculator_api/internal/service/calculator"
	"github.com/wDRxxx/yandex_calculator_api/pkg/utils"
)

type App struct {
	port              string
	httpServer        *http.ServeMux
	calculatorService service.CalculatorService
}

func (a *App) CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.Input
	err := utils.ReadJSON(r.Body, &input)
	if err != nil {
		utils.WriteJSON(
			&models.ErrorOutput{Error: "Expression is not valid"},
			w,
			http.StatusBadRequest,
		)
	}

	result, err := a.calculatorService.Calculate(input.Expression)
	if err != nil {
		if err.Error() == "wrong expression format" {
			utils.WriteJSON(
				&models.ErrorOutput{Error: "Expression is not valid"},
				w,
				http.StatusBadRequest,
			)
			return
		} else {
			utils.WriteJSON(
				&models.ErrorOutput{Error: "Internal server error"},
				w,
				http.StatusInternalServerError,
			)
			return
		}
	}

	utils.WriteJSON(
		&models.SuccessOutput{Result: result},
		w,
		http.StatusOK,
	)
}

func NewApp(port string) *App {
	app := &App{
		port:              port,
		calculatorService: calculator.NewCalculatorService(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", app.CalculateHandler)

	app.httpServer = mux

	return app
}

func (a *App) Start() error {
	log.Printf("Starting server on %s", "127.0.0.1:"+a.port)
	err := http.ListenAndServe("127.0.0.1:"+a.port, a.httpServer)
	if err != nil {
		return err
	}

	return nil
}
