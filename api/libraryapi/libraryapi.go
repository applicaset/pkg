// Package libraryapi is the wire shape of xonyagar's library service: types and paths only, so the
// service, its client and the site share one definition without importing each other.
package libraryapi

import "time"

// Every request names the acting user. The service trusts it: only sibling services on the private
// network reach these paths, and the site has already resolved the session. An empty ActorRef is a
// visitor without an account.
const (
	PathLibraries        = "/v1/libraries"
	PathGetLibrary       = "/v1/get-library"
	PathCreateLibrary    = "/v1/create-library"
	PathRenameLibrary    = "/v1/rename-library"
	PathDeleteLibrary    = "/v1/delete-library"
	PathMembers          = "/v1/members"
	PathInvite           = "/v1/invite"
	PathMyInvitations    = "/v1/my-invitations"
	PathAnswerInvitation = "/v1/answer-invitation"
	PathCancelInvitation = "/v1/cancel-invitation"
	PathSetMemberRole    = "/v1/set-member-role"
	PathRemoveMember     = "/v1/remove-member"
	PathGetAccess        = "/v1/get-access"
	PathSetCurator       = "/v1/set-curator"
	PathCheckUpload      = "/v1/check-upload"
	PathImportTrack      = "/v1/import-track"
	PathListTracks       = "/v1/list-tracks"
	PathGetTracks        = "/v1/get-tracks"
	PathEditTracks       = "/v1/edit-tracks"
	PathTrashTracks      = "/v1/trash-tracks"
	PathRestoreTracks    = "/v1/restore-tracks"
	PathEmptyTrash       = "/v1/empty-trash"
	PathCopyTracks       = "/v1/copy-tracks"
	PathStream           = "/v1/stream"
	PathRecordPlay       = "/v1/record-play"
	PathDownloadTrack    = "/v1/download-track"
	PathHome             = "/v1/home"
	PathListArtists      = "/v1/list-artists"
	PathGetArtist        = "/v1/get-artist"
	PathEditArtist       = "/v1/edit-artist"
	PathMergeArtists     = "/v1/merge-artists"
	PathSetArtistPhoto   = "/v1/set-artist-photo"
	PathListAlbums       = "/v1/list-albums"
	PathGetAlbum         = "/v1/get-album"
	PathEditAlbum        = "/v1/edit-album"
	PathMergeAlbums      = "/v1/merge-albums"
	PathSetAlbumCover    = "/v1/set-album-cover"
	PathListPlaylists    = "/v1/list-playlists"
	PathCreatePlaylist   = "/v1/create-playlist"
	PathGetPlaylist      = "/v1/get-playlist"
	PathEditPlaylist     = "/v1/edit-playlist"
	PathDeletePlaylist   = "/v1/delete-playlist"
	PathAddToPlaylist    = "/v1/add-to-playlist"
	PathArrangePlaylist  = "/v1/arrange-playlist"
	PathLookupProviders  = "/v1/lookup-providers"
	PathLookupSearch     = "/v1/lookup-search"
	PathLookupRelease    = "/v1/lookup-release"
	PathApplyRelease     = "/v1/apply-release"
	// PathEvents answers with a text/event-stream of changed track ids, empty when something other
	// than one track changed, until the caller hangs up.
	PathEvents = "/v1/events"
)

