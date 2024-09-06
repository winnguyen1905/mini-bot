# Blitz Bot

## 1. Information

### Exchange Support

- Binance
- Bitget
- Bybit
- OKX
- Kucoin

### Payment Gateway Support

- USDT (BEP20)
- BNB (BEP20)
- Credit Card

### AI Support

- **Model:** gpt-4o-2024-08-06
- **Support:** Text, Image
- **Cost Optimization:** Compress images and scale them down before sending to OpenAI

### Requirements

- **Response Time:** Must be as fast as possible
- **Bot Technology:** Golang + MongoDB
- **Frontend Technology:** React (or Next.js)
- **Payment Gateway:** Lemonsqueezy for credit card, manual block scan for crypto
- **Deployment:** Docker
- **Error Monitoring:** Sentry
- **Design Patterns:** Clean Architecture, Strategy Pattern for crypto exchange

## 2. Features ✨

### Open Future Order

- Accepts text or image from the user, processes it through the AI API to extract JSON in the following format:

    ```json
    {
        "entry": [], // []number or Nil or "market"
        "takeProfit": [], //[]number or Nil
        "stopLoss": [], //[]number or Nil
        "lerverage": 0, //number or Nil
        "volume": 100, // USDT, number or Nil
    }
    ```

- Displays extracted information to the user for verification and allows them to modify the data.
- Lets users select the crypto exchange account to open the order, showing a list of logged-in accounts.
- On order placement, checks for success or failure, and if failed, sends an error to Sentry.
- Supports RR or Fixed Leverage order types

### View and Edit Order

- Allows users to view and edit orders in the selected crypto account, including entry price, Take Profit (TP), Stop Loss (SL), and volume.

>[!note]
> Users have two configuration types:
>
> - **Global Configuration**: Accessible via the settings button.
> - **Order-Specific Configuration**:
>   - Volume
>   - 1R Ratio (if order type is RR)
>   - Leverage (if order type is fixed leverage)
>   - TP Strategy (Limit (Percent distance), Market)
>   - SL Strategy (details depend on user config)
>   - Entry Strategy

### Change SL Price

- **Cron Job:** Runs every 30 seconds (modifiable in code)
- Scans all user orders and adjusts SL if the current Return on Equity (ROE) matches user configuration

### Add SL/TP When Order is Open

- When an order is open (has a position), adds additional TP/SL based on user configuration

### Notify

- Sends messages to users when an order:
  - Is opened
  - Hits stop loss
  - Hits take profit

### Create and Share ROE Image

- Allows users to create and share images showing Return on Equity

### Feature Request Site

- Provides a platform for users to request new features

### Referral System

- Enables users to invite friends and earn a configurable 5% commission when a friend buys a subscription
- Commissions are paid through BEP20 USDT

### Analytics

- PNL by DMY of (each/all) crypto exchange
- Balance of (each/all) crypto exchange
- Total order of (each/all) crypto exchange
- Win ratio of (each/all) crypto exchange

### Premium Feature

Must be configurable

| Feature / Plan          | Free   | Starter                   | Pro                       | Elite                     |
| ----------------------- | ------ | ------------------------- | ------------------------- | ------------------------- |
| Price                   | Free   | $10 Monthly / $100 Yearly | $20 Monthly / $200 Yearly | $50 Monthly / $450 Yearly |
| Link Crypto account     | 1      | 3                         | 5                         | Unlimited                 |
| Limit Open order by day | 3      | 10                        | 50                        | Unlimited                 |
| Limit token per message | 500    | 500                       | 500                       | 500                       |
| Analytics <br>storage   | 3 days | 7 days                    | 1 month                   | unlimited                 |
| Commissions             | 2%     | 3%                        | 5%                        | 15%                       |

Policy: If user expired then lock user account:

- Purchase to upgrade the current plan again
- Choose plan & account to keep.

### Voucher feature

- type: amount, percentage
- plan: free, starter, pro, elite, all
- expired date
- use left: unlimited or number

### Admin Dashboard

- Provides an administrative dashboard for management and monitoring
