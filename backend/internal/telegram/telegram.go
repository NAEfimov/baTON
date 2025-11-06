package telegram

import (
    "baton/backend/internal/repository"
    // "baton/backend/internal/blockchain"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
)

var (
    botToken  string
    webAppURL string
)

var awaitingVerifier = make(map[int64]struct {
    expIndex int
    expName  string
})

type verificationRequest struct {
    CandidateID       int64
    CandidateUsername string
    ExperienceLabel   string
    VerifierID        int64
}

var (
    pendingVerifications = make(map[string]verificationRequest)
)

func init() {
    botToken = os.Getenv("BOT_TOKEN")
    if botToken == "" {
        log.Fatal("FATAL: BOT_TOKEN environment variable not set.")
    }

    webAppURL = os.Getenv("WEBAPP_URL")
    if webAppURL == "" {
        log.Fatal("FATAL: WEBAPP_URL environment variable not set.")
    }

    log.Printf("INFO: Telegram package initialized. Web App URL is: '%s'", webAppURL)
}

func sendReplyKeyboard(chatID int64, text string, buttons [][]map[string]interface{}) {
    payload := map[string]interface{}{
        "chat_id": chatID,
        "text":    text,
        "reply_markup": map[string]interface{}{
            "keyboard":        buttons,
            "resize_keyboard": true, // Makes the buttons fit nicely
        },
    }
    sendTelegramRequest("sendMessage", payload)
}

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()

    var update struct {
        Message *struct {
            From struct {
                ID       int64  `json:"id"`
                Username string `json:"username"`
            } `json:"from"`
            Text string `json:"text"`
        } `json:"message"`
        CallbackQuery *struct {
            ID   string `json:"id"`
            From struct {
                ID       int64  `json:"id"`
                Username string `json:"username"`
            } `json:"from"`
            Data string `json:"data"`
        } `json:"callback_query"`
    }

    if err := json.Unmarshal(body, &update); err != nil {
        log.Printf("Failed to unmarshal: %v", err)
        w.WriteHeader(200)
        return
    }

    if update.Message != nil {
        handleMessage(update.Message.From.ID, update.Message.From.Username, update.Message.Text)
    }

    if update.CallbackQuery != nil {
        handleCallbackQuery(update.CallbackQuery)
    }

    w.WriteHeader(200)
}
func handleMessage(userID int64, username, text string) {
    _ = repository.CreateUser(userID, username)

    if text == "/start" || text == "/register" {
        // This is the structure for a ReplyKeyboardMarkup button
        buttons := [][]map[string]interface{}{
            { // A single row of buttons
                { // The first button
                    "text": "Open Web App",
                    "web_app": map[string]string{
                        "url": webAppURL,
                    },
                },
            },
        }
        // --- CALL THE NEW FUNCTION HERE ---
        sendReplyKeyboard(userID, "Welcome! Click below to open the app:", buttons)
        return
    }

    if text == "/verify" {
        candidate, err := repository.GetCandidateByTelegramID(userID)
        if err != nil || candidate == nil || len(candidate.Experience) == 0 {
            sendMessage(userID, "No experience found. Please add your profile first.")
            return
        }

        buttons := [][]map[string]interface{}{}
        for i, exp := range candidate.Experience {
            buttons = append(buttons, []map[string]interface{}{
                {
                    "text":          fmt.Sprintf("%s @ %s", exp.Vacancy, exp.Company),
                    "callback_data": fmt.Sprintf("verify:%d", i),
                },
            })
        }
        sendInlineKeyboard(userID, "Select experience to verify:", buttons)
        return
    }

    if info, ok := awaitingVerifier[userID]; ok {
        if !strings.HasPrefix(text, "@") {
            sendMessage(userID, "Please send verifier's username starting with @")
            return
        }
        verifierUsername := strings.TrimPrefix(text, "@")
        delete(awaitingVerifier, userID)
        // log.Printf("Send message to: %v\n", userID)
        sendMessage(userID, fmt.Sprintf("Verification request for '%s' will be sent to @%s\n(Smart contract integration pending)", info.expName, verifierUsername))
        
        {
            verifierID, err := repository.GetTelegramIDByUsername(verifierUsername)
            if err != nil {
                log.Printf("WARN: verifier lookup failed: %v", err)
                sendMessage(userID, "Technical error. Try again later.")
                return
            }
            if verifierID == 0 {
                sendMessage(userID, fmt.Sprintf("User @%s must start the bot first.", verifierUsername))
                delete(awaitingVerifier, userID)
                return
            }
            // sendMessage(verifierID, fmt.Sprintf("Verification request for '%s' will be sent to @%s\n(Smart contract integration pending)", info.expName, verifierUsername))

            delete(awaitingVerifier, userID)

            // verifierID, err := repository.GetTelegramIDByUsername(verifierUsername)
            // if err != nil {
            //     sendMessage(userID, "Technical error. Try again later.")
            //     return
            // }
            // if verifierID == 0 {
            //     sendMessage(userID, fmt.Sprintf("User @%s must start the bot first.", verifierUsername))
            //     return
            // }

            reqID := fmt.Sprintf("%d_%d_%d", userID, verifierID, info.expIndex)
            pendingVerifications[reqID] = verificationRequest{
                CandidateID:       userID,
                CandidateUsername: username,
                ExperienceLabel:   info.expName,
                VerifierID:        verifierID,
            }

            sendInlineKeyboard(
                verifierID,
                fmt.Sprintf("Approve verification for @%s (%s)?", username, info.expName),
                [][]map[string]interface{}{
                    {
                        {"text": "✅ Yes", "callback_data": "approve:" + reqID},
                        {"text": "❌ No", "callback_data": "reject:" + reqID},
                    },
                },
            )
            sendMessage(userID, fmt.Sprintf("Verification request sent to @%s.", verifierUsername))
        }
    }
}

