package queries

const (
	AddWordList    = "INSERT INTO word_list (word, is_filtred) VALUES ($1, $2) RETURNING ID;"
	GetWordList    = "SELECT ID, word, is_filtred FROM word_list WHERE ID = $1;"
	UpdateWordList = "UPDATE word_list SET word = $1, is_filtred = $2 WHERE id = $3;"
	DeleteWordList = "DELETE FROM word_list WHERE ID = $1;"

	AddUrlList    = "INSERT INTO url_list (link, is_parsed) VALUES ($1, $2) RETURNING ID;"
	GetUrlList    = "SELECT ID, link FROM url_list WHERE link = $1;"
	UpdateUrlList = "UPDATE url_list SET link = $1, is_parsed = $2 WHERE id = $3;"
	DeleteUrlList = "DELETE FROM url_list WHERE ID = $1;"

	AddWordLocation    = "INSERT INTO word_location (fk_word_ID, fk_url_ID, location) VALUES ($1, $2, $3) RETURNING ID;"
	GetWordLocation    = "SELECT ID, fk_word_ID, fk_url_ID, location FROM word_location WHERE ID = $1;"
	UpdateWordLocation = "UPDATE word_location SET word_id = $1, url_id = $2, location = $3 WHERE id = $4;"
	DeleteWordLocation = "DELETE FROM word_location WHERE ID = $1;"

	AddLinkBetweenURL    = "INSERT INTO link_between_url (fk_from_url_ID, fk_to_url_ID) VALUES ($1, $2) RETURNING ID;"
	GetLinkBetweenURL    = "SELECT ID, fk_from_url_ID, fk_to_url_ID FROM link_between_url WHERE ID = $1;"
	UpdateLinkBetweenURL = "UPDATE link_between_url SET from_url_id = $1, to_url_id = $2 WHERE id = $3;"
	DeleteLinkBetweenURL = "DELETE FROM link_between_url WHERE ID = $1;"

	AddLinkWord    = "INSERT INTO link_word (fk_wordId, fk_linkId) VALUES ($1, $2) RETURNING ID;"
	GetLinkWord    = "SELECT ID, fk_wordId, fk_linkId FROM link_word WHERE ID = $1;"
	UpdateLinkWord = "UPDATE link_word SET word_id = $1, link_id = $2 WHERE id = $3;"
	DeleteLinkWord = "DELETE FROM link_word WHERE ID = $1;"
)
