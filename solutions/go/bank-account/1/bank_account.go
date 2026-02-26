package account
import "sync"
// Define the Account type here.
type Account struct{
    mu sync.Mutex
    Money int64
    Active bool  
}

func Open(amount int64) *Account {
	if amount <0{
        return nil
    }
    return &Account{Money: amount, Active: true}
}

func (a *Account) Balance() (int64, bool) {
    if !a.Active{
        return 0, false
    }
    return a.Money, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
    if amount <0{
        if amount < -a.Money{
            return 0, false
        }
    }
    a.Money += amount
    
    return a.Balance()
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
    if !a.Active{
        return 0, false
    }
    a.Active = false
    return a.Money, true
}
