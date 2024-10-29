package handlers

import (
	"log"
	"strings"
	"time"

	"github.com/aleksjjk/hookah-bot/internal/localization"
	"github.com/aleksjjk/hookah-bot/internal/models"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Handler struct {
  bot         *tgbotapi.BotAPI
  UserState   map[int64]string
  UserData    map[int64]*models.User
  AdminChatID int64
}
var(
  messages = localization.Russian
  changeFlavours bool = false
)

func NewHandler(bot *tgbotapi.BotAPI, adminChatID int64) *Handler {
  return &Handler{
    bot:         bot,
    UserState:   make(map[int64]string),
    UserData:    make(map[int64]*models.User),
    AdminChatID: adminChatID,
  }
}

func (h *Handler) HandleUpdate(update tgbotapi.Update) {
  if update.Message == nil {
    return
  }

  chatID := update.Message.Chat.ID

  // Обработка состояний
  switch h.UserState[chatID] {
  case "":
    h.UserState[chatID] = "waiting_for_name"
    h.sendMessage(chatID,messages.WelcomeMessage + messages.AskName)
  case "waiting_for_name":
    user := h.UserData[chatID]
  	if user == nil {
    	user = &models.User{}
    	h.UserData[chatID] = user
  	}
    h.UserData[chatID].Name = update.Message.Text
    h.UserState[chatID] = "waiting_for_address"
    h.sendMessage(chatID,h.UserData[chatID].Name + messages.AskAddress)
  case "waiting_for_address":
    h.UserData[chatID].Address = update.Message.Text
    h.UserState[chatID] = "waiting_for_phone"
    h.sendMessage(chatID, h.UserData[chatID].Name + messages.AskPhone)
  case "waiting_for_phone":
    if !localization.PhoneRegexp.MatchString(update.Message.Text) {
      h.sendMessage(chatID, messages.CorrectingPhone)
      return
    }
    h.UserData[chatID].Phone = update.Message.Text
    h.UserState[chatID] = "waiting_for_delivery"
    h.askForDeliveryDate(chatID)
  case "waiting_for_delivery":
    switch update.Message.Text {
    case "На сегодня":
        h.UserState[chatID] = "today_delivery_time"
        h.askTodayDeliveryTime(chatID)
    case "На завтра":
        h.UserState[chatID] = "tomorrow_delivery_time"
        h.askTomorrowDeliveryTime(chatID)
    case "На этой неделе":
        h.UserState[chatID] = "this_week_delivery_details"
        h.askThisWeekDeliveryDetails(chatID)
    case "Другое":
        h.UserState[chatID] = "other_delivery_info"
        msg := tgbotapi.NewMessage(chatID, messages.AskForMoreDetails)
        msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // Убираем клавиатуру
        h.bot.Send(msg)
    default:
        h.sendMessage(chatID, messages.WrongChoice)
    }
  case "today_delivery_time":
    // Handle delivery time for today
    // You can save the input if needed and move to confirmation state
    h.UserData[chatID].Delivery = strings.ToLower(update.Message.Text + " " + messages.Today)// Assuming you have a DeliveryTime field
    h.UserState[chatID] = "waiting_for_tobacco_choice"
    h.sendMenu(chatID)
  case "tomorrow_delivery_time":
    // Handle delivery time for tomorrow
    h.UserData[chatID].Delivery = strings.ToLower(update.Message.Text+ " " + messages.Tomorrow)
    h.UserState[chatID] = "waiting_for_tobacco_choice"
    h.sendMenu(chatID)
  case "this_week_delivery_details":
    // Handle delivery time for this week
    h.UserData[chatID].Delivery = strings.ToLower(update.Message.Text + " " + messages.ThisWeek)
    h.UserState[chatID] = "waiting_for_tobacco_choice"
    h.sendMenu(chatID)
  case "other_delivery_info":
    // Handle additional details for delivery
    h.UserData[chatID].Delivery = update.Message.Text
    h.UserState[chatID] = "waiting_for_tobacco_choice"
    h.sendMenu(chatID)
  case "waiting_for_tobacco_choice":
    h.Flavour(chatID,update,changeFlavours)

case "waiting_for_additional_tobacco":
    h.Flavour2(chatID,update,changeFlavours)

case "waiting_for_flavor_choice":
    h.Flavour3(chatID,update,changeFlavours)

  case "confirm_order":
    if update.Message.Text == messages.Yes {
      username := update.Message.From.UserName
      h.sendOrderConfirmation(chatID,username)
    } else if update.Message.Text == messages.No {
      h.UserState[chatID] = "choosing_correction"
      h.askForCorrection(chatID)
    }
  case "choosing_correction":
    // Обработка выбора параметра для изменения
    switch update.Message.Text {
    case messages.CorrectName:
      h.UserState[chatID] = "correcting_name"
      h.sendMessage(chatID, messages.AskName)
    case messages.CorrectAddress:
      h.UserState[chatID] = "correcting_address"
      h.sendMessage(chatID, h.UserData[chatID].Name + messages.AskAddress)
    case messages.CorrectPhone:
      h.UserState[chatID] = "correcting_phone"
      h.sendMessage(chatID, h.UserData[chatID].Name+ messages.AskPhone)
    case messages.CorrectFlavours:
      h.UserData[chatID].SelectedTobaccos = []string{}
      changeFlavours= true
      h.UserState[chatID] = "correcting1"
      h.sendMenu(chatID)
    default:
      h.sendMessage(chatID, messages.WrongChoice)
    }
  
  case "correcting_name":
    if update.Message.Text == messages.No || update.Message.Text == messages.Yes {
      h.askForCorrection(chatID)
      return
    }
    h.UserData[chatID].Name = update.Message.Text
    h.UserState[chatID] = "confirm_order"
    h.confirmOrder(chatID)
  case "correcting_address":
    if update.Message.Text == messages.No || update.Message.Text == messages.Yes {
      h.askForCorrection(chatID)
      return
    }
    h.UserData[chatID].Address = update.Message.Text
    h.UserState[chatID] = "confirm_order"
    h.confirmOrder(chatID)
  case "correcting_phone":
    if update.Message.Text == messages.No || update.Message.Text == messages.Yes {
      h.askForCorrection(chatID)
      return
    }
    if !localization.PhoneRegexp.MatchString(update.Message.Text) {
      h.sendMessage(chatID, messages.CorrectingPhone)
      return
    }
    h.UserData[chatID].Phone = update.Message.Text
    h.UserState[chatID] = "confirm_order"
    h.confirmOrder(chatID)
  case "correcting1":
    h.Flavour(chatID,update,changeFlavours)

case "correcting2":
    h.Flavour2(chatID,update,changeFlavours)

case "correcting3":
    h.Flavour3(chatID,update,changeFlavours) 

  case "new_order":
    h.UserState[chatID] = "waiting_for_delivery"
    h.askForDeliveryDate(chatID)
  default:
    h.sendMessage(chatID, messages.WrongChoice)
  }
}

func (h *Handler) sendMessage(chatID int64, text string) {
  msg := tgbotapi.NewMessage(chatID, text)
  _, err := h.bot.Send(msg)
  if err != nil {
    log.Printf("Ошибка при отправке сообщения: %v", err)
  }
}




func (h *Handler) Flavour2(chatID int64,update tgbotapi.Update,changeFlavour bool){
  if update.Message.Text == messages.Yes {
        // Если пользователь хочет дополнительную забивку, показываем варианты
        h.sendFlavorOptions(chatID, "Ask for additional tobacco")
        // Устанавливаем состояние ожидания выбора дополнительного табака
        if changeFlavour{
          h.UserState[chatID] = "correcting3" 
          return
        } else{
          h.UserState[chatID] = "waiting_for_flavor_choice"
          return
        }
        
    } else if update.Message.Text == messages.No {
        // Если пользователь не хочет дополнительную забивку
        if !changeFlavour{
          h.sendMessage(chatID, messages.NoAdditionalTobaccoDetails)
          time.Sleep(2*time.Second)
        }
        username := update.Message.From.UserName
        if exist,_:=h.userExists(username);!exist{
          h.sendMessage(chatID, messages.TenPercentsFor2)
        }else if exist{
          h.sendMessage(chatID, messages.FivePercentsFor2)
        }
        // Переходим к подтверждению заказа
        h.UserState[chatID] = "confirm_order"
        h.confirmOrder(chatID) // Вызываем функцию для подтверждения заказа
    }
}
func (h *Handler) Flavour(chatID int64,update tgbotapi.Update, changeFlavour bool){
      // Проверяем, выбрано ли 2 табака
    if len(h.UserData[chatID].SelectedTobaccos) < 2 {
        // Добавляем выбранный табак
        h.UserData[chatID].SelectedTobaccos = append(h.UserData[chatID].SelectedTobaccos, update.Message.Text)
        
        // Проверяем, выбрано ли 2 табака
        if len(h.UserData[chatID].SelectedTobaccos) == 2 {
            // Запрашиваем о дополнительной забивке
            h.askForAdditionalTobacco(chatID,changeFlavours)
            // Здесь мы не добавляем еще один табак, потому что мы ожидаем ответа пользователя о дополнительной забивке
        } else {
            // Запрашиваем второй табак
            h.sendMessage(chatID, messages.SecondTobacco)
        }
    }
}
func (h *Handler) Flavour3(chatID int64, update tgbotapi.Update, changeFlavour bool) {
    // Добавляем выбранный дополнительный табак
    h.UserData[chatID].SelectedTobaccos = append(h.UserData[chatID].SelectedTobaccos, update.Message.Text)


    if !changeFlavour {
        h.sendMessage(chatID, messages.YesAdditionalTobaccoDetails)
        time.Sleep(2 * time.Second)
    }
    
    // Получаем имя пользователя
    username := update.Message.From.UserName

    // Проверка наличия пользователя и отправка скидки
    if exist, _ := h.userExists(username); !exist {
        h.sendMessage(chatID, messages.TenPercentsFor3)
    } else {
        h.sendMessage(chatID, messages.FivePercentsFor3)
    }
    
    // Переход к подтверждению заказа
    h.UserState[chatID] = "confirm_order"
    h.confirmOrder(chatID)
}


