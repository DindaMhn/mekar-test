package tools

const (
	DATE_FORMAT        = `2006-01-02`
	CREATE_USER        = `INSERT INTO USER (id_user,nama_user,tanggal_lahir,no_ktp,pekerjaan,pendidikan) VALUES (?,?,?,?,?,?)`
	SELECT_ALL_USER    = `SELECT * FROM USER WHERE status='A'`
	SELECT_USER_BY_ID  = `SELECT * FROM USER WHERE id_user=?`
	UPDATE_DATA_USER   = `UPDATE USER SET nama_user=?,tanggal_lahir=?,no_ktp=?,pekerjaan=?,pendidikan=? WHERE id_user=?`
	SELECT_ADMIN_LOGIN = `SELECT * FROM admin WHERE username=?`
	DELETE_USER        = `UPDATE USER SET status='P' WHERE id_user=?`
)
