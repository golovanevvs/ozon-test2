var dfs func(currency string, amount float64, usedBanks map[int]bool)
   dfs = func(currency string, amount float64, usedBanks map[int]bool) {
       // Базовый случай: если валюта — USD, обновляем максимум
       if currency == "USD" {
           if amount > maxDollars {
               maxDollars = amount
           }
           return
       }

       // Мемоизация: если состояние уже встречалось, пропускаем
       key := fmt.Sprintf("%s_%.10f", currency, amount)
       if val, ok := memo[key]; ok && val >= amount {
           return
       }
       memo[key] = amount

       // Перебираем все банки и их курсы
       for bankIdx, bank := range banks {
           if usedBanks[bankIdx] { // Банк уже использован
               continue
           }
           usedBanks[bankIdx] = true // Помечаем банк как использованный

           // Перебираем все возможные обмены в этом банке
           for _, exchange := range bank.exchanges {
               if exchange.from == currency {
                   // Вычисляем новую сумму после обмена
                   newAmount := amount / float64(exchange.n) * float64(exchange.m)
                   // Рекурсивный вызов для новой валюты и суммы
                   dfs(exchange.to, newAmount, usedBanks)
               }
           }
           usedBanks[bankIdx] = false // Возвращаем банк в неиспользованное состояние
       }
   }