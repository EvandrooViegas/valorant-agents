package comments_handler

import (
	"encoding/json"
	"fmt"
	"valorant-agents/services"
	"valorant-agents/utils"

	"github.com/gofiber/fiber/v2"
)

func GetCommentsHandler(c *fiber.Ctx) error {
	comments, err := services.GetComments()
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "Someting went wrong",
			Error:   err,
		})
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{
		Status:  fiber.StatusOK,
		Message: "Fetched all comments successfully",
		Data: map[string]interface{}{
			"comments": comments,
			"length":   len(comments),
		},
	})
}

func GetCommentRepliesHandler(c *fiber.Ctx) error {
	parentID := c.Params("id")
	fmt.Println("parentID: ", parentID)
	replies, err := services.GetCommentReplies(parentID)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "Someting went wrong",
			Error:   err,
		})
	}
	return utils.WriteJSON(c, utils.WriteJSONpayload{
		Status:  fiber.StatusOK,
		Message: "Fetched Replies Successfully",
		Data: map[string]interface{}{
			"has_replies": len(replies) > 0,
			"replies": replies,
		},
	})

}

func CreateCommentHandler(c *fiber.Ctx) error {
	authToken := c.Cookies("token")
	// read the token and get the player
	player, err := services.ReadTokenAndReturnPlayer(authToken)
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusUnauthorized,
			Message: "Unauthorized",
			Error:   err,
		})
	}
	// read the body
	var request services.CreateCommentRequest
	body := c.Body()
	if err := json.Unmarshal(body, &request); err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusBadRequest,
			Message: "Invalid request",
			Error:   err,
		})
	}
	reqComment := request.Comment
	// insert to the database
	nComment, err := services.CreateComment(services.DatabaseComment{
		AgentID:  reqComment.AgentID,
		ParentID: reqComment.ParentID,
		Text:     reqComment.Text,
		Author:   player,
	})
	if err != nil {
		return utils.WriteJSON(c, utils.WriteJSONpayload{
			Status:  fiber.StatusInternalServerError,
			Message: "Something went wrong",
			Error:   err,
		})
	}

	return utils.WriteJSON(c, utils.WriteJSONpayload{
		Status:  fiber.StatusCreated,
		Message: "Comment Created",
		Data: map[string]interface{}{
			"comment":    nComment,
			"wasCreated": true,
		},
	})

}