type Library struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Kind      string    `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Membership struct {
	Library Library `json:"library"`
	Role    string  `json:"role"`
}

type Member struct {
	UserRef  string    `json:"user_ref"`
	Role     string    `json:"role"`
	JoinedAt time.Time `json:"joined_at"`
}

type Invitation struct {
	ID         string    `json:"id"`
	LibraryID  string    `json:"library_id"`
	InviteeRef string    `json:"invitee_ref"`
	Role       string    `json:"role"`
	InvitedBy  string    `json:"invited_by"`
	CreatedAt  time.Time `json:"created_at"`
}

type InvitationView struct {
	Invitation Invitation `json:"invitation"`
	Library    Library    `json:"library"`
}

type Artist struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Bio         string   `json:"bio"`
	Links       []string `json:"links"`
	PhotoFileID string   `json:"photo_file_id"`
}

type Album struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	ArtistID    string    `json:"artist_id"`
	Year        int       `json:"year"`
	Kind        string    `json:"kind"`
	CoverFileID string    `json:"cover_file_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type Track struct {
	ID             string     `json:"id"`
	LibraryID      string     `json:"library_id"`
	Title          string     `json:"title"`
	ArtistIDs      []string   `json:"artist_ids"`
	AlbumID        string     `json:"album_id"`
	TrackNo        int        `json:"track_no"`
	DiscNo         int        `json:"disc_no"`
	Year           int        `json:"year"`
	Genre          string     `json:"genre"`
	Lyrics         string     `json:"lyrics"`
	CoverFileID    string     `json:"cover_file_id"`
	DurationMS     int64      `json:"duration_ms"`
	Status         string     `json:"status"`
	Error          string     `json:"error"`
	OriginalFileID string     `json:"original_file_id"`
	OriginalName   string     `json:"original_name"`
	Size           int64      `json:"size"`
	Codec          string     `json:"codec"`
	Bitrate        int64      `json:"bitrate"`
	SampleRate     int        `json:"sample_rate"`
	Channels       int        `json:"channels"`
	Revision       int        `json:"revision"`
	CreatedAt      time.Time  `json:"created_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
	DeleteAfter    *time.Time `json:"delete_after,omitempty"`
}

// TrackList is tracks with every artist and album they name, keyed by id.
type TrackList struct {
	Tracks  []Track           `json:"tracks"`
	Total   int               `json:"total"`
	Artists map[string]Artist `json:"artists"`
	Albums  map[string]Album  `json:"albums"`
}

type Rendition struct {
	DurationMS int64     `json:"duration_ms"`
	Variants   []Variant `json:"variants"`
}

type Variant struct {
	Bitrate    int       `json:"bitrate"`
	InitFileID string    `json:"init_file_id"`
	Segments   []Segment `json:"segments"`
}

type Segment struct {
	FileID     string `json:"file_id"`
	DurationMS int64  `json:"duration_ms"`
}

type Playlist struct {
	ID          string    `json:"id"`
	LibraryID   string    `json:"library_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TrackCount  int       `json:"track_count"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ActorRequest struct {
	ActorRef string `json:"actor_ref"`
}

// Access names the actor and the library they act in.
type Access struct {
	ActorRef  string `json:"actor_ref"`
	LibraryID string `json:"library_id"`
}

type AccessRequest struct {
	Access Access `json:"access"`
}

type Empty struct{}

type LibrariesResponse struct {
	Memberships []Membership `json:"memberships"`
}

type MembershipResponse struct {
	Membership Membership `json:"membership"`
}

type CreateLibraryRequest struct {
	ActorRef string `json:"actor_ref"`
	Name     string `json:"name"`
}

type LibraryResponse struct {
	Library Library `json:"library"`
}

type RenameLibraryRequest struct {
	Access Access `json:"access"`
	Name   string `json:"name"`
}

type MembersResponse struct {
	Members     []Member     `json:"members"`
	Invitations []Invitation `json:"invitations"`
}

type InviteRequest struct {
	Access     Access `json:"access"`
	InviteeRef string `json:"invitee_ref"`
	Role       string `json:"role"`
}

type InvitationsResponse struct {
	Invitations []InvitationView `json:"invitations"`
}

type AnswerInvitationRequest struct {
	ActorRef     string `json:"actor_ref"`
	InvitationID string `json:"invitation_id"`
	Accept       bool   `json:"accept"`
}

type InvitationRequest struct {
	Access       Access `json:"access"`
	InvitationID string `json:"invitation_id"`
}

type MemberRequest struct {
	Access  Access `json:"access"`
	UserRef string `json:"user_ref"`
	Role    string `json:"role"`
}

type GetAccessRequest struct {
	ActorRef   string `json:"actor_ref"`
	SubjectRef string `json:"subject_ref"`
}

type AccessResponse struct {
	Admin   bool `json:"admin"`
	Curator bool `json:"curator"`
}

