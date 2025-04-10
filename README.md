# Technical Test Sharing Vision | Golang Post Article

Aplikasi sederhana untuk CRUD (Create, Read, Update, Delete) artikel ini dibangun dengan arsitektur modular menggunakan framework Echo untuk backend. Backend ini dirancang dengan struktur service-layer untuk memisahkan logika bisnis dari akses data.

Seluruh aplikasi dikemas ke dalam Docker untuk memastikan lingkungan yang konsisten dan kemudahan dalam pengelolaan serta penyebaran aplikasi.

## Fitur Backend 

- **Validation**: validasi data yang diinput pengguna, agar sesuai dengan ketentuan yang berlaku.  
- **Error Handler**: Penanganan kesalahan yang efektif untuk setiap permintaan.  
- **Migration**: Migrasi tabel data otomatis dari sistem ke database mysql.

## CORS (Yes)
Domain dan Port yang terdaftar:
- http://localhost
- http://localhost:8080
- http://localhost:8888

## Endpoint API 
URL for published documentation : https://documenter.getpostman.com/view/23897308/2sB2cX9h2P
- **[GET]** /api/article/{limit}/{offset}
- **[GET]** /api/article/{id}
- **[POST]** /api/article
- **[POST]** /api/article/{id}
- **[DELETE]** /api/article/{id}

## Prasyarat  

Pastikan Anda memiliki:  

- [Docker](https://www.docker.com/get-started) terpasang di komputer Anda.  
- [Docker Compose](https://docs.docker.com/compose/install/) terinstal.  

## Cara Menjalankan Aplikasi  

1. **Clone repositori**:  

   ```bash  
   git clone https://github.com/Dstar18/GoArticle
   cd GoArticle  

2. **Jalankan docker compose**:  

   ```bash  
   docker compose -f 'docker-compose.yml' up -d --build 'mysql'
 
   docker compose -f 'docker-compose.yml' up -d --build 'adminer'

   docker compose -f 'docker-compose.yml' up -d --build 'app'

## Access Aplikasi dan Port  

- Apps: localhost:3000
- MySQL: 3306
- Adminer: localhost:8081