package queries

const (
	ADD_CLUB = `
		INSERT INTO 
		club (club_name, league_id)
		VALUES (?,?)
	`
	ADD_PLAYER_TO_CLUB = `
		INSERT INTO 
		player_ext_ref (player_id, club_id)
		VALUES (?,?)
	`
	REMOVE_PLAYER_FROM_CLUB = `
		delete from player_ext_ref where player_id = ? and club_id = ?
	`

	FIND_CLUB_BASE_QUERY = `
		select cl.club_id, cl.club_name,
			l.league_id, l.league_name,
			c.country_id, c.country_name,
			c1.continent_id, c1.continent_name,
			cf.confederation_id, cf.confederation_name,cf.confederation_description,			
			c2.continent_id, c2.continent_name
		from club cl
			inner join league l on l.league_id = cl.league_id
			inner join country c on c.country_id = l.country_id
			inner join confederation cf on cf.continent_id = c.continent_id
			inner join continent c1 on c1.continent_id = c.continent_id
			inner join continent c2 on c2.continent_id = cf.continent_id
	`
	CLUB_ID_CLAUSE = `
		cl.club_id = ?
	`
	LEAGUE_ID_CLAUSE = `
		cl.league_id = ?
	`
	FIND_PLAYER_REFERENCE_BY_CLUB_ID = `
		select player_id from player_ext_ref
		where club_id = ?
	`
	FIND_CLUB_BY_PLAYER_ID = `
		select cl.club_id, cl.club_name,
			l.league_id, l.league_name,
			c.country_id, c.country_name,
			c1.continent_id, c1.continent_name,
			cf.confederation_id, cf.confederation_name,cf.confederation_description,			
			c2.continent_id, c2.continent_name
		from club cl
			inner join league l on l.league_id = cl.league_id
			inner join country c on c.country_id = l.country_id
			inner join confederation cf on cf.continent_id = c.continent_id
			inner join continent c1 on c1.continent_id = c.continent_id
			inner join continent c2 on c2.continent_id = cf.continent_id
			inner join player_ext_ref ref on ref.club_id = cl.club_id
		where ref.player_id = ?
	`
)
