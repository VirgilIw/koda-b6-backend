package repository

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
	"github.com/joho/godotenv"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/config"
	"github.com/virgiIw/koda-b6-coffeshopdb/internal/dto"
)

// assert dipakai untuk mengecek apakah hasil sesuai dengan yang diharapkan
// karena test butuh data → jadi kita setup sendiri
// TDD (Test Driven Development):
// 1. Tulis test terlebih dahulu → test harus gagal (RED)
// 2. Tulis code secukupnya untuk membuat test lulus (GREEN)
// 3. Refactor code tanpa mengubah behavior (REFACTOR)

// kenapa butuh tdd?
// - memastikan code sesuai dengan kebutuhan
// - mengurangi bug
// - membuat design code lebih baik

func TestCreateUser_Success(t *testing.T) {

	// Load file .env karena Go tidak otomatis membaca .env
	// Path "../../.env" karena test dijalankan dari folder internal/repository
	godotenv.Load("../../.env")

	// Context kosong untuk kebutuhan function repository
	// karena di test tidak ada request HTTP
	ctx := context.Background()

	// Inisialisasi koneksi database dari config
	db, err := config.InitDB()

	// Pastikan koneksi DB berhasil
	// kalau gagal → test tetap jalan (karena pakai assert)
	assert.NoError(t, err)

	// Buat instance repository
	// Redis diisi nil karena tidak dipakai di test
	repo := NewUserRepository(db, nil)

	// Simulasi request user register
	req := dto.AuthRegisterRequest{
		FullName: "Bor Success",
		Email:    "success@gmail.com", //harus unik kalau test dijalankan berkali-kali
		Password: "123",
	}

	// Panggil function yang ingin dites
	// → akan insert data ke database
	user, err := repo.CreateUser(ctx, req)

	// CLEANUP (PENTING)
	// defer = dijalankan di akhir function test
	// digunakan untuk menghapus data agar DB tetap bersih
	// tanpa ini → data akan numpuk dan bisa kena duplicate error
	defer db.Exec(ctx, "DELETE FROM users WHERE email = $1", req.Email)

	// Validasi tidak ada error saat insert
	assert.NoError(t, err)

	// Validasi ID terisi (berarti data berhasil masuk DB)
	assert.NotEmpty(t, user.Id)

	// Validasi data yang disimpan sesuai dengan input
	assert.Equal(t, req.Email, user.Email)
}

func TestGetUserById_Success(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()

	db, _ := config.InitDB()
	repo := NewUserRepository(db, nil)

	// insert dulu (setup data)
	req := dto.AuthRegisterRequest{
		FullName: "Test User",
		Email:    fmt.Sprintf("user_%d@test.com", time.Now().UnixNano()),
		Password: "123",
	}

	user, _ := repo.CreateUser(ctx, req)

	defer db.Exec(ctx, "DELETE FROM users WHERE id = $1", user.Id)

	// test ambil by id
	result, err := repo.GetById(ctx, user.Id)

	assert.NoError(t, err)
	assert.Equal(t, user.Id, result.Id)
	assert.Equal(t, req.Email, result.Email)
}

func TestGetUserByEmail_Success(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()

	db, _ := config.InitDB()
	repo := NewUserRepository(db, nil)

	req := dto.AuthRegisterRequest{
		FullName: "Test Email",
		Email:    fmt.Sprintf("email_%d@test.com", time.Now().UnixNano()),
		Password: "123",
	}

	user, _ := repo.CreateUser(ctx, req)

	defer db.Exec(ctx, "DELETE FROM users WHERE id = $1", user.Id)

	result, err := repo.GetByEmail(ctx, req.Email)

	assert.NoError(t, err)
	assert.Equal(t, req.Email, result.Email)
}

func TestUpdateUser_Success(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()

	db, _ := config.InitDB()
	repo := NewUserRepository(db, nil)

	req := dto.AuthRegisterRequest{
		FullName: "Before Update",
		Email:    fmt.Sprintf("update_%d@test.com", time.Now().UnixNano()),
		Password: "123",
	}

	userReg, _ := repo.CreateUser(ctx, req)

	defer db.Exec(ctx, "DELETE FROM users WHERE id = $1", userReg.Id)

	// ambil data lengkap dulu
	user, _ := repo.GetById(ctx, userReg.Id)

	// update
	user.FullName = "After Update"

	err := repo.UpdateUser(ctx, user)

	assert.NoError(t, err)

	// cek lagi ke DB
	updated, _ := repo.GetById(ctx, user.Id)

	assert.Equal(t, "After Update", updated.FullName)
}

func TestDeleteUser_Success(t *testing.T) {
	godotenv.Load("../../.env")
	ctx := context.Background()

	db, _ := config.InitDB()
	repo := NewUserRepository(db, nil)

	req := dto.AuthRegisterRequest{
		FullName: "Delete User",
		Email:    fmt.Sprintf("delete_%d@test.com", time.Now().UnixNano()),
		Password: "123",
	}

	user, _ := repo.CreateUser(ctx, req)
	// Test harus self-contained (berdiri sendiri)
	// Artinya:
	// tidak bergantung data yang sudah ada di DB
	// tidak asumsi user tertentu sudah ada
	// pakai data yang dibuat khusus untuk test ini (unique)

	// karna test delete tidak perlu cleanup dengan defer, berbeda dengan test yang lain
	err := repo.DeleteUser(ctx, user.Id)

	assert.NoError(t, err)

	// pastikan sudah tidak ada
	_, err = repo.GetById(ctx, user.Id)

	assert.Error(t, err)
}
