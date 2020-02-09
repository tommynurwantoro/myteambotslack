package utility

import (
	"fmt"
)

// Start _
func Start() string {
	return "Halo!!\n" +
		"Bot ini punya beberapa fitur yang bisa kamu pakai untuk tim kamu.\n" +
		"Coba gunakan command `help` untuk melihat perintah-perintah yang tersedia ya.\n"
}

// Help _
func Help(commands string) string {
	return "Kamu bisa gunakan perintah-perintah ini loh:\n" + commands
}

// InvalidCommand _
func InvalidCommand() string {
	return "Aku gak ngerti perintah itu, coba perintah yang lain ya."
}

// InvalidParameter _
func InvalidParameter() string {
	return "Parameternya belum bener tuh, coba dicek lagi ya"
}

// SuccessInsertData _
func SuccessInsertData() string {
	return "OK!"
}

// SuccessUpdateData _
func SuccessUpdateData() string {
	return "Updated!"
}

// InvalidSequece _
func InvalidSequece() string {
	return "Gak bisa, gak ada di list"
}

func GreetingNewJoinedUser(username string) string {
	return fmt.Sprintf("Welcome @%s!!! GLHF 😁", username)
}

func CustomCommandNotFound() string {
	return "Belum ada custom command nih, pakai command /simpan_command dulu aja"
}
