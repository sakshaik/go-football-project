package queries

const (
	FIND_ALL_POSITIONS = `
		select * from position
	`

	ADD_PLAYER = `
		INSERT INTO 
		player (first_name, last_name, age, city_id, position_id)
		VALUES (?,?,?,?,?)
	`
	REMOVE_PLAYER = `
		delete from player where player_id = ?
	`
	FIND_PLAYER_BASE_QUERY = `
		select 
			p.player_id,p.first_name, p.last_name, p.age,
			pos.position_id, pos.position_short_name, pos.position_long_name, 
			c.city_id, c.city_name,			
			cntry.country_id,cntry.country_name,
			cont.continent_id, cont.continent_name 
		from player p 
			inner join city c on c.city_id = p.city_id
			inner join position pos on pos.position_id = p.position_id
			inner join country cntry on cntry.country_id = c.country_id
			inner join continent cont on cont.continent_id = cntry.continent_id		
	`
	PLAYER_ID_CLAUSE = `
		p.player_id = ?
	`
	PLAYER_FIRST_NAME_CLAUSE = `
		p.first_name like ?
	`
	PLAYER_LAST_NAME_CLAUSE = `
		p.last_name like ?
	`
	PLAYER_IDS_CLAUSE = `
		p.player_id IN (?)
	`
)
