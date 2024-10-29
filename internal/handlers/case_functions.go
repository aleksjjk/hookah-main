package handlers

import (
	"github.com/aleksjjk/hookah-bot/internal/localization"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Функция для запроса даты доставки
func (h *Handler) askForDeliveryDate(chatID int64) {

    // Клавиатура с вариантами выбора даты доставки
    msg := tgbotapi.NewMessage(chatID, messages.AskDeliveryDate)
    msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton(messages.Today),
            tgbotapi.NewKeyboardButton(messages.Tomorrow),
            tgbotapi.NewKeyboardButton(messages.ThisWeek),
            tgbotapi.NewKeyboardButton(messages.Other),
        ),
    )
    h.UserState[chatID] = "waiting_for_delivery" // Состояние для выбора даты доставки
    h.bot.Send(msg)
}
// Функция для запроса времени доставки на сегодня
func (h *Handler) askTodayDeliveryTime(chatID int64) {
    msg := tgbotapi.NewMessage(chatID, messages.AskTimeForToday)
    msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    h.bot.Send(msg)
}

// Функция для запроса времени доставки на завтра
func (h *Handler) askTomorrowDeliveryTime(chatID int64) {
    msg := tgbotapi.NewMessage(chatID, messages.AskTimeForTomorrow)
    msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    h.bot.Send(msg)
}

// Функция для запроса деталей доставки на этой неделе
func (h *Handler) askThisWeekDeliveryDetails(chatID int64) {
    msg := tgbotapi.NewMessage(chatID, messages.AskDetailsForThisWeek)
    msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
    h.bot.Send(msg)
}

func (h *Handler) sendMenu(chatID int64) {
    msg := tgbotapi.NewMessage(chatID, messages.ChooseTobacco)
    
    photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath("pari_menu.jpg"))
    h.bot.Send(photo)
    h.bot.Send(msg)
    h.sendFlavorOptions(chatID,"Choose tobacco")
    

}
func (h *Handler) askForAdditionalTobacco(chatID int64, changeFlavour bool) {
    msg := tgbotapi.NewMessage(chatID, messages.AskForAdditionalTobacco)
    msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
        tgbotapi.NewKeyboardButtonRow(
            tgbotapi.NewKeyboardButton(messages.Yes),
            tgbotapi.NewKeyboardButton(messages.No),
        ),
    )
    if changeFlavour{
        h.UserState[chatID] = "correcting2"
    }else {
        h.UserState[chatID] = "waiting_for_additional_tobacco"}
    h.bot.Send(msg)
}
func (h *Handler) sendFlavorOptions(chatID int64, message string) {
    // Получаем сообщения на нужном языке

    // Используем switch для обработки разных случаев
    switch message {
    case "Choose tobacco":
        // Отправляем сообщение для выбора первого табака
        msg := tgbotapi.NewMessage(chatID, messages.FirstTobacco)
        msg.ReplyMarkup = localization.KeyboardFlavor // Добавляем клавиатуру с вариантами табаков
        h.bot.Send(msg)

    case "Ask for additional tobacco":
        // Отправляем сообщение для выбора второго табака
        msg := tgbotapi.NewMessage(chatID, messages.SecondTobacco)
        msg.ReplyMarkup = localization.KeyboardFlavor // Добавляем клавиатуру с вариантами табаков
        h.bot.Send(msg)
        
    default:
        // Можно обработать случай, если сообщение не соответствует ни одному из вариантов
        msg := tgbotapi.NewMessage(chatID, messages.WrongChoice)
        h.bot.Send(msg)
    }
}