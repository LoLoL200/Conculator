# Conculator

**Conculator** is a financial accumulation calculator written in Go.  
It calculates the future value of a savings plan based on initial deposit, monthly contributions, interest rate, and duration.  
Additionally, it converts results into USD and EUR using fixed exchange rates.

---

## 💡 Features

- Calculates:
  - Total contributions over time
  - Final amount with compound interest
  - Interest earned (profit)
- Converts results into:
  - US Dollars (USD)
  - Euros (EUR)
- User-friendly terminal output
- No external dependencies (just Go standard library)

---

## 🧮 Calculation Logic

- **Compound interest formula**:
  - Initial deposit grows with monthly compounding
  - Monthly contributions are also compounded as a series
- Exchange rates:
  - USD: 1 USD = 38.5 UAH
  - EUR: 1 EUR = 42.1 UAH

---

## 📦 Installation

### Requirements

- Go 1.16 or newer

### Clone and Build

```bash
git clone https://github.com/LoLoL200/Conculator.git
cd Conculator
go build -o conculator
