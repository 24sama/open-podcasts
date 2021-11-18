package broadcast

const (
	FreeAlbumType = iota
	VipAlbumType
	PaidAlbumType
)

type AlbumInfo struct {
	Msg               string `json:"msg"`
	Ret               int    `json:"ret"`
	CostTimeAlbumInfo int    `json:"costTimeAlbumInfo"`
	Data              struct {
		Album struct {
			AlbumID     int    `json:"albumId"`
			Title       string `json:"title"`
			IsPaid      bool   `json:"isPaid"`
			TrackCount  int    `json:"tracks"`
			VipFreeType int    `json:"vipFreeType"`
			PriceTypes  []struct {
				FreeTrackIds   string `json:"freeTrackIds"`
				FreeTrackCount int    `json:"freeTrackCount"`
			} `json:"priceTypes"`
			IsFinished int `json:"isFinished"`
		} `json:"album"`
	}
}

func (ai *AlbumInfo) AlbumType() int {
	if ai.Data.Album.IsPaid {
		if ai.Data.Album.VipFreeType == 0 {
			return PaidAlbumType
		}
		return VipAlbumType
	} else {
		return FreeAlbumType
	}
}

type TrackInfo struct {
	TrackID                int     `json:"trackId"`
	TrackRecordID          int     `json:"trackRecordId"`
	UID                    int     `json:"uid"`
	PlayURL64              string  `json:"playUrl64"`
	PlayURL32              string  `json:"playUrl32"`
	PlayPathAacv164        string  `json:"playPathAacv164"`
	PlayPathAacv224        string  `json:"playPathAacv224"`
	Title                  string  `json:"title"`
	Duration               int     `json:"duration"`
	AlbumID                int     `json:"albumId"`
	IsPaid                 bool    `json:"isPaid"`
	IsFree                 bool    `json:"isFree"`
	IsVideo                bool    `json:"isVideo"`
	IsDraft                bool    `json:"isDraft"`
	IsRichAudio            bool    `json:"isRichAudio"`
	IsAuthorized           bool    `json:"isAuthorized"`
	Price                  float64 `json:"price"`
	DiscountedPrice        float64 `json:"discountedPrice"`
	PriceTypeID            int     `json:"priceTypeId"`
	SampleDuration         int     `json:"sampleDuration"`
	PriceTypeEnum          int     `json:"priceTypeEnum"`
	DisplayPrice           string  `json:"displayPrice"`
	DisplayDiscountedPrice string  `json:"displayDiscountedPrice"`
	Type                   int     `json:"type"`
	RelatedID              int     `json:"relatedId"`
	OrderNo                int     `json:"orderNo"`
	IsHoldCopyright        bool    `json:"isHoldCopyright"`
	VipFirstStatus         int     `json:"vipFirstStatus"`
	PaidType               int     `json:"paidType"`
	IsSample               bool    `json:"isSample"`
	ProcessState           int     `json:"processState"`
	CreatedAt              int64   `json:"createdAt"`
	CoverSmall             string  `json:"coverSmall"`
	CoverMiddle            string  `json:"coverMiddle"`
	CoverLarge             string  `json:"coverLarge"`
	Nickname               string  `json:"nickname"`
	SmallLogo              string  `json:"smallLogo"`
	UserSource             int     `json:"userSource"`
	OpType                 int     `json:"opType"`
	IsPublic               bool    `json:"isPublic"`
	Likes                  int     `json:"likes"`
	Playtimes              int     `json:"playtimes"`
	Comments               int     `json:"comments"`
	Shares                 int     `json:"shares"`
	Status                 int     `json:"status"`
	IsTrailer              bool    `json:"isTrailer"`
}

type TrackList struct {
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Data struct {
		List       []*TrackInfo `json:"list"`
		PageID     int          `json:"pageId"`
		PageSize   int          `json:"pageSize"`
		MaxPageID  int          `json:"maxPageId"`
		TotalCount int          `json:"totalCount"`
	} `json:"data"`
}

