package handler

import (
	"fmt"
	"strconv"

	"github.com/ccns/quiz-server/db"
	"github.com/gin-gonic/gin"
)

// Answer describe the structue of an answer from player to quiz.
type Answer struct {
	PlayerName string `json:"player_name"`
	QuizNumber int    `json:"quiz_number"`
	Correct    bool   `json:"correct"`
}

// GetAnswersHandler handlers get requests on route /answers.
func GetAnswersHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/answers"},
	}

	playerName := ctx.DefaultQuery("player", "")
	quizNumber := ctx.DefaultQuery("quiz", "")

	if playerName != "" && quizNumber != "" {

		quizNumber, err := strconv.Atoi(quizNumber)
		if err != nil {
			resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
			ctx.String(400, resp)
			return
		}

		data, err := db.GetAnswer(playerName, quizNumber)
		if err != nil {
			resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
			ctx.String(500, resp)
			return
		}
		links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quizNumber)}
		links["player"] = LinkDetail{fmt.Sprintf("/v1/players/%s", playerName)}

		status := Status{200, "answers listed successfully."}
		resp, _ := BuildHATEOAS(links, status, data, nil)
		ctx.String(200, resp)
		return
	}

	if playerName != "" {

		data, err := db.QueryQuizAnswersByPlayer(playerName)
		if err != nil {
			status := Status{500, err.Error()}
			resp, _ := BuildHATEOAS(links, status, nil, nil)
			ctx.String(500, resp)
			return
		}
		links["player"] = LinkDetail{fmt.Sprintf("/v1/players/%s", playerName)}

		status := Status{200, "answers listed successfully."}
		resp, _ := BuildHATEOAS(links, status, data, nil)
		ctx.String(200, resp)
		return

	} else if quizNumber != "" {

		quizNumber, err := strconv.Atoi(quizNumber)
		if err != nil {
			resp, _ := BuildHATEOAS(links, Status{400, err.Error()}, nil, nil)
			ctx.String(400, resp)
			return
		}

		data, err := db.QueryPlayerAnswersByQuiz(quizNumber)
		if err != nil {
			resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
			ctx.String(500, resp)
			return
		}
		links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", quizNumber)}

		status := Status{200, "answers listed successfully."}
		resp, _ := BuildHATEOAS(links, status, data, nil)
		ctx.String(200, resp)
		return
	}

	data, err := db.ListAnswers()
	if err != nil {
		resp, _ := BuildHATEOAS(links, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}

	status := Status{200, "answers listed successfully."}
	resp, _ := BuildHATEOAS(links, status, data, nil)
	ctx.String(200, resp)
}

// PostAnswersHandler handlers post requests on route /answers.
func PostAnswersHandler(ctx *gin.Context) {

	links := map[string]LinkDetail{
		"self": LinkDetail{"/v1/answers"},
	}

	var answer Answer
	err := ctx.BindJSON(&answer)
	if err != nil {
		resp, _ := BuildHATEOAS(nil, Status{400, err.Error()}, nil, nil)
		ctx.String(400, resp)
		return
	}

	err = db.RegisterAnswer(answer.PlayerName, answer.QuizNumber, answer.Correct)
	if err != nil {
		tpl := "answer from player %s to quiz number %d already existed"
		if err.Error() == fmt.Sprintf(tpl, answer.PlayerName, answer.QuizNumber) {
			resp, _ := BuildHATEOAS(nil, Status{409, err.Error()}, nil, nil)
			ctx.String(409, resp)
			return
		}
		resp, _ := BuildHATEOAS(nil, Status{500, err.Error()}, nil, nil)
		ctx.String(500, resp)
		return
	}
	links["quiz"] = LinkDetail{fmt.Sprintf("/v1/quizzes/%d", answer.QuizNumber)}
	links["player"] = LinkDetail{fmt.Sprintf("/v1/players/%s", answer.PlayerName)}

	status := Status{201, "answer created successfully."}
	resp, _ := BuildHATEOAS(nil, status, answer, nil)
	ctx.String(201, resp)
}
