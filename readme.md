TPGbot it's a simple bot for Telegram to logging messages in local Postgres database.
Set config in /config/config.go like this:

const (
	DB_NAME   = "your database name"
	DB_USER   = "your database login"
	DB_PASS   = "your database password"
	BOT_TOKEN = "your telegram bot API from Thebotfather"
)

And then run bot this command "go run main.go"
Also you need local postgres database with table "logs" with columns "login","text","chat","time"
"Anlayzer" contain experimental functional  to research users messages and automatically answer on it.
You can prepare answers in /analyzer/answers.json file like this:

[
   {
      "text":"Some message",
      "answer":"some answer"
   }...
]
