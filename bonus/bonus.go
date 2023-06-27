package bonus

import "errors"

func CalculateDamage(characterLevel, enemyLevel int) (int, error) {
	if characterLevel <= 0 || enemyLevel <= 0 {
		return 0, errors.New("nível inválido")
	}

	damage := 0

	if characterLevel > 100 {
		if enemyLevel > 100 {
			damage = characterLevel * 2
		} else {
			damage = characterLevel * 20
		}
	} else if enemyLevel > 100 {
		damage = characterLevel * 2
	} else if characterLevel > enemyLevel {
		damage = characterLevel * 10
	} else if characterLevel < enemyLevel {
		damage = characterLevel * 5
	} else {
		damage = characterLevel * 7
	}

	levelDifference := characterLevel - enemyLevel
	if levelDifference > 20 {
		damage *= 5
	} else if levelDifference < -20 {
		damage /= 2
	}

	return damage, nil
}
