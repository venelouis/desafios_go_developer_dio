package main

import "testing"

func TestSoma(t *testing.T) {
	// Arrange
	teste := soma(3, 2, 1)
	// Act
	resultado := 6
	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSoma2(t *testing.T) {
	// Arrange
	teste := soma(3, 2, 1)
	// Act
	resultado := 7
	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult(t *testing.T) {
	// Arrange
	teste := multiplica(10, 10)
	// Act
	resultado := 100
	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMult2(t *testing.T) {
	// Arrange
	teste := multiplica(10, 10)
	// Act
	resultado := 2560
	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSub(t *testing.T) {
	// Arrange
	teste := subtrai(10, 5)
	// Act
	resultado := -5
	// Assert
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSub2(t *testing.T) {
	// Arrange
	// Act
	// Assert
}
	teste := s