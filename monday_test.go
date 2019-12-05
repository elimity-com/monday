package monday

import "testing"

func TestStructures(t *testing.T) {
	tests := map[string]map[string]struct {
		query Query
		str   string
	}{
		"Boards": {
			"Simple": {
				query: NewBoards(nil),
				str:   `boards{id}`,
			},
			"Args": {
				query: NewBoardsWithArguments(
					nil,
					[]BoardsArgument{
						NewBoardsIDsArgument([]int{1}),
						NewBoardsPageArgument(1),
						NewBoardsLimitArgument(1),
						NewBoardsStateArgument(AllState()),
						NewBoardsKindArgument(BoardsKindPublic()),
						NewBoardsNewestFirstArgument(true),
					},
				),
				str: `boards(ids:1,page:1,limit:1,state:all,board_kind:public,newest_first:true){id}`,
			},
			"Fields1": {
				query: NewBoards(
					[]BoardsField{
						BoardsIDField(),
						BoardsPositionField(),
						BoardsNameField(),
						NewBoardsTagsField(nil, nil),
						NewBoardsItemsFields(nil, nil),
						NewBoardsOwnerField(nil, nil),
						BoardsStateField(),
						NewBoardsGroupsFields(nil, nil),
						NewBoardsColumnField(nil),
					},
				),
				str: `boards{id pos name tags{id} items{id} owner{id} state groups{id} columns{id}}`,
			},
			"Fields2": {
				query: NewBoards(
					[]BoardsField{
						NewBoardsUpdatesField(nil, nil),
						BoardsKindField(),
						BoardsDescriptionField(),
						BoardsPermissionsField(),
						NewBoardsSubscribersField(nil, nil),
						BoardsFolderIDField(),
					},
				),
				str: `boards{updates{id} board_kind description permissions subscribers{id} board_folder_id}`,
			},
		},
		"Columns": {
			"Simple": {
				query: NewColumns(nil),
				str:   `columns{id}`,
			},
			"Fields": {
				query: NewColumns(
					[]ColumnsField{
						ColumnsIDField(),
						ColumnsTypeField(),
						ColumnsTitleField(),
						ColumnsWidthField(),
						ColumnsArchivedField(),
						ColumnsSettingsStrField(),
					},
				),
				str: `columns{id type title width archived settings_str}`,
			},
		},
		"Groups": {
			"Simple": {
				query: NewGroups(nil),
				str:   `groups{id}`,
			},
			"Args": {
				query: NewGroupWithArguments(
					nil,
					[]GroupsArgument{
						NewGroupsIDsArgument([]string{"x", "y"}),
					},
				),
				str: `groups(ids:["x","y"]){id}`,
			},
			"Fields": {
				query: NewGroups(
					[]GroupsField{
						GroupsIDField(),
						GroupsColorField(),
						NewGroupsItemsField(nil, nil),
						GroupsTitleField(),
						GroupsDeletedField(),
						GroupsArchivedField(),
						GroupsPositionField(),
					},
				),
				str: `groups{id color items{id} title deleted archived position}`,
			},
		},
		"Items": {
			"Simple": {
				query: NewItems(nil),
				str:   `items{id}`,
			},
			"Args": {
				query: NewItemsWithArguments(
					nil,
					[]ItemsArgument{
						NewItemsIDsArgument([]int{1}),
						NewItemsPageArgument(1),
						NewItemsLimitArgument(1),
						NewItemsNewestFirst(true),
					},
				),
				str: `items(ids:1,page:1,limit:1,newest_first:true){id}`,
			},
			"Fields1": {
				query: NewItems(
					[]ItemsField{
						ItemsIDField(),
						ItemsNameField(),
						NewItemsBoardField(nil, nil),
						NewItemsColumnValuesField(nil, nil),
						NewItemsGroupField(nil, nil),
						ItemsStateField(),
						NewItemsCreatorField(nil, nil),
						NewItemsUpdatesField(nil, nil),
					},
				),
				str: `items{id name board{id} column_values{id} group{id} state creator{id} updates{id}}`,
			},
			"Fields2": {
				query: NewItems(
					[]ItemsField{
						ItemsCreatedAtField(),
						ItemsCreatorIDField(),
						ItemsUpdatedAtField(),
						NewItemsSubscribersField(nil, nil),
					},
				),
				str: `items{created_at creator_id updated_at subscribers{id}}`,
			},
		},
		"ItemsByColumnValues": {
			"Simple": {
				query: NewItemsByColumnValues(1, "x", NewTextValue("x"), nil),
				str:   `items_by_column_values(board_id:1,column_id:"x",column_value:"\"x\""){id}`,
			},
			"Args": {
				query: NewItemsByColumnValuesWithArguments(
					1, "x", NewTextValue("x"),
					nil,
					[]ItemsByColumnValuesArgument{
						NewItemsByColumnValuesPageArgument(1),
						NewItemsByColumnValuesLimitArgument(1),
						NewItemsByColumnValuesStateArgument(ActiveState()),
						NewItemsByColumnValuesColumnTypeArgument("x"),
					},
				),
				str: `items_by_column_values(board_id:1,column_id:"x",column_value:"\"x\"",page:1,limit:1,state:active,column_type:"x"){id}`,
			},
			"Fields1": {
				query: NewItemsByColumnValues(
					1, "x", NewTextValue("x"),
					[]ItemsByColumnValuesField{
						ItemsByColumnValuesIDField(),
						ItemsByColumnValuesNameField(),
						NewItemsByColumnValuesBoardField(nil, nil),
						NewItemsByColumnValuesColumnValuesField(nil, nil),
						NewItemsByColumnValuesGroupField(nil, nil),
						ItemsByColumnValuesStateField(),
						NewItemsByColumnValuesCreatorField(nil, nil),
						NewItemsByColumnValuesUpdatesField(nil, nil),
						ItemsByColumnValuesCreatedAtField(),
						ItemsByColumnValuesCreatorIDField(),
						ItemsByColumnValuesUpdatedAtField(),
						NewItemsByColumnValuesSubscribersField(nil, nil),
					},
				),
				str: `items_by_column_values(board_id:1,column_id:"x",column_value:"\"x\""){id name board{id} column_values{id} group{id} state creator{id} updates{id} created_at creator_id updated_at subscribers{id}}`,
			},
		},
		"Updates": {
			"Simple": {
				query: NewUpdates(nil),
				str:   `updates{id}`,
			},
			"Args": {
				query: NewUpdatesWithArguments(
					nil,
					[]UpdatesArgument{
						NewUpdatesPageArgument(1),
						NewUpdatesLimitArgument(1),
					},
				),
				str: `updates(page:1,limit:1){id}`,
			},
			"Fields": {
				query: NewUpdates(
					[]UpdatesField{
						UpdatesBodyField(),
						UpdatesCreatedAtField(),
						NewUpdatesCreatorField(nil, nil),
						UpdatesCreatorIDField(),
						UpdatesIDField(),
						UpdatesItemIDField(),
						NewUpdatesRepliesField(nil),
						UpdatesTextBodyField(),
						UpdatesUpdatedAtField(),
					},
				),
				str: `updates{body updated_at creator{id} creator_id id item_id replies{id} text_body updated_at}`,
			},
		},
		"Tags": {
			"Simple": {
				query: NewTags(nil),
				str:   `tags{id}`,
			},
			"Args": {
				query: NewTagsWithArguments(
					nil,
					[]TagsArgument{
						NewTagsIDsArgument([]int{1}),
					},
				),
				str: `tags(ids:1){id}`,
			},
			"Fields": {
				query: NewTags(
					[]TagsField{
						TagsColorField(),
						TagsIDField(),
						TagsNameField(),
					},
				),
				str: `tags{color id name}`,
			},
		},
		"Users": {
			"Simple": {
				query: NewUsers(nil),
				str:   `users{id}`,
			},
			"Args": {
				query: NewUsersWithArguments(
					nil,
					[]UsersArgument{
						NewUsersIDsArgument([]int{1}),
						NewUsersKindArgument(AllUsersKind()),
						NewUsersNewestFirstArgument(true),
						NewUsersLimitArgument(1),
					},
				),
				str: `users(ids:1,kind:all,newest_first:true,limit:1){id}`,
			},
			"Fields1": {
				query: NewUsers(
					[]UsersField{
						NewUsersAccountField(nil),
						UsersBirthDayField(),
						UsersCountryCodeField(),
						UsersCreatedAtField(),
						UsersEmailField(),
						UsersEnabledField(),
						UsersIDField(),
						UsersIsGuestField(),
						UsersIsPendingField(),
						UsersJoinDateField(),
						UsersLocationField(),
						UsersMobilePhoneField(),
						UsersNameField(),
						UsersPhoneField(),
					},
				),
				str: `users{account{id} birthday country_code created_at email enabled id is_guest is_pending join_date location mobile_phone name phone}`,
			},
			"Fields2": {
				query: NewUsers(
					[]UsersField{
						UsersPhotoOriginalField(),
						UsersPhotoSmallField(),
						UsersPhotoThumbField(),
						UsersPhotoThumbSmallField(),
						UsersPhotoTinyField(),
						NewUsersTeamsField(nil, nil),
						UsersTimeZoneIdentifierField(),
						UsersTitleField(),
						UsersURLField(),
						UsersUTCHoursDifferenceField(),
					},
				),
				str: `users{photo_original photo_small photo_thumb photo_thumb_small photo_tiny teams{id} time_zone_identifier title url utc_hours_diff}`,
			},
		},
		"Teams": {
			"Simple": {
				query: NewTeams(nil),
				str:   `teams{id}`,
			},
			"Args": {
				query: NewTeamsWithArguments(
					nil,
					[]TeamsArgument{
						NewTeamsIDsArgument([]int{1}),
					},
				),
				str: `teams(ids:1){id}`,
			},
			"Fields": {
				query: NewTeams(
					[]TeamsField{
						TeamsIDField(),
						TeamsNameField(),
						TeamsPictureURLField(),
						NewTeamsUsersField(nil, nil),
					},
				),
				str: `teams{id name picture_url users{id}}`,
			},
		},
		"Complexity": {
			"Simple": {
				query: NewComplexity(nil),
				str:   `complexity{after before query}`,
			},
		},
		"Account": {
			"Simple": {
				query: NewAccount(nil),
				str:   `account{id}`,
			},
			"Fields": {
				query: NewAccount(
					[]AccountField{
						AccountFirstDayOfTheWeekField(),
						AccountIDField(),
						AccountLogoField(),
						AccountNameField(),
						AccountPlanField(nil),
						AccountShowTimelineWeekendsField(),
						AccountSlugField(),
					},
				),
				str: `account{first_day_of_the_week id logo name plan{max_users period tier version} show_timeline_weekends slug}`,
			},
		},
	}

	for entityName, entityTests := range tests {
		t.Run(entityName, func(t *testing.T) {
			for nameTest, test := range entityTests {
				t.Run(nameTest, func(t *testing.T) {
					if str := test.query.stringify(); str != test.str {
						t.Errorf("got: %s, exprected: %s\n", str, test.str)
					}
				})
			}
		})
	}
}
