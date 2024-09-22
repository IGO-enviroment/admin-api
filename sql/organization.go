package sql

const GetOrgProfile = `
select s.id, s.email, s.first_name, s.last_name, s.middle_name, s.phone_number, o.name, o.description, o.inn, o.kpp, o.ogrn, o.place  from public.users s
join public.organizations o on o.id = s.organization_id                                                            
where s.email = '%s'
`
