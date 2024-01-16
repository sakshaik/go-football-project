insert or ignore into city (city_name, country_id) values ('London',(select country_id from country where country_name='England'))
insert or ignore into city (city_name, country_id) values ('Manchester',(select country_id from country where country_name='England'))
insert or ignore into city (city_name, country_id) values ('Liverpool',(select country_id from country where country_name='England'))
insert or ignore into city (city_name, country_id) values ('Newcastle',(select country_id from country where country_name='England'))