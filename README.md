# Technical Test Sharing Vision | Golang Post Article

Aplikasi sederhana CRUD Article, dengan Backend menggunakan framework echo dan modular service-layer structure sedangkan Frontend menggunakan ... dikemas didalam Docker.

## Fitur Backend 

- **Validation**: validasi data yang diinput pengguna, agar sesuai dengan ketentuan yang berlaku.  
- **Error Handler**: Penanganan kesalahan yang efektif untuk setiap permintaan.  
- **Migration**: Migrasi tabel data otomatis dari sistem ke database mysql.

## Endpoint API 


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