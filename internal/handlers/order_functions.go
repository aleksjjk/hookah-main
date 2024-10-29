package handlers



import (
	"fmt"
	"time"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Функция для подтверждения заказа
func (h *Handler) confirmOrder(chatID int64) {
  time.Sleep(2*time.Second)
  orderDetails := fmt.Sprintf(messages.OrderDetails,
    h.UserData[chatID].Name,
    h.UserData[chatID].Address,
    h.UserData[chatID].Phone,
    h.UserData[chatID].Delivery,
    h.UserData[chatID].SelectedTobaccos,
  )
  h.sendMessage(chatID, orderDetails)

  // Клавиатура с кнопками "Да" и "Нет"
  msg := tgbotapi.NewMessage(chatID, messages.AskIfCorrect)
  msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
      tgbotapi.NewKeyboardButton(messages.Yes),
      tgbotapi.NewKeyboardButton(messages.No),
    ),
  )
  h.bot.Send(msg)
}

// Функция для изменения данных
func (h *Handler) askForCorrection(chatID int64) {

  // Клавиатура с кнопками для исправления данных
  msg := tgbotapi.NewMessage(chatID, h.UserData[chatID].Name+messages.WhatToCorrect)
  msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
      tgbotapi.NewKeyboardButton(messages.CorrectName),
      tgbotapi.NewKeyboardButton(messages.CorrectAddress),
      tgbotapi.NewKeyboardButton(messages.CorrectPhone),
      tgbotapi.NewKeyboardButton(messages.CorrectFlavours),
    ),
  )
  h.UserState[chatID] = "choosing_correction" // Состояние для выбора параметра
  h.bot.Send(msg)
}

// Отправка подтверждения заказа админу и пользователю
func (h *Handler) sendOrderConfirmation(chatID int64,username string ) {
  if username == "" {
      username = "Не указан"
  }

  orderDetails := fmt.Sprintf(messages.OrderDetailsForAdmin,
    username,
    h.UserData[chatID].Name,
    h.UserData[chatID].Address,
    h.UserData[chatID].Phone,
    h.UserData[chatID].Delivery,
    h.UserData[chatID].SelectedTobaccos,

  )
  msg := tgbotapi.NewMessage(chatID,h.UserData[chatID].Name+messages.OrderConfirmed )
  msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) 
  h.bot.Send(msg)

  h.sendMessage(h.AdminChatID, orderDetails)
  if exist,_:=h.userExists(username);!exist{
    h.addUserToDatabase(username)}
  

  // Очищаем данные заказа, но сохраняем выбранный язык
  h.UserState[chatID] = "new_order"
  changeFlavours = false
  h.handleNewOrder(chatID)
  

  
  // Кнопка для начала нового заказа
  time.Sleep(10*time.Second)
  msg = tgbotapi.NewMessage(chatID,h.UserData[chatID].Name + messages.NewOrder)
  msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
      tgbotapi.NewKeyboardButton(messages.StartNewOrder),
    ),
  )
  h.bot.Send(msg)
  h.handleNewOrder(chatID)
}

// Метод для обработки нового заказа
func (h *Handler) handleNewOrder(chatID int64) {
    // Сбрасываем данные заказа, оставляя личную информацию
    user := h.UserData[chatID]
    user.ResetOrderData() // Сбрасываем данные о заказе
}