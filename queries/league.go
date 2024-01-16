package queries

const (
	INSERT_LEAGUE = `
		insert into league(league_name, country_id, confederation_id)
		values (?, ?, ?)
	`

	FIND_LEAGUES_BY_COUNTRY_ID = `
		select l.league_id, l.league_name,
			c.country_id, c.country_name,
			c1.continent_id, c1.continent_name,
			cf.confederation_id, cf.confederation_name,			
			c2.continent_id, c2.continent_name
		from league l
			inner join country c on c.country_id = l.country_id
			inner join confederation cf on cf.continent_id = c.continent_id
			inner join continent c1 on c1.continent_id = c.continent_id
			inner join continent c2 on c2.continent_id = cf.continent_id
		where l.country_id = ?
	`
)