type AudioInfo struct {
	TrackID         int    `json:"trackId"`
	TrackRecordID   int    `json:"trackRecordId"`
	UID             int    `json:"uid"`
	PlayURL64       string `json:"playUrl64"` //mp3 64kbps
	PlayURL32       string `json:"playUrl32"` //mp3 32kbps
	PlayPathHq      string `json:"playPathHq"`
	PlayPathAacv164 string `json:"playPathAacv164"` //m4a 64kbps
	PlayPathAacv224 string `json:"playPathAacv224"` //m4a 24kbps
	Title           string `json:"title"`
	Duration        int    `json:"duration"`
	AlbumID         int    `json:"albumId"`
	AlbumTitle      string `json:"albumTitle"`
	AlbumImage      string `json:"albumImage"`
	IsPaid          bool   `json:"isPaid"`
	IsFree          bool   `json:"isFree"`
	IsVideo         bool   `json:"isVideo"`
	IsDraft         bool   `json:"isDraft"`
	IsRichAudio     bool   `json:"isRichAudio"`
	IsAuthorized    bool   `json:"isAuthorized"`
	PriceTypeID     int    `json:"priceTypeId"`
	PriceTypeEnum   int    `json:"priceTypeEnum"`
	Type            int    `json:"type"`
	RelatedID       int    `json:"relatedId"`
	OrderNo         int    `json:"orderNo"`
	VipFirstStatus  int    `json:"vipFirstStatus"`
	PaidType        int    `json:"paidType"`
}

type Playlist struct {
	Msg        string       `json:"msg"`
	Ret        int          `json:"ret"`
	MaxPageID  int          `json:"maxPageId"`
	PageSize   int          `json:"pageSize"`
	List       []*AudioInfo `json:"list"`
	PageID     int          `json:"pageId"`
	TotalCount int          `json:"totalCount"`
}

type ChargeTrackInfo struct {
	Ret                  int    `json:"ret"`
	Msg                  string `json:"msg"`
	TrackID              int    `json:"trackId"`
	UID                  int    `json:"uid"`
	AlbumID              int    `json:"albumId"`
	Title                string `json:"title"`
	Domain               string `json:"domain"`
	TotalLength          int    `json:"totalLength"`
	SampleDuration       int    `json:"sampleDuration"`
	SampleLength         int    `json:"sampleLength"`
	IsAuthorized         bool   `json:"isAuthorized"`
	APIVersion           string `json:"apiVersion"`
	Seed                 int    `json:"seed"`
	FileID               string `json:"fileId"`
	BuyKey               string `json:"buyKey"`
	Duration             int    `json:"duration"`
	Ep                   string `json:"ep"`
	HighestQualityLevel  int    `json:"highestQualityLevel"`
	DownloadQualityLevel int    `json:"downloadQualityLevel"`
}

type VipAudioInfo struct {
	Ret                  int    `json:"ret"`
	Msg                  string `json:"msg"`
	TrackID              int    `json:"trackId"`
	UID                  int    `json:"uid"`
	AlbumID              int    `json:"albumId"`
	Title                string `json:"title"`
	Domain               string `json:"domain"`
	TotalLength          int    `json:"totalLength"`
	SampleDuration       int    `json:"sampleDuration"`
	SampleLength         int    `json:"sampleLength"`
	IsAuthorized         bool   `json:"isAuthorized"`
	APIVersion           string `json:"apiVersion"`
	Seed                 int    `json:"seed"`
	FileID               string `json:"fileId"`
	BuyKey               string `json:"buyKey"`
	Duration             int    `json:"duration"`
	Ep                   string `json:"ep"`
	HighestQualityLevel  int    `json:"highestQualityLevel"`
	DownloadQualityLevel int    `json:"downloadQualityLevel"`
	AuthorizedType       int    `json:"authorizedType"`
}

