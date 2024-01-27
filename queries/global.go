package queries

const (
	WHERE = " where "
	AND   = " and "

	GET_ALL_CITIES = `
		select 
			c.city_id, c.city_name, 
			cntry.country_id,cntry.country_name, 
			cont.continent_id, cont.continent_name 
		from city c 
			inner join country cntry on cntry.country_id = c.country_id
			inner join continent cont on cont.continent_id = cntry.continent_id
	`

	ADD_CITY = `
		INSERT INTO 
		city (city_name, country_id)
		VALUES (?,?)
	`

	GET_ALL_CONFEDERATIONS = `
		select 
			conf.confederation_id,conf.confederation_name,conf.confederation_description, 
			cont.continent_id, cont.continent_name 
		from confederation conf 
			inner join continent cont on cont.continent_id = conf.continent_id
	`

	GET_ALL_CONTINENTS = `
		select * from continent
	`

	GET_ALL_COUNTRIES = `
		select 
			cntry.country_id,cntry.country_name, 
			cont.continent_id, cont.continent_name 
		from country cntry inner 
			join continent cont on cont.continent_id = cntry.continent_id
	`

	GET_COUNTRY = `
		select * from country where country_id = ? and country_name = ?
	`

	GET_COUNTRY_BY_NAME = `
		select * from country where country_name = ?
	`

	GET_COUNTRY_BY_ID = `
		select * from country where country_id = ?
	`
)
