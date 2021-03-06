package cli

import (
	"fmt"
	"strconv"

	"github.com/VitJRBOG/CleaningVk/ccb"
	"github.com/VitJRBOG/CleaningVk/cwp"
	"github.com/VitJRBOG/CleaningVk/cwpc"
)

// ShowUI отображает консольный пользовательский интерфейс (КПИ)
func ShowUI() {
	var ui UI
	ui.showMainMenu()
	for {
		ui.setUserSelection()
		success := ui.showSelected()
		if success {
			break
		} else {
			ui.showMessageOfWrongInput()
		}
	}
	ui.showMessageOfExit()
}

// UI хранит информацию о выборе пользователя и служит для вызова методов отображения разделов КПИ
type UI struct {
	UserSelection string
}

func (ui *UI) showMainMenu() {
	fmt.Print("\n[Main menu]\n" +
		"1. Deleting wallposts;\n" +
		"2. Deleting comments of wallposts;\n" +
		"3. Unbanning subjects of community's blacklist.\n" +
		"--- Enter number of menu item and press «Enter» ---\n" +
		"> ")
}

func (ui *UI) setUserSelection() {
	_, err := fmt.Scan(&ui.UserSelection)
	if err != nil {
		panic(err.Error())
	}
}

func (ui *UI) showSelected() bool {
	switch ui.UserSelection {
	case "1":
		showDeletingWallPostsUI()
		return true
	case "2":
		showDeletingWallPostsCommentsUI()
		return true
	case "3":
		showUnbanningCommunityBlacklistSubjectsUI()
		return true
	default:
		return false
	}
}

func (ui *UI) showMessageOfWrongInput() {
	fmt.Print("ERROR! Your input is wrong. Please try again...\n" +
		"> ")
}

func (ui *UI) showMessageOfExit() {
	fmt.Print("Exit...\n")
}

func showDeletingWallPostsUI() {
	var cwpUI CleanWallPostsUI
	cwpUI.init()
	fmt.Print("[Deleting wallposts]\n")
	cwpUI.setAccessToken()
	cwpUI.setOwnerID()
	cwpUI.setAuthorID()
	go cwp.RunWallPostsCleaning(cwpUI.AccessToken, cwpUI.OwnerID, cwpUI.AuthorID, cwpUI.msgChannel)
	cwpUI.outputtingMessages()
}

// CleanWallPostsUI хранит информацию для модуля удаления постов со стены
type CleanWallPostsUI struct {
	AccessToken string
	OwnerID     int
	AuthorID    int
	msgChannel  chan string
}

func (c *CleanWallPostsUI) init() {
	c.msgChannel = make(chan string)
}

func (c *CleanWallPostsUI) setAccessToken() {
	fmt.Print("--- Enter your access token and press «Enter» ---\n" +
		"> ")
	var accessToken string
	_, err := fmt.Scan(&accessToken)
	if err != nil {
		panic(err.Error())
	}
	c.AccessToken = accessToken
}

func (c *CleanWallPostsUI) setOwnerID() {
	fmt.Print("--- Now enter ID of owner of wallposts and press «Enter» ---\n" +
		"> ")
	var ownerID string
	_, err := fmt.Scan(&ownerID)
	if err != nil {
		panic(err.Error())
	}
	c.OwnerID, err = strconv.Atoi(ownerID)
	if err != nil {
		panic(err.Error())
	}
}

func (c *CleanWallPostsUI) setAuthorID() {
	fmt.Print("--- And enter ID of author of wallposts (or enter 0" +
		" if you need deleting all wallposts) and press «Enter» ---\n" +
		"> ")
	var authorID string
	_, err := fmt.Scan(&authorID)
	if err != nil {
		panic(err.Error())
	}
	c.AuthorID, err = strconv.Atoi(authorID)
	if err != nil {
		panic(err.Error())
	}
}

func (c *CleanWallPostsUI) outputtingMessages() {
	for {
		msg := <-c.msgChannel
		fmt.Printf("%v\n", msg)
		if msg == "Done!" ||
			msg == "No wallposts from this author..." ||
			msg == "No posts on the wall..." {
			break
		}
	}
}

