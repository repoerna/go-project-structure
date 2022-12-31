package config

import (
	"github.com/spf13/viper"
)

type Configurations struct {
	AppEnv,
	AppName,
	AppPort,
	LogFormat,
	LogLevel,
	DBHost,
	DBPort,
	DBName,
	DBUser,
	DBPassword string
}

func New() *Configurations {

	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_PORT", "5432")

	// viper.SetDefault("JWT_SECRET", "bE8fsjU^BD$n%7")
	// viper.SetDefault("JWT_EXPIRATION", 30)
	// viper.SetDefault("PASSWORD_RESET_CODE_EXPIRATION", 15)
	// viper.SetDefault("PASSWORD_RESET_TEMPLATE_PATH", "templates/password_reset.html")
	// viper.SetDefault("MAIL_VERIFICATION_CODE_EXPIRATION", 24)
	// viper.SetDefault("MAIL_VERIFICATION_TEMPLATE_PATH", "templates/confirm_mail.html")
	// viper.SetDefault("MAIL_VERIF_TEMPLATE_ID", "d-33f5050a59604892b76f37450f476f12")
	// viper.SetDefault("PASS_RESET_TEMPLATe_ID", "d-86a003810c6341fd9d4442738e7b95a8")

	configs := &Configurations{
		AppEnv:  viper.GetString("APP_ENV"),
		AppPort: viper.GetString("APP_PORT"),

		LogLevel:  viper.GetString("LOG_LEVEL"),
		LogFormat: viper.GetString("LOG_FORMAT"),

		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		// JwtSecret:               viper.GetString("JWT_SECRET"),
		// JwtExpiration:           viper.GetInt("JWT_EXPIRATION"),
		// PassResetCodeExpiration: viper.GetInt("PASSWORD_RESET_CODE_EXPIRATION"),
		// PassResetTemplatePath:   viper.GetString("PASSWORD_RESET_TEMPLATE_PATH"),
		// MailVerifCodeExpiration: viper.GetInt("MAIL_VERIFICATION_CODE_EXPIRATION"),
		// MailVerifTemplatePath:   viper.GetString("MAIL_VERIFICATION_TEMPLATE_PATH"),
		// MailVerifTemplateID:     viper.GetString("MAIL_VERIF_TEMPLATE_ID"),
		// PassResetTemplateID:     viper.GetString("PASS_RESET_TEMPLATe_ID"),
	}
	return configs
}
