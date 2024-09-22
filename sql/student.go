package sql

const GetStudentByEmail = `
	select s.id, s.email, s.password_digest, s.is_admin, s.is_student, s.is_company, s.is_university from public.users s
	where s.email = '%s'
`
const GetStudentProfile = `
select s.id, s.email, s.first_name, s.last_name, s.middle_name, s.phone_number, p.birthdate, p.sex, p.description, p.competencies  from public.users s
join public.profiles p on p.user_id = s.id                                                            
where s.email = '%s'
`
