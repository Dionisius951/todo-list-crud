# Todo List API
API ini dirancang untuk memenuhi kebutuhan pengelolaan tugas seperti menambahkan, menghapus, memperbarui, dan melihat daftar tugas.

## Cara Menjalankan Server
```bash
go run server.go
```

## Endpoints
### 1. Mendapatkan Data Todos
Untuk mendapatkan data todos, gunakan metode **GET** dan buka URL berikut:
[http://localhost:9000/todos](http://localhost:9000/todos)

Jika data tersedia, maka tampilan akan seperti berikut ini:

```json
[
    {
        "id": 1,
        "value": "Belajar Fisika",
        "status": false
    }
]
```

---

### 2. Menambahkan Data Todos
Untuk menambahkan data todos, gunakan metode **POST** dan buka URL berikut:
[http://localhost:9000/todos](http://localhost:9000/todos)

Inputkan data value sesuai dengan tugas yang ingin ditambahkan. Sebagai contoh:

```json
{
  "value": "Menyelesaikan tutorial dasar Golang",
}
```

Ketika data dikirim, output akan tampak seperti berikut:

```json
  "Berhasil Menambahkan Todo"
```

### 3. Memperbarui Data Todos
Untuk memperbarui data todos, gunakan metode **PUT** dan buka URL berikut:
[http://localhost:9000/todos?id=[id_todos]](http://localhost:9000/todos?id=[id_todos])

Inputkan data value yang ingin diubah. Jika hanya ingin mengubah status, kirim data seperti berikut:

```json
{
  "status": true
}
```

Jika ingin memperbarui seluruh informasi tugas, kirim data seperti ini:

```json
{
    "value": "Input value todos baru",
}
```

Respons jika berhasil memperbarui data:

```json
"Berhasil Memperbarui Todo"
```

---

### 4. Menghapus Data Todos
Untuk menghapus data todos, gunakan metode **DELETE** dan buka URL berikut:
[http://localhost:9000/todos?id=[id_todos]](http://localhost:9000/todos?id=[id_todos])

Jika permintaan berhasil, respons akan seperti berikut:

```json
"Berhasil Menghapus Todo"
```

---

## Catatan
- **ID** adalah parameter query yang wajib digunakan untuk metode **PUT** dan **DELETE**.
- API ini dirancang untuk dijalankan di server lokal dengan port **9000**. Pastikan tidak ada konflik dengan aplikasi lain yang berjalan di port yang sama.
