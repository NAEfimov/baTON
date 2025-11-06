package repository

import (
	"baton/backend/internal/database"
	"baton/backend/internal/models"
	"encoding/json"
)

func CreateVacancy(
	telegramID int64,
	username string,
	skills []string,
) (int64, error) {
	bSkills, _ := json.Marshal(skills)
	res, err := database.DB.Exec(
		`INSERT OR REPLACE INTO vacancies (telegram_id, username, skills)
         VALUES (?, ?, ?)`,
		telegramID,
		username,
		string(bSkills),
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetVacancyByTelegramID(telegramID int64) (*models.Vacancy, error) {
    var v models.Vacancy
    var skillsJSON string
    err := database.DB.QueryRow(
        `SELECT telegram_id, username, skills FROM vacancies WHERE telegram_id = ?`,
        telegramID,
    ).Scan(&v.TelegramID, &v.Username, &skillsJSON)
    if err != nil {
        return nil, err
    }
    _ = json.Unmarshal([]byte(skillsJSON), &v.Skills)
    return &v, nil
}