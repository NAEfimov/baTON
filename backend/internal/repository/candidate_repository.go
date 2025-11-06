package repository

import (
	"baton/backend/internal/database"
	"baton/backend/internal/models"
	"database/sql"
	"encoding/json"
)

func CreateCandidate(
	telegramID int64,
	username string,
	name string,
	skills []string,
	matchingScore float64,
	years int,
	education string,
	experienceJSON string,
	location string,
) (int64, error) {
	bSkills, _ := json.Marshal(skills)
	res, err := database.DB.Exec(
		`INSERT OR REPLACE INTO candidates (telegram_id, username, name, skills, matching_score, years, education, experience, location)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		telegramID,
		username,
		name,
		string(bSkills),
		matchingScore,
		years,
		education,
		experienceJSON,
		location,
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetCandidateByTelegramID(telegramID int64) (*models.Candidate, error) {
    var c models.Candidate
    var skillsJSON, expJSON string

    err := database.DB.QueryRow(
        `SELECT telegram_id, username, name, skills, matching_score, years, education, experience, location
         FROM candidates WHERE telegram_id = ? LIMIT 1`,
        telegramID,
    ).Scan(&c.TelegramID, &c.Username, &c.Name, &skillsJSON, &c.MatchingScore, &c.Years, &c.Education, &expJSON, &c.Location)
    if err == sql.ErrNoRows {
        return nil, nil 
    }
    if err != nil {
        return nil, err
    }
    _ = json.Unmarshal([]byte(skillsJSON), &c.Skills)
    _ = json.Unmarshal([]byte(expJSON), &c.Experience)
    return &c, nil
}

func GetAllCandidates() ([]*models.Candidate, error) {
    rows, err := database.DB.Query(
        `SELECT telegram_id, username, name, skills, matching_score, years, education, experience, location FROM candidates`,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var candidates []*models.Candidate
    for rows.Next() {
        var c models.Candidate
        var skillsJSON, expJSON string
        if err := rows.Scan(&c.TelegramID, &c.Username, &c.Name, &skillsJSON, &c.MatchingScore, &c.Years, &c.Education, &expJSON, &c.Location); err != nil {
            continue
        }
        _ = json.Unmarshal([]byte(skillsJSON), &c.Skills)
        _ = json.Unmarshal([]byte(expJSON), &c.Experience)
        candidates = append(candidates, &c)
    }
    return candidates, nil
}