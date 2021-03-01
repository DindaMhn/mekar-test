LINKAJA TEST

1. Teknologi yang Digunakan
    - Golang
    - MySQL
2. Cara Menjalankan Program
        1. Import file db_mekar di dalam folder database 
        2. Ubah DB_USER dan DB_PASSWORD di dalam file .env pada folder files dan sesuaikan
        2. Ketik di terminal dengan command "go run main/main.go"
3. Output Program
    - UNTUK USER
        1. CREATE USER
            - URL => /user
            - Method : POST
            - Body :
                - Contoh Body
                {
                    "nama_user": "Dinda Maharani",
                    "tanggal_lahir": "1998-11-01",
                    "no_ktp": "16710941198003",
                    "pekerjaan": "4",
                    "pendidikan": "5"
                }
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/user
                - Code : 202
                - Contoh Content : 
                    {
                        "status": 201,
                        "message": "Success",
                        "data": {
                            "id_user": "223b1535-4e24-4580-a7c1-3f3788ca8222",
                            "nama_user": "Dinda Maani",
                            "tanggal_lahir": "1998-11-01",
                            "no_ktp": "16710941198003",
                            "pekerjaan": "4",
                            "pendidikan": "5",
                            "status": "A"
                        }
                    }
            - Response jika Gagal :
                - Contoh URL :
                    localhost:8080/user
                - Code : 404 
                - Method : POST
                - Contoh Content :
                    {
                        "status": 404
                        "message": "Failed",
                        "data : null
                    }
        2. GET ALL DATA USER
            - URL => /user
            - Method : GET
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/user
                - Code : 200
                - Contoh Content : 
                {
                        "status": 200,
                        "message": "Success",
                        "data": [
                            {
                                "id_user": "223b1535-4e24-4580-a7c1-3f3788ca8222",
                                "nama_user": "Dinda Maani",
                                "tanggal_lahir": "1998-11-01",
                                "no_ktp": "16710941198003",
                                "pekerjaan": "4",
                                "pendidikan": "5",
                                "status": "A"
                            },
                            {
                                "id_user": "4bfb21b6-4bee-4265-bd24-2e00906b5cf5",
                                "nama_user": "Dinda Maharani",
                                "tanggal_lahir": "1998-11-01",
                                "no_ktp": "16710941198003",
                                "pekerjaan": "4",
                                "pendidikan": "5",
                                "status": "A"
                            }
                        ]
                    }
            - Response jika Akun Tidak Ditemukan :
                - Contoh URL :
                    localhost:8080/user
                - Code : 404
                - Contoh Content : 
                    {
                        "status": 404,
                        "message": "Not Found",
                        "data": null
                    }
    3. GET USER BY ID
        - URL => /user/{id}
        - Method : GET
        - Response jika Sukses :
            - Contoh URL :
                localhost:8080/user/4bfb21b6-4bee-4265-bd24-2e00906b5cf5
            - Code : 200
            - Contoh Content : 
               {
                    "status": 200,
                    "message": "Success",
                    "data": {
                        "id_user": "4bfb21b6-4bee-4265-bd24-2e00906b5cf5",
                        "nama_user": "Dinda Maharani",
                        "tanggal_lahir": "1998-11-01",
                        "no_ktp": "16710941198003",
                        "pekerjaan": "4",
                        "pendidikan": "5",
                        "status": "A"
                    }
                }
        - Response jika Akun Tidak Ditemukan :
            - Contoh URL :
                localhost:8080/user/123
            - Code : 404
            - Contoh Content : 
                {
                    "status": 404,
                    "message": "Not Found",
                    "data": null
                }
    4. UPDATE USER 
            - URL => /user/{id}
            - Method : PUT
            - Body :
                - Contoh Body
                {
                    "nama_user": "Dinda Maharani",
                    "tanggal_lahir": "1998-11-01",
                    "no_ktp": "16710941198003",
                    "pekerjaan": "4",
                    "pendidikan": "5"
                }
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/user/4bfb21b6-4bee-4265-bd24-2e00906b5cf5
                - Code : 200
                - Contoh Content : 
                    {
                        "status": 200,
                        "message": "Success",
                        "data": {
                            "id_user": "4bfb21b6-4bee-4265-bd24-2e00906b5cf5",
                            "nama_user": "Dinda Maharani",
                            "tanggal_lahir": "1998-11-01",
                            "no_ktp": "16710941198003",
                            "pekerjaan": "4",
                            "pendidikan": "5"
                        }
                    }
            - Response jika Gagal :
                - Contoh URL :
                    localhost:8080/user/123
                - Code : 404 
                - Method : POST
                - Contoh Content :
                {
                        "status": 404,
                        "message": "Failed",
                        "data": null
                    }
    - UNTUK ADMIN
        1. LOGIN
            - URL => /admin/login
            - Method : POST
            - Body :
                - Contoh Body
                {
                    "username": "admin",
                    "password": "admin
                }
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/admin/login
                - Code : 200
                - Contoh Content : 
                {
                        "status": 200,
                        "message": "Success",
                        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21LZXkiOiJyYWhhc2lhZG9uZyIsImV4cGlyZWRBdCI6IjIwMjEtMDMtMTIgMjM6MTQ6NDMiLCJuYW1lIjoiYWRtaW4ifQ.WbIWu06N6Ljk25pY2jhfOVeFSM_sXx5h5C5ZgD_6eV8",
                        "data": {
                            "username": "admin",
                            "password": "$2y$14$I82sSL0y4ZvGq6K1KKbPXe.Kyjh67he17Xu714GrlLd2nWK.JtCZG"
                        }
                    }
            - Response jika username/password salah :
                - Contoh URL :
                    localhost:8080/admin/login
                - Code : 400 
                - Body :
                - Contoh Body
                    {
                        "username": "admin",
                        "password": "admin
                    }
                - Contoh Content : 
                {
                        "status": 400,
                        "message": "Login User Failed",
                        "data": null
                    }
            
        2. GET ALL DATA USER
            - URL => /admin/user
            - Method : GET
            - Response jika Sukses :
                - Contoh URL : localhost:8080/admin/user
                - Code : 200
                - Contoh Content : 
                    sama seperti pada GET ALL DATA USER di atas
            - Response jika Akun Tidak Ditemukan :
                - Contoh URL : localhost:8080/user
                - Code : 404
                - Contoh Content : 
                    sama seperti pada GET ALL DATA USER di atas                           
            - Response jika token tidak sesuai :
                You are not allowed to access this service
        3. GET USER BY ID
            - URL => /admin/user/{id}
            - Method : GET
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/admin/user/4bfb21b6-4bee-4265-bd24-2e00906b5cf5
                - Code : 200
                - Contoh Content : 
                sama seperti pada GET USER BY ID di atas
            - Response jika Akun Tidak Ditemukan :
                - Contoh URL :
                    localhost:8080/user/123
                - Code : 404
                - Contoh Content : 
                    sama seperti pada GET USER BY ID di atas
                - Response jika token tidak sesuai :
                    You are not allowed to access this service
        4. UPDATE USER 
            - URL => /admin/user/{id}
            - Method : PUT
            - Body :
                - Contoh Body
                {
                    "nama_user": "Dinda Maharani",
                    "tanggal_lahir": "1998-11-01",
                    "no_ktp": "16710941198003",
                    "pekerjaan": "4",
                    "pendidikan": "5"
                }
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/admin/user/4bfb21b6-4bee-4265-bd24-2e00906b5cf5
                - Code : 200
                - Contoh Content : 
                    sama seperti pada UPDATE USER di atas
            - Response jika Gagal :
                - Contoh URL :
                    localhost:8080/admin/user/123
                - Code : 404 
                - Contoh Content :
                sama seperti pada UPDATE USER di atas
            - Response jika token tidak sesuai :
                    You are not allowed to access this service
        5. DELETE USER
            - URL => /admin/user/{id}
            - Method : DELETE
            - Response jika Sukses :
                - Contoh URL :
                    localhost:8080/admin/user/223b1535-4e24-4580-a7c1-3f3788ca8222
                - Code : 200
                - Contoh Content : 
                    {
                        "status": 200,
                        "message": "User with ID 223b1535-4e24-4580-a7c1-3f3788ca8222 Success Deleted",
                        "data": "223b1535-4e24-4580-a7c1-3f3788ca8222"
                    }
            - Response jika Gagal :
                - Contoh URL :
                    localhost:8080/admin/user/123
                - Code : 404 
                - Contoh Content :
                    {
                        "status": 404,
                        "message": "Failed",
                        "data": null
                    }
            - Response jika token tidak sesuai :
                    You are not allowed to access this service
    