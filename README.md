# Hookah Rental Telegram Bot

This is a Telegram bot written in Go that allows users to order hookahs for delivery. The bot supports multiple languages (Russian, English, and Czech) and offers features such as tobacco flavor selection, delivery scheduling, and user order confirmation.

## Features

- Multi-language support (Russian, English, Czech)
- Tobacco flavor selection
- Delivery date and time scheduling
- User data collection (name, address, phone)
- Order confirmation with optional corrections
- Admin notification of new orders
- Database storage for user information (SQLite)

## Technologies Used

- **Go**: Main programming language
- **Telegram Bot API**: Used for interacting with Telegram users
- **SQLite**: Database for storing user data
- **tgbotapi**: Go library for working with Telegram bots

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install) (version 1.17 or later)
- [SQLite](https://www.sqlite.org/download.html)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/hookah-rental-bot.git
   cd hookah-rental-bot
2. Install dependencies: go mod download
3. Create a Telegram bot and obtain an API token.
4. Set up environment variables:
Create a .env file and add your Telegram API token:
  TELEGRAM_TOKEN=your-telegram-bot-token
  ADMIN_CHAT_ID=your-admin-chat-id
5. Database Setup
Make sure that you have an SQLite database file named hookah_bot.db in the project directory with the following structure:  
  CREATE TABLE users (
  id TEXT PRIMARY KEY,
  );
6. How It Works
Language Selection: When a user starts the bot, they choose a language.
Order Process: The bot collects the user’s name, address, phone number, and preferred delivery time.
Tobacco Selection: Users can select one or more tobacco flavors.
Order Confirmation: The user is asked to confirm their order. If any details are incorrect, they can update them.
Admin Notification: Once the user confirms the order, the admin receives the details.
Database Storage: If the user is placing an order for the first time, their information is stored in the SQLite database.
7. License
MIT License

Copyright (c) [2024] [aleksjjk]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

