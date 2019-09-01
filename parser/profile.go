package parser

type Profile struct {
	Config struct {
		CsrfToken string `json:"csrf_token"`
		Viewer    struct {
			Biography        string `json:"biography"`
			ExternalURL      string `json:"external_url"`
			FullName         string `json:"full_name"`
			HasPhoneNumber   bool   `json:"has_phone_number"`
			HasProfilePic    bool   `json:"has_profile_pic"`
			ID               string `json:"id"`
			IsJoinedRecently bool   `json:"is_joined_recently"`
			IsPrivate        bool   `json:"is_private"`
			ProfilePicURL    string `json:"profile_pic_url"`
			ProfilePicURLHd  string `json:"profile_pic_url_hd"`
			Username         string `json:"username"`
			BadgeCount       string `json:"badge_count"`
		} `json:"viewer"`
		ViewerID string `json:"viewerId"`
	} `json:"config"`
	CountryCode  string `json:"country_code"`
	LanguageCode string `json:"language_code"`
	Locale       string `json:"locale"`
	EntryData    struct {
		ProfilePage []struct {
			LoggingPageID         string      `json:"logging_page_id"`
			ShowSuggestedProfiles bool        `json:"show_suggested_profiles"`
			ShowFollowDialog      bool        `json:"show_follow_dialog"`
			ToastContentOnLoad    interface{} `json:"toast_content_on_load"`
			Graphql               struct {
				User struct {
					Biography              string      `json:"biography"`
					BlockedByViewer        bool        `json:"blocked_by_viewer"`
					CountryBlock           bool        `json:"country_block"`
					ExternalURL            interface{} `json:"external_url"`
					ExternalURLLinkshimmed interface{} `json:"external_url_linkshimmed"`
					EdgeFollowedBy         struct {
						Count int `json:"count"`
					} `json:"edge_followed_by"`
					FollowedByViewer bool `json:"followed_by_viewer"`
					EdgeFollow       struct {
						Count int `json:"count"`
					} `json:"edge_follow"`
					FollowsViewer        bool        `json:"follows_viewer"`
					FullName             string      `json:"full_name"`
					HasChannel           bool        `json:"has_channel"`
					HasBlockedViewer     bool        `json:"has_blocked_viewer"`
					HighlightReelCount   int         `json:"highlight_reel_count"`
					HasRequestedViewer   bool        `json:"has_requested_viewer"`
					ID                   string      `json:"id"`
					IsBusinessAccount    bool        `json:"is_business_account"`
					IsJoinedRecently     bool        `json:"is_joined_recently"`
					BusinessCategoryName interface{} `json:"business_category_name"`
					IsPrivate            bool        `json:"is_private"`
					IsVerified           bool        `json:"is_verified"`
					EdgeMutualFollowedBy struct {
						Count int           `json:"count"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_mutual_followed_by"`
					ProfilePicURL          string      `json:"profile_pic_url"`
					ProfilePicURLHd        string      `json:"profile_pic_url_hd"`
					RequestedByViewer      bool        `json:"requested_by_viewer"`
					Username               string      `json:"username"`
					ConnectedFbPage        interface{} `json:"connected_fb_page"`
					EdgeFelixVideoTimeline struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_felix_video_timeline"`
					EdgeOwnerToTimelineMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool   `json:"has_next_page"`
							EndCursor   string `json:"end_cursor"`
						} `json:"page_info"`
						Edges []struct {
							Node struct {
								Typename           string `json:"__typename"`
								ID                 string `json:"id"`
								EdgeMediaToCaption struct {
									Edges []struct {
										Node struct {
											Text string `json:"text"`
										} `json:"node"`
									} `json:"edges"`
								} `json:"edge_media_to_caption"`
								Shortcode          string `json:"shortcode"`
								EdgeMediaToComment struct {
									Count int `json:"count"`
								} `json:"edge_media_to_comment"`
								CommentsDisabled bool `json:"comments_disabled"`
								TakenAtTimestamp int  `json:"taken_at_timestamp"`
								Dimensions       struct {
									Height int `json:"height"`
									Width  int `json:"width"`
								} `json:"dimensions"`
								DisplayURL  string `json:"display_url"`
								EdgeLikedBy struct {
									Count int `json:"count"`
								} `json:"edge_liked_by"`
								EdgeMediaPreviewLike struct {
									Count int `json:"count"`
								} `json:"edge_media_preview_like"`
								Location     interface{} `json:"location"`
								GatingInfo   interface{} `json:"gating_info"`
								MediaPreview string      `json:"media_preview"`
								Owner        struct {
									ID       string `json:"id"`
									Username string `json:"username"`
								} `json:"owner"`
								ThumbnailSrc       string `json:"thumbnail_src"`
								ThumbnailResources []struct {
									Src          string `json:"src"`
									ConfigWidth  int    `json:"config_width"`
									ConfigHeight int    `json:"config_height"`
								} `json:"thumbnail_resources"`
								IsVideo              bool        `json:"is_video"`
								FelixProfileGridCrop interface{} `json:"felix_profile_grid_crop"`
								VideoViewCount       int         `json:"video_view_count"`
							} `json:"node,omitempty"`
						} `json:"edges"`
					} `json:"edge_owner_to_timeline_media"`
					EdgeSavedMedia struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_saved_media"`
					EdgeMediaCollections struct {
						Count    int `json:"count"`
						PageInfo struct {
							HasNextPage bool        `json:"has_next_page"`
							EndCursor   interface{} `json:"end_cursor"`
						} `json:"page_info"`
						Edges []interface{} `json:"edges"`
					} `json:"edge_media_collections"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
	Hostname        string  `json:"hostname"`
	DeploymentStage string  `json:"deployment_stage"`
	Platform        string  `json:"platform"`
	Nonce           string  `json:"nonce"`
	MidPct          float64 `json:"mid_pct"`
	ZeroData        struct {
	} `json:"zero_data"`
	CacheSchemaVersion int `json:"cache_schema_version"`
	ServerChecks       struct {
		Hfe bool `json:"hfe"`
	} `json:"server_checks"`
	Knobx struct {
		Num4  bool `json:"4"`
		Num16 bool `json:"16"`
		Num17 bool `json:"17"`
	} `json:"knobx"`
	ToCache struct {
		Gatekeepers struct {
			Num4  bool `json:"4"`
			Num5  bool `json:"5"`
			Num6  bool `json:"6"`
			Num7  bool `json:"7"`
			Num8  bool `json:"8"`
			Num9  bool `json:"9"`
			Num10 bool `json:"10"`
			Num11 bool `json:"11"`
			Num12 bool `json:"12"`
			Num13 bool `json:"13"`
			Num14 bool `json:"14"`
			Num15 bool `json:"15"`
			Num16 bool `json:"16"`
			Num18 bool `json:"18"`
			Num19 bool `json:"19"`
			Num23 bool `json:"23"`
			Num24 bool `json:"24"`
			Num26 bool `json:"26"`
			Num27 bool `json:"27"`
			Num28 bool `json:"28"`
			Num29 bool `json:"29"`
			Num31 bool `json:"31"`
			Num32 bool `json:"32"`
			Num34 bool `json:"34"`
			Num35 bool `json:"35"`
			Num38 bool `json:"38"`
			Num40 bool `json:"40"`
			Num41 bool `json:"41"`
			Num43 bool `json:"43"`
			Num45 bool `json:"45"`
			Num59 bool `json:"59"`
		} `json:"gatekeepers"`
		Qe struct {
			Num0 struct {
				P struct {
					Num4 bool `json:"4"`
					Num7 bool `json:"7"`
					Num8 bool `json:"8"`
					Num9 bool `json:"9"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"0"`
			Num2 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"2"`
			Num4 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"4"`
			Num5 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"5"`
			Num6 struct {
				P struct {
					Num1  bool `json:"1"`
					Num5  bool `json:"5"`
					Num6  bool `json:"6"`
					Num7  bool `json:"7"`
					Num9  bool `json:"9"`
					Num10 bool `json:"10"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"6"`
			Num7 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"7"`
			Num10 struct {
				P struct {
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"10"`
			Num12 struct {
				P struct {
					Num0 int `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"12"`
			Num13 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"13"`
			Num16 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"16"`
			Num17 struct {
				P struct {
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"17"`
			Num18 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"18"`
			Num19 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"19"`
			Num21 struct {
				P struct {
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"21"`
			Num22 struct {
				P struct {
					Num0 bool    `json:"0"`
					Num1 bool    `json:"1"`
					Num2 float64 `json:"2"`
					Num3 float64 `json:"3"`
					Num4 float64 `json:"4"`
					Num5 float64 `json:"5"`
					Num7 bool    `json:"7"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"22"`
			Num23 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"23"`
			Num25 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"25"`
			Num26 struct {
				P struct {
					Num0 string `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"26"`
			Num27 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"27"`
			Num28 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"28"`
			Num29 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"29"`
			Num30 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"30"`
			Num31 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"31"`
			Num33 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"33"`
			Num34 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"34"`
			Num36 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
					Num3 bool `json:"3"`
					Num4 bool `json:"4"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"36"`
			Num37 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"37"`
			Num38 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"38"`
			Num39 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"39"`
			Num40 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"40"`
			Num41 struct {
				P struct {
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"41"`
			Num42 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"42"`
			Num43 struct {
				P struct {
					Num0 bool `json:"0"`
					Num1 bool `json:"1"`
					Num2 bool `json:"2"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"43"`
			Num44 struct {
				P struct {
					Num1 string `json:"1"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"44"`
			Num45 struct {
				P struct {
					Num0 bool   `json:"0"`
					Num1 string `json:"1"`
					Num2 bool   `json:"2"`
					Num4 int    `json:"4"`
					Num5 bool   `json:"5"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"45"`
			Num46 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"46"`
			Num47 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"47"`
			Num48 struct {
				P struct {
					Num0 bool `json:"0"`
				} `json:"p"`
				Qex bool `json:"qex"`
			} `json:"48"`
			AppUpsell struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"app_upsell"`
			IglAppUpsell struct {
				G string `json:"g"`
				P struct {
					HasIgliteContentAndLink string `json:"has_iglite_content_and_link"`
					HasNoUpsell             string `json:"has_no_upsell"`
					HasOnlyIgliteLink       string `json:"has_only_iglite_link"`
				} `json:"p"`
			} `json:"igl_app_upsell"`
			Notif struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"notif"`
			LogCont struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"log_cont"`
			Onetaplogin struct {
				G string `json:"g"`
				P struct {
					AfterLogin     string `json:"after_login"`
					StorageVersion string `json:"storage_version"`
				} `json:"p"`
			} `json:"onetaplogin"`
			MultiregIter struct {
				G string `json:"g"`
				P struct {
					HasBackRemoved string `json:"has_back_removed"`
				} `json:"p"`
			} `json:"multireg_iter"`
			FelixClearFbCookie struct {
				G string `json:"g"`
				P struct {
					IsEnabled string `json:"is_enabled"`
					Blacklist string `json:"blacklist"`
				} `json:"p"`
			} `json:"felix_clear_fb_cookie"`
			FelixCreationDurationLimits struct {
				G string `json:"g"`
				P struct {
					MaximumLengthSeconds string `json:"maximum_length_seconds"`
					MinimumLengthSeconds string `json:"minimum_length_seconds"`
				} `json:"p"`
			} `json:"felix_creation_duration_limits"`
			FelixCreationFbCrossposting struct {
				G string `json:"g"`
				P struct {
					IsEnabled string `json:"is_enabled"`
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting"`
			FelixCreationFbCrosspostingV2 struct {
				G string `json:"g"`
				P struct {
					IsEnabled      string `json:"is_enabled"`
					DisplayVersion string `json:"display_version"`
				} `json:"p"`
			} `json:"felix_creation_fb_crossposting_v2"`
			FelixCreationValidation struct {
				G string `json:"g"`
				P struct {
					EditVideoControls                  string `json:"edit_video_controls"`
					DescriptionMaximumLength           string `json:"description_maximum_length"`
					MaxVideoSizeInBytes                string `json:"max_video_size_in_bytes"`
					MinimumLengthForFeedPreviewSeconds string `json:"minimum_length_for_feed_preview_seconds"`
					TitleMaximumLength                 string `json:"title_maximum_length"`
					ValidCoverMimeTypes                string `json:"valid_cover_mime_types"`
					ValidVideoExtensions               string `json:"valid_video_extensions"`
					ValidVideoMimeTypes                string `json:"valid_video_mime_types"`
				} `json:"p"`
			} `json:"felix_creation_validation"`
			MwebTopicalExplore struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"mweb_topical_explore"`
			PostOptions struct {
				G string `json:"g"`
				P struct {
					UseRefactor     string `json:"use_refactor"`
					EnableIgtvEmbed string `json:"enable_igtv_embed"`
				} `json:"p"`
			} `json:"post_options"`
			Iglscioi struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"iglscioi"`
			StickerTray struct {
				G string `json:"g"`
				P struct {
				} `json:"p"`
			} `json:"sticker_tray"`
			WebSentry struct {
				G string `json:"g"`
				P struct {
					ShowFeedback string `json:"show_feedback"`
				} `json:"p"`
			} `json:"web_sentry"`
		} `json:"qe"`
		ProbablyHasApp bool `json:"probably_has_app"`
		Cb             bool `json:"cb"`
	} `json:"to_cache"`
	DeviceID      interface{} `json:"device_id"`
	RolloutHash   string      `json:"rollout_hash"`
	BundleVariant string      `json:"bundle_variant"`
	IsCanary      bool        `json:"is_canary"`
}