func handleCallbackQuery(cb *struct {
    ID   string `json:"id"`
    From struct {
        ID       int64  `json:"id"`
        Username string `json:"username"`
    } `json:"from"`
    Data string `json:"data"`
}) {
    answerCallbackQuery(cb.ID)
    parts := strings.Split(cb.Data, ":")
    if len(parts) < 2 {
        return
    }

    action := parts[0]
    param := parts[1]
    key := parts[1]

    // if action == "verify" {
    switch action {
    case "verify":
        expIndex, _ := strconv.Atoi(param)
        candidate, _ := repository.GetCandidateByTelegramID(cb.From.ID)
        if candidate != nil && expIndex < len(candidate.Experience) {
            exp := candidate.Experience[expIndex]
            awaitingVerifier[cb.From.ID] = struct {
                expIndex int
                expName  string
            }{expIndex, fmt.Sprintf("%s @ %s", exp.Vacancy, exp.Company)}
            sendMessage(cb.From.ID, "Send verifier's username (e.g., @username):")
        }
    case "approve", "reject":
        req, ok := pendingVerifications[key]
        if !ok {
            sendMessage(cb.From.ID, "Request already handled or expired.")
            return
        }
        if req.VerifierID != cb.From.ID {
            sendMessage(cb.From.ID, "This request is not assigned to you.")
            return
        }
        delete(pendingVerifications, key)

        if action == "approve" {
            sendMessage(cb.From.ID, "You approved the experience. Thanks!")
            sendMessage(
                req.CandidateID,
                fmt.Sprintf("🎉 @%s approved your experience \"%s\".", cb.From.Username, req.ExperienceLabel),    
            )

            createContract(req.CandidateID)

        } else {
            sendMessage(cb.From.ID, "You rejected the experience.")
            sendMessage(
                req.CandidateID,
                fmt.Sprintf("⛔ @%s rejected your experience \"%s\".", cb.From.Username, req.ExperienceLabel),
            )
        }
    }
}

func sendMessage(chatID int64, text string) {
    payload := map[string]interface{}{"chat_id": chatID, "text": text}
    sendTelegramRequest("sendMessage", payload)
}

func sendInlineKeyboard(chatID int64, text string, buttons [][]map[string]interface{}) {
    payload := map[string]interface{}{
        "chat_id":      chatID,
        "text":         text,
        "reply_markup": map[string]interface{}{"inline_keyboard": buttons},
    }
    sendTelegramRequest("sendMessage", payload)
}

func answerCallbackQuery(callbackQueryID string) {
    payload := map[string]interface{}{"callback_query_id": callbackQueryID}
    sendTelegramRequest("answerCallbackQuery", payload)
}

func sendTelegramRequest(method string, payload map[string]interface{}) {
    url := fmt.Sprintf("https://api.telegram.org/bot%s/%s", botToken, method)
    body, _ := json.Marshal(payload)
    resp, err := http.Post(url, "application/json", bytes.NewReader(body))
    if err != nil {
        log.Printf("Telegram API error: %v", err)
        return
    }
    defer resp.Body.Close()
}

func createContract(CandidateID int64) {
	sendMessage(CandidateID, "Starting smart contract deploing\n")
}

func HandleVerify(w http.ResponseWriter, r *http.Request) {
    var req struct {
        InitData string `json:"initData"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "bad request", 400)
        return
    }
    
    params := strings.Split(req.InitData, "&")
    for _, param := range params {
        if strings.HasPrefix(param, "user=") {
            userJSON := strings.TrimPrefix(param, "user=")
            var user struct {
                ID       int64  `json:"id"`
                Username string `json:"username"`
            }
            if err := json.Unmarshal([]byte(userJSON), &user); err == nil {
                _ = repository.CreateUser(user.ID, user.Username)
                log.Printf("User registered via webapp: %d @%s", user.ID, user.Username)
            }
            break
        }
    }
    
    resp := map[string]interface{}{
        "ok":       true,
        "initData": req.InitData,
    }
    json.NewEncoder(w).Encode(resp)
}