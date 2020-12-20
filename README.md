# Coffeebucks - pay off that coffee machine

This projact started when I invested in a quite expensive coffee machine for my home. I took the amount out of my savings account as an investment.
Then I wondered, how long would it take to pay it back to myself if I were to pay Starbucks prices for my coffees every day.

This project uses [Bunq](https://bunq.com/invite/Maartje) and an ESP32 (with giant button for fun!) to pay myself 2 euros every time I drink a coffee.
Due to the quite important part of having a good TLS stack (and lack of a good C library) I added a backend service in Go that does the actual transaction. This code is in Go and in `backend/`