type SetCuratorRequest struct {
	ActorRef string `json:"actor_ref"`
	UserRef  string `json:"user_ref"`
	Allowed  bool   `json:"allowed"`
}

type ImportTrackRequest struct {
	Access       Access `json:"access"`
	FileID       string `json:"file_id"`
	OriginalName string `json:"original_name"`
}

type TrackResponse struct {
	Track Track `json:"track"`
}

type TrackQuery struct {
	AlbumID  string `json:"album_id"`
	ArtistID string `json:"artist_id"`
	Search   string `json:"search"`
	Trashed  bool   `json:"trashed"`
	Sort     string `json:"sort"`
	Desc     bool   `json:"desc"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type ListTracksRequest struct {
	Access Access     `json:"access"`
	Query  TrackQuery `json:"query"`
}

type TrackListResponse struct {
	List TrackList `json:"list"`
}

type TrackIDsRequest struct {
	Access   Access   `json:"access"`
	TrackIDs []string `json:"track_ids"`
}

// TrackPatch changes the fields that are set.
type TrackPatch struct {
	Title       *string   `json:"title,omitempty"`
	Artists     *[]string `json:"artists,omitempty"`
	Album       *string   `json:"album,omitempty"`
	AlbumArtist *string   `json:"album_artist,omitempty"`
	TrackNo     *int      `json:"track_no,omitempty"`
	DiscNo      *int      `json:"disc_no,omitempty"`
	Year        *int      `json:"year,omitempty"`
	Genre       *string   `json:"genre,omitempty"`
	Lyrics      *string   `json:"lyrics,omitempty"`
}

type EditTracksRequest struct {
	Access   Access     `json:"access"`
	TrackIDs []string   `json:"track_ids"`
	Patch    TrackPatch `json:"patch"`
}

type CopyTracksRequest struct {
	Access          Access   `json:"access"`
	TrackIDs        []string `json:"track_ids"`
	TargetLibraryID string   `json:"target_library_id"`
}

type CountResponse struct {
	Count int `json:"count"`
}

type TrackRequest struct {
	Access  Access `json:"access"`
	TrackID string `json:"track_id"`
}

type StreamResponse struct {
	Track     Track     `json:"track"`
	Rendition Rendition `json:"rendition"`
}

type DownloadResponse struct {
	FileID   string `json:"file_id"`
	FileName string `json:"file_name"`
}

type HomeResponse struct {
	Recent       *TrackList `json:"recent,omitempty"`
	MostPlayed   TrackList  `json:"most_played"`
	LatestTracks TrackList  `json:"latest_tracks"`
	LatestAlbums []Album    `json:"latest_albums"`
	Artists      []Artist   `json:"artists"`
}

type ListArtistsRequest struct {
	Access Access `json:"access"`
	Search string `json:"search"`
	Limit  int    `json:"limit"`
}

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
}

type ArtistRequest struct {
	Access   Access `json:"access"`
	ArtistID string `json:"artist_id"`
}

type ArtistPageResponse struct {
	Artist    Artist    `json:"artist"`
	TopTracks TrackList `json:"top_tracks"`
	Albums    []Album   `json:"albums"`
	Singles   []Album   `json:"singles"`
	AppearsOn []Album   `json:"appears_on"`
}

type ArtistPatch struct {
	Name  *string   `json:"name,omitempty"`
	Bio   *string   `json:"bio,omitempty"`
	Links *[]string `json:"links,omitempty"`
	// AllowMerge lets a name another artist has merge the two. Without it, that edit is a conflict.
	AllowMerge bool `json:"allow_merge,omitempty"`
}

type EditArtistRequest struct {
	Access   Access      `json:"access"`
	ArtistID string      `json:"artist_id"`
	Patch    ArtistPatch `json:"patch"`
}

type ArtistResponse struct {
	Artist Artist `json:"artist"`
}

type MergeRequest struct {
	Access    Access   `json:"access"`
	TargetID  string   `json:"target_id"`
	SourceIDs []string `json:"source_ids"`
}

// PictureRequest hands the library an uploaded image for an album or artist. The library owns the
// file afterwards.
type PictureRequest struct {
	Access   Access `json:"access"`
	TargetID string `json:"target_id"`
	FileID   string `json:"file_id"`
}

type AlbumQuery struct {
	ArtistID string `json:"artist_id"`
	Search   string `json:"search"`
	Sort     string `json:"sort"`
	Limit    int    `json:"limit"`
}

type ListAlbumsRequest struct {
	Access Access     `json:"access"`
	Query  AlbumQuery `json:"query"`
}

type AlbumsResponse struct {
	Albums []Album `json:"albums"`
}

type AlbumRequest struct {
	Access  Access `json:"access"`
	AlbumID string `json:"album_id"`
}

type AlbumPageResponse struct {
	Album  Album     `json:"album"`
	Artist *Artist   `json:"artist,omitempty"`
	Tracks TrackList `json:"tracks"`
}

type AlbumPatch struct {
	Title  *string `json:"title,omitempty"`
	Artist *string `json:"artist,omitempty"`
	Year   *int    `json:"year,omitempty"`
	Kind   *string `json:"kind,omitempty"`
	// AllowMerge lets a title and artist another album has merge the two. Without it, that edit
	// is a conflict.
	AllowMerge bool `json:"allow_merge,omitempty"`
}

type EditAlbumRequest struct {
	Access  Access     `json:"access"`
	AlbumID string     `json:"album_id"`
	Patch   AlbumPatch `json:"patch"`
}

type AlbumResponse struct {
	Album Album `json:"album"`
}

type PlaylistsResponse struct {
	Playlists []Playlist `json:"playlists"`
}

type CreatePlaylistRequest struct {
	Access      Access `json:"access"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PlaylistResponse struct {
	Playlist Playlist `json:"playlist"`
}

type PlaylistRequest struct {
	Access     Access `json:"access"`
	PlaylistID string `json:"playlist_id"`
}

type PlaylistPageResponse struct {
	Playlist Playlist  `json:"playlist"`
	EntryIDs []string  `json:"entry_ids"`
	Tracks   TrackList `json:"tracks"`
}

type EditPlaylistRequest struct {
	Access      Access `json:"access"`
	PlaylistID  string `json:"playlist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PlaylistTracksRequest struct {
	Access     Access   `json:"access"`
	PlaylistID string   `json:"playlist_id"`
	TrackIDs   []string `json:"track_ids"`
}

type ArrangePlaylistRequest struct {
	Access     Access   `json:"access"`
	PlaylistID string   `json:"playlist_id"`
	EntryIDs   []string `json:"entry_ids"`
}

// Provider is an online catalog a user may look tags up in.
type Provider struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProvidersResponse struct {
	Providers []Provider `json:"providers"`
}

type LookupSearchRequest struct {
	Access     Access `json:"access"`
	ProviderID string `json:"provider_id"`
	Artist     string `json:"artist"`
	Album      string `json:"album"`
}

// Release is one candidate an online catalog offers.
type Release struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Artist  string         `json:"artist"`
	Year    int            `json:"year"`
	Country string         `json:"country"`
	Tracks  []ReleaseTrack `json:"tracks"`
	// HasCover says the catalog holds a front cover for the release.
	HasCover bool `json:"has_cover"`
}

type ReleaseTrack struct {
	DiscNo  int    `json:"disc_no"`
	TrackNo int    `json:"track_no"`
	Title   string `json:"title"`
	Artist  string `json:"artist"`
}

type ReleasesResponse struct {
	Releases []Release `json:"releases"`
}

type LookupReleaseRequest struct {
	Access     Access `json:"access"`
	ProviderID string `json:"provider_id"`
	ReleaseID  string `json:"release_id"`
}

type ReleaseResponse struct {
	Release Release `json:"release"`
}

// ApplyReleaseRequest writes a release's tags to the album's tracks, matched by disc and track
// number, and optionally its cover.
type ApplyReleaseRequest struct {
	Access     Access `json:"access"`
	ProviderID string `json:"provider_id"`
	ReleaseID  string `json:"release_id"`
	AlbumID    string `json:"album_id"`
	Cover      bool   `json:"cover"`
}
