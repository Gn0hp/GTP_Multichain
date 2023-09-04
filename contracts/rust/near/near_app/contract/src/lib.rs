mod common;

use std::collections::HashMap;
use near_sdk::{AccountId, env, near_bindgen};
use near_sdk::borsh::{BorshDeserialize, BorshSerialize};
use crate::common::NEP20;

#[near_bindgen]
#[derive(BorshSerialize, BorshDeserialize)]
pub struct GToken {
    pub name: String,
    pub symbol: String,
    pub decimals: u8,
    pub total_supply: u128,
    pub owner: AccountId,
    pub zero_account: AccountId,
    pub balances: HashMap<AccountId, u128>,
    pub allowances: HashMap<AccountId, HashMap<AccountId, u128>>
}

#[near_bindgen]
impl GToken {
    pub fn new(name: String, symbol: String, decimals: u8, total_supply: u128) -> Self {
        let mut balances = HashMap::new();
        let mut allowances = HashMap::new();
        Self {
            name,
            symbol,
            decimals,
            total_supply,
            owner: env::signer_account_id(),
            zero_account: env::predecessor_account_id(),
            balances,
            allowances,
        }
    }

    fn _mint(&mut self, account_id: AccountId, amount: u128, log_event: bool) {
        assert_eq!(env::signer_account_id(), self.owner, "Only owner can mint");
        assert!(amount > 0, "Amount must be greater than 0");

        self._before_token_transfer(*self.zero_account, account_id.clone(), amount);
        let balance = self.balances.get(&account_id).unwrap_or(&0);
        self.balances.insert(account_id.clone(), balance + amount);
        self.total_supply += amount;

        if log_event {
            near_sdk::log!("Minted {} to {}", amount, account_id);
        }
        self._after_token_transfer(*self.zero_account, account_id, amount);
    }

    fn _burn(&mut self, account_id: AccountId, amount: u128, log_event: bool) {
        assert_eq!(env::predecessor_account_id(), self.owner, "Only owner can burn");
        assert!(amount > 0, "Amount must be greater than 0");

        self._before_token_transfer(account_id.clone(), *self.zero_account, amount);
        let balance = self.balances.get(&account_id).unwrap_or(&0);
        assert!(balance >= &amount, "Not enough balance");
        self.balances.insert(account_id, balance - amount);
        self.total_supply -= amount;

        if log_event{
            near_sdk::log!("Burned {} from {}", amount, account_id);
        }
        self._after_token_transfer(account_id.clone(), *self.zero_account, amount);
    }
    pub fn mint(&mut self, account_id: AccountId, amount: u128) {
        // assert some conditions
        self._mint(account_id, amount, true);
    }
    pub fn burn(&mut self, account_id: AccountId, amount: u128) {
        self._burn(account_id, amount, true);
    }

    fn _before_token_transfer(&mut self, from: AccountId, to: AccountId, amount: u128) {
        todo!()
    }
    fn _after_token_transfer(&mut self, from: AccountId, to: AccountId, amount: u128) {
        todo!()
    }
}
#[near_bindgen]
impl NEP20 for GToken {
    fn name(&self) -> String {
        self.name.clone()
    }

    fn symbol(&self) -> String {
        self.symbol.clone()
    }

    fn decimals(&self) -> u8 {
        self.decimals
    }

    fn total_supply(&self) -> u128 {
        self.total_supply
    }

    fn balance_of(&self, account_id: AccountId) -> u128 {
        *self.balances.get(&account_id).unwrap_or(&0)
    }

    fn transfer(&mut self, to: AccountId, amount: u128) -> bool {
        todo!()
    }

    fn transfer_from(&mut self, from: AccountId, to: AccountId, amount: u128) -> bool {
        todo!()
    }

    fn allowance(&self, owner: AccountId, spender: AccountId) -> u128 {
        todo!()
    }

    fn approve(&mut self, spender: AccountId, amount: u128) -> bool {
        todo!()
    }

}


#[cfg(test)]
mod tests {
    #[test]
    fn test_transfer() {}
}