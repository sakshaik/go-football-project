package queries

const (
	INSERT_PLAYER = `
		INSERT INTO 
		player (first_name, last_name, age, city_id)
		VALUES (?,?,?,?)
	`
	FIND_PLAYER_BY_ID = `
		select 
			p.player_id,p.first_name, p.last_name, p.age,
			c.city_id, c.city_name, 
			cntry.country_id,cntry.country_name,
			cont.continent_id, cont.continent_name 
		from player p 
			inner join city c on c.city_id = p.city_id
			inner join country cntry on cntry.country_id = c.country_id
			inner join continent cont on cont.continent_id = cntry.continent_id
		where p.player_id = ?
	`
)