type AudioInfoMobile struct {
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Data struct {
		List []struct {
			TrackID                int     `json:"trackId"`
			TrackRecordID          int     `json:"trackRecordId"`
			UID                    int     `json:"uid"`
			PlayURL64              string  `json:"playUrl64"`
			PlayURL32              string  `json:"playUrl32"`
			PlayPathHq             string  `json:"playPathHq"`
			PlayPathAacv164        string  `json:"playPathAacv164"`
			PlayPathAacv224        string  `json:"playPathAacv224"`
			Title                  string  `json:"title"`
			Duration               int     `json:"duration"`
			AlbumID                int     `json:"albumId"`
			IsPaid                 bool    `json:"isPaid"`
			IsFree                 bool    `json:"isFree"`
			IsVideo                bool    `json:"isVideo"`
			IsDraft                bool    `json:"isDraft"`
			IsRichAudio            bool    `json:"isRichAudio"`
			IsAuthorized           bool    `json:"isAuthorized"`
			Price                  float64 `json:"price"`
			DiscountedPrice        float64 `json:"discountedPrice"`
			PriceTypeID            int     `json:"priceTypeId"`
			SampleDuration         int     `json:"sampleDuration"`
			PriceTypeEnum          int     `json:"priceTypeEnum"`
			DisplayPrice           string  `json:"displayPrice"`
			DisplayDiscountedPrice string  `json:"displayDiscountedPrice"`
			VipPrice               float64 `json:"vipPrice"`
			DisplayVipPrice        string  `json:"displayVipPrice"`
			Type                   int     `json:"type"`
			RelatedID              int     `json:"relatedId"`
			OrderNo                int     `json:"orderNo"`
			IsHoldCopyright        bool    `json:"isHoldCopyright"`
			VipFirstStatus         int     `json:"vipFirstStatus"`
			PaidType               int     `json:"paidType"`
			IsSample               bool    `json:"isSample"`
			ProcessState           int     `json:"processState"`
			CreatedAt              int64   `json:"createdAt"`
			CoverSmall             string  `json:"coverSmall"`
			CoverMiddle            string  `json:"coverMiddle"`
			CoverLarge             string  `json:"coverLarge"`
			Nickname               string  `json:"nickname"`
			SmallLogo              string  `json:"smallLogo"`
			UserSource             int     `json:"userSource"`
			OpType                 int     `json:"opType"`
			IsPublic               bool    `json:"isPublic"`
			Likes                  int     `json:"likes"`
			Playtimes              int     `json:"playtimes"`
			Comments               int     `json:"comments"`
			Shares                 int     `json:"shares"`
			Status                 int     `json:"status"`
			ExpireTime             int64   `json:"expireTime"`
			IsTrailer              bool    `json:"isTrailer"`
		} `json:"list"`
		PageID     int `json:"pageId"`
		PageSize   int `json:"pageSize"`
		MaxPageID  int `json:"maxPageId"`
		TotalCount int `json:"totalCount"`
	} `json:"data"`
}

type UserInfo struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		UID           int         `json:"uid"`
		RealUID       int         `json:"realUid"`
		Nickname      string      `json:"nickname"`
		LogoPic       string      `json:"logoPic"`
		IsLoginBan    bool        `json:"isLoginBan"`
		IsVerified    bool        `json:"isVerified"`
		Ptitle        interface{} `json:"ptitle"`
		Mobile        string      `json:"mobile"`
		IsRobot       bool        `json:"isRobot"`
		VerifyType    int         `json:"verifyType"`
		IsVip         bool        `json:"isVip"`
		VipExpireTime int         `json:"vipExpireTime"`
		AnchorGrade   int         `json:"anchorGrade"`
		UserGrade     int         `json:"userGrade"`
		UserTitle     interface{} `json:"userTitle"`
		LogoType      int         `json:"logoType"`
	} `json:"data"`
}

type QRCode struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	QrID string `json:"qrId"`
	Img  string `json:"img"`
}

type QRCodeStatus struct {
	Ret                int         `json:"ret"`
	Msg                string      `json:"msg"`
	BizKey             interface{} `json:"bizKey"`
	UID                int         `json:"uid"`
	Token              interface{} `json:"token"`
	UserType           interface{} `json:"userType"`
	IsFirst            bool        `json:"isFirst"`
	ToSetPwd           bool        `json:"toSetPwd"`
	LoginType          string      `json:"loginType"`
	MobileMask         string      `json:"mobileMask"`
	MobileCipher       string      `json:"mobileCipher"`
	CaptchaInfo        interface{} `json:"captchaInfo"`
	Avatar             interface{} `json:"avatar"`
	ThirdpartyAvatar   interface{} `json:"thirdpartyAvatar"`
	ThirdpartyNickname interface{} `json:"thirdpartyNickname"`
	SmsKey             interface{} `json:"smsKey"`
	ThirdpartyID       interface{} `json:"thirdpartyId"`
}

