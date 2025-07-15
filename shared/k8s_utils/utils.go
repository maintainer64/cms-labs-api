package k8s_utils

import (
	"fmt"
	"strings"
	"time"
)

func UsernameByEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return email
}

func NormalizeK8SEntityName(entityName string) string {
	var result strings.Builder

	for _, r := range strings.ToLower(entityName) {
		switch {
		case r >= 'a' && r <= 'z':
			result.WriteRune(r)
		case r >= '0' && r <= '9':
			result.WriteRune(r)
		case r == '-':
			// Дефисы разрешены, но не в начале/конце и не подряд
			if result.Len() > 0 && result.String()[result.Len()-1] != '-' {
				result.WriteRune(r)
			}
		}
	}

	// Удаляем дефисы с конца
	normalized := strings.TrimRight(result.String(), "-")

	// Если после нормализации имя пустое, генерируем случайное
	if normalized == "" {
		normalized = fmt.Sprintf("user-%x", time.Now().UnixNano())
	}

	// Ограничиваем максимальную длину
	//(63 символа - ограничение Kubernetes)
	// (40 символов - чтобы ещё дописать можно было что-то)
	if len(normalized) > 63 {
		normalized = normalized[:63]
		normalized = strings.TrimRight(normalized, "-")
	}

	return normalized
}
