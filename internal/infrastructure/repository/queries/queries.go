package queries

const (
	AddWordList    = "INSERT INTO word_list (word, is_filtred) VALUES ($1, $2) RETURNING ID;"
	DeleteWordList = "DELETE FROM word_list WHERE ID = $1;"
	GetWordList    = "SELECT ID, word, is_filtred FROM word_list WHERE ID = $1;"

	AddUrlList    = "INSERT INTO url_list (link) VALUES ($1) RETURNING ID;"
	DeleteUrlList = "DELETE FROM url_list WHERE ID = $1;"
	GetUrlList    = "SELECT ID, link FROM url_list WHERE ID = $1;"

	AddWordLocation    = "INSERT INTO word_location (fk_word_ID, fk_url_ID, location) VALUES ($1, $2, $3) RETURNING ID;"
	DeleteWordLocation = "DELETE FROM word_location WHERE ID = $1;"
	GetWordLocation    = "SELECT ID, fk_word_ID, fk_url_ID, location FROM word_location WHERE ID = $1;"

	AddLinkBetweenURL    = "INSERT INTO link_between_url (fk_from_url_ID, fk_to_url_ID) VALUES ($1, $2) RETURNING ID;"
	DeleteLinkBetweenURL = "DELETE FROM link_between_url WHERE ID = $1;"
	GetLinkBetweenURL    = "SELECT ID, fk_from_url_ID, fk_to_url_ID FROM link_between_url WHERE ID = $1;"

	AddLinkWord    = "INSERT INTO link_word (fk_wordId, fk_linkId) VALUES ($1, $2) RETURNING ID;"
	DeleteLinkWord = "DELETE FROM link_word WHERE ID = $1;"
	GetLinkWord    = "SELECT ID, fk_wordId, fk_linkId FROM link_word WHERE ID = $1;"
)