type AlbumSearchResult struct {
	Ret  int `json:"ret"`
	Data struct {
		AlbumViews struct {
			Albums []struct {
				AlbumInfo struct {
					Play              int           `json:"play"`
					UserSource        string        `json:"user_source"`
					CoverPath         string        `json:"cover_path"`
					Title             string        `json:"title"`
					UID               int           `json:"uid"`
					CategoryID        int           `json:"category_id"`
					Intro             string        `json:"intro"`
					ID                int           `json:"id"`
					IsPaid            bool          `json:"is_paid"`
					IsFinished        int           `json:"is_finished"`
					CategoryTitle     string        `json:"category_title"`
					CreatedAt         int64         `json:"created_at"`
					AlbumType         int           `json:"albumType"`
					Type              string        `json:"type"`
					IsV               bool          `json:"is_v"`
					RefundSupportType int           `json:"refund_support_type"`
					CountComment      int           `json:"count_comment"`
					Score             float64       `json:"score"`
					PriceTypeID       int           `json:"price_type_id"`
					UpdatedAt         int64         `json:"updated_at"`
					IsVipFree         bool          `json:"isVipFree"`
					Nickname          string        `json:"nickname"`
					VerifyType        int           `json:"verify_type"`
					VipFreeType       int           `json:"vipFreeType"`
					AlbumVipPayType   int           `json:"albumVipPayType"`
					SerialState       int           `json:"serialState"`
					PriceTypeEnum     int           `json:"price_type_enum"`
					Tracks            int           `json:"tracks"`
					IsNoCopyright     bool          `json:"isNoCopyright"`
					CommentsCount     int           `json:"comments_count"`
					PriceTypes        []interface{} `json:"price_types"`
				} `json:"albumInfo"`
				PageURIInfo struct {
					ID           int    `json:"id"`
					CategoryID   int    `json:"categoryId"`
					CategoryName string `json:"categoryName"`
					CategoryCode string `json:"categoryCode"`
					Pinyin       string `json:"pinyin"`
					URL          string `json:"url"`
				} `json:"pageUriInfo"`
			} `json:"albums"`
			TotalPage   int `json:"totalPage"`
			Start       int `json:"start"`
			PageSize    int `json:"pageSize"`
			CurrentPage int `json:"currentPage"`
			Total       int `json:"total"`
		} `json:"albumViews"`
		TrackViews struct {
			Tracks      []interface{} `json:"tracks"`
			PageSize    int           `json:"pageSize"`
			CurrentPage int           `json:"currentPage"`
		} `json:"trackViews"`
		UserViews struct {
			Users       []interface{} `json:"users"`
			PageSize    int           `json:"pageSize"`
			CurrentPage int           `json:"currentPage"`
		} `json:"userViews"`
		User1Views struct {
			Users []interface{} `json:"users"`
		} `json:"user1Views"`
		User2Views struct {
			Users []interface{} `json:"users"`
		} `json:"user2Views"`
		SpecialViews struct {
			Specials    []interface{} `json:"specials"`
			Categories  []interface{} `json:"categories"`
			PageSize    int           `json:"pageSize"`
			CurrentPage int           `json:"currentPage"`
		} `json:"specialViews"`
		RecommendItems []interface{} `json:"recommendItems"`
		IsIllegal      bool          `json:"isIllegal"`
	} `json:"data"`
	Context struct {
		CurrentUser struct {
			ID           int    `json:"id"`
			Nickname     string `json:"nickname"`
			IsVip        bool   `json:"isVip"`
			IsNewCreated bool   `json:"isNewCreated"`
			Logo         string `json:"logo"`
		} `json:"currentUser"`
		BasicRequestContext struct {
			IsHybrid                    bool `json:"isHybrid"`
			IsEmbedded3RdPartner        bool `json:"isEmbedded3rdPartner"`
			IsKnowAmbassadorDistributor bool `json:"isKnowAmbassadorDistributor"`
			IsM2WapHost                 bool `json:"isM2WapHost"`
		} `json:"basicRequestContext"`
	} `json:"context"`
}
