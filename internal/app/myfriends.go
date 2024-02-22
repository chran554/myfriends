package app

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/0x6flab/namegenerator"
	"math/rand"
	"myfriends/internal/app/model"
	widget2 "myfriends/internal/app/widget"
	"strings"
)

const amountFriends = 50

var boundFriends []binding.ExternalUntyped

var friendDetailForm *widget2.FriendDetailForm
var friendList *widget.List

func ApplicationContent() *fyne.Container {
	// Setup data model

	friends := loadFriends() // "Load" our application model (rather generate model data), the list of Friend

	// Map the friends model list to list of bound data friends
	boundFriends = make([]binding.ExternalUntyped, len(friends)) // Do not use slice append to build boundFriends list, for some reason! Strange effects...
	for i := 0; i < len(friends); i++ {
		boundFriend := binding.BindUntyped(&friends[i])
		boundFriends[i] = boundFriend
	}

	// Setup UI

	friendListWidget := createFriendListWidget()
	friendDetailsWidget := createFriendDetailsWidget()
	split := container.NewHSplit(friendListWidget, friendDetailsWidget)
	split.SetOffset(0.33333)

	printModelAction := widget.NewToolbarAction(theme.DocumentPrintIcon(), func() { printFriends(friends) })
	toolbar := widget.NewToolbar(printModelAction)

	mainLayout := container.NewBorder(toolbar, nil, nil, nil, split)

	return mainLayout
}

func printFriends(friends []model.Friend) {
	for i, friend := range friends {
		fmt.Printf("%2d: %s\n", i, friend.String())
	}
}

func loadFriends() []model.Friend {
	var friends []model.Friend

	// Generate random friends
	generator := namegenerator.NewGenerator()
	for i := 0; i < amountFriends; i++ {
		generatedName := strings.Split(generator.Generate(), "-")
		friend := model.Friend{FirstName: generatedName[0], FamilyName: generatedName[1], Age: 18 + rand.Intn(50)}
		friends = append(friends, friend)
	}

	return friends
}

func createFriendListWidget() fyne.CanvasObject {
	friendList = widget.NewList(
		func() int {
			return len(boundFriends)
		},
		func() fyne.CanvasObject {
			return widget2.NewFriendListItem()
		},
		func(id widget.ListItemID, object fyne.CanvasObject) {
			boundFriend := boundFriends[id]
			friendListItem := object.(*widget2.FriendListItem)
			friendListItem.Bind(boundFriend)
		},
	)

	friendList.OnSelected = func(id widget.ListItemID) {
		boundFriend := boundFriends[id]
		friendDetailForm.Bind(boundFriend)
	}

	caption := widget.NewLabel("Friends")
	caption.TextStyle.Bold = true
	caption.Importance = widget.HighImportance

	return container.NewBorder(caption, nil, nil, nil, friendList)
}

func createFriendDetailsWidget() fyne.CanvasObject {
	caption := widget.NewLabel("Friend details")
	caption.TextStyle.Bold = true
	caption.Importance = widget.HighImportance

	emptyFriend := binding.BindUntyped(&model.Friend{}) // how sad :(

	friendDetailForm = widget2.NewFriendDetailForm()
	friendDetailForm.Bind(emptyFriend)

	locked := true
	friendDetailForm.SetEditable(!locked)
	padLockButton := widget2.NewPadLockButton(locked, func(locked bool) {
		friendDetailForm.SetEditable(!locked)
	})

	captionPanel := container.NewHBox(caption, layout.NewSpacer(), padLockButton)
	panel := container.NewVBox(captionPanel, friendDetailForm)

	return container.NewBorder(panel, nil, nil, nil)
}
