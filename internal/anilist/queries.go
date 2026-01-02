package anilist

const (
	QueryGetMediaList = `
query GetMediaList ($usrId: Int, $type: MediaType, $statusIn: [MediaListStatus]) {
  MediaListCollection (userId: $usrId, type: $type, status_in: $statusIn) {
    lists {
      entries {
        media {
          title {
            userPreferred
          }
          airingSchedule {
            nodes {
              airingAt
              timeUntilAiring
              episode
            }
          }
          status
        }
      }
    }
  }
}`
	QueryGetUserId = `
query GetUserId ($name: String) {
  User (name: $name) {
    name
    id
  }
}`
)
