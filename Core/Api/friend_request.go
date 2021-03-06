package Api

import (
	"golang-restAPI-FR/Core/Structs"
	"golang-restAPI-FR/Core/Models"
)

func FriendRequest(frq_req Structs.FriendRequestRequest) interface{} {
	success_response.Success = true
	error_response.Success = false
	var friend_requestor Models.User
	var user Models.User
	var friend Models.Friend
	friend.Status = "pending"

	// store user email requestor and receiver
	friend_requestor.Email = frq_req.Requestor
	if err := Models.CreateUser(&friend_requestor); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}
	user.Email = frq_req.To
	if err := Models.CreateUser(&user); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}
	// ----------

	// create friend request
	friend.UserID = user.ID
	friend.FriendRequestID = friend_requestor.ID

	if err := Models.FindFriendRequest(&friend, []string{"pending"}); err == nil {
		error_response.Msg = "Friend request already send"
		return error_response
	}

	if err := Models.FindFriendRequest(&friend, []string{"rejected"}); err == nil {
		error_response.Msg = "Friend request being rejected"
		return error_response
	}

	if err := Models.FindFriendRequest(&friend, []string{"accepted"}); err == nil {
		error_response.Msg = "Already friend with " + user.Email
		return error_response
	}

	if err := Models.FindFriendRequest(&friend, []string{"blocked"}); err == nil {
		error_response.Msg = "Friend request invalid"
		return error_response
	}

	if err := Models.CreateFriendRequest(&friend); err != nil {
		error_response.Msg = err.Error()
		return error_response
	}

	return success_response
}
