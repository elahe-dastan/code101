class Solution:
    def maxProfit(self, prices) -> int:
        max_profit = 0
        has_stock = False
        for i in range(len(prices)-1):
            if prices[i] < prices[i + 1]:
                if has_stock:
                    continue
                buy_price = prices[i]
                has_stock = True

            if prices[i] > prices[i+1]:
                if not has_stock:
                    continue
                max_profit += prices[i] - buy_price
                has_stock = False

        if has_stock:
            max_profit += prices[-1] - buy_price

        return max_profit


S = Solution()
print(S.maxProfit([7,1,5,3,6,4]))
print(S.maxProfit([1,2,3,4,5]))
print(S.maxProfit([7,6,4,3,1]))

