SELECT id, nik, CONCAT(first_name, ' ', last_name) AS fullname, date_of_birth, weight, address from people where gender = 'laki-laki' ORDER BY weight DESC LIMIT 5