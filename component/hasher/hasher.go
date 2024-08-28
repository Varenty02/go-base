package hasher

import "golang.org/x/crypto/bcrypt"
type hasher struct {
}

func NewHasher() *hasher {
	return &hasher{}
}
func(*hasher) HashPassword(password string) (string, error) {
	// Cost là mức độ phức tạp của hash, có thể điều chỉnh từ 4 đến 31.
	// Thường dùng 10 như một giá trị mặc định hợp lý.
	cost := 10

	// Tạo hash mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func(*hasher) CheckPasswordAndHash(password, hash string) bool {
    // So sánh mật khẩu và hash
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
        return false
    }
    return true
}