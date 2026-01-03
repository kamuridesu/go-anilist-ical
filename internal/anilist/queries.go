package anilist

const (
	QueryGetMediaList = `
query GetMediaList ($usrId: Int, $type: MediaType, $forceSingleCompletedList: Boolean, $statusIn: [MediaListStatus]) {
  MediaListCollection (userId: $usrId, type: $type, forceSingleCompletedList: $forceSingleCompletedList, status_in: $statusIn) {
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
