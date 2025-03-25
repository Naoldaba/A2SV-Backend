import os
from dotenv import load_dotenv
from telegram import Bot
from telegram.ext import Updater, CommandHandler, MessageHandler, Filters

load_dotenv()

BOT_TOKEN = os.getenv('BOT_TOKEN')

def start(update, context):
    update.message.reply_text("Hey Yo! I'm astranaL, How may I be in service?")

def handle_message(update, context):
    update.message.reply_text(f'You said: {update.message.text}')

def main():
    updater = Updater(TOKEN, use_context=True)

    # Dispatcher
    dispatcher = updater.dispatcher
    dispatcher.add_handler(CommandHandler("start", start))

    # Handlers
    dispatcher.add_handler(MessageHandler(Filters.text & ~Filters.command, handle_message))
    updater.start_polling()
    updater.idle()

if __name__ == '__main__':
    main()