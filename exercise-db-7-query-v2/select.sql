SELECT id, CONCAT(first_name, ' ', last_name) AS student_name, student_class, final_score, absent from reports WHERE final_score < 70