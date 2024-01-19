package queries

const (
	ADD_LEAGUE = `
		insert into league(league_name, country_id, confederation_id)
		values (?, ?, ?)
	`

	FIND_LEAGUES_BASE_QUERY = `
		select l.league_id, l.league_name,
			c.country_id, c.country_name,
			c1.continent_id, c1.continent_name,
			cf.confederation_id, cf.confederation_name,cf.confederation_description,			
			c2.continent_id, c2.continent_name
		from league l
			inner join country c on c.country_id = l.country_id
			inner join confederation cf on cf.continent_id = c.continent_id
			inner join continent c1 on c1.continent_id = c.continent_id
			inner join continent c2 on c2.continent_id = cf.continent_id
	`
	DEFAULT_SEARCH_CLAUSE = `
		l.country_id = 0
	`
	LEAGUE_CLAUSE = `
		l.league_id = ?
	`
	COUNTRY_AND_CONFEDERATION_CLAUSE = `
		c.country_id = ? and cf.confederation_id = ?
	`
	COUNTRY_CLAUSE = `
		c.country_id = ?
	`
	CONFEDERATION_CLAUSE = `
		cf.confederation_id = ?
	`
)
