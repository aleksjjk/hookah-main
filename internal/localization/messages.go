package localization
import (
	"regexp"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
const OrderAdmin = "Данные клиента:\nID клиента: @%s\nИмя: %s\nАдрес: %s\nТелефон:%s \nДоставка:%s \nТовары:%s"

// Структура для хранения сообщений на разных языках
type Messages struct {
	WelcomeMessage string
	AskName         string
	AskAddress      string
	AskPhone        string
	InvalidPhone   string 
	OrderConfirmed  string
	OrderDetails    string
	OrderDetailsForAdmin string
	NewOrder        string
	AskIfCorrect    string
	WhatToCorrect   string
	CorrectName     string
	CorrectAddress  string
	CorrectPhone    string
	CorrectDate string
	CorrectFlavours string
	StartNewOrder   string
	Yes string
	No string
	CorrectingPhone string
	WrongChoice string
	AskDeliveryDate string
	AskTimeForToday string
	AskTimeForTomorrow string
	AskDetailsForThisWeek string
	AskForMoreDetails string
	AskForAdditionalTobacco string
	FirstTobacco string 
	SecondTobacco string
	ChooseTobacco string 
	YesAdditionalTobaccoDetails string
	NoAdditionalTobaccoDetails string
	TenPercentsFor3 string
	TenPercentsFor2 string
	FivePercentsFor3 string
	FivePercentsFor2 string
	Today string
	Tomorrow string
	ThisWeek string
	Other string



}

// Сообщения на русском языке
var Russian = Messages{
	WelcomeMessage: "Добрый день, вас приветствует бот Pari Prague 👋.\n",
	AskName:        "Пожалуйста, введите ваше имя:",
	AskAddress:     ", пожалуйста, введите адрес доставки:",
	AskPhone:       ", пожалуйста, введите ваш номер телефона в формате +420XXXXXXXXX:",
	InvalidPhone:   "Неверный номер телефона. Пожалуйста, введите номер в формате +420XXXXXXXXX:", // Новое сообщение
	OrderConfirmed: ", спасибо! Ваш заказ подтвержден. Мы свяжемся с вами скоро 💌.",
	OrderDetails:   "Ваши данные:\nИмя: %s\nАдрес: %s\nТелефон: %s \nИнформация о доставке: %s\nВыбранные вкусы: %s\n",
	OrderDetailsForAdmin:OrderAdmin,
	NewOrder:       ", хотите сделать новый заказ?",
	AskIfCorrect:   "Проверьте данные: всё верно?",
	WhatToCorrect:  ", что вы хотите исправить?",
	CorrectName:    "Имя",
	CorrectAddress: "Адрес",
	CorrectPhone:   "Телефон",
	CorrectFlavours:"Вкусы",
	StartNewOrder:  "Новый заказ",
	Yes: "Да",
	No: "Нет",
  CorrectingPhone: "Пожалуйста, введите номер телефона в формате +420XXXXXXXXX (9 цифр):",
	WrongChoice: "Неверный выбор.",
	AskDeliveryDate:"На какой день вы бы хотели заказать доставку?",
	AskTimeForToday: "В какое время вы хотите доставку сегодня?",
	AskTimeForTomorrow: "Во сколько вы хотите доставку завтра?",
	AskDetailsForThisWeek: "Какое время и день вы предпочитаете на этой неделе?",
	AskForMoreDetails: "Пожалуйста, укажите более точную информацию.",
	AskForAdditionalTobacco : "Вы бы хотели дополнительную забивку табака?",
	FirstTobacco: "В комплект входит 2 забивки табака.Выберите, пожалуйста, ваш первый вкус.",
  SecondTobacco: "Выберите, пожалуйста, ваш следующий вкус.",
	ChooseTobacco: "Давайте перейдём к выбору наших вкусов.",
	YesAdditionalTobaccoDetails:`🚗 Курьер привезёт комплект кальяна на ваш адрес в указанное время

💎 В комплект входит: 
- кальян 
- плита для разогрева углей 
- калауд 
- 3 забивки табака (на Ваш выбор)
- 12 углей

Стоимость комплекта + доставка на ваш адрес 1000 Kč`,
NoAdditionalTobaccoDetails: `🚗 Курьер привезёт комплект кальяна на ваш адрес в указанное время

💎 В комплект входит: 
- кальян 
- плита для разогрева углей 
- калауд 
- 2 забивки табака (на Ваш выбор)
- 8 углей

Стоимость комплекта + доставка на ваш адрес 900 Kč`,
TenPercentsFor3: "У вас есть скидка на первый заказ - 10% ✅️. Итоговая стоимость комплекта + доставка 900 Kč",
TenPercentsFor2:"У вас есть скидка на первый заказ - 10% ✅️. Итоговая стоимость комплекта + доставка 810 Kč" ,
FivePercentsFor3:"У вас есть скидка на заказ - 5% ✅️. Итоговая стоимость комплекта + доставка 950 Kč" ,
FivePercentsFor2: "У вас есть скидка на заказ - 5% ✅️. Итоговая стоимость комплекта + доставка 855 Kč",
Today: "На сегодня",
Tomorrow: "На завтра",
ThisWeek: "На этой неделе",
Other: "Другое",
}

var KeyboardFlavor = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
      tgbotapi.NewKeyboardButton("1.Forest Smoke🌳💨"),
      tgbotapi.NewKeyboardButton("2.Melon Kiss🍈💋"),
      tgbotapi.NewKeyboardButton("3.Berries Flow🍓🏝️🍍🥭"),
      tgbotapi.NewKeyboardButton("4.Madagascar🍍🥭"),
      tgbotapi.NewKeyboardButton("5.Irish creme☕🍮"),
      tgbotapi.NewKeyboardButton("6.Berries Flow(strawberry)🍓"),
      tgbotapi.NewKeyboardButton("7.Sour Exotic🍍"),
      tgbotapi.NewKeyboardButton("8.Juicy Story🍊🍏🍇"),
        ),
      )

 	var PhoneRegexp = regexp.MustCompile(`^\+420\d{9}$`) // Регулярное выражение для проверки номера) // Регулярное выражение для проверки номера