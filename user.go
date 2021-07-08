package wildduck

type User struct {
	Id               string                 `json:"id,omitempty"`
	Username         string                 `json:"username,omitempty"`
	Password         string                 `json:"password,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Address          string                 `json:"address,omitempty"`
	Retention        int                    `json:"retention,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Targets          []string               `json:"targets,omitempty"`
	Enabled2Fa       []string               `json:"enabled2fa,omitempty"`
	AutoReply        bool                   `json:"autoreply,omitempty"`
	EncryptMessages  bool                   `json:"encryptMessages,omitempty"`
	EncryptForwarded bool                   `json:"encryptForwarded,omitempty"`
	PubKey           string                 `json:"pubKey,omitempty"`
	KeyInfo          KeyInfo                `json:"keyInfo,omitempty"`
	MetaData         map[string]interface{} `json:"metaData,omitempty"`
	InternalData     map[string]interface{} `json:"internalData,omitempty"`
	SpamLevel        int                    `json:"spamLevel,omitempty"`
	Limits           Limits                 `json:"limits,omitempty"`
	FromWhiteList    []string               `json:"fromWhiteList,omitempty"`
	DisabledScopes   []string               `json:"disabledScopes,omitempty"`
	HasPasswordSet   bool                   `json:"hasPasswordSet,omitempty"`
	Disabled         bool                   `json:"disabled,omitempty"`
	Suspended        bool                   `json:"suspended,omitempty"`
	Success          bool                   `json:"success,omitempty"`
}

type UserParams struct {
	Id               string                 `json:"id,omitempty"`
	Username         string                 `json:"username,omitempty"`
	Password         string                 `json:"password,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Address          string                 `json:"address,omitempty"`
	Retention        int                    `json:"retention,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Targets          []string               `json:"targets,omitempty"`
	Enabled2Fa       bool                   `json:"enabled2fa,omitempty"`
	AutoReply        bool                   `json:"autoreply,omitempty"`
	EncryptMessages  bool                   `json:"encryptMessages,omitempty"`
	EncryptForwarded bool                   `json:"encryptForwarded,omitempty"`
	PubKey           string                 `json:"pubKey,omitempty"`
	MetaData         map[string]interface{} `json:"metaData,omitempty"`
	InternalData     map[string]interface{} `json:"internalData,omitempty"`
	SpamLevel        int                    `json:"spamLevel,omitempty"`
	FromWhiteList    []string               `json:"fromWhiteList,omitempty"`
	DisabledScopes   []string               `json:"disabledScopes,omitempty"`
	HasPasswordSet   bool                   `json:"hasPasswordSet,omitempty"`
	Disabled         bool                   `json:"disabled,omitempty"`
	Suspended        bool                   `json:"suspended,omitempty"`
	Quota            int64                  `json:"quota,omitempty"`
}

type KeyInfo struct {
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
}
type Quota struct {
	Allowed int64 `json:"allowed,omitempty"`
	Used    int64 `json:"used,omitempty"`
	TTL     int   `json:"ttl,omitempty"`
}

type Limits struct {
	Quota              Quota `json:"quota,omitempty"`
	Recipients         Quota `json:"recipients,omitempty"`
	Forwards           Quota `json:"forwards,omitempty"`
	Received           Quota `json:"received,omitempty"`
	ImapUpload         Quota `json:"imapUpload,omitempty"`
	ImapDownload       Quota `json:"imapDownload,omitempty"`
	Pop3Download       Quota `json:"pop3Download,omitempty"`
	ImapMaxConnections Quota `json:"imapMaxConnections,omitempty"`
}

type AllUsersResponse struct {
	Success        bool        `json:"success,omitempty"`
	Total          int         `json:"total,omitempty"`
	Page           int         `json:"page,omitempty"`
	PreviousCursor interface{} `json:"previousCursor,omitempty"` // false if none otherwise string
	NextCursor     interface{} `json:"nextCursor,omitempty"`     // false if none otherwise string
	Results        []User      `json:"results,omitempty"`
	Error          string      `json:"error,omitempty"`
}
