package payments

import "errors"

// Payouts returns (winnerKobo, loserKobo, commissionKobo, totalPotKobo)
func Payouts(stakeKobo int64) (winnerKobo, loserKobo, commissionKobo, totalPotKobo int64, err error) {
	if stakeKobo <= 0 {
		return 0, 0, 0, 0, errors.New("invalid stake")
	}
	totalPotKobo = stakeKobo * 2

	// commission = 10% of total pot
	commissionKobo = (totalPotKobo * 10) / 100

	// loser keeps 50% of their stake
	loserKobo = (stakeKobo * 50) / 100

	// winner = stake + 30% of stake
	winnerKobo = stakeKobo + (stakeKobo*30)/100

	// ensure sums add up
	remainder := (totalPotKobo - commissionKobo) - (winnerKobo + loserKobo)
	if remainder > 0 {
		commissionKobo += remainder
	} else if remainder < 0 {
		return 0, 0, 0, 0, errors.New("payout sums mismatch")
	}
	return winnerKobo, loserKobo, commissionKobo, totalPotKobo, nil
}