func showDeletingWallPostsCommentsUI() {
	var cwpcUI CleanWallPostCommentsUI
	cwpcUI.init()
	fmt.Print("[Deleting comments of wallposts]\n")
	cwpcUI.setAccessToken()
	cwpcUI.setOwnerID()
	cwpcUI.setAuthorID()
	go cwpc.RunWallPostCommentsCleaning(cwpcUI.AccessToken, cwpcUI.OwnerID, cwpcUI.AuthorID,
		cwpcUI.msgChannel)
	cwpcUI.outputtingMessages()
}

// CleanWallPostCommentsUI хранит информацию для модуля удаления комментариев из-под постов на стене
type CleanWallPostCommentsUI struct {
	AccessToken string
	OwnerID     int
	AuthorID    int
	msgChannel  chan string
}

func (c *CleanWallPostCommentsUI) init() {
	c.msgChannel = make(chan string)
}

func (c *CleanWallPostCommentsUI) setAccessToken() {
	fmt.Print("--- Enter your access token and press «Enter» ---\n" +
		"> ")
	var accessToken string
	_, err := fmt.Scan(&accessToken)
	if err != nil {
		panic(err.Error())
	}
	c.AccessToken = accessToken
}

func (c *CleanWallPostCommentsUI) setOwnerID() {
	fmt.Print("--- Now enter ID of owner of comments and press «Enter» ---\n" +
		"> ")
	var ownerID string
	_, err := fmt.Scan(&ownerID)
	if err != nil {
		panic(err.Error())
	}
	c.OwnerID, err = strconv.Atoi(ownerID)
	if err != nil {
		panic(err.Error())
	}
}

func (c *CleanWallPostCommentsUI) setAuthorID() {
	fmt.Print("--- And enter ID of author of comments (or enter 0" +
		" if you need deleting all comments) and press «Enter» ---\n" +
		"> ")
	var authorID string
	_, err := fmt.Scan(&authorID)
	if err != nil {
		panic(err.Error())
	}
	c.AuthorID, err = strconv.Atoi(authorID)
	if err != nil {
		panic(err.Error())
	}
}

func (c *CleanWallPostCommentsUI) outputtingMessages() {
	for {
		msg := <-c.msgChannel
		fmt.Printf("%v\n", msg)
		if msg == "Done!" || msg == "No comments of wallpost from this author..." {
			break
		}
	}
}

func showUnbanningCommunityBlacklistSubjectsUI() {
	var cbUI CleanBlacklistUI
	cbUI.init()
	fmt.Print("[Unbanning subjects of community's blacklist]\n")
	cbUI.setAccessToken()
	cbUI.setOwnerID()
	go ccb.RunCommunityBlacklistCleaning(cbUI.AccessToken, cbUI.OwnerID, cbUI.msgChannel)
	cbUI.outputtingMessages()
}

// CleanBlacklistUI хранит информацию для модуля удаления субъектов из черного списка сообщества
type CleanBlacklistUI struct {
	AccessToken string
	OwnerID     int
	msgChannel  chan string
}

func (c *CleanBlacklistUI) init() {
	c.msgChannel = make(chan string)
}

func (c *CleanBlacklistUI) setAccessToken() {
	fmt.Print("--- Enter your access token and press «Enter» ---\n" +
		"> ")
	var accessToken string
	_, err := fmt.Scan(&accessToken)
	if err != nil {
		panic(err.Error())
	}
	c.AccessToken = accessToken
}

func (c *CleanBlacklistUI) setOwnerID() {
	fmt.Print("--- Now enter ID of owner of blacklist and press «Enter» ---\n" +
		"> ")
	var ownerID string
	_, err := fmt.Scan(&ownerID)
	if err != nil {
		panic(err.Error())
	}
	c.OwnerID, err = strconv.Atoi(ownerID)
	if err != nil {
		panic(err.Error())
	}
}

func (c *CleanBlacklistUI) outputtingMessages() {
	for {
		msg := <-c.msgChannel
		fmt.Printf("%v\n", msg)
		if msg == "Done!" || msg == "No banned subjects in the blacklist..." {
			break
		}
	}
}
