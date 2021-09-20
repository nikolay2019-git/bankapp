package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	Card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return Card
}

func Withdraw(card types.Card, amount types.Money) types.Card {

	// Карта не активна
	// сумма снятия больше баланса
	//Сумма снятия меньше 0
	//Сумма снятия больше лимита в 20000 единиц  за один вызов функции
	// на карте сумма в копейках

	if card.Active == false {
		return card
	}
	if card.Balance == 0 {
		return card
	}
	if amount < 0 {
		return card
	}
	if amount > 20_000_00 {
		return card
	}
	if card.Balance > amount {
		newBalance := card.Balance - amount
		card.Balance = newBalance
	}

	return card
}

func Deposit(card *types.Card, amount types.Money) {
	//1. Денþги заùислāĀтсā толþко на активнуĀ карту
	//2. Денþги заùислāĀтсā даже при отриøателþном балансе
	//3. На карту можно заùислитþ не более 50_000 сомони за раз

	if !card.Active {
		return
	}
	if amount > 50_000_00 {
		return
	}
	if amount < 0 {
		return
	}

	card.Balance += amount

}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	if card.Active == false {
		return
	}
	if card.Balance <= 0 {
		return
	}
	profit := types.Money(int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear)
	if profit > 500000 {
		return
	}

	card.Balance += profit

}

func Total(cards []types.Card) types.Money {

	//Высчитывает сумму денег на всех ваших картах
	//Отрицательные суммы и суммы на заблокированных картах игнорируются
	sum := types.Money(0)
	for _, card := range cards {
		if card.Active {
			if card.Balance > 0 {
				sum += card.Balance
			}
		}

	}
	return sum
}
