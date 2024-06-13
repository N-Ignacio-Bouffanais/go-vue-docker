package chatbot

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"go-vue-docker/config"

	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	userLastCommand map[int]map[string]time.Time
	mu              sync.Mutex
)

func StartBot(cfg config.Config) {
	userLastCommand = make(map[int]map[string]time.Time)

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	//bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			handleUpdate(bot, update.Message, cfg)
		}
	}
}

func handleUpdate(bot *tgbotapi.BotAPI, message *tgbotapi.Message, cfg config.Config) {
	switch message.Command() {
	case "comandos":
		handleCommands(bot, message)
	case "mle":
		handleMLE(bot, message, cfg)
	case "test":
		handleTest(bot, message, cfg)
	default:
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Comando no reconocido. Por favor, intenta nuevamente."))
	}
}

func handleCommands(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	commands := map[string]string{
		"/mle":  "Grafana MLE FLR",
		"/test": "Pagina de ejemplo",
	}

	var reply strings.Builder
	reply.WriteString("Comandos disponibles:\n")
	for cmd, desc := range commands {
		reply.WriteString(cmd)
		reply.WriteString(" - ")
		reply.WriteString(desc)
		reply.WriteString("\n")
	}
	bot.Send(tgbotapi.NewMessage(message.Chat.ID, reply.String()))
}

func handleMLE(bot *tgbotapi.BotAPI, message *tgbotapi.Message, cfg config.Config) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	command := strings.ToLower(strings.ReplaceAll(message.Text, " ", ""))

	mu.Lock()
	if userLastCommand[int(message.From.ID)] == nil {
		userLastCommand[int(message.From.ID)] = make(map[string]time.Time)
	}
	lastExecTime := userLastCommand[int(message.From.ID)]["/mle"]
	mu.Unlock()

	if time.Since(lastExecTime).Seconds() < 15 {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Debes esperar al menos 15 segundos entre ejecuciones de este comando."))
		return
	}

	if command == "/mle" {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Tomando captura"))

		screenshot, err := MLE(ctx, "http://10.115.43.118:3008/il/grafana/login", cfg.User, cfg.Password)
		if err != nil {
			log.Printf("Error al tomar captura de pantalla: %v", err)
			return
		}
		PhotosResponse(screenshot, message, bot)
	} else {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Comando no reconocido. Por favor, intenta nuevamente."))
	}

	mu.Lock()
	userLastCommand[int(message.From.ID)]["/mle"] = time.Now()
	mu.Unlock()
}

func handleTest(bot *tgbotapi.BotAPI, message *tgbotapi.Message, cfg config.Config) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	command := strings.ToLower(strings.ReplaceAll(message.Text, " ", ""))

	mu.Lock()
	if userLastCommand[int(message.From.ID)] == nil {
		userLastCommand[int(message.From.ID)] = make(map[string]time.Time)
	}
	lastExecTime := userLastCommand[int(message.From.ID)]["/test"]
	mu.Unlock()

	if time.Since(lastExecTime).Seconds() < 15 {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Debes esperar al menos 15 segundos entre ejecuciones de este comando."))
		return
	}

	if command == "/test" {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Tomando captura"))

		screenshot, err := Test(ctx, "https://www.youtube.com/", cfg.User, cfg.Password)
		if err != nil {
			log.Printf("Error al tomar captura de pantalla: %v", err)
			return
		}
		PhotosResponse(screenshot, message, bot)
	} else {
		bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Comando no reconocido. Por favor, intenta nuevamente."))
	}

	mu.Lock()
	userLastCommand[int(message.From.ID)]["/test"] = time.Now()
	mu.Unlock()
}

func PhotosResponse(screenshot []byte, m *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	// Send the screenshot as a photo
	photo := tgbotapi.FileBytes{Name: "screenshot.png", Bytes: screenshot}
	photoMsg := tgbotapi.NewPhoto(m.Chat.ID, photo)
	bot.Send(photoMsg)
}

func MLE(ctx context.Context, url, user, password string) ([]byte, error) {
	var buf []byte

	task := chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(1920, 1080, 1, false),
		chromedp.Navigate(url),
		chromedp.WaitVisible("input[name=password]", chromedp.BySearch),
		chromedp.SendKeys("input[name=user]", user, chromedp.BySearch),
		chromedp.SendKeys("input[name=password]", password, chromedp.BySearch),
		chromedp.Click(`button[aria-label="Login button"]`, chromedp.BySearch),
		chromedp.WaitVisible("body", chromedp.BySearch),
		chromedp.Sleep(1 * time.Second),
		chromedp.Navigate("http://10.115.43.118:3008/il/grafana/?orgId=1"),
		chromedp.Navigate("http://10.115.43.118:3008/il/grafana/d/sDmADcSIk/mle-flr?orgId=1&refresh=30s"),
		chromedp.WaitVisible("body", chromedp.BySearch),
		chromedp.Sleep(6 * time.Second),
		chromedp.FullScreenshot(&buf, 90),
	}

	err := chromedp.Run(ctx, task)
	if err != nil {
		log.Fatal(err)
	}

	return buf, nil
}

func Test(ctx context.Context, url, user, password string) ([]byte, error) {
	var buf []byte

	task := chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(1920, 1080, 2, false),
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.Sleep(3),
		chromedp.FullScreenshot(&buf, 100),
	}

	err := chromedp.Run(ctx, task)
	if err != nil {
		log.Fatal(err)
	}

	return buf, nil
}